package main

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func (s Sphere) Hit(ray Ray) (bool, float64, Vec3) {
	oc := s.Center.Sub(ray.Origin)
	a := ray.Direction.Dot(ray.Direction)
	b := -2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		// No intersection
		return false, 0, Vec3{}
	} else {
		// Closest intersection point
		t := (-b - math.Sqrt(discriminant)) / (2.0 * a)

		hitPoint := ray.At(t)
		// Compute the normal vector
		normal := hitPoint.Sub(s.Center).Normalize()
		return true, t, normal
	}
}
