package main

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

// At method computes the position along the ray at distance t
func (r Ray) At(t float64) Vec3 {
	return r.Origin.Add(r.Direction.Scale(t))
}
