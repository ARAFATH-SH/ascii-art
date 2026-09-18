package main

import (
	"ascii-art/internal/ascii"
	"ascii-art/internal/image"
	"ascii-art/internal/output"
	"flag"
	"fmt"
	"os"
)

const defaultWidth = 100

func main() {
	width := flag.Int("width", defaultWidth, "output width")
	outputPath := flag.String("output", "", "output file path")
	input := flag.String("input", "", "ASCII text input file")

	flag.Parse()

	if *input != "" {
		if *outputPath == "" {
			fmt.Println("Error: --output is required when using --input")
			return
		}
		data, err := os.ReadFile(*input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if err := output.SaveASCIIAsImage(string(data), *outputPath); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("ASCII image saved to %s\n", *outputPath)
		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Usage: go run . [options] <image-path>")
		fmt.Println("  go run . --input <ascii-file> --output <image.png|image.jpg>\n")
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

	targetWidth, targetHeight := image.CalculateDimensions(img, *width)

	resized := image.Resize(img, targetWidth, targetHeight)

	gray := image.GrayScale(resized)

	art := ascii.ConvertWithPatterns(gray)

	if *outputPath == "" {
		output.PrintASCII(art)
		return
	}

	if err := os.WriteFile(*outputPath, []byte(art), 0644); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("ASCII art saved to %s\n", *outputPath)
}
