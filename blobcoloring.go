package main

import (
	"fmt"
	"sort"
)

type rect struct {
	yMin, xMin, yMax, xMax int
}

func updateRect(old, new rect) rect {
	if new.xMax > old.xMax {
		old.xMax = new.xMax
	}

	if new.xMin < old.xMin {
		old.xMin = new.xMin
	}

	if new.yMax > old.yMax {
		old.yMax = new.yMax
	}

	if new.yMin < old.yMin {
		old.yMin = new.yMin
	}

	return old
}

func getMapRoot(maptable map[int]int, val int) int {
	for val != maptable[val] {
		val = maptable[val]
	}

	return val
}

// SegmentChar - Segment characters
func SegmentChar(imgArr [][][]uint8) {
	grass := make([][]int, len(imgArr))
	maptable := make(map[int]int)
	recttable := make(map[int]rect)
	num := 0

	// Init grass
	for r := range imgArr {
		grass[r] = make([]int, len(imgArr[r]))

		for c := range imgArr[r] {
			if imgArr[r][c][0] != 255 {
				grass[r][c] = 0
			} else {
				grass[r][c] = -1
			}
		}
	}

	// fmt.Printf("%v\n", grass)

	// Start a fire
	for y := range grass {
		for x := range grass[y] {
			if grass[y][x] >= 0 {
				found := make([]int, 0)

				// Contour search
				searchArea := [][]int{{y, x - 1}, {y - 1, x - 1}, {y - 1, x}, {y - 1, x + 1}}

				for s := range searchArea {
					j := searchArea[s][0]
					i := searchArea[s][1]

					if j >= 0 && j < len(grass) && i >= 0 && i < len(grass[y]) && grass[j][i] > 0 {
						found = append(found, grass[j][i])
					}
				}

				if len(found) == 0 {
					// New object
					num++
					grass[y][x] = num
					maptable[num] = num
					recttable[num] = rect{y, x, y, x}
					// fmt.Printf("Add %v: %v\n", num, recttable[num])

				} else {
					// Same object
					sort.Ints(found)

					rootNode := getMapRoot(maptable, found[0])
					grass[y][x] = rootNode
					recttable[rootNode] = updateRect(recttable[rootNode], rect{y, x, y, x})

					// Update maptable and recttable
					for k := 1; k < len(found); k++ {
						if newRect, ok := recttable[found[k]]; ok && found[k] != rootNode {
							maptable[found[k]] = rootNode
							// fmt.Printf("Update %v: %v <-> %v: %v ", rootNode, recttable[rootNode], found[k], newRect)
							recttable[rootNode] = updateRect(recttable[rootNode], newRect)
							// fmt.Printf("--> %v \n", recttable[rootNode])
							delete(recttable, found[k])
						}
					}
				}
			}
		}
	}

	// fmt.Printf("%v\n", grass)
	fmt.Printf("%v\n", maptable)
	fmt.Printf("%v\n", recttable)
}
