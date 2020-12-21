package mat

import (
	"raytracing/pkg/color"
	"raytracing/pkg/geom"
	"raytracing/pkg/hit"
	"raytracing/pkg/random"
)

type Lambertian color.Color

func (l Lambertian) Scatter(
	ray geom.Ray,
	record *hit.Record,
	attenuation *color.Color,
	scattered *geom.Ray,
) bool {
	direction := record.Normal.Add(random.RandomUnitVector())

	// Catch degenerate scatter direction
	if direction.Zero() {
		direction = record.Normal
	}

	*scattered = geom.Ray{Origin: record.P, Direction: direction}
	*attenuation = color.Color{X: l.X, Y: l.Y, Z: l.Z}
	return true
}
