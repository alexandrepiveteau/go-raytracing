package camera

import (
	"math"
	"raytracing/pkg/geom"
)

type Camera struct {
	Origin geom.Point

	LowerLeftCorner geom.Point
	Horizontal      geom.Vec
	Vertical        geom.Vec
}

func NewCamera(
	from, at geom.Vec,
	up geom.Vec,
	fov, aspectRatio float64,
) Camera {
	theta := fov * (math.Pi / 180)
	h := math.Tan(theta / 2)
	height := 2 * h
	width := aspectRatio * height

	w := from.Sub(at).Normalize()
	u := up.Cross(w).Normalize()
	v := w.Cross(u)

	return Camera{
		Origin:     from,
		Horizontal: u.Times(width),
		Vertical:   v.Times(height),
		LowerLeftCorner: from.
			Sub(u.Times(width).Div(2)).
			Sub(v.Times(height).Div(2)).
			Sub(from.Sub(at).Normalize()),
	}
}

func (c Camera) Ray(u float64, v float64) geom.Ray {
	return geom.Ray{
		Origin: c.Origin,
		Direction: c.LowerLeftCorner.
			Add(c.Horizontal.Times(u)).
			Add(c.Vertical.Times(v)).
			Sub(c.Origin),
	}
}
