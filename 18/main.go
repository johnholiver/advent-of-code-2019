package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/18/astar"
	"github.com/johnholiver/advent-of-code-2019/pkg/graph"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := input.Load("2019", "18")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Result part1: %v\n", part1(file))

	file.Seek(0, io.SeekStart)
	fmt.Printf("Result part2: %v\n", part2(file))
}

func part1(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += fmt.Sprintln(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	at, ks, _ := fetchKeys(buildGrid(input))
	kPlusAt := append(ks, at)

	paths := astar.NewAllPaths(kPlusAt)

	depTree := buildDependencyTree(at, ks, paths)
	_, c := leastCostyPath(paths, depTree.Roots[at.Kind])

	return fmt.Sprint(c)
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += fmt.Sprintln(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//k,d:=fetchKeys(buildGrid(input))

	return "implement me"
}

func buildGrid(input string) *grid.Grid {
	h := 0
	w := 0
	for _, c := range input {
		if unicode.IsSpace(c) {
			h++
		}
		if h == 0 {
			w++
		}
	}
	g := grid.NewGrid(w, h)
	g.SetFormatter(tileFormatter)

	x := 0
	y := 0
	for _, c := range input {
		if unicode.IsSpace(c) {
			x = 0
			y++
			continue
		}
		g.Get(x, y).Value = &astar.Tile{c, x, y, g}
		x++
	}
	return g
}

func asciiFormatter(e interface{}) string {
	return string(e.(rune))
}

func tileFormatter(e interface{}) string {
	return string(e.(*astar.Tile).Kind)
}

func fetchKeys(g *grid.Grid) (*astar.Tile, []*astar.Tile, []*astar.Tile) {
	var at *astar.Tile
	var keys []*astar.Tile
	var doors []*astar.Tile

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			vp := g.Get(x, y)
			c := vp.Value.(*astar.Tile).Kind
			if c == '@' {
				at = vp.Value.(*astar.Tile)
			}
			if unicode.IsLower(c) {
				keys = append(keys, vp.Value.(*astar.Tile))
			}
			if unicode.IsUpper(c) {
				doors = append(doors, vp.Value.(*astar.Tile))
			}

		}
	}
	return at, keys, doors
}

func buildDependencyTree(at *astar.Tile, ks []*astar.Tile, paths astar.AllPaths) *graph.Graph {
	depGraph := graph.NewGraph()
	depGraph.SetFormatter(asciiFormatter)
	atN := depGraph.BuildBranch(at.Kind, nil)
	leftKs := make([]rune, len(ks))
	for i, k := range ks {
		leftKs[i] = k.Kind
	}
	level := 1
	recursiveAddTreeBranch(paths, depGraph, leftKs, atN, level)
	return depGraph
}

func recursiveAddTreeBranch(paths astar.AllPaths, depGraph *graph.Graph, leftK []rune, parent *graph.Node, level int) {
	indexOf := func(slice []rune, val rune) (int, bool) {
		for i, item := range slice {
			if item == val {
				return i, true
			}
		}
		return -1, false
	}

	for iK, k := range leftK {
		dependsOnK := false
		for _, dep := range paths[k][parent.Value.(rune)].Dependencies {
			if _, dependsOnK = indexOf(leftK, dep); dependsOnK {
				break
			}
		}
		if !dependsOnK {
			//fmt.Printf("%v: Down [%c->%c] - leftover: %c\n",level, parent.Value.(rune),k, leftK)
			kN := depGraph.BuildBranch(k, parent)
			notKs := append(make([]rune, 0), leftK[:iK]...)
			notKs = append(notKs, leftK[iK+1:]...)
			recursiveAddTreeBranch(paths, depGraph, notKs, kN, level+1)
			//fmt.Printf("%v: Up [%c->%c] - leftover: %c\n",level, parent.Value.(rune),k, leftK)
		}
	}
}

func leastCostyPath(paths astar.AllPaths, parent *graph.Node) ([]rune, int) {
	if len(parent.Children) == 0 {
		return []rune{parent.Value.(rune)}, 0
	}
	var lcPath []rune
	lcCost := int(^uint(0) >> 1)
	for _, child := range parent.Children {
		p, c := leastCostyPath(paths, child)
		if c < lcCost {
			lcPath = p
			lcCost = c
		}
	}
	newLcCost := lcCost + int(paths[parent.Value.(rune)][lcPath[0]].Distance)
	newLcPath := append([]rune{parent.Value.(rune)}, lcPath...)
	return newLcPath, newLcCost
}
