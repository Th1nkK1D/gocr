package main

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"gocv.io/x/gocv"
)

const binSize = 1
const binAmount = 5

// Template structure
type Template struct {
	char   string
	arr    [][][]uint8
	height int
	width  int
}

// MatchRes - Result from matching
type MatchRes struct {
	char string
	sse  int
	i    int
}

// GetRatioBin - Calculate ratio bin
func GetRatioBin(height, width int) int {
	return int((float32(height) / float32(width)) / binSize)
}

// ReadTemplate - read template from file
func ReadTemplate(templateChar, templateDir string) [][]Template {
	// Init templateIndex
	templateArr := make([][]Template, binAmount)

	// Fetch each character
	for i, str := range strings.Split(templateChar, " ") {
		imgArr := GetImgArray(gocv.IMRead(templateDir+strconv.Itoa(i+1)+".png", gocv.IMReadGrayScale))
		height, width := len(imgArr), len(imgArr[0])

		temp := Template{str, imgArr, height, width}
		bin := GetRatioBin(height, width)

		// fmt.Printf("%v -> %v x %v (%v)\n", str, height, width, bin)

		templateArr[bin] = append(templateArr[bin], temp)
	}

	return templateArr
}

// MatchTemplate - template matching
func MatchTemplate(blob [][][]uint8, templates []Template) []MatchRes {
	resArr := make([]MatchRes, 0)

	for t := range templates {
		// fmt.Printf("%v -> %v x %v\n", templates[t].char, templates[t].height, templates[t].width)

		// comp := Resize(blob, templates[t].height, templates[t].width)
		comp := Resize(templates[t].arr, len(blob), len(blob[0]))
		sse := 0

		for r := range comp {
			for c := range comp[r] {
				// sse += int(math.Pow(float64(comp[r][c][0]-templates[t].arr[r][c][0]), 2))
				sse += int(math.Pow(float64(comp[r][c][0]-blob[r][c][0]), 2))
			}
		}

		resArr = append(resArr, MatchRes{templates[t].char, sse, t})
	}

	sort.Slice(resArr, func(i, j int) bool { return resArr[i].sse < resArr[j].sse })

	// gocv.IMWrite("resized.png", GetImgMat(Resize(blob, templates[resArr[0].i].height, templates[resArr[0].i].width)))
	// gocv.IMWrite("matchtemp.png", GetImgMat(templates[resArr[0].i].arr))

	return resArr
}
