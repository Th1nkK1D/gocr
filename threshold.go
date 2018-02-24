// gOCR - Automatic Thresholding Module

package main

import (
	"gocv.io/x/gocv"
)

func getAvgIn(hist []int, lb int) int {
	sum := 0
	count := 0

	for i := range hist {
		count += hist[i]
		sum += (lb + i) * hist[i]
	}

	return sum / count
}

// Automatic threshold algorithm
func getThresholdValue(img gocv.Mat) float32 {
	hist := GetHistArray(img)

	t := getAvgIn(hist, 0)
	ou1, ou2 := 0, 0
	u1, u2 := 255, 255

	for {
		u1 = getAvgIn(hist[:t], 0)
		u2 = getAvgIn(hist[t:], t-1)

		t = (u1 + u2) / 2

		if u1 == ou1 && u2 == ou2 {
			break
		}

		ou1, ou2 = u1, u2
	}

	return float32(t)
}

// AutoThreshold - Return Apply Auto Threshold
func AutoThreshold(img gocv.Mat) gocv.Mat {
	imgOut := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	gocv.Threshold(img, imgOut, getThresholdValue(img), 255, gocv.ThresholdBinary)

	return imgOut
}
