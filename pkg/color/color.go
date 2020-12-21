package color

import (
	"fmt"
	"raytracing/pkg/geom"
)

type Color geom.Vec

func (c Color) ToString() string {
	r := uint64(c.X * 255.999)
	g := uint64(c.Y * 255.999)
	b := uint64(c.Z * 255.999)
	return fmt.Sprintf("%d %d %d", r, g, b)
}
