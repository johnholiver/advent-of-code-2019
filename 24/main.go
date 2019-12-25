package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"github.com/johnholiver/advent-of-code-2019/pkg/life"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := input.Load("2019", "24")
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
	lines := ""
	for scanner.Scan() {
		lines += fmt.Sprintln(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	w := buildWorld(lines, life.NewMonoverseCell)

	worldStates := make(map[string]*life.World)
	for {
		wKey := w.String()
		if _, found := worldStates[wKey]; found {
			break
		}
		worldStates[wKey] = w

		w = w.Tick()
	}

	return fmt.Sprint(w.BiodiversityRating())
}

func buildWorld(input string, cellBuilder func(bool, int, int, *life.World) life.Cell) *life.World {
	lines := strings.Split(input, "\n")
	w := life.NewWorld(5, 5, 0)
	for y := 0; y < len(lines)-1; y++ {
		line := lines[y]
		for x := 0; x < len(line); x++ {
			var hasBug bool
			switch line[x] {
			case '.':
				hasBug = false
			case '#':
				hasBug = true
			}
			c := cellBuilder(hasBug, x, y, w)
			w.SetCell(x, y, c)
		}
	}
	return w
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	lines := ""
	for scanner.Scan() {
		lines += fmt.Sprintln(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	w := buildWorld(lines, life.NewMultiverseCell)

	mv := life.NewMultiverse(w)

	for tick := 0; tick < 200; tick++ {
		mv.Tick()
	}

	return fmt.Sprint(mv.CountBugs())
}
