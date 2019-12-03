package grid

type Vector struct {
	D string
	F int
}

func NewVector(direction string, force int) *Vector {
	return &Vector{direction, force}
}

func (a *Vector) Equals(b *Vector) bool {
	return a.D == b.D && a.F == b.F
}
