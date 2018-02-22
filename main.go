package main

import (
	"fmt"
	"image"
	"os"

	"gocv.io/x/gocv"
)

const imgPath = "image6.png"

const templateChar = "ล ู ก ค ิ ด ม า เ ล ้ ว อ ฟ ห ่ ไ โ บ ซ บ ใ จ"
const templateDir = "templates/"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--gentemp" {
		// Generate template
		GenTemplate(templateChar, templateDir)
	} else {
		// OCR

		// Read image
		img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

		// Apply auto threshold
		newImg := AutoThreshold(img)

		// Show(newImg)

		// Get image array
		imgArr := GetImgArray(newImg)

		// Read template
		templates := ReadTemplate(templateChar, templateDir)

		// Row segmentation
		start, end := SplitLine(imgArr)
		// fmt.Printf("%v\n", start)
		// fmt.Printf("%v\n", end)

		// DrawRowSegment(newImg, start, end)

		for i := range start {
			// Character segmentation
			row := CropImgArr(imgArr, image.Rectangle{image.Point{0, start[i]}, image.Point{len(imgArr[0]), end[i]}})
			rectTable := GetSegmentChar(row)

			// testImg := GetImgMat(row)

			for b := range rectTable {
				// fmt.Println(rectTable[b])
				// gocv.Rectangle(testImg, rectTable[b], color.RGBA{255, 0, 0, 0}, 1)

				cropImg := CropImgArr(row, rectTable[b])

				// gocv.IMWrite("out"+strconv.Itoa(b)+".jpg", GetImgMat(cropImg))

				fmt.Printf("%v", MatchTemplate(cropImg, templates[GetRatioBin(len(cropImg), len(cropImg[b]))])[0].char)
			}
			// gocv.IMWrite("out.jpg", testImg)

			println()
		}

	}
}
