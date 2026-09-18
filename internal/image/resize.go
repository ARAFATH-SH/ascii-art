package image

import (
	"image"
	"image/color"
)

func Resize(img image.Image, targetWidth, targetHeight int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))

	srcBounds := img.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()

	for y := 0; y < targetHeight; y++ {
		for x := 0; x < targetWidth; x++ {
			srcX := x * srcWidth / targetWidth
			srcY := y * srcHeight / targetHeight

			c := img.At(
				srcBounds.Min.X+srcX,
				srcBounds.Min.Y+srcY,
			)
			dst.Set(x, y, color.RGBAModel.Convert(c))
		}
	}
	return dst
}
