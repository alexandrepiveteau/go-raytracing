package hit

import (
	"raytracing/pkg/color"
	"raytracing/pkg/geom"
)

type Material interface {
	Scatter(
		ray geom.Ray,
		record *Record,
		attenuation *color.Color,
		scattered *geom.Ray,
	) bool
}
