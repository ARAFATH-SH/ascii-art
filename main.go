package main

import (
	"ascii-art/internal/ascii"
	"ascii-art/internal/image"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run. <image-path>")
		return
	}

	img, err := image.Load(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Original: %d x %d\n", img.Bounds().Dx(), img.Bounds().Dy())

	width, height := image.CalculateDimensions(img, 100)

	resized := image.Resize(img, width, height)

	fmt.Printf("Resized: %d x %d\n", resized.Bounds().Dx(), resized.Bounds().Dy())

	gray := image.GrayScale(resized)

	fmt.Printf("Grayscale: %d x %d\n", gray.Bounds().Dx(), gray.Bounds().Dy())

	art := ascii.Convert(gray, ascii.DefaultCharset)

	fmt.Print(art)
}
