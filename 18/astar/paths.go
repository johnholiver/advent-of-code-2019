package astar

import (
	"fmt"
	goastar "github.com/beefsack/go-astar"
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

func NewAllPaths(kPlusAt []*Tile) AllPaths {
	paths := make(AllPaths, len(kPlusAt))

	for i := 0; i < len(kPlusAt); i++ {
		tile := kPlusAt[i]
		paths[tile.Kind] = make(map[rune]*OnePath, len(kPlusAt)-1)
	}

	for i := 0; i < len(kPlusAt); i++ {
		tileFrom := kPlusAt[i]
		for j := i + 1; j < len(kPlusAt); j++ {
			tileTo := kPlusAt[j]
			p, dist, found := goastar.Path(tileFrom, tileTo)
			if !found {
				panic("AHHH!")
			}

			pathResult := &OnePath{p, dist, calculateDependencies(p)}

			paths[tileFrom.Kind][tileTo.Kind] = pathResult
			paths[tileTo.Kind][tileFrom.Kind] = pathResult
		}
	}
	return paths
}

func calculateDependencies(p []goastar.Pather) []rune {
	var dd []rune
	for _, e := range p {
		tile := e.(*Tile)
		if unicode.IsUpper(tile.Kind) {
			dd = append(dd, unicode.ToLower(tile.Kind))
		}
	}
	return dd
}
