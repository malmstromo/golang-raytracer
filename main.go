package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Prompt the user for the blue value
	var input string
	fmt.Println("Enter a blue value (between 0 and 1):")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Failed to read input. Using default value of 0.25.")
		input = "0.25" // Default blue value
	}

	// Parse the input
	blue, err := strconv.ParseFloat(input, 64)
	if err != nil || blue < 0 || blue > 1 {
		fmt.Println("Invalid input. Please provide a number between 0 and 1.")
		return
	}

	// Image dimensions
	imageWidth := 256
	imageHeight := 256

	// Open a file to save the image
	file, err := os.Create("image.ppm")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the PPM header
	fmt.Fprintf(file, "P3\n%d %d\n255\n", imageWidth, imageHeight)

	// Generate the image
	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)
			b := blue

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
		}
	}

	fmt.Println("Image generated: image.ppm")
}
