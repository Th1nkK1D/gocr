package main

import (
	"gocv.io/x/gocv"
)

const imgPath = "image.jpg"

func main() {
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

	newImg := AutoThreshold(img)

	Show(newImg)
}
