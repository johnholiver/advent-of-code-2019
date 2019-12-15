package painter

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
	"strings"
)

type Painter struct {
	DirectionFacing     string
	Path                []*grid.ValuedPoint
	p                   *computer.Processor
	overrideProcessFunc machine.ProcessingStepFunc
	debugMode           bool
}

func NewPainter(program string) *Painter {
	buildComputer := func(program string) *computer.Processor {
		i := computer_io.NewTape()
		p := computer.NewProcessor(i, nil, nil)
		m := computer_mem.NewRelative(p, program)
		p.Memory = m
		o := computer_io.NewInterruptingTape(p)
		p.Output = o
		return p
	}

	return &Painter{
		"U",
		[]*grid.ValuedPoint{
			grid.NewValuedPoint(0, 0, 0),
		},
		buildComputer(program),
		nil,
		false,
	}
}

func (r *Painter) SetDebugMode(d bool) {
	r.debugMode = d
}

func (r *Painter) Exec() {
	cnt := 0
	for r.p.IsHalted == false {
		if r.debugMode {
			fmt.Printf("%v%v: ", r.Path[len(r.Path)-1].Point, r.DirectionFacing)
		}
		r.ExecOneStep()
		cnt++
		if r.debugMode {
			fmt.Printf(" %v -> %v%v | %v\n",
				r.Path[len(r.Path)-2],
				r.Path[len(r.Path)-1].Point, r.DirectionFacing,
				r.Path)
		}
	}
	if r.debugMode {
		fmt.Println("Robot executed", cnt, "steps")
		fmt.Printf("Input (Tape: %+v)\n", r.p.Input.(*computer_io.Tape))
		fmt.Printf("Output (Tape: %+v)\n", r.p.Output.(*computer_io.InterruptingTape))
	}
}

func (r *Painter) ExecOneStep() {
	lastPosition := r.Path[len(r.Path)-1]
	nextColor := lastPosition.Value

	if r.debugMode {
		fmt.Printf("%v#", nextColor)
	}

	processingFunc := r.processOne
	if r.overrideProcessFunc != nil {
		processingFunc = r.overrideProcessFunc
	}

	output, done := processingFunc(&nextColor)
	color := output[0]
	dir := output[1]
	if done {
		return
	}

	if r.debugMode {
		fmt.Printf("%v,%v |", color, dir)
	}

	r.Paint(color)
	r.Move(dir)
}

func (r *Painter) processOne(input *int) ([]int, bool) {
	output := make([]int, 2)
	r.p.Input.Append(*input)
	r.p.Process()
	if r.p.IsHalted {
		//Emergency break :D
		return output, true
	}

	output[0] = r.p.Output.Read()
	r.p.Process()
	output[1] = r.p.Output.Read()
	return output, false
}

func (r *Painter) Paint(color int) {
	lastPosition := r.Path[len(r.Path)-1]
	lastPosition.Value = color
}

func (r *Painter) Move(dir int) *grid.ValuedPoint {
	lastPosition := r.Path[len(r.Path)-1]

	r.DirectionFacing = transformDir(r.DirectionFacing, dir)

	newPoint := grid.NewPoint(lastPosition.X, lastPosition.Y)
	walker := grid.NewWalker(newPoint, grid.NewVector(r.DirectionFacing, 1))
	walker.WalkOne()

	color := r.ColorOfPoint(*newPoint)
	newVP := grid.NewValuedPoint(newPoint.X, newPoint.Y, color)
	r.Path = append(r.Path, newVP)
	return newVP
}

func (r *Painter) ColorOfPoint(newPoint grid.Point) int {
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
