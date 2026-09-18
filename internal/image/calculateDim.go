package image

import "image"

func CalculateDimensions(img image.Image, targetWidth int) (int, int) {
	bounds := img.Bounds()

	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	aspectRatio := float64(originalHeight) / float64(originalWidth)

	targetHeight := int(float64(targetWidth) * aspectRatio)

	return targetWidth, targetHeight

}
