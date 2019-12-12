package grid

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%v,%v)", p.X, p.Y)
}

func (a Point) Equals(b Point) bool {
	return a.X == b.X && a.Y == b.Y
}

func (p *Point) Transform(x, y int) {
	p.X += x
	p.Y += y
}

func (p *Point) MirrorX() {
	p.X *= -1
}

func (p *Point) MirrorY() {
	p.Y *= -1
}

// Distance finds the length of the hypotenuse between two points.
// Forumula is the square root of (x2 - x1)^2 + (y2 - y1)^2
func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)
	second := math.Pow(float64(p2.Y-p.Y), 2)
	return math.Sqrt(first + second)
}

type ValuedPoint struct {
	*Point
	Value int
}

func NewValuedPoint(x, y, value int) *ValuedPoint {
	return &ValuedPoint{
		NewPoint(x, y),
		value,
	}
}

func (a ValuedPoint) Equals(b ValuedPoint) bool {
	return a.X == b.X && a.Y == b.Y && a.Value == b.Value
}

func (p ValuedPoint) String() string {
	return fmt.Sprintf("(%v,%v|%v)", p.X, p.Y, p.Value)
}
