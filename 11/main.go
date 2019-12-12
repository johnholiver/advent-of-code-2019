package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"log"
	"os"
	"strconv"
	"strings"
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

	robot := NewRobot(program)

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

type Robot struct {
	DirectionFacing string
	Path            []*grid.ValuedPoint
	p               *computer.Processor
}

func NewRobot(program string) *Robot {
	buildComputer := func(program string) *computer.Processor {
		i := computer_io.NewTape()
		p := computer.NewProcessor(i, nil, nil)
		m := computer_mem.NewRelative(p, program)
		p.Memory = m
		o := computer_io.NewInterruptingTape(p)
		p.Output = o
		return p
	}

	return &Robot{
		"U",
		[]*grid.ValuedPoint{
			grid.NewValuedPoint(0, 0, 0),
		},
		buildComputer(program),
	}
}

func (r *Robot) Exec() {
	cnt := 0
	for r.p.IsHalted == false {
		//fmt.Printf("%v%v: ", r.Path[len(r.Path)-1].Point, r.DirectionFacing)
		r.ExecOneStep(r.processOne)
		cnt++
		//fmt.Printf(" %v -> %v%v | %v\n",
		//	r.Path[len(r.Path)-2],
		//	r.Path[len(r.Path)-1].Point, r.DirectionFacing,
		//	r.Path)
	}
	//fmt.Println("Robot executed",cnt,"steps")

	//fmt.Printf("Input (Tape: %+v)\n",r.p.Input.(*computer_io.Tape))
	//fmt.Printf("Output (Tape: %+v)\n",r.p.Output.(*computer_io.InterruptingTape))
}

func (r *Robot) ExecOneStep(processingFunc func(nextColor int) (int, int, bool)) {
	lastPosition := r.Path[len(r.Path)-1]
	nextColor := lastPosition.Value

	//fmt.Printf("%v#", nextColor)

	color, dir, done := processingFunc(nextColor)
	if done {
		return
	}

	//fmt.Printf("%v,%v |", color, dir)

	r.Paint(color)
	r.Move(dir)
}

func (r *Robot) processOne(nextColor int) (int, int, bool) {
	r.p.Input.Append(nextColor)
	r.p.Process()
	if r.p.IsHalted {
		//Emergency break :D
		return 0, 0, true
	}

	color := r.p.Output.Read()
	r.p.Process()
	dir := r.p.Output.Read()
	return color, dir, false
}

func (r *Robot) Paint(color int) {
	lastPosition := r.Path[len(r.Path)-1]
	lastPosition.Value = color
}

func (r *Robot) Move(dir int) int {
	lastPosition := r.Path[len(r.Path)-1]

	r.DirectionFacing = transformDir(r.DirectionFacing, dir)

	newPoint := grid.NewPoint(lastPosition.X, lastPosition.Y)
	walker := grid.NewWalker(newPoint, grid.NewVector(r.DirectionFacing, 1))
	walker.WalkOne()

	color := r.ColorOfPoint(*newPoint)
	r.Path = append(r.Path, grid.NewValuedPoint(newPoint.X, newPoint.Y, color))
	return color
}

func (r *Robot) ColorOfPoint(newPoint grid.Point) int {
	lP := lastPoint(r.Path, newPoint)
	if lP != nil {
		return lP.Value
	}
	return 0
}

func lastPoint(path []*grid.ValuedPoint, nP grid.Point) *grid.ValuedPoint {
	idx := len(path) - 1
	for idx > -1 {
		lP := path[idx]
		if lP.Point.Equals(nP) {
			return lP
		}
		idx--
	}
	return nil
}

func transformDir(currentFacing string, dir int) string {
	facing := "URDL"
	idx := strings.Index(facing, currentFacing)

	switch dir {
	case 0:
		idx = idx - 1
		if idx == -1 {
			idx = len(facing) - 1
		}
	case 1:
		idx = idx + 1
		if idx == len(facing) {
			idx = 0
		}
	}

	return string(facing[idx])
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

	robot := NewRobot(program)
	robot.Path[0] = grid.NewValuedPoint(0, 6, 1)

	robot.Exec()

	g := grid.NewGrid(43, 8)

	for _, vp := range robot.Path {
		g.Get(vp.X, vp.Y).Value = vp.Value
	}

	return "\n" + g.String()
}
