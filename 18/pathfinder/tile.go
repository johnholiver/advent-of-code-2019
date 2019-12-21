package pathfinder

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/astar"
)

type MazeTileKind struct {
	Value rune
}

func (t MazeTileKind) String() string {
	return fmt.Sprintf("%c", t.Value)
}

func (t MazeTileKind) IsTraversable(t2 astar.Traversable) bool {
	return t2.(MazeTileKind).Value != '#'
}
