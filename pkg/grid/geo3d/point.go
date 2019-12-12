package geo3d

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
	Z int
}

func NewPoint(x, y, z int) *Point {
	return &Point{x, y, z}
}

func (p Point) String() string {
	return fmt.Sprintf("(%v,%v,%v)", p.X, p.Y, p.Z)
}

func (a Point) Equals(b Point) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

func (p *Point) Transform(x, y, z int) {
	p.X += x
	p.Y += y
	p.Z += z
}

func (a *Point) TransformByPoint(b Point) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
}

func (p *Point) Abs() *Point {
	xAbs := int(math.Abs(float64(p.X)))
	yAbs := int(math.Abs(float64(p.Y)))
	zAbs := int(math.Abs(float64(p.Z)))
	return NewPoint(xAbs, yAbs, zAbs)
}
