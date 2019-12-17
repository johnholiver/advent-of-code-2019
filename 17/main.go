package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
	robot "github.com/johnholiver/advent-of-code-2019/pkg/machine/screen"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := input.Load("2019", "17")
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
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	r := robot.NewScreen(program)
	ai := robot.NewScannerAI()
	r.SetAI(ai)
	r.Exec()

	ps := getIntersections(ai.(*robot.ScannerAI).Map)
	return fmt.Sprint(sumAlignmentParameters(ps))
}

func getIntersections(g *grid.Grid) []grid.Point {
	ps := make([]grid.Point, 0)
	for y := 1; y < g.Height-1; y++ {
		for x := 1; x < g.Width-1; x++ {
			cell := g.Get(x, y)
			if cell.Value == int('#') &&
				g.Get(x+1, y).Value == int('#') &&
				g.Get(x-1, y).Value == int('#') &&
				g.Get(x, y+1).Value == int('#') &&
				g.Get(x, y-1).Value == int('#') {
				ps = append(ps, *cell.Point)
			}
		}
	}
	return ps
}

func sumAlignmentParameters(ps []grid.Point) int {
	sum := 0
	for _, p := range ps {
		sum += p.X * p.Y
	}
	return sum
}

var aiAlgorithm = `A,B,A,B,A,C,A,C,B,C
R,6,L,10,R,10,R,10
L,10,L,12,R,10
R,6,L,12,L,10
y
`

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	r := robot.NewScreen(program)
	r.SetAI(machine.NewNoopAI())
	r.StartShow(programToInput(aiAlgorithm))
	r.Exec()

	return fmt.Sprint(r.GetLastOutput())
}

var (
	Up    = grid.NewPoint(0, -1)
	Down  = grid.NewPoint(0, 1)
	Left  = grid.NewPoint(-1, 0)
	Right = grid.NewPoint(1, 0)
)

var (
	turnLeft  = map[*grid.Point]*grid.Point{Up: Left, Left: Down, Down: Right, Right: Up}
	turnRight = map[*grid.Point]*grid.Point{Up: Right, Right: Down, Down: Left, Left: Up}
)

func yieldPath(g *grid.Grid) []string {
	var pos, dir *grid.Point
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			p := g.Get(x, y)
			switch p.Value {
			case int('^'): //up
				pos = p.Point
				dir = Up
			case int('v'): //down
				pos = p.Point
				dir = Down
			case int('<'): //left
				pos = p.Point
				dir = Left
			case int('>'): //right
				pos = p.Point
				dir = Right
			}
		}
	}

	isScaffold := func(pos *grid.Point) bool {
		return pos.X >= 0 && pos.Y >= 0 && pos.X < g.Width && pos.Y < g.Height && g.Get(pos.X, pos.Y).Value == 35
	}

	// Gather commands to follow the path.
	// Walk straight for as long as possible, then check if we can turn left or right.
	// If we cannot do either, we have reached the end of the path.
	var path []string
	for {
		length := 0
		for isScaffold(pos.Plus(dir)) {
			pos = pos.Plus(dir)
			length++
		}
		if length != 0 {
			path = append(path, strconv.Itoa(length))
		}

		if newDir := turnLeft[dir]; isScaffold(pos.Plus(newDir)) {
			dir = newDir
			path = append(path, "L")
		} else if newDir := turnRight[dir]; isScaffold(pos.Plus(newDir)) {
			dir = newDir
			path = append(path, "R")
		} else {
			break
		}
	}
	return path
}

func programToInput(command string) []int {
	runes := []rune(command)
	input := make([]int, len(runes))
	for i, c := range runes {
		input[i] = int(c)
	}
	return input
}
