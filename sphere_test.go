package main

import (
	"math"
	"reflect"
	"testing"
)

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

func TestSphere_Front_Face_Hit(t *testing.T) {
	// Sphere setup
	sphere := Sphere{
		Center: Vec3{0, 0, -2},
		Radius: 1.0,
	}

	ray := Ray{
		Origin:    Vec3{0, 0, 0},
		Direction: Vec3{0, 0, -1}.Normalize(),
	}

	hit, rec := sphere.Hit(ray, 0, math.MaxFloat64)

	t.Logf("Ray direction: %v", ray.Direction)
	t.Logf("Hit point: %v", rec.P)
	t.Logf("Normal: %v", rec.Normal)
	t.Logf("FrontFace: %v", rec.FrontFace)
	if !hit {
		t.Errorf("Expected front face hit, but no hit detected.")
	}

	if !rec.FrontFace {
		t.Errorf("Expected front face, but got back face.")
	}

	expectedNormal := Vec3{0, 0, 1}
	if !reflect.DeepEqual(rec.Normal, expectedNormal) {
		t.Errorf("Expected normal %v, got %v", expectedNormal, rec.Normal)
	}

}

func TestSphere_Back_Face_Hit(t *testing.T) {
	// Sphere setup
	sphere := Sphere{
		Center: Vec3{0, 0, -2},
		Radius: 1.0,
	}

	ray := Ray{
		Origin:    Vec3{0, 0, -1.5},
		Direction: Vec3{0, 0, 1},
	}

	hit, rec := sphere.Hit(ray, 0.001, math.MaxFloat64)
	if !hit {
		t.Errorf("Expected back face hit, but no hit detected.")
	}

	if rec.FrontFace {
		t.Errorf("Expected back face, but got front face.")
	}

	expectedNormal := Vec3{0, 0, -1}
	if !reflect.DeepEqual(rec.Normal, expectedNormal) {
		t.Errorf("Expected normal %v, got %v", expectedNormal, rec.Normal)
	}
}
