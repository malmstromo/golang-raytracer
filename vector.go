package main

type Vector struct {
	X, Y, Z float64
}

func (v Vector) MultiplyScalar(t float64) Vector {
	return Vector{v.X * t, v.Y * t, v.Z * t}
}

func (p Vector) Add(v Vector) Vector {
	return Vector{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func (p Point) Add(v Vector) Vector {
	return Vector{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}
