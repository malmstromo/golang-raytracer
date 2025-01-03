package main

type HitRecord struct {
	T         float64
	P, Normal Vec3
}

type Hittable interface {
	Hit(ray Ray, rayTmin float64, rayTmax float64, rec HitRecord) (bool, HitRecord)
}
