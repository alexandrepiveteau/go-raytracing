package mat

import (
	"raytracing/pkg/color"
	"raytracing/pkg/geom"
	"raytracing/pkg/hit"
	"raytracing/pkg/random"
)

type Metal struct {
	Color color.Color
	Fuzz  float64
}

func reflect(v geom.Vec, n geom.Vec) geom.Vec {
	return v.Sub(n.Times(v.Dot(n)).Times(2))
}

func (m Metal) Scatter(
	ray geom.Ray,
	record *hit.Record,
	attenuation *color.Color,
	scattered *geom.Ray,
) bool {
	reflected := reflect(ray.Direction.Normalize(), record.Normal)
	*scattered = geom.Ray{
		Origin:    record.P,
		Direction: reflected.Add(random.RandomUnitSphere().Times(m.Fuzz)),
	}
	*attenuation = color.Color{X1: m.Color.X1, X2: m.Color.X2, X3: m.Color.X2}
	return true
}
