package grid

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (a Point) Equals(b Point) bool {
	return a.X == b.X && a.Y == b.Y
}
