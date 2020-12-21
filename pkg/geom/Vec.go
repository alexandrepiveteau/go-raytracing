package geom

import (
	"fmt"
	"math"
)

type Vec struct {
	X, Y, Z float64
}

type Point = Vec

// UNARY

func (u Vec) Inv() Vec {
	return Vec{
		X: -u.X,
		Y: -u.Y,
		Z: -u.Z,
	}
}

// BINARY

func (u Vec) Add(v Vec) Vec {
	return Vec{
		X: u.X + v.X,
		Y: u.Y + v.Y,
		Z: u.Z + v.Z,
	}
}

func (u Vec) Sub(v Vec) Vec {
	return Vec{
		X: u.X - v.X,
		Y: u.Y - v.Y,
		Z: u.Z - v.Z,
	}
}

func (u Vec) Mul(v Vec) Vec {
	return Vec{
		X: u.X * v.X,
		Y: u.Y * v.Y,
		Z: u.Z * v.Z,
	}
}

func (u Vec) Dot(v Vec) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

func (u Vec) Cross(v Vec) Vec {
	return Vec{
		X: u.Y*v.Z - u.Z*v.Y,
		Y: u.Z*v.X - u.X*v.Z,
		Z: u.X*v.Y - u.Y*v.X,
	}
}

// SCALARS

func (u Vec) Times(scalar float64) Vec {
	return Vec{
		X: u.X * scalar,
		Y: u.Y * scalar,
		Z: u.Z * scalar,
	}
}

func (u Vec) Div(scalar float64) Vec {
	return u.Times(1 / scalar)
}

// NORMS

func (u Vec) Length() float64 {
	return math.Sqrt(u.LengthSquared())
}

func (u Vec) LengthSquared() float64 {
	return u.X*u.X + u.Y*u.Y + u.Z*u.Z
}

func (u Vec) Unit() Vec {
	return u.Div(u.Length())
}

// UTILS

func (u Vec) ToString(samples uint64) string {
	scale := 1.0 / float64(samples)

	r := u.X * scale
	g := u.Y * scale
	b := u.Z * scale

	clamp := func(v float64) float64{
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
		uint64(256 * clamp(r)),
		uint64(256 * clamp(g)),
		uint64(256 * clamp(b)),
	)
}
