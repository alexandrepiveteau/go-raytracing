package mat

import (
	"raytracing/pkg/color"
	"raytracing/pkg/geom"
	"raytracing/pkg/hit"
)

type Metal color.Color

func reflect(v geom.Vec, n geom.Vec) geom.Vec {
	return v.Sub(n.Times(v.Dot(n)).Times(2))
}

func (m Metal) Scatter(
	ray geom.Ray,
	record *hit.Record,
	attenuation *color.Color,
	scattered *geom.Ray,
) bool {
	reflected := reflect(ray.Direction.Unit(), record.Normal)
	*scattered = geom.Ray{Origin: record.P, Direction: reflected}
	*attenuation = color.Color{X: m.X, Y: m.Y, Z: m.Y}
	return true
}
