package main

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const fontFile = "templates/angsanaNew.ttf"
const fontSize = 80

func getGlypBound(img image.Image) image.Rectangle {
	white := color.RGBA{255, 255, 255, 255}
	imgBound := img.Bounds()
	yRange := make([]int, 0)
	xRange := make([]int, 0)

	for y := 0; y < imgBound.Max.Y; y++ {
		for x := 0; x < imgBound.Max.X; x++ {
			if img.At(x, y) != white {
				yRange = append(yRange, y)
				xRange = append(xRange, x)
			}
		}
	}

	sort.Ints(yRange)
	sort.Ints(xRange)

	return image.Rectangle{image.Point{xRange[0], yRange[0]}, image.Point{xRange[len(xRange)-1] + 1, yRange[len(yRange)-1] + 1}}
}

// Write character to the file
func writeGlyp(str string, count int, font *truetype.Font) {
	background := image.NewRGBA(image.Rect(0, 0, fontSize*3/2, fontSize*3/2))

	draw.Draw(background, background.Bounds(), image.NewUniform(color.RGBA{255, 255, 255, 255}), image.ZP, draw.Src)

	// Set context value
	ctx := freetype.NewContext()
	ctx.SetDPI(72)
	ctx.SetFont(font)
	ctx.SetFontSize(fontSize)
	ctx.SetClip(background.Bounds())
	ctx.SetDst(background)
	ctx.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))

	// Draw the text to the background
	_, err := ctx.DrawString(str, freetype.Pt(fontSize/2, fontSize))

	if err != nil {
		panic(err)
	}

	count++

	filename := strconv.Itoa(count) + ".png"
	glypBound := getGlypBound(background)

	// Save
	outFile, err := os.Create(templateDir + filename)
	if err != nil {
		panic(err)
	}

	buff := bufio.NewWriter(outFile)

	err = png.Encode(buff, background.SubImage(glypBound))
	if err != nil {
		panic(err)
	}

	// flush everything out to file
	err = buff.Flush()
	if err != nil {
		panic(err)
	}

	outFile.Close()
}

// GenTemplate - Generate Template file
func GenTemplate(templateChar, templateDir string) {
	fontBytes, err := ioutil.ReadFile(fontFile)

	if err != nil {
		panic(err)
	}

	font, err := freetype.ParseFont(fontBytes)

	if err != nil {
		panic(err)
	}

	count := 0

	// Each glyps
	for _, str := range strings.Split(templateChar, " ") {
		writeGlyp(str, count, font)
		count++
	}
}
