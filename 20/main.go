package main

import (
	"bufio"
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"github.com/johnholiver/advent-of-code-2019/20/pathfinder"
	"github.com/johnholiver/advent-of-code-2019/pkg/astar"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
)

func main() {
	file, err := input.Load("2019", "20")
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

	_, pp := buildWorld(input)
	allpaths := pathfinder.NewAllPaths(pp)

	graph := buildGraph(pp, allpaths)

	vAA, _ := graph.GetMapping("AAe")
	vZZ, _ := graph.GetMapping("ZZe")
	bestPath, _ := graph.Shortest(vAA, vZZ)
	return fmt.Sprintln(bestPath.Distance)
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

	_, pp := buildWorld(input)
	allpaths := pathfinder.NewAllPaths(pp)

	graph := buildGraph2(pp, allpaths)

	vAA, _ := graph.GetMapping("AAe_0")
	vZZ, _ := graph.GetMapping("ZZe_0")
	bestPath, _ := graph.Shortest(vAA, vZZ)

	return fmt.Sprintln(bestPath.Distance)
}

func buildWorld(input string) (*grid.Grid, []*astar.Tile) {
	h := 0
	w := 0
	for _, c := range input {
		if c == '\n' {
			h++
		}
		if h == 0 {
			w++
		}
	}
	g := grid.NewGrid(w, h)
	g.SetFormatter(asciiFormatter)

	i := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := rune(input[i])
			g.Get(x, y).Value = c
			i++
		}
		i++
	}

	tileG := grid.NewGrid(w-4, h-4)
	tileG.SetFormatter(tileFormatter)
	var portals []*astar.Tile
	for y := 2; y < h-2; y++ {
		for x := 2; x < w-2; x++ {
			tileKind := pathfinder.NewMazeTileKind(g, x, y)
			tile := &astar.Tile{tileKind, x - 2, y - 2, tileG}

			if tileKind.IsPortal {
				portals = append(portals, tile)
			}

			tileG.Get(x-2, y-2).Value = tile
			i++
		}
		i++
	}

	return tileG, portals
}

func asciiFormatter(e interface{}) string {
	return string(e.(rune))
}

func tileFormatter(e interface{}) string {
	s := e.(*astar.Tile).Kind.(pathfinder.MazeTileKind).Name
	if len(s) > 1 {
		switch s {
		case "AA":
			s = "F"
		case "ZZ":
			s = "T"
		default:
			s = "@"
		}
	}

	return s
}

func buildGraph(portals []*astar.Tile, allpaths pathfinder.AllPaths) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	for _, p := range portals {
		if p.Kind.(pathfinder.MazeTileKind).Name == "AA" || p.Kind.(pathfinder.MazeTileKind).Name == "ZZ" {
			graph.AddMappedVertex(p.Kind.(pathfinder.MazeTileKind).String())
			continue
		}

		from := p.Kind.(pathfinder.MazeTileKind).String()
		//reverse String()
		iS := "i"
		if p.Kind.(pathfinder.MazeTileKind).Internal {
			iS = "e"
		}
		to := fmt.Sprintf("%v%v", p.Kind.(pathfinder.MazeTileKind).Name, iS)

		graph.AddMappedArc(from, to, int64(1))
	}

	for from, pp := range allpaths {
		for to, p := range pp {
			graph.AddMappedArc(from, to, int64(p.Distance))
		}
	}

	return graph
}

func buildGraph2(portals []*astar.Tile, allpaths pathfinder.AllPaths) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	maxLevel := 30
	for level := 0; level < maxLevel; level++ {
		for _, p := range portals {
			if p.Kind.(pathfinder.MazeTileKind).Name == "AA" || p.Kind.(pathfinder.MazeTileKind).Name == "ZZ" {
				if level == 0 {
					graph.AddMappedVertex(fmt.Sprintf("%v_%v", p.Kind.(pathfinder.MazeTileKind).String(), level))
				}
				continue
			}

			if level == 0 && !p.Kind.(pathfinder.MazeTileKind).Internal {
				continue
			}

			from := fmt.Sprintf("%v_%v", p.Kind.(pathfinder.MazeTileKind).String(), level)
			//reverse String()
			iS := "i"
			levelDiff := -1
			if p.Kind.(pathfinder.MazeTileKind).Internal {
				levelDiff = 1
				iS = "e"
			}
			to := fmt.Sprintf("%v%v_%v", p.Kind.(pathfinder.MazeTileKind).Name, iS, level+levelDiff)

			graph.AddMappedArc(from, to, int64(1))
		}
	}

	for level := 0; level < maxLevel; level++ {
		for from, pp := range allpaths {
			for to, p := range pp {
				graph.AddMappedArc(
					fmt.Sprintf("%v_%v", from, level),
					fmt.Sprintf("%v_%v", to, level),
					int64(p.Distance),
				)
			}
		}
	}

	return graph
}
