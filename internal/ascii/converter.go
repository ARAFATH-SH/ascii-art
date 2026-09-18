package ascii

import (
	"image"
	"strings"
)

func Convert(img *image.Gray, charset string) string {
	bounds := img.Bounds()

	var result strings.Builder

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := img.GrayAt(x, y)

			index := int(gray.Y) * (len(charset) - 1) / 255

			result.WriteByte(charset[index])
		}
		result.WriteByte('\n')
	}
	return result.String()
}

func ConvertWithPatterns(img *image.Gray) string {
	bounds := img.Bounds()

	var result strings.Builder

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 3 {
		for x := bounds.Min.X; x < bounds.Max.X; x += 3 {
			pattern := ExtractPattern(img, x, y)

			character := FindBestCharacter(pattern)

			result.WriteByte(character)
		}

		result.WriteByte('\n')
	}

	return result.String()
}
