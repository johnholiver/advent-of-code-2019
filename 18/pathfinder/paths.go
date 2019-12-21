package pathfinder

import (
	"fmt"
	goastar "github.com/beefsack/go-astar"
	"github.com/johnholiver/advent-of-code-2019/pkg/astar"
	"unicode"
)

type AllPaths map[rune]map[rune]*OnePath
type OnePath struct {
	Path         []goastar.Pather
	Distance     float64
	Dependencies []rune
}

func (pp AllPaths) String() string {
	s := ""
	for tileFrom, fromV := range pp {
		for tileTo, onePath := range fromV {
			s += fmt.Sprintf("%c -> %c = %v %c\n", tileFrom, tileTo, onePath.Distance, onePath.Dependencies)
		}
	}
	return s
}

func NewAllPaths(tiles []*astar.Tile) AllPaths {
	paths := make(AllPaths, len(tiles))

	for i := 0; i < len(tiles); i++ {
		tile := tiles[i]
		paths[tile.Kind.(MazeTileKind).Value] = make(map[rune]*OnePath, len(tiles)-1)
	}

	for i := 0; i < len(tiles); i++ {
		tileFrom := tiles[i]
		for j := i + 1; j < len(tiles); j++ {
			tileTo := tiles[j]
			p, dist, found := goastar.Path(tileFrom, tileTo)
			if !found {
				panic("AHHH!")
			}

			pathResult := &OnePath{p, dist, calculateDependencies(p)}

			paths[tileFrom.Kind.(MazeTileKind).Value][tileTo.Kind.(MazeTileKind).Value] = pathResult
			paths[tileTo.Kind.(MazeTileKind).Value][tileFrom.Kind.(MazeTileKind).Value] = pathResult
		}
	}
	return paths
}

func calculateDependencies(p []goastar.Pather) []rune {
	var dd []rune
	for _, e := range p {
		tile := e.(*astar.Tile)
		if unicode.IsUpper(tile.Kind.(MazeTileKind).Value) {
			dd = append(dd, unicode.ToLower(tile.Kind.(MazeTileKind).Value))
		}
	}
	return dd
}
