package main

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func (s Sphere) Hit(ray Ray) (bool, float64) {
	oc := ray.Origin.Sub(s.Center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		// No intersection
		return false, 0
	} else {
		// Closest intersection point
		t := (-b - math.Sqrt(discriminant)) / (2.0 * a)
		return true, t
	}
}
