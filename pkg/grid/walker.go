package grid

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

	switch w.V.D {
	case "U":
		w.P.Y += steps
	case "D":
		w.P.Y -= steps
	case "R":
		w.P.X += steps
	case "L":
		w.P.X -= steps
	}

	return w.P
}

func (w *Walker) WalkOne() *Point {
	return w.Walk(1)
}

func (w *Walker) WalkAll() *Point {
	return w.Walk(w.V.F)
}
