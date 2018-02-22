package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

const imgPath = "image4.png"

const templateChar = "ฟ ห ก ด เ ้ ่ า ส ว ง ๆ ไ พ ั ี ร น ย บ ล ค"
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
		// start, end := SplitLine(imgArr)

		// for i := range start {
		// Character segmentation
		rectTable := GetSegmentChar(imgArr) //[start[0]:end[0]]

		for b := range rectTable {
			cropImg := CropImgArr(imgArr, rectTable[b])

			fmt.Printf("%v", MatchTemplate(cropImg, templates[GetRatioBin(len(cropImg), len(cropImg[b]))])[0].char)
		}

		// }

		// for i := range rectTable {
		// 	gocv.Rectangle(newImg, rectTable[i], color.RGBA{255, 0, 0, 0}, 1)
		// }

	}
}
