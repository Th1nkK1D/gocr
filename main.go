package main

import (
	"os"

	"gocv.io/x/gocv"
)

const imgPath = "templates/1.png"

const templateChar = "ฟ ห ก ด เ ้ ่ า ส ว ง ๆ ไ พ ั ี ร น ย บ ล"
const templateDir = "templates/"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--gentemp" {
		// Generate template
		GenTemplate(templateChar, templateDir)
	} else {
		// OCR

		img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

		newImg := AutoThreshold(img)

		Show(newImg)

		imgArr := GetImgArray(newImg)

		// start, end := SplitLine(imgArr)

		// for i := range start {
		// 	Show(GetImgMat(imgArr[start[i]:end[i]]))
		// }

		// rectTable := GetSegmentChar(imgArr)

		// fmt.Printf("%v\n", rectTable)

		// for i := range rectTable {
		// 	gocv.Rectangle(newImg, rectTable[i], color.RGBA{255, 0, 0, 0}, 1)
		// }

		// gocv.IMWrite("out.jpg", newImg)

		// templates := ReadTemplate(templateChar, templateDir)

		// fmt.Println(templates)

		// test := [][][]uint8{{{1}, {2}, {1}, {1}, {2}, {1}}, {{2}, {3}, {2}, {2}, {3}, {2}}, {{1}, {2}, {1}, {1}, {2}, {1}}}

		Show(GetImgMat(Resize(imgArr, len(imgArr)/2, len(imgArr[0])/2)))

	}
}
