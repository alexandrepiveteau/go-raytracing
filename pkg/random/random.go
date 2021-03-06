package random

import (
	"math/rand"
	"raytracing/pkg/geom"
)

func random(from float64, until float64) float64 {
	size := until - from
	return from + rand.Float64()*size
}

func randomVector(from, until float64) geom.Vec {
	return geom.Vec{
		X1: random(from, until),
		X2: random(from, until),
		X3: random(from, until),
	}
}

func RandomUnitSphere() geom.Vec {
	for {
		vec := randomVector(-1, 1)
		if vec.LengthSquared() < 1 {
			return vec
		}
	}
}

func RandomUnitVector() geom.Vec  {
	return RandomUnitSphere().Normalize()
}
