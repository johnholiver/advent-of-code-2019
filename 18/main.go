package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/18/pathfinder"
	"github.com/johnholiver/advent-of-code-2019/pkg/astar"
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

	paths := pathfinder.NewAllPaths(kPlusAt)

	atK, leftKs := simplifyDependencyTreeInput(at, ks)

	_, root := buildDependencyTree(atK, leftKs, paths)
	_, c := leastCostyPath(paths, root)

	return fmt.Sprint(c)
}

func simplifyDependencyTreeInput(at *astar.Tile, ks []*astar.Tile) (rune, []rune) {
	atK := at.Kind.(pathfinder.MazeTileKind).Value
	leftKs := make([]rune, len(ks))
	for i, k := range ks {
		leftKs[i] = k.Kind.(pathfinder.MazeTileKind).Value
	}
	return atK, leftKs
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
		g.Get(x, y).Value = &astar.Tile{pathfinder.MazeTileKind{c}, x, y, g}
		x++
	}
	return g
}

func tileFormatter(e interface{}) string {
	return e.(*astar.Tile).Kind.(pathfinder.MazeTileKind).String()
}

func fetchKeys(g *grid.Grid) (*astar.Tile, []*astar.Tile, []*astar.Tile) {
	var at *astar.Tile
	var keys []*astar.Tile
	var doors []*astar.Tile

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			vp := g.Get(x, y)
			c := vp.Value.(*astar.Tile).Kind.(pathfinder.MazeTileKind).Value
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

type DependencyNode struct {
	Value   rune
	MinCost int
}

func NewDependencyNode(c rune) *DependencyNode {
	return &DependencyNode{
		c,
		int(^uint(0) >> 1),
	}
}

func (n *DependencyNode) String() string {
	return fmt.Sprintf("%c", n.Value)
}

func depTreeFormatter(e interface{}) string {
	return e.(*DependencyNode).String()
}

func buildDependencyTree(atK rune, leftKs []rune, paths pathfinder.AllPaths) (*graph.Graph, *graph.Node) {
	depGraph := graph.NewGraph()
	depGraph.SetFormatter(depTreeFormatter)
	atN := depGraph.BuildBranch(NewDependencyNode(atK), nil)
	recursiveAddTreeBranch(paths, depGraph, leftKs, atN, atN.Value.(*DependencyNode).MinCost)
	return depGraph, atN
}

func recursiveAddTreeBranch(paths pathfinder.AllPaths, depGraph *graph.Graph, leftK []rune, parent *graph.Node, maxCost int) int {
	indexOf := func(slice []rune, val rune) (int, bool) {
		for i, item := range slice {
			if item == val {
				return i, true
			}
		}
		return -1, false
	}

	if len(leftK) == 0 {
		parent.Value.(*DependencyNode).MinCost = 0
		return 0
	}

	var nodeMinPath *graph.Node
	costCap := maxCost

	for iK, k := range leftK {
		dependsOnK := false
		for _, dep := range paths[k][parent.Value.(*DependencyNode).Value].Dependencies {
			if _, dependsOnK = indexOf(leftK, dep); dependsOnK {
				break
			}
		}
		if !dependsOnK {
			//this is one possible branch
			kN := graph.NewNode(NewDependencyNode(k))
			kN.SetFormatter(depTreeFormatter)

			//up cost
			branchCost := int(paths[parent.Value.(*DependencyNode).Value][kN.Value.(*DependencyNode).Value].Distance)
			if branchCost > maxCost {
				return int(^uint(0) >> 1)
			}

			//fmt.Printf("%v: Down [%c->%c] - leftover: %c\n",maxCost, parent.Value.(rune),k, leftK)
			notKs := append(make([]rune, 0), leftK[:iK]...)
			notKs = append(notKs, leftK[iK+1:]...)

			//down cost
			if nodeMinPath != nil {
				costCap = maxCost - nodeMinPath.Value.(*DependencyNode).MinCost
			}
			branchCost += recursiveAddTreeBranch(paths, depGraph, notKs, kN, costCap)
			//fmt.Printf("%v: Up [%c->%c] - leftover: %c\n",maxCost, parent.Value.(rune),k, leftK)

			if nodeMinPath == nil ||
				branchCost < nodeMinPath.Value.(*DependencyNode).MinCost {
				kN.Value.(*DependencyNode).MinCost = branchCost
				nodeMinPath = kN
			}
		}
	}

	parent.AddChild(nodeMinPath)
	parent.Value.(*DependencyNode).MinCost = nodeMinPath.Value.(*DependencyNode).MinCost
	return parent.Value.(*DependencyNode).MinCost
}

func leastCostyPath(paths pathfinder.AllPaths, parent *graph.Node) ([]rune, int) {
	if len(parent.Children) == 0 {
		return []rune{parent.Value.(*DependencyNode).Value}, 0
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
	newLcCost := lcCost + int(paths[parent.Value.(*DependencyNode).Value][lcPath[0]].Distance)
	newLcPath := append([]rune{parent.Value.(*DependencyNode).Value}, lcPath...)
	return newLcPath, newLcCost
}
