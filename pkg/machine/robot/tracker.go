package robot

import (
	"fmt"

	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
)

type Tracker struct {
	p         *computer.Processor
	Path      []*grid.ValuedPoint
	ai        machine.AI
	debugMode bool
}

func NewTracker(program string) *Tracker {
	buildComputer := func(program string) *computer.Processor {
		i := computer_io.NewTape()
		p := computer.NewProcessor(i, nil, nil)
		m := computer_mem.NewRelative(p, program)
		p.Memory = m
		o := computer_io.NewInterruptingTape(p)
		p.Output = o
		return p
	}

	return &Tracker{
		buildComputer(program),
		[]*grid.ValuedPoint{
			grid.NewValuedPoint(0, 0, 0),
		},
		nil,
		false,
	}
}

func (r *Tracker) SetDebugMode(d bool) {
	r.debugMode = d
}

func (r *Tracker) Exec() {
	cnt := 0
	for r.p.IsHalted == false {
		if r.debugMode {
			fmt.Printf("%v: ", r.Path[len(r.Path)-1].Point)
		}
		r.ExecOneStep()
		cnt++
		if r.debugMode {
			fmt.Printf(" %v -> %v | %v\n",
				r.Path[len(r.Path)-2],
				r.Path[len(r.Path)-1].Point,
				r.Path)
		}
	}
	if r.debugMode {
		fmt.Println("Robot executed", cnt, "steps")
		fmt.Printf("Input (Tape: %+v)\n", r.p.Input.(*computer_io.Tape))
		fmt.Printf("Output (Tape: %+v)\n", r.p.Output.(*computer_io.InterruptingTape))
	}
}

func (r *Tracker) ExecOneStep() {
	//HEAVY LOGIC
	panic("implement me")
}

func (r *Tracker) processOne(input *int) ([]int, bool) {
	output := make([]int, 1)
	if input != nil {
		r.p.Input.Append(*input)
	}

	r.p.Process()
	if r.p.IsHalted {
		//Emergency break :D
		return output, true
	}

	output[0] = r.p.Output.Read()
	return output, false
}

func (r *Tracker) Move(dir int) *grid.ValuedPoint {
	lastPosition := r.Path[len(r.Path)-1]

	direction := ""
	switch dir {
	case 1:
		direction = "U"
	case 2:
		direction = "D"
	case 3:
		direction = "L"
	case 4:
		direction = "R"
	}

	newPoint := grid.NewPoint(lastPosition.X, lastPosition.Y)
	walker := grid.NewWalker(newPoint, grid.NewVector(direction, 1))
	walker.WalkOne()

	newVP := grid.NewValuedPoint(newPoint.X, newPoint.Y, 0)
	r.Path = append(r.Path, newVP)
	return newVP
}
