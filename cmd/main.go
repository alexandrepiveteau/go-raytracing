package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"raytracing/pkg/color"
	"raytracing/pkg/geom"
	"raytracing/pkg/hit"
)

func rayColor(ray geom.Ray, world hit.Hittable) color.Color {
	record := hit.Record{}
	if world.Hit(ray, 0, math.MaxFloat64, &record) {
		return (record.Normal.Add(color.Color{X: 1, Y: 1, Z: 1})).
			Times(0.5)
	}
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

	// World
	world := hit.NewHittables()
	world.Add(hit.Sphere{Center: geom.Point{Z: -1}, Radius: 0.5})
	world.Add(hit.Sphere{Center: geom.Point{Y: -100.5, Z: -1}, Radius: 100})

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
			c := rayColor(ray, world)
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
