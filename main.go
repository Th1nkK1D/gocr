package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

const imgPath = "image3.png"

const templateChar = "ล ู ก ค ิ ด ม า เ ล ้ ว อ ฟ ห ่ ไ โ บ"
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
		// templates := ReadTemplate(templateChar, templateDir)

		// Row segmentation
		start, end := SplitLine(imgArr)
		fmt.Printf("%v\n", start)
		fmt.Printf("%v\n", end)

		DrawRowSegment(newImg, start, end)

		gocv.IMWrite("out.jpg", newImg)

		// for i := range start {
		// 	Show(GetImgMat(imgArr[start[i]:end[i]]))
		// }

		// for i := range start {
		// Character segmentation
		// rectTable := GetSegmentChar(imgArr) //[start[0]:end[0]]

		// for b := range rectTable {
		// 	cropImg := CropImgArr(imgArr, rectTable[b])

		// 	fmt.Printf("%v", MatchTemplate(cropImg, templates[GetRatioBin(len(cropImg), len(cropImg[b]))])[0].char)
		// }

		// }

		// for i := range rectTable {
		// 	gocv.Rectangle(newImg, rectTable[i], color.RGBA{255, 0, 0, 0}, 1)
		// }

	}
}
