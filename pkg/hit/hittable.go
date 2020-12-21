package hit

import "raytracing/pkg/geom"

type Record struct {
	P         geom.Point
	Normal    geom.Vec
	T         float64
	FrontFace bool
}

func (r *Record) SetFaceNormal(ray geom.Ray, outward geom.Vec) {
	r.FrontFace = ray.Direction.Dot(outward) < 0
	if r.FrontFace {
		r.Normal = outward
	} else {
		r.Normal = outward.Inv()
	}
}

type Hittable interface {
	Hit(ray geom.Ray, tMin float64, tMax float64, record *Record) bool
}
