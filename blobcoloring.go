package main

import (
	"image"
	"sort"
)

func updateRect(old, new image.Rectangle) image.Rectangle {
	if new.Max.X > old.Max.X {
		old.Max.X = new.Max.X
	}

	if new.Min.X < old.Min.X {
		old.Min.X = new.Min.X
	}

	if new.Max.Y > old.Max.Y {
		old.Max.Y = new.Max.Y
	}

	if new.Min.Y < old.Min.Y {
		old.Min.Y = new.Min.Y
	}

	return old
}

func getMapRoot(maptable map[int]int, val int) int {
	for val != maptable[val] {
		val = maptable[val]
	}

	return val
}

// GetSegmentChar - Segment characters
func GetSegmentChar(imgArr [][][]uint8) []image.Rectangle {
	grass := make([][]int, len(imgArr))
	maptable := make(map[int]int)
	recttable := make(map[int]image.Rectangle)
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
					recttable[num] = image.Rectangle{image.Point{x, y}, image.Point{x, y}}
				} else {
					// Same object
					sort.Ints(found)

					rootNode := getMapRoot(maptable, found[0])
					grass[y][x] = rootNode
					recttable[rootNode] = updateRect(recttable[rootNode], image.Rectangle{image.Point{x, y}, image.Point{x, y}})

					// Update maptable and recttable
					for k := 1; k < len(found); k++ {
						if newRect, ok := recttable[found[k]]; ok && found[k] != rootNode {
							maptable[found[k]] = rootNode
							recttable[rootNode] = updateRect(recttable[rootNode], newRect)
							delete(recttable, found[k])
						}
					}
				}
			}
		}
	}

	// Map to array
	rectArray := make([]image.Rectangle, 0)

	for _, r := range recttable {
		rectArray = append(rectArray, r)
	}

	sort.Slice(rectArray, func(i, j int) bool { return rectArray[i].Min.X < rectArray[j].Min.X })

	return rectArray
}
