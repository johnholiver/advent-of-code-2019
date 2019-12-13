package arcade

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
)

type Arcade struct {
	originalProgram string
	p               *computer.Processor
	g               *grid.Grid
	Coins           int
	Score           int
	debugMode       bool
}

func New(program string) *Arcade {
	a := &Arcade{
		program,
		nil,
		grid.NewGrid(36, 21),
		0,
		0,
		false,
	}
	a.Reset()
	return a
}

func (a *Arcade) SetDebugMode(d bool) {
	a.debugMode = d
}

func (a *Arcade) Reset() {
	buildComputer := func(program string) *computer.Processor {
		i := computer_io.NewTape()
		p := computer.NewProcessor(i, nil, nil)
		m := computer_mem.NewRelative(p, program)
		p.Memory = m
		o := computer_io.NewInterruptingTape(p)
		p.Output = o
		return p
	}

	a.p = buildComputer(a.originalProgram)
}

func (a *Arcade) Exec(ai Player) {
	for a.p.IsHalted == false {
		var stepInput *int
		if a.Coins > 0 {
			stepInput = ai.Play()
			if a.debugMode && stepInput != nil {
				fmt.Println("Input Given:", *stepInput)
			}
		}

		x, y, tile, _ := a.ExecOneStep(stepInput)

		if x == -1 && y == 0 {
			a.Score = tile
		} else {
			a.g.Get(x, y).Value = tile

			switch tile {
			case 3:
				ai.UpdatePaddle(*grid.NewPoint(x, y))
				if a.debugMode {
					fmt.Print(a.g.Print(arcadeFormatter))
					fmt.Println("Score:", a.Score)
				}
			case 4:
				ai.UpdateBall(*grid.NewPoint(x, y))
				if a.debugMode {
					fmt.Print(a.g.Print(arcadeFormatter))
					fmt.Println("Score:", a.Score)
				}
			}
		}
	}
}

func (r *Arcade) ExecOneStep(stepInput *int) (int, int, int, bool) {
	if stepInput != nil {
		r.p.Input.Append(*stepInput)
	}

	r.p.Process()
	if r.p.IsHalted {
		//Emergency break :D
		return 0, 0, 0, true
	}

	x := r.p.Output.Read()
	r.p.Process()
	y := r.p.Output.Read()
	r.p.Process()
	tile := r.p.Output.Read()
	return x, y, tile, false
}

func (a *Arcade) PutCoin() {
	a.Reset()
	a.Coins++
	a.p.Memory.(*computer_mem.RelativeMemory).Variables[0] = 2
}

func arcadeFormatter(e int) string {
	switch e {
	case 0:
		return "."
	case 1:
		return "#"
	case 2:
		return "X"
	case 3:
		return "="
	case 4:
		return "o"
	}
	return "e"
}
