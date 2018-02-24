// gOCR - Median Filter Module

package main

import (
	"sort"
)

func getAbs(num int) int {
	if num < 0 {
		num = -num
	}

	return num
}

func getMedian(list []uint8) uint8 {
	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })

	return list[(len(list)-1)/2]
}

func getMirror(p, bound int) int {
	if p < 0 {
		p = getAbs(p + 1)
	} else if p >= bound {
		p = bound*2 - p - 1
	}

	return p
}

// MedianFilter - 3x3 median filter
func MedianFilter(imgArr [][][]uint8, rad int) [][][]uint8 {
	height := len(imgArr)
	width := len(imgArr[0])

	// Init new array
	newArr := make([][][]uint8, height)

	for r := 0; r < height; r++ {
		newArr[r] = make([][]uint8, width)

		for c := 0; c < width; c++ {
			newArr[r][c] = make([]uint8, 1)

			list := make([]uint8, 0)

			// Matrix loop
			for j := -rad; j <= rad; j++ {
				for i := -rad; i <= rad; i++ {
					list = append(list, imgArr[getMirror(r+j, height)][getMirror(c+i, width)][0])
				}
			}

			newArr[r][c][0] = getMedian(list)
		}
	}

	return newArr
}
