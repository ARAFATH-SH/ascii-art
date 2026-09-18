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

	bounds := img.Bounds()

	fmt.Println("Image Loaded successfully!\n")
	fmt.Printf("Width: %d\n", bounds.Dx())
	fmt.Printf("Height: %d\n", bounds.Dy())
}
