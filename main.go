package main

import (
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

	resized := image.Resize(img, 100, 75)

	fmt.Printf("Resized: %d x %d\n", resized.Bounds().Dx(), resized.Bounds().Dy())
}
