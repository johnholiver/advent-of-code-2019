package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/18/pathfinder"
	"github.com/johnholiver/advent-of-code-2019/pkg/astar"
	"github.com/johnholiver/advent-of-code-2019/pkg/graph"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"log"
	"os"
	"sort"
	"unicode"
)

func main() {
	file, err := input.Load("2019", "18")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//fmt.Printf("Result part1: %v\n", part1(file))
	//
	//file.Seek(0, io.SeekStart)
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

	ats, ks, _ := fetchKeys(buildGrid(input))
	kPlusAt := append(ks, ats...)

	paths := pathfinder.NewAllPaths(kPlusAt)

	atK, leftKs := simplifyDependencyTreeInput(ats, ks)

	_, root := buildDependencyTree(atK, leftKs, paths)

	return fmt.Sprint(root.Value.(*DependencyNode).MinCost)
}

func simplifyDependencyTreeInput(at []*astar.Tile, ks []*astar.Tile) (rune, []rune) {
	atK := at[0].Kind.(pathfinder.MazeTileKind).Value
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
	g := buildGrid(input)

	fixGrid(g)

	ats, ks, _ := fetchKeys(g)
	kPlusAt := append(ks, ats...)

	paths := pathfinder.NewAllPaths(kPlusAt)

	atK, leftKs := simplifyDependencyTreeInput(ats, ks)

	fmt.Println(paths)

	_, root := buildDependencyTree(atK, leftKs, paths)

	return fmt.Sprint(root.Value.(*DependencyNode).MinCost)
}

func fixGrid(g *grid.Grid) {
	var atPos *[2]int
	atPos = nil
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Get(x, y).Value.(*astar.Tile).Kind.(pathfinder.MazeTileKind).Value == '@' {
				atPos = &[2]int{x, y}
				break
			}
		}
		if atPos != nil {
			break
		}
	}

	vaultFixer := map[[2]int]rune{
		{-1, -1}: '@',
		{-1, 0}:  '#',
		{-1, 1}:  '@',
		{0, -1}:  '#',
		{0, 0}:   '#',
		{0, 1}:   '#',
		{1, -1}:  '@',
		{1, 0}:   '#',
		{1, 1}:   '@',
	}
	for offset, c := range vaultFixer {
		g.Get(atPos[0]+offset[0], atPos[1]+offset[1]).Value.(*astar.Tile).Kind = pathfinder.MazeTileKind{c}
	}
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

func fetchKeys(g *grid.Grid) ([]*astar.Tile, []*astar.Tile, []*astar.Tile) {
	var ats []*astar.Tile
	var keys []*astar.Tile
	var doors []*astar.Tile

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			vp := g.Get(x, y)
			c := vp.Value.(*astar.Tile).Kind.(pathfinder.MazeTileKind).Value
			if c == '@' {
				ats = append(ats, vp.Value.(*astar.Tile))
			}
			if unicode.IsLower(c) {
				keys = append(keys, vp.Value.(*astar.Tile))
			}
			if unicode.IsUpper(c) {
				doors = append(doors, vp.Value.(*astar.Tile))
			}

		}
	}
	return ats, keys, doors
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
	cache := make(map[string]int)
	recursiveAddTreeBranch(paths, depGraph, cache, leftKs, atN, atN.Value.(*DependencyNode).MinCost)
	return depGraph, atN
}

func recursiveAddTreeBranch(paths pathfinder.AllPaths, depGraph *graph.Graph, cache map[string]int, leftK []rune, parent *graph.Node, maxCost int) int {
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

	leftKCacheKey := make([]rune, len(leftK))
	copy(leftKCacheKey, leftK)
	sort.Slice(leftKCacheKey, func(i, j int) bool { return leftKCacheKey[i] < leftKCacheKey[j] })
	cacheKey := string(parent.Value.(*DependencyNode).Value)
	for _, c := range leftKCacheKey {
		cacheKey += string(c)
	}
	if cached, ok := cache[cacheKey]; ok {
		return cached
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
			branchCost += recursiveAddTreeBranch(paths, depGraph, cache, notKs, kN, costCap)
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

	cache[cacheKey] = parent.Value.(*DependencyNode).MinCost

	return parent.Value.(*DependencyNode).MinCost
}
