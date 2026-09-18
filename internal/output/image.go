package output

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func SaveASCIIAsImage(asciiArt, outputPath string) error {
	asciiArt = strings.TrimRight(asciiArt, "\n")

	if asciiArt == "" {
		return fmt.Errorf("ASCII art is empty")
	}

	lines := strings.Split(asciiArt, "\n")

	face := basicfont.Face7x13

	metrics := face.Metrics()

	advance, ok := face.GlyphAdvance('M')
	if !ok {
		return fmt.Errorf("failed to determine character width")
	}

	charWidth := advance.Round()
	lineHeight := metrics.Height.Round()

	const padding = 10

	// Find the longest line.
	maxWidth := 0

	for _, line := range lines {
		width := len([]rune(line))

		if width > maxWidth {
			maxWidth = width
		}
	}

	imageWidth := maxWidth*charWidth + padding*2
	imageHeight := len(lines)*lineHeight + padding*2

	img := image.NewRGBA(
		image.Rect(0, 0, imageWidth, imageHeight),
	)

	// Fill background with black.
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			img.Set(x, y, color.Black)
		}
	}

	// Configure font drawer.
	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: face,
	}

	// Draw every line.
	for row, line := range lines {
		x := padding
		y := padding + row*lineHeight + metrics.Ascent.Round()

		drawer.Dot = fixed.P(x, y)
		drawer.DrawString(line)
	}

	// Determine output format from extension.
	extension := strings.ToLower(filepath.Ext(outputPath))

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	defer file.Close()

	switch extension {
	case ".png":
		err := png.Encode(file, img)
		if err != nil {
			return fmt.Errorf("failed to encode PNG: %w", err)
		}

	case ".jpg", ".jpeg":
		err := jpeg.Encode(file, img, &jpeg.Options{
			Quality: 95,
		})
		if err != nil {
			return fmt.Errorf("failed to encode JPEG: %w", err)
		}

	default:
		return fmt.Errorf(
			"unsupported image format %q; use .png, .jpg, or .jpeg",
			extension,
		)
	}

	return nil
}
