package geom

type Ray struct {
	Origin    Point
	Direction Vec
}

func (r Ray) At(t float64) Point {
	return r.Origin.Add(r.Direction.Times(t))
}
