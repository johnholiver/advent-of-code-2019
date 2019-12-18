package main

import (
	"bufio"
	"fmt"
	astar "github.com/beefsack/go-astar"
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

	k, d := fetchKeys(buildGrid(input))

	return fmt.Sprint(k, d)
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
	h := 1
	w := 0
	for _, c := range input {
		if unicode.IsSpace(c) {
			h++
		}
		if h == 1 {
			w++
		}
	}
	g := grid.NewGrid(w, h)
	g.SetFormatter(asciiFormatter)

	x := 0
	y := 0
	for _, c := range input {
		if unicode.IsSpace(c) {
			x = 0
			y++
			continue
		}
		g.Get(x, y).Value = int(c)
		x++
	}
	return g
}

func asciiFormatter(e interface{}) string {
	cast := e.(int)
	return string(rune(cast))
}

func fetchKeys(g *grid.Grid) ([]*grid.ValuedPoint, []*grid.ValuedPoint) {
	var keys []*grid.ValuedPoint
	var doors []*grid.ValuedPoint

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			vp := g.Get(x, y)
			c := rune(vp.Value.(int))
			if unicode.IsLower(c) {
				keys = append(keys, vp)
			}
			if unicode.IsUpper(c) {
				doors = append(doors, vp)
			}

		}
	}
	return keys, doors
}
