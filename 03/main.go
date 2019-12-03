package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"github.com/juliangruber/go-intersect"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := input.Load("2019", "3")
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
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	smallValue := smallestManhattan(lines[0], lines[1])

	return strconv.Itoa(smallValue)
}

func smallestManhattan(wire1, wire2 string) int {
	w1 := wireToCoords(wire1)
	w2 := wireToCoords(wire2)

	sI := intersect.Hash(w1, w2)
	intersects := sI.([]interface{})
	var recast []grid.Point
	for i, p := range intersects {
		if i == 0 {
			continue
		}
		recast = append(recast, p.(grid.Point))
	}

	_, smallValue := smallestOf(recast, predicateManhattan)
	return smallValue
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	smallValue := smallestSteps(lines[0], lines[1])

	return strconv.Itoa(smallValue)
}

func smallestSteps(wire1, wire2 string) int {
	w1 := wireToCoords(wire1)
	w2 := wireToCoords(wire2)

	sI := intersect.Hash(w1, w2)
	intersects := sI.([]interface{})
	var recast []grid.Point
	for i, p := range intersects {
		if i == 0 {
			continue
		}
		recast = append(recast, p.(grid.Point))
	}

	_, smallValue := smallestOf(recast, predicateStepsBuilder(w1, w2))
	return smallValue
}

func wireToCoords(wire string) []grid.Point {
	ss := strings.Split(wire, ",")
	vectors := make([]*grid.Vector, len(ss))
	for i, v := range ss {
		direction := v[0:1]
		value, _ := strconv.Atoi(v[1:])

		vectors[i] = grid.NewVector(direction, value)
	}

	var wCoords []grid.Point
	origin := grid.NewPoint(0, 0)
	wCoords = append(wCoords, *origin)
	cCoord := origin
	for _, v := range vectors {
		walker := grid.NewWalker(cCoord, v)

		for {
			if walker.Finished() {
				break
			}
			newCoord := walker.WalkOne()
			wCoords = append(wCoords, *grid.NewPoint(newCoord.X, newCoord.Y))
		}
	}
	return wCoords
}

// This resembles functional programming (map/filter functions)
// Code has been refactored to this because of how the puzzle asks very similar things from part 1 & 2, i.e to find the
// smallest result of a given method (Manhattan distance or step count) when applied over multiple elements of an array
func smallestOf(slice []grid.Point, predicate func(grid.Point) int) (p *grid.Point, smallest int) {
	smallest = int(^uint(0) >> 1)
	var pSmallest *grid.Point

	for i, e := range slice {
		if i == 0 {
			continue
		}
		man := predicate(e)
		if man < smallest {
			smallest = man
			pSmallest = &e
		}
	}

	return pSmallest, smallest
}

func predicateManhattan(p grid.Point) int {
	return grid.Manhattan(grid.Point{0, 0}, p)
}

func predicateStepsBuilder(w1, w2 []grid.Point) func(p grid.Point) int {
	return func(p grid.Point) int {
		return indexOf(p, w1) + indexOf(p, w2)
	}
}

func indexOf(element grid.Point, data []grid.Point) int {
	for k, v := range data {
		if element.Equals(v) {
			return k
		}
	}
	return -1 //not found.
}
