package pathfinder

import (
	"fmt"
	goastar "github.com/beefsack/go-astar"
	"github.com/johnholiver/advent-of-code-2019/pkg/astar"
)

type AllPaths map[string]map[string]*OnePath
type OnePath struct {
	Path     []goastar.Pather
	Distance float64
}

func (pp AllPaths) String() string {
	s := ""
	for tileFrom, fromV := range pp {
		for tileTo, onePath := range fromV {
			s += fmt.Sprintf("%v -> %v = %v\n", tileFrom, tileTo, onePath.Distance)
		}
	}
	return s
}

func NewAllPaths(tiles []*astar.Tile) AllPaths {
	paths := make(AllPaths)

	for i := 0; i < len(tiles); i++ {
		tile := tiles[i]
		paths[tile.Kind.(MazeTileKind).String()] = make(map[string]*OnePath)
	}

	for i := 0; i < len(tiles); i++ {
		tileFrom := tiles[i]
		for j := i + 1; j < len(tiles); j++ {
			tileTo := tiles[j]
			p, dist, found := goastar.Path(tileFrom, tileTo)
			if found {
				////Hacky, add 1 to the distance to account for the warp
				//if tileTo.Kind.(MazeTileKind).Name != "ZZ" {dist++}
				pathResult := &OnePath{p, dist}
				paths[tileFrom.Kind.(MazeTileKind).String()][tileTo.Kind.(MazeTileKind).String()] = pathResult
				paths[tileTo.Kind.(MazeTileKind).String()][tileFrom.Kind.(MazeTileKind).String()] = pathResult
			}
		}
	}
	return paths
}
