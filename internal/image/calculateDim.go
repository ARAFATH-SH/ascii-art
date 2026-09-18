package image

import "image"

func CalculateDimensions(img image.Image, targetWidth int) (int, int) {
	bounds := img.Bounds()

	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	aspectRatio := float64(originalHeight) / float64(originalWidth)

	const characterAspectRatio = 0.5

	targetHeight := int(float64(targetWidth) * aspectRatio * characterAspectRatio)

	if targetHeight < 1 {
		targetHeight = 1
	}
	return targetWidth, targetHeight

}
