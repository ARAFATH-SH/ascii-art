package main

import (
	"ascii-art/internal/ascii"
	"ascii-art/internal/image"
	"fmt"
	"os"
	"strconv"
)

const defaultWidth = 100

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run. <image-path>")
		return
	}

	width := defaultWidth

	if len(os.Args) >= 3 {
		parseWidth, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: width must be a positive integer")
			return
		}
		width = parseWidth
	}

	img, err := image.Load(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// fmt.Printf("Original: %d x %d\n", img.Bounds().Dx(), img.Bounds().Dy())

	targetWidth, targetHeight := image.CalculateDimensions(img, width)

	resized := image.Resize(img, targetWidth, targetHeight)

	// fmt.Printf("Resized: %d x %d\n", resized.Bounds().Dx(), resized.Bounds().Dy())

	gray := image.GrayScale(resized)

	// fmt.Printf("Grayscale: %d x %d\n", gray.Bounds().Dx(), gray.Bounds().Dy())

	art := ascii.Convert(gray, ascii.DefaultCharset)

	fmt.Print(art)
}
