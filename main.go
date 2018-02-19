package main

import (
	"gocv.io/x/gocv"
)

const imgPath = "image3.png"

func main() {
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

	newImg := AutoThreshold(img)

	Show(newImg)

	imgArr := GetImgArray(newImg)
	start, end := SplitLine(imgArr)

	for i := range start {
		Show(GetImgMat(imgArr[start[i]:end[i]]))
	}
}
