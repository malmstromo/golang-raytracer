package main

import "testing"

func TestVec3_Add(t *testing.T) {
	v1 := Vec3{X: 1, Y: 2, Z: 3}
	v2 := Vec3{X: 4, Y: 5, Z: 6}
	expected := Vec3{X: 5, Y: 7, Z: 9}

	result := v1.Add(v2)
	if result != expected {
		t.Errorf("Add: expected %v, got %v", expected, result)
	}
}

func TestVec3_Dot(t *testing.T) {
	v1 := Vec3{X: 1, Y: 2, Z: 3}
	v2 := Vec3{X: 4, Y: 5, Z: 6}
	expected := float64(32) // 1*4 + 2*5 + 3*6

	result := v1.Dot(v2)
	if result != expected {
		t.Errorf("Dot: expected %v, got %v", expected, result)
	}
}

func TestVec3_Length(t *testing.T) {
	v := Vec3{X: 3, Y: 4, Z: 0}
	expected := float64(5) // sqrt(3^2 + 4^2)

	result := v.Length()
	if result != expected {
		t.Errorf("Length: expected %v, got %v", expected, result)
	}
}
