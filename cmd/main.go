package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"raytracing/pkg/camera"
	"raytracing/pkg/color"
	"raytracing/pkg/geom"
	"raytracing/pkg/hit"
	"raytracing/pkg/mat"
)

func rayColor(ray geom.Ray, world hit.Hittable, depth int) color.Color {
	record := hit.Record{}

	if depth <= 0 {
		return color.Color{}
	}

	if world.Hit(ray, 0.001, math.MaxFloat64, &record) {
		scattered := geom.Ray{}
		attenuation := color.Color{}
		if (*record.Material).Scatter(ray, &record, &attenuation, &scattered) {
			return rayColor(scattered, world, depth-1).Mul(attenuation)
		}
		return color.Color{}
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
	width := 300
	height := int(float64(width) / aspectRatio)
	samples := 100
	depth := 50

	// World
	world := hit.NewHittables()

	materialGround := mat.Lambertian{X: 0.8, Y: 0.8}
	materialCenter := mat.Lambertian{X: 0.7, Y: 0.3, Z: 0.3}
	materialLeft := mat.Metal{X: 0.8, Y: 0.8, Z: 0.8}
	materialRight := mat.Metal{X: 0.8, Y: 0.6, Z: 0.2}

	world.Add(hit.Sphere{
		Center:   geom.Point{Y: -100.5, Z: -1},
		Radius:   100,
		Material: materialGround,
	})
	world.Add(hit.Sphere{
		Center:   geom.Point{Z: -1},
		Radius:   0.5,
		Material: materialCenter,
	})
	world.Add(hit.Sphere{
		Center:   geom.Point{X: -1, Z: -1},
		Radius:   0.5,
		Material: materialLeft,
	})
	world.Add(hit.Sphere{
		Center:   geom.Point{X: 1, Z: -1},
		Radius:   0.5,
		Material: materialRight,
	})

	// Camera
	cam := camera.NewCamera()

	// Rendering
	fmt.Fprintf(f, "P3\n %d %d\n255\n", width, height)

	for j := height - 1; j >= 0; j-- {
		fmt.Printf("%d lines remaining\n", j)
		for i := 0; i < width; i++ {
			c := color.Color{}
			for s := 0; s < samples; s++ {
				u := (float64(i) + rand.Float64()) / float64(width-1)
				v := (float64(j) + rand.Float64()) / float64(height-1)
				ray := cam.Ray(u, v)
				c = c.Add(rayColor(ray, world, depth))
			}
			fmt.Fprintf(f, "%s\n", c.ToString(uint64(samples)))
		}
	}

	file, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	f.WriteTo(file)
}
