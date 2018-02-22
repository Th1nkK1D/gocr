package main

// Resize - resize the image
func Resize(imgArr [][][]uint8, height, width int) [][][]uint8 {
	oldHeight := len(imgArr)
	oldWidth := len(imgArr[0])

	// Init newImage
	newImg := make([][][]uint8, height)

	for r := range newImg {
		newImg[r] = make([][]uint8, width)

		for c := range newImg[r] {
			newImg[r][c] = make([]uint8, 1)

			sum := 0

			for j := r * oldHeight; j < oldHeight*(r+1); j++ {
				yIndex := j / height

				for i := c * oldWidth; i < oldWidth*(c+1); i++ {
					xIndex := i / width

					sum += int(imgArr[yIndex][xIndex][0])
				}
			}
			newImg[r][c][0] = uint8((float64(sum))/float64(oldHeight*oldWidth) + 0.5)
		}
	}

	return newImg
}
