package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

const imgPath = "image5.png"

const templateChar = "ล ู ก ค ิ ด ม า เ ล ้ ว อ ฟ ห ่ ไ โ บ ซ บ"
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
		fmt.Printf("%v\n", start)
		fmt.Printf("%v\n", end)

		// DrawRowSegment(newImg, start, end)

		gocv.IMWrite("out.jpg", newImg)

		for i := range start {
			// Character segmentation
			rectTable := GetSegmentChar(imgArr[start[i]:end[i]])

			// testImg := GetImgMat(imgArr[start[i]:end[i]])

			for b := range rectTable {
				// gocv.Rectangle(testImg, rectTable[b], color.RGBA{255, 0, 0, 0}, 1)

				cropImg := CropImgArr(imgArr[start[i]:end[i]], rectTable[b])

				fmt.Printf("%v", MatchTemplate(cropImg, templates[GetRatioBin(len(cropImg), len(cropImg[b]))])[0].char)
			}

			// gocv.IMWrite("out.jpg", testImg)

			println()

		}
	}
}
