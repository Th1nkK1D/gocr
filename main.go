package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

const imgPath = "image_c.jpg"

func main() {
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)

	height := img.Rows()
	width := img.Cols()
	channels := img.Channels()

	fmt.Printf("%v x %v (%v)\n", width, height, channels)

	arr := GetImgArray(img)
	fmt.Printf("%v x %v x %v", len(arr), len(arr[0]), len(arr[0][0]))

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			fmt.Printf("%v\n", arr[r][c])

		}
	}

	fmt.Printf("%#v\n", img.Mean())

	// Show(img)

}
