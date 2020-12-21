package hit

import (
	"math"
	"raytracing/pkg/geom"
)

type Sphere struct {
	Center   geom.Point
	Radius   float64
	Material Material
}

func (s Sphere) Hit(
	ray geom.Ray,
	tMin float64,
	tMax float64,
	record *Record,
) bool {
	oc := ray.Origin.Sub(s.Center)
	a := ray.Direction.LengthSquared()
	halfB := oc.Dot(ray.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrt := math.Sqrt(discriminant)

	root := (-halfB - sqrt) / a
	if root < tMin || root > tMax {
		root = (-halfB + sqrt) / a
		if root < tMin || root > tMax {
			return false
		}
	}

	record.T = root
	record.P = ray.At(record.T)
	outward := (record.P.Sub(s.Center)).Div(s.Radius)
	record.SetFaceNormal(ray, outward)
	record.Material = &s.Material

	return true
}
