package main

import (
	"image/color"

	"gocv.io/x/gocv"
)

const imgPath = "image3.png"

func main() {
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

	newImg := AutoThreshold(img)

	// Show(newImg)

	imgArr := GetImgArray(newImg)

	// start, end := SplitLine(imgArr)

	// for i := range start {
	// 	Show(GetImgMat(imgArr[start[i]:end[i]]))
	// }

	rectTable := GetSegmentChar(imgArr)

	// fmt.Printf("%v\n", rectTable)

	for i := range rectTable {
		gocv.Rectangle(newImg, rectTable[i], color.RGBA{255, 0, 0, 0}, 1)
	}

	gocv.IMWrite("out.jpg", newImg)
}
