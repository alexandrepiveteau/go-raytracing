package camera

import "raytracing/pkg/geom"

type Camera struct {
	Origin geom.Point

	LowerLeftCorner geom.Point
	Horizontal      geom.Vec
	Vertical        geom.Vec
}

func NewCamera() Camera {
	aspectRatio := 16.0 / 9.0
	height := 2.0
	width := aspectRatio * height
	focal := 1.0

	return Camera{
		Origin:     geom.Point{},
		Horizontal: geom.Vec{X: width},
		Vertical:   geom.Vec{Y: height},
		LowerLeftCorner: geom.Point{}.
			Sub(geom.Vec{X: width}.Div(2)).
			Sub(geom.Vec{Y: height}.Div(2)).
			Sub(geom.Vec{Z: focal}),
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
