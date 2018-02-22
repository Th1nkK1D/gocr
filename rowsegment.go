package main

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

const cutoffTh = 2
const countTh = 2
const mergeTh = 2

func getRowFreq(arr [][][]uint8) []int {
	freq := make([]int, len(arr))

	for r := range arr {
		for c := range arr[r] {
			if arr[r][c][0] == 0 {
				freq[r]++
			}
		}
	}

	return freq
}

// SplitLine - return start/end mark array of each line
func SplitLine(arr [][][]uint8) ([]int, []int) {
	freq := getRowFreq(arr)

	startMark := make([]int, 0)
	endMark := make([]int, 0)
	wCount := 0
	bCount := 0
	start := 0
	end := 0
	sumSize := 0
	lineCount := 0

	for r := range arr {
		if freq[r] >= countTh {
			// Spot black
			if wCount > 0 {
				if wCount >= cutoffTh {
					start = r
				}

				wCount = 0
			}

			bCount++

		} else {
			// Spot white
			if bCount > 0 {
				if bCount >= cutoffTh {
					end = r - 1

					if lineCount > 0 && sumSize/lineCount/(end-start) > mergeTh {
						// Merge line if size is suspectect
						endMark[len(endMark)-1] = end
					} else {
						// New line
						startMark = append(startMark, start)
						endMark = append(endMark, end)
						sumSize += end - start
						lineCount++
					}

				}

				bCount = 0
			}

			wCount++
		}
	}

	return startMark, endMark
}

// DrawRowSegment - draw line segment on mat
func DrawRowSegment(img gocv.Mat, start, end []int) {
	width := img.Cols()

	for i := range start {
		gocv.Line(img, image.Point{0, start[i]}, image.Point{width, start[i]}, color.RGBA{0, 0, 0, 255}, 1)
		gocv.Line(img, image.Point{0, end[i]}, image.Point{width, end[i]}, color.RGBA{0, 0, 0, 255}, 2)
	}
}
