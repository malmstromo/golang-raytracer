package main

import "math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func (s Sphere) Hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
	oc := s.Center.Sub(ray.Origin)
	a := ray.Direction.LengthSquared()

	// represents the distance along the ray's direction to the point closest to the sphere's center.
	h := ray.Direction.Dot(oc)

	// Measures the offset between the ray's origin and the sphere’s surface.
	c := oc.Dot(oc) - s.Radius*s.Radius

	var rec HitRecord

	discriminant := h*h - a*c

	// no hit
	if discriminant < 0 {
		return false, rec
	}

	// Find the nearest root that lies in the acceptable range.
	t := (h - math.Sqrt(discriminant)) / a

	if t <= tMin || tMax <= t {
		t = (h + math.Sqrt(discriminant)) / a
		if t <= tMin || tMax <= t {
			return false, rec
		}
	}

	hitPoint := ray.At(t)
	normal := hitPoint.Sub(s.Center).Scale(1.0 / s.Radius)

	return true, HitRecord{
		T:      t,
		P:      hitPoint,
		Normal: normal,
	}
}
