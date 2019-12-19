package astar

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
)

// A Tile is a tile in a grid which implements Pather.
type Tile struct {
	// Kind is the kind of tile, potentially affecting movement.
	Kind rune
	// X and Y are the coordinates of the tile.
	X, Y int
	// W is a reference to the World that the tile is a part of.
	W *grid.Grid
}

func (t *Tile) String() string {
	return fmt.Sprintf("%c(%v,%v)", t.Kind, t.X, t.Y)
}

// PathNeighbors returns the neighbors of the tile, excluding blockers and
// tiles off the edge of the board.
func (t *Tile) PathNeighbors() []astar.Pather {
	var neighbors []astar.Pather
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := t.W.Get(t.X+offset[0], t.Y+offset[1]); n != nil &&
			n.Value.(*Tile).Kind != '#' {
			neighbors = append(neighbors, n.Value.(*Tile))
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

// PathEstimatedCost uses Manhattan Distance to estimate orthogonal Distance
// between non-adjacent nodes.
func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}
