package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"raytracing/pkg/color"
	"raytracing/pkg/geom"
)

func hitSphere(center geom.Point, radius float64, ray geom.Ray) float64 {
	oc := ray.Origin.Sub(center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1
	}
	res := (-b - math.Sqrt(discriminant)) / (2 * a)
	return res
}

func rayColor(ray geom.Ray) color.Color {
	t := hitSphere(geom.Point{Z: -1}, 0.5, ray)
	if t > 0 {
		N := ray.At(t).Sub(geom.Point{Z: -1}).Unit()
		return color.Color{
			X: N.X + 1,
			Y: N.Y + 1,
			Z: N.Z + 1,
		}.Times(0.5)
	}
	unit := ray.Direction.Unit()
	t = 0.5 * (unit.Y + 1.0)
	from := color.Color{X: 1.0, Y: 1.0, Z: 1.0}
	to := color.Color{X: 0.5, Y: 0.7, Z: 1.0}
	return from.Times(1.0 - t).Add(to.Times(t))
}

func main() {
	f := new(bytes.Buffer)

	// Image
	aspectRatio := 16.0 / 9.0
	width := 1000
	height := int(float64(width) / aspectRatio)

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := geom.Vec{}
	horizontal := geom.Vec{X: viewportWidth}
	vertical := geom.Vec{Y: viewportHeight}
	lowerLeft := origin.
		Sub(horizontal.Div(2)).
		Sub(vertical.Div(2)).
		Sub(geom.Vec{Z: focalLength})

	// Rendering
	fmt.Fprintf(f, "P3\n %d %d\n255\n", width, height)

	for j := height - 1; j >= 0; j-- {
		fmt.Printf("%d lines remaining\n", j)
		for i := 0; i < width; i++ {
			u := float64(i) / float64(width-1)
			v := float64(j) / float64(height-1)
			ray := geom.Ray{
				Origin: origin,
				Direction: lowerLeft.
					Add(horizontal.Times(u)).
					Add(vertical.Times(v)).
					Sub(origin),
			}
			c := rayColor(ray)
			fmt.Fprintf(f, "%s\n", c.ToString())
		}
	}

	file, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	f.WriteTo(file)
}
