package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine/robot"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := input.Load("2019", "11")
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
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	robot := robot.NewPainter(program)

	robot.Exec()

	//fmt.Printf("Robot Path size: %v\n", len(robot.Path))
	//fmt.Printf("RobotPath:%v\n",robot.Path)
	uniquePoints := uniquePointsInPath(robot.Path[:len(robot.Path)-1])
	return strconv.Itoa(len(uniquePoints))
}

func uniquePointsInPath(path []*grid.ValuedPoint) []grid.Point {
	uniquePoints := []grid.Point{}

	//fmt.Printf("Robot Path size: %v\n", len(Path))
	for _, vp := range path {
		contains := !containsPoint(uniquePoints, *vp.Point)
		if contains {
			uniquePoints = append(uniquePoints, *vp.Point)
		}
		//fmt.Printf("%v | %v | newSlice:%v\n", contains, *vp.Point, uniquePoints)
	}
	return uniquePoints
}

func containsPoint(s []grid.Point, e grid.Point) bool {
	for _, a := range s {
		if a.Equals(e) {
			return true
		}
	}
	return false
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	robot := robot.NewPainter(program)
	robot.Path[0] = grid.NewValuedPoint(0, 6, 1)

	robot.Exec()

	g := grid.NewGrid(43, 8)

	for _, vp := range robot.Path {
		g.Get(vp.X, vp.Y).Value = vp.Value
	}

	return "\n" + g.Print(painterFormatter)
}

func painterFormatter(e int) string {
	switch e {
	case 0:
		return "."
	case 1:
		return "#"
	}
	return fmt.Sprintf("%v", e)
}
