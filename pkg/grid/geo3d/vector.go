package geo3d

import "fmt"

type Vector struct {
	P Point
	F int
}

func (p Vector) String() string {
	return fmt.Sprintf("%v|%v", p.P, p.F)
}

func NewVector(direction Point, force int) *Vector {
	return &Vector{direction, force}
}

func (a *Vector) Equals(b *Vector) bool {
	return a.P.Equals(b.P) && a.F == b.F
}
