package geom

import (
	"fmt"
	"math"
)

type Vec struct {
	X1, X2, X3 float64
}

type Point = Vec

// UNARY

func (u Vec) Inv() Vec {
	return Vec{
		X1: -u.X1,
		X2: -u.X2,
		X3: -u.X3,
	}
}

// BINARY

func (u Vec) Add(v Vec) Vec {
	return Vec{
		X1: u.X1 + v.X1,
		X2: u.X2 + v.X2,
		X3: u.X3 + v.X3,
	}
}

func (u Vec) Sub(v Vec) Vec {
	return Vec{
		X1: u.X1 - v.X1,
		X2: u.X2 - v.X2,
		X3: u.X3 - v.X3,
	}
}

func (u Vec) Mul(v Vec) Vec {
	return Vec{
		X1: u.X1 * v.X1,
		X2: u.X2 * v.X2,
		X3: u.X3 * v.X3,
	}
}

func (u Vec) Dot(v Vec) float64 {
	return u.X1*v.X1 + u.X2*v.X2 + u.X3*v.X3
}

func (u Vec) Cross(v Vec) Vec {
	return Vec{
		X1: u.X2*v.X3 - u.X3*v.X2,
		X2: u.X3*v.X1 - u.X1*v.X3,
		X3: u.X1*v.X2 - u.X2*v.X1,
	}
}

// SCALARS

func (u Vec) Times(scalar float64) Vec {
	return Vec{
		X1: u.X1 * scalar,
		X2: u.X2 * scalar,
		X3: u.X3 * scalar,
	}
}

func (u Vec) Div(scalar float64) Vec {
	return u.Times(1 / scalar)
}

// NORMS

func (u Vec) Zero() bool {
	// TODO : Fix approximate rounding.
	delta := 0.0000001
	return u.X1 < delta && u.X2 < 0 && u.X3 < 0
}

func (u Vec) Length() float64 {
	return math.Sqrt(u.LengthSquared())
}

func (u Vec) LengthSquared() float64 {
	return u.X1*u.X1 + u.X2*u.X2 + u.X3*u.X3
}

func (u Vec) Normalize() Vec {
	return u.Div(u.Length())
}

// UTILS

func (u Vec) ToString(samples uint64) string {
	scale := 1.0 / float64(samples)

	r := math.Sqrt(u.X1 * scale)
	g := math.Sqrt(u.X2 * scale)
	b := math.Sqrt(u.X3 * scale)

	clamp := func(v float64) float64 {
		if v <= 0 {
			return 0
		}
		if v >= 1 {
			return 0.999
		}
		return v
	}

	return fmt.Sprintf(
		"%d %d %d",
		uint64(256*clamp(r)),
		uint64(256*clamp(g)),
		uint64(256*clamp(b)),
	)
}
