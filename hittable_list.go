package main

type HittableList struct {
	Objects []Hittable
}

func (hl *HittableList) Add(object Hittable) {
	hl.Objects = append(hl.Objects, object)
}

func (hl HittableList) Hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closestSoFar := tMax
	var rec HitRecord

	for _, object := range hl.Objects {
		if hit, tempRec := object.Hit(ray, tMin, closestSoFar); hit {
			hitAnything = true
			closestSoFar = tempRec.T
			rec = tempRec
		}
	}

	return hitAnything, rec
}
