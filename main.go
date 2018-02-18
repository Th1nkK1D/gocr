package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

const imgPath = "image_c.jpg"

func main() {
	img := gocv.IMRead(imgPath, gocv.IMReadColor)

	height := img.Rows()
	width := img.Cols()
	channels := img.Channels()

	fmt.Printf("%v x %v (%v)\n", width, height, channels)

	arr := GetImgArray(img)
	fmt.Printf("%v x %v x %v", len(arr), len(arr[0]), len(arr[0][0]))

	newImg := GetImgMat(arr)
	fmt.Printf("new: %v x %v (%v)\n", newImg.Rows(), newImg.Cols(), newImg.Channels())

	gocv.IMWrite("out.jpg", newImg)
}
