// gOCR - Row Segmentation Module

package main

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

const cutoffTh = 2
const countTh = 2
const mergeTh = 4

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
	sumText := 0
	sumSpace := 0
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

					if len(endMark) > 0 {
						sumSpace += start - endMark[len(endMark)-1]
					}

					startMark = append(startMark, start)
					endMark = append(endMark, end)

					lineCount++
					sumText += end - start

				}

				bCount = 0
			}

			wCount++
		}
	}

	if lineCount > 0 {
		// Need to merge check
		avgSpace := sumSpace / (lineCount - 1)
		avgText := sumText / lineCount

		if avgText/avgSpace > mergeTh {
			// Merge to single line
			startMark = startMark[:1]
			endMark = endMark[len(endMark)-1:]
		} else {
			// Merge to multiple line
			for i := len(startMark) - 1; i > 0; i-- {
				if startMark[i]-endMark[i-1] < avgSpace {
					// Merge line
					endMark[i-1] = endMark[i]

					startMark = startMark[:i+copy(startMark[i:], startMark[i+1:])]
					endMark = endMark[:i+copy(endMark[i:], endMark[i+1:])]
				}
			}
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
