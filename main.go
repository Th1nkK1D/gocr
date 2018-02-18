package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

const imgPath = "image_c.jpg"

func main() {
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

	hist := GetHistArray(img)

	fmt.Printf("%v", hist)
}
