package main

import (
	"fmt"
	"os"
)

func checkError(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func main() {

	image_width := 256
	image_height := 256
	color := float64(255.99)

	f, err := os.Create("out.ppm")
	checkError(err, "error opening file")

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", image_width, image_height)
	checkError(err, "error writing to file")

	defer f.Close()

	for j := 0; j < image_height; j++ {
		fmt.Printf("\rScanlines remaining: %d ", image_height-j)
		for i := 0; i < image_width; i++ {

			r := float64(i) / float64(image_width-1)
			g := float64(j) / float64(image_height-1)
			b := float64(0)

			ir := int(color * r)
			ig := int(color * g)
			ib := int(color * b)

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)

			checkError(err, "Error writting to file: %v\n")
		}
	}
}
