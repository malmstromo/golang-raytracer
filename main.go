package main

import (
	"fmt"
	"os"
)

func main() {
	// Image dimensions
	imageWidth := 400
	imageHeight := 200

	// Sphere
	sphere := Sphere{Center: Vec3{X: 0, Y: 0, Z: -1}, Radius: 0.5}

	// Camera
	origin := Vec3{0, 0, 0}

	// Viewport
	lowerLeftCorner := Vec3{-2.0, -1.0, -1.0}
	horizontal := Vec3{4.0, 0.0, 0.0}
	vertical := Vec3{0.0, 2.0, 0.0}

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
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)

			direction := lowerLeftCorner.Add(horizontal.Scale(u)).Add(vertical.Scale(v)).Sub(origin)
			ray := Ray{Origin: origin, Direction: direction}

			fmt.Fprintln(file, rayColor(ray, sphere).ToPPM())
		}
	}

	fmt.Println("Rendered output to image.ppm")
}

/*
If the Ray hits the Sphere, calculate its color from the surface normal.
If no hit, create a blue-white gradient.
*/
func rayColor(r Ray, sphere Sphere) Vec3 {
	if hit, _, normal := sphere.Hit(r); hit {
		return Vec3{
			0.5 * (normal.X + 1.0),
			0.5 * (normal.Y + 1.0),
			0.5 * (normal.Z + 1.0),
		}
	}
	return lerp(r)
}

/*
From the book:

This function will linearly blend white and blue depending on the height of the 𝑦 coordinate after scaling the ray direction to unit length (so −1.0<𝑦<1.0).
Because we're looking at the 𝑦 height after normalizing the vector, you'll notice a horizontal gradient to the color in addition to the vertical gradient.

I'll use a standard graphics trick to linearly scale 0.0≤𝑎≤1.0.
When 𝑎=1.0, I want blue.
When 𝑎=0.0, I want white.
In between, I want a blend. This forms a “linear blend”, or “linear interpolation”. This is commonly referred to as a lerp between two values.
A lerp is always of the form:

blendedValue=(1−𝑎)⋅startValue+𝑎⋅endValue, with 𝑎 going from zero to one.
*/
func lerp(r Ray) Vec3 {
	unitDirection := r.Direction.Normalize()
	a := 0.5 * (unitDirection.Y + 1.0)
	white := Vec3{1.0, 1.0, 1.0}
	blue := Vec3{0.5, 0.7, 1.0}
	return white.Scale(1.0 - a).Add(blue.Scale(a))
}
