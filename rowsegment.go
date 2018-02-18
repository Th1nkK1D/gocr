package main

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
				wCount = 0
				startMark = append(startMark, r)
			}

			bCount++

		} else {
			// Spot white
			if bCount > 0 {
				bCount = 0
				endMark = append(endMark, r)
			}

			wCount++
		}
	}

	return startMark, endMark
}
