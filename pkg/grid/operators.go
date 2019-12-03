package grid

import "math"

//if 𝑥=(𝑎,𝑏) and 𝑦=(𝑐,𝑑), the Manhattan distance between 𝑥 and 𝑦 is
//|𝑎−𝑐|+|𝑏−𝑑|.
func Manhattan(p1, p2 Point) int {
	return int(math.Abs(float64(p1.X-p2.X))) + int(math.Abs(float64(p1.Y-p2.Y)))
}
