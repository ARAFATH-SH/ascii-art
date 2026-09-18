package ascii

import "image"

func ExtractPattern(img *image.Gray, startX, startY int) [3][3]float64 {
	var pattern [3][3]float64

	bounds := img.Bounds()

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			px := startX + x
			py := startY + y

			if px >= bounds.Max.X || py >= bounds.Max.Y {
				continue
			}

			gray := img.GrayAt(px, py)

			pattern[y][x] = float64(gray.Y) / 255.0
		}
	}

	return pattern
}
