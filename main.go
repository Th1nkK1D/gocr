package main

import (
	"gocv.io/x/gocv"
)

const imgPath = "quiz2_test2.png"

func main() {
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

	newImg := AutoThreshold(img)

	// Show(newImg)

	imgArr := GetImgArray(newImg)

	SegmentChar(imgArr)
	// start, end := SplitLine(imgArr)

	// for i := range start {
	// 	Show(GetImgMat(imgArr[start[i]:end[i]]))
	// }
}
