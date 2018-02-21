package main

import (
	"strconv"
	"strings"

	"gocv.io/x/gocv"
)

const binSize = 0.8
const binNum = 5

// Template structure
type Template struct {
	char   string
	arr    [][][]uint8
	height int
	width  int
}

// ReadTemplate - read template from file
func ReadTemplate(templateChar, templateDir string) [][]Template {
	// Init templateIndex
	templateArr := make([][]Template, binNum)

	// Fetch each character
	for i, str := range strings.Split(templateChar, " ") {
		imgArr := GetImgArray(gocv.IMRead(templateDir+strconv.Itoa(i+1)+".png", gocv.IMReadGrayScale))
		height, width := len(imgArr), len(imgArr[0])

		temp := Template{str, imgArr, height, width}
		binNum := int((float32(height) / float32(width)) / binSize)

		templateArr[binNum] = append(templateArr[binNum], temp)
	}

	return templateArr
}
