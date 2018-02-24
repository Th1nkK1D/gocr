package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"

	"gocv.io/x/gocv"
)

const templateDir = "templates/"
const outputDir = "outputs/"
const mfSize = 3

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Invalid argument: gocr [mode, filename]")
	} else {
		if os.Args[1] == "--gentemp" {
			// --- Generate template ---
			if len(os.Args) < 4 {
				fmt.Println("Invalid argument: gocr --gentemp [charfile] [fontfile]")
			} else {
				fmt.Println("Generating template...")

				GenTemplate(ReadCharList(os.Args[2]), os.Args[3], templateDir)

				fmt.Println("DONE!")
			}

		} else {
			// --- OCR ---

			if len(os.Args) < 3 {
				fmt.Println("Invalid argument: gocr [charfile] [imagefile]")
			} else {
				// Read image
				fmt.Printf("Opening %v...\n", os.Args[2])

				img := gocv.IMRead(os.Args[2], gocv.IMReadGrayScale)

				// Apply auto threshold
				fmt.Println("Applying auto threshold...")

				img = AutoThreshold(img)

				gocv.IMWrite(outputDir+"01_auto_threshold.jpg", img)

				// Apply median filter
				fmt.Println("Applying median filter...")

				imgArr := MedianFilter(GetImgArray(img), (mfSize-1)/2)
				img = GetImgMat(imgArr)

				gocv.IMWrite(outputDir+"02_median_filter.jpg", img)

				// Read template
				fmt.Println("Loading templates...")

				templates := ReadTemplate(ReadCharList(os.Args[1]), templateDir)

				// Row segmentation
				fmt.Println("Rows segmenting...")

				start, end := SplitLine(imgArr)

				DrawRowSegment(img, start, end)

				gocv.IMWrite(outputDir+"03_row_segment.jpg", img)

				// Open output file
				output, err := os.Create(outputDir + "text.txt")
				check(err)
				writer := bufio.NewWriter(output)

				// Character segmentation
				fmt.Println("Characters segmenting and template mathching...")
				fmt.Println(">>")

				for i := range start {
					row := CropImgArr(imgArr, image.Rectangle{image.Point{0, start[i]}, image.Point{len(imgArr[0]), end[i]}})
					rectTable := GetSegmentChar(row)

					rowImg := GetImgMat(row)

					for _, rect := range rectTable {
						gocv.Rectangle(rowImg, rect, color.RGBA{255, 0, 0, 0}, 1)
					}

					gocv.IMWrite(outputDir+"04_character_segment_"+strconv.Itoa(i)+".jpg", rowImg)

					// Template Matching
					for b := range rectTable {
						cropImg := CropImgArr(row, rectTable[b])
						ratioBin := GetRatioBin(len(cropImg), len(cropImg[b]))

						if ratioBin >= 0 && ratioBin < len(templates) {
							// Valid object
							res := MatchTemplate(cropImg, templates[ratioBin])

							fmt.Printf("%v", res[0].char)
							_, err = fmt.Fprintf(writer, "%v", res[0].char)
							check(err)
						} else {
							// Invalid object
							fmt.Printf("?")
							_, err = fmt.Fprintf(writer, "?")
							check(err)
						}
					}

					println()
					_, err = fmt.Fprintf(writer, "\n")
					check(err)
				}

				fmt.Println("<<")

				// Flush buffer and close file
				writer.Flush()
				output.Close()

				fmt.Println("DONE!")

			}
		}
	}
}
