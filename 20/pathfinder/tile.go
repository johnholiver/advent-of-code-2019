package pathfinder

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/astar"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
)

type MazeTileKind struct {
	Name     string
	IsPortal bool
	Internal bool
}

func NewMazeTileKind(g *grid.Grid, x, y int) MazeTileKind {
	isPortal := false
	internal := true
	c := string(g.Get(x, y).Value.(rune))

	if portalCode, yes := tileIsPortal(g, x, y); yes {
		c = portalCode
		isPortal = true

		if x == 2 || x == g.Width-3 || y == 2 || y == g.Height-3 {
			internal = false
		}
	}

	return MazeTileKind{c, isPortal, internal}
}

func (t MazeTileKind) String() string {
	iS := "e"
	if t.Internal {
		iS = "i"
	}
	return fmt.Sprintf("%v%v", t.Name, iS)
}

func tileIsPortal(g *grid.Grid, x, y int) (string, bool) {
	portalCode := ""
	isPortal := false
	if g.Get(x, y).Value.(rune) != '.' {
		return portalCode, isPortal
	}

	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := g.Get(x+offset[0], y+offset[1]); n != nil &&
			n.Value.(rune) != '#' && n.Value.(rune) != '.' && n.Value.(rune) != ' ' {
			isPortal = true
			if offset[0] == -1 || offset[1] == -1 {
				portalCode += string(g.Get(n.X+offset[0], n.Y+offset[1]).Value.(rune)) + string(n.Value.(rune))
			} else if offset[0] == 1 || offset[1] == 1 {
				portalCode += string(n.Value.(rune)) + string(g.Get(n.X+offset[0], n.Y+offset[1]).Value.(rune))
			} else {
				panic("BUG")
			}
		}
	}
	return portalCode, isPortal
}

func (t MazeTileKind) IsTraversable(t2 astar.Traversable) bool {
	return t2.(MazeTileKind).Name == "." || t2.(MazeTileKind).IsPortal
}
