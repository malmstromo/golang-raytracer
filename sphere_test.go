package main

import "testing"

func TestSphere_Hit(t *testing.T) {
	// Sphere centered at (0, 0, -5) with radius 1
	sphere := Sphere{Center: Vec3{X: 0, Y: 0, Z: -5}, Radius: 1}

	// Ray starting at (0, 0, 0) and pointing directly at the sphere
	ray := Ray{Origin: Vec3{X: 0, Y: 0, Z: 0}, Direction: Vec3{X: 0, Y: 0, Z: -1}}

	// Check for intersection
	hit, tValue := sphere.Hit(ray)
	if !hit {
		t.Errorf("Expected ray to hit the sphere")
	}
	if tValue <= 0 {
		t.Errorf("Expected positive t-value, got %v", tValue)
	}
}
