package main

import (
	"sort"
)

const cutoffTh = 2
const countTh = 2

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

	for r := range arr {
		if freq[r] >= countTh {
			// Spot black
			if wCount > 0 {
				// fmt.Printf("b -> %v\n", )
				if wCount >= cutoffTh {
					startMark = append(startMark, r-cutoffTh)
				}

				wCount = 0
			}

			bCount++

		} else {
			// Spot white
			if bCount > 0 {
				if bCount >= cutoffTh {
					endMark = append(endMark, r)
				}

				bCount = 0
			}

			wCount++
		}
	}

	// fmt.Printf("%v\n", startMark)
	// fmt.Printf("%v\n", endMark)

	rangeArr := make([]int, len(startMark))

	for i := range startMark {
		rangeArr[i] = endMark[i] - startMark[i]
	}

	// fmt.Printf("value: %v\n", rangeArr)

	sort.Ints(rangeArr)

	// fmt.Printf("value: %v\n", rangeArr)

	return startMark, endMark
}
