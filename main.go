package main

import (
	"ascii-art/internal/ascii"
	"ascii-art/internal/image"
	"flag"
	"fmt"
)

const defaultWidth = 100

func main() {
	width := flag.Int("width", defaultWidth, "output width")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: go run . [options] <image-path>")
		flag.PrintDefaults()
		return
	}

	if *width <= 0 {
		fmt.Println("Error: width must be a positive integer")
		return
	}

	imagePath := flag.Arg(0)

	img, err := image.Load(imagePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// fmt.Printf("Original: %d x %d\n", img.Bounds().Dx(), img.Bounds().Dy())

	targetWidth, targetHeight := image.CalculateDimensions(img, *width)

	resized := image.Resize(img, targetWidth, targetHeight)

	// fmt.Printf("Resized: %d x %d\n", resized.Bounds().Dx(), resized.Bounds().Dy())

	gray := image.GrayScale(resized)

	// fmt.Printf("Grayscale: %d x %d\n", gray.Bounds().Dx(), gray.Bounds().Dy())

	art := ascii.Convert(gray, ascii.DefaultCharset)

	fmt.Print(art)
}
