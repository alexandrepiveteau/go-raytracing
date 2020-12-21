package main

import (
	"bytes"
	"fmt"
	"os"
	"raytracing/pkg/color"
	"raytracing/pkg/geom"
)

func rayColor(ray geom.Ray) color.Color {
	unit := ray.Direction.Unit()
	t := 0.5 * (unit.Y + 1.0)
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
