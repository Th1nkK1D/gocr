// gOCR - Resizer Module

package main

// Greatest common divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Least common factor
func lcf(a, b int) int {
	if a < b {
		a, b = b, a
	}

	return a * b / gcd(a, b)
}

// Resize - resize the image
func Resize(imgArr [][][]uint8, height, width int) [][][]uint8 {
	lcfHeight := lcf(height, len(imgArr))
	lcfWidth := lcf(width, len(imgArr[0]))

	newHeightFrac := lcfHeight / height
	oldHeightFrac := lcfHeight / len(imgArr)
	newWidthFrac := lcfWidth / width
	oldWidthFrac := lcfWidth / len(imgArr[0])

	// Init newImage
	newImg := make([][][]uint8, height)

	for r := range newImg {
		newImg[r] = make([][]uint8, width)

		for c := range newImg[r] {
			newImg[r][c] = make([]uint8, 1)

			sum := 0

			// Sampling
			for j := r * newHeightFrac; j < newHeightFrac*(r+1); j++ {
				yIndex := j / oldHeightFrac

				for i := c * newWidthFrac; i < newWidthFrac*(c+1); i++ {
					xIndex := i / oldWidthFrac

					sum += int(imgArr[yIndex][xIndex][0])
				}
			}

			newImg[r][c][0] = uint8((float64(sum))/float64(newHeightFrac*newWidthFrac) + 0.5)
		}
	}

	return newImg
}
