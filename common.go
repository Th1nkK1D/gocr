// gOCR - Common Module

package main

import (
	"image"

	"gocv.io/x/gocv"
)

// Show - show image dialog
func Show(img gocv.Mat) {
	window := gocv.NewWindow("Display")

	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}

// GetImgArray - Convert cvMat to 3-dimension array
func GetImgArray(img gocv.Mat) [][][]uint8 {
	height := img.Rows()
	width := img.Cols()
	channels := img.Channels()
	i := 0

	flatArr := img.ToBytes()

	arr := make([][][]uint8, height)

	for row := 0; row < height; row++ {
		arr[row] = make([][]uint8, width)

		for col := 0; col < width; col++ {
			arr[row][col] = make([]uint8, channels)

			for ch := 0; ch < channels; ch++ {
				arr[row][col][ch] = flatArr[i]
				i++
			}
		}
	}

	return arr
}

// CropImgArr - Crop image array
func CropImgArr(imgArr [][][]uint8, bound image.Rectangle) [][][]uint8 {
	newImg := make([][][]uint8, bound.Max.Y-bound.Min.Y)

	for r := range newImg {
		newImg[r] = make([][]uint8, bound.Max.X-bound.Min.X)

		for c := range newImg[r] {
			newImg[r][c] = make([]uint8, 1)
			newImg[r][c][0] = imgArr[bound.Min.Y+r][bound.Min.X+c][0]
		}
	}

	return newImg
}

// GetImgMat - Convert cvMat to 3-dimension array
func GetImgMat(arr [][][]uint8) gocv.Mat {
	flag := [...]gocv.MatType{gocv.MatTypeCV8UC1, gocv.MatTypeCV8UC2, gocv.MatTypeCV8UC3, gocv.MatTypeCV8UC4}
	height := len(arr)
	width := len(arr[0])
	channels := len(arr[0][0])
	i := 0

	flatArr := make([]byte, height*width*channels)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			for ch := 0; ch < channels; ch++ {
				flatArr[i] = arr[row][col][ch]
				i++
			}
		}
	}

	return gocv.NewMatFromBytes(height, width, flag[channels-1], flatArr)

}

// GetHistArray - Get histrogram array
func GetHistArray(img gocv.Mat) []int {
	hist := make([]int, 256)

	flatArr := img.ToBytes()

	for i := range flatArr {
		hist[flatArr[i]]++
	}

	return hist
}
