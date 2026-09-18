package image

import (
	"image"
	"image/color"
)

func GrayScale(img image.Image) image.Image {
	bounds := img.Bounds()

	gray := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y))

			gray.Set(x, y, c.(color.Gray))
		}
	}
	return gray
}
