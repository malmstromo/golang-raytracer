package main

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func (s Sphere) Hit(ray Ray) (bool, float64, Vec3) {
	oc := s.Center.Sub(ray.Origin)
	a := ray.Direction.LengthSquared()
	h := ray.Direction.Dot(oc)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := h*h - a*c

	if discriminant < 0 {
		// No intersection
		return false, 0, Vec3{}
	} else {
		// Closest intersection point
		t := (h - math.Sqrt(discriminant)) / a

		hitPoint := ray.At(t)
		// Compute the normal vector
		normal := hitPoint.Sub(s.Center).Normalize()
		return true, t, normal
	}
}
