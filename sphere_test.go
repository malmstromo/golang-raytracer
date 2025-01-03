package main

import "testing"

func TestSphere_Hit(t *testing.T) {
	// Sphere centered at (0, 0, -5) with radius 1
	sphere := Sphere{Center: Vec3{X: 0, Y: 0, Z: -5}, Radius: 1}

	// Ray starting at (0, 0, 0) and pointing directly at the sphere
	ray := Ray{Origin: Vec3{X: 0, Y: 0, Z: 0}, Direction: Vec3{X: 0, Y: 0, Z: -1}}

	// Check for intersection
	hit, rec := sphere.Hit(ray, 0, 10)
	if !hit {
		t.Errorf("Expected ray to hit the sphere")
	}
	if rec.T <= 0 {
		t.Errorf("Expected positive t-value, got %v", rec.T)
	}
}

func TestSphere_Miss_Far(t *testing.T) {
	sphereZ := -10
	tMax := 9
	sphere := Sphere{Center: Vec3{X: 0, Y: 0, Z: float64(sphereZ)}, Radius: 1}

	// Ray starting at (0, 0, 0) and pointing directly at the sphere
	ray := Ray{Origin: Vec3{X: 0, Y: 0, Z: 0}, Direction: Vec3{X: 0, Y: 0, Z: -1}}

	hit, rec := sphere.Hit(ray, 0, float64(tMax))
	if hit {
		t.Errorf("Did not expect hit. Sphere center is at %v and intersection should occur T < %v", sphereZ, tMax)
	}
	if rec.T > 0 {
		t.Errorf("Expected 0 t-value, got %v", rec.T)
	}
}

func TestSphere_Miss_Close(t *testing.T) {
	sphereZ := -1
	tMin := 3
	sphere := Sphere{Center: Vec3{X: 0, Y: 0, Z: float64(sphereZ)}, Radius: 1}

	// Ray starting at (0, 0, 0) and pointing directly at the sphere
	ray := Ray{Origin: Vec3{X: 0, Y: 0, Z: 0}, Direction: Vec3{X: 0, Y: 0, Z: -1}}

	hit, rec := sphere.Hit(ray, float64(tMin), 10)
	if hit {
		t.Errorf("Did not expect hit. Sphere center is at %v and intersection should occur T > %v", sphereZ, tMin)
	}
	if rec.T > 0 {
		t.Errorf("Expected 0 t-value, got %v", rec.T)
	}
}
