package main

type Ray struct {
	Direction Vector
	Origin    Point
}

func (r Ray) At(t float64) Vector {
	b := r.Direction.MultiplyScalar(t)
	a := r.Origin
	return a.Add(b)
}
