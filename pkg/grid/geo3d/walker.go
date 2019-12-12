package geo3d

type Walker struct {
	P      *Point
	V      *Vector
	Walked int
}

func NewWalker(p *Point, v *Vector) *Walker {
	return &Walker{p, v, 0}
}

func (w *Walker) Finished() bool {
	return w.V.F == w.Walked
}

func (w *Walker) Walk(steps int) *Point {
	w.Walked += steps

	for n := 0; n < steps; n++ {
		w.P.X += w.V.P.X * w.V.F
		w.P.Y += w.V.P.Y * w.V.F
		w.P.Z += w.V.P.Z * w.V.F
	}

	return w.P
}

func (w *Walker) WalkOne() *Point {
	return w.Walk(1)
}

func (w *Walker) WalkAll() *Point {
	return w.Walk(w.V.F)
}
