package hit

import (
	"raytracing/pkg/geom"
)

type Hittables struct {
	elements []Hittable
}

func NewHittables() Hittables {
	return Hittables{
		elements: make([]Hittable, 0),
	}
}

func (h *Hittables) Add(hittable Hittable) {
	h.elements = append(h.elements, hittable)
}

func (h *Hittables) Clear() {
	h.elements = make([]Hittable, 0)
}

func (h Hittables) Hit(
	ray geom.Ray,
	tMin float64,
	tMax float64,
	record *Record,
) bool {
	rec := new(Record)
	hit := false
	closest := tMax

	for i := 0; i < len(h.elements); i++ {
		if h.elements[i].Hit(ray, tMin, closest, rec) {
			hit = true
			closest = rec.T
			*record = *rec
		}
	}

	return hit
}
