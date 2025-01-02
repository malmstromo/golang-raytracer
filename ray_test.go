package main

import "testing"

func TestRay_At(t *testing.T) {
	origin := Vec3{X: 1, Y: 2, Z: 3}
	direction := Vec3{X: 4, Y: 5, Z: 6}
	ray := Ray{Origin: origin, Direction: direction}

	t1 := 0.0
	expected1 := Vec3{X: 1, Y: 2, Z: 3}
	result1 := ray.At(t1)
	if result1 != expected1 {
		t.Errorf("At(%v): expected %v, got %v", t1, expected1, result1)
	}

	t2 := 2.0
	expected2 := Vec3{X: 9, Y: 12, Z: 15}
	result2 := ray.At(t2)
	if result2 != expected2 {
		t.Errorf("At(%v): expected %v, got %v", t2, expected2, result2)
	}
}
