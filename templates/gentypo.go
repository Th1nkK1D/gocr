package main

import (
	"bufio"
	"fmt"
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
)

const templateChar = "ส ก น ด"
const fontFile = "angsanaNew.ttf"
const fontSize = 20

func getGlypBound(img image.Image) image.Rectangle {
	white := color.RGBA{255, 255, 255, 255}
	imgBound := img.Bounds()
	// typoBound := image.Rectangle{image.Point{0, 0}, image.Point{imgBound.Max.X, imgBound.Max.Y}}
	yRange := make([]int, 0)
	xRange := make([]int, 0)

	// fmt.Println(typoBound)

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

	// fmt.Println(yRange)
	// fmt.Println(xRange)

	return image.Rectangle{image.Point{xRange[0], yRange[0]}, image.Point{xRange[len(xRange)-1], yRange[len(yRange)-1]}}
}

func main() {
	fontBytes, err := ioutil.ReadFile(fontFile)

	if err != nil {
		panic(err)
	}

	font, err := freetype.ParseFont(fontBytes)

	if err != nil {
		panic(err)
	}

	count := 0

	for _, str := range strings.Split(templateChar, " ") {
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
		pt := freetype.Pt(fontSize/2, fontSize)

		_, err := ctx.DrawString(str, pt)

		if err != nil {
			fmt.Println(err)
			return
		}

		getGlypBound(background)

		count++

		// Save
		outFile, err := os.Create(strconv.Itoa(count) + ".png")
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		buff := bufio.NewWriter(outFile)

		err = png.Encode(buff, background)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		// flush everything out to file
		err = buff.Flush()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		outFile.Close()
	}
}
