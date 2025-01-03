package main

import (
	"fmt"
	"math"
)

// Vec3 represents a 3D vector or point in space.
type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vec3) Scale(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v Vec3) Dot(u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec3) Normalize() Vec3 {
	mag := v.Length()
	if mag == 0 {
		return Vec3{0, 0, 0} // Avoid division by zero
	}
	return Vec3{v.X / mag, v.Y / mag, v.Z / mag}
}

func (v Vec3) ToPPM() string {
	ir := int(255.999 * math.Min(1.0, math.Max(0.0, v.X)))
	ig := int(255.999 * math.Min(1.0, math.Max(0.0, v.Y)))
	ib := int(255.999 * math.Min(1.0, math.Max(0.0, v.Z)))
	return fmt.Sprintf("%d %d %d", ir, ig, ib)
}
