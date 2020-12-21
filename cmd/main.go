package main

import (
	"bytes"
	"fmt"
	"os"
	"raytracing/pkg/color"
)

func main() {
	f := new(bytes.Buffer)

	width := 1000
	height := 1000

	fmt.Fprintf(f, "P3\n %d %d\n255\n", width, height)

	for j := height - 1; j >= 0; j-- {
		for i := 0; i < width; i++ {
			c := color.Color{
				X: float64(i) / float64(width-1),
				Y: float64(j) / float64(height-1),
				Z: 0.25,
			}
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
