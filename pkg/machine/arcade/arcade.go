package arcade

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
)

type Arcade struct {
	originalProgram string
	p               *computer.Processor
	g               *grid.Grid
	player          machine.AI
	Coins           int
	Score           int
	debugMode       bool
}

func New(program string) *Arcade {
	a := &Arcade{
		program,
		nil,
		grid.NewGrid(36, 21).SetFormatter(arcadeFormatter),
		nil,
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

func (a *Arcade) Exec() {
	for a.p.IsHalted == false {
		a.ExecOneStep()
	}
}

func (a *Arcade) ExecOneStep() {
	var stepInput *int
	if a.Coins > 0 {
		stepInput = a.player.GetNextInput()
		if a.debugMode && stepInput != nil {
			fmt.Println("Input Given:", *stepInput)
		}
	}

	output, done := a.ProcessOne(stepInput)
	x := output[0]
	y := output[1]
	tile := output[2]
	if done {
		return
	}

	if x == -1 && y == 0 {
		a.Score = tile
	} else {
		a.g.Get(x, y).Value = tile
		a.player.LastOutput(output)
		if a.debugMode {
			switch tile {
			case 3:
				fmt.Print(a.g.Print())
				fmt.Println("Score:", a.Score)
			case 4:
				fmt.Print(a.g.Print())
				fmt.Println("Score:", a.Score)
			}
		}
	}
}

func (a *Arcade) ProcessOne(input *int) ([]int, bool) {
	output := make([]int, 3)
	if input != nil {
		a.p.Input.Append(*input)
	}

	a.p.Process()
	if a.p.IsHalted {
		//Emergency break :D
		return output, true
	}

	output[0] = a.p.Output.Read()
	a.p.Process()
	output[1] = a.p.Output.Read()
	a.p.Process()
	output[2] = a.p.Output.Read()
	return output, false
}

func (a *Arcade) PutCoin(p machine.AI) {
	a.Reset()
	a.Coins++
	a.player = p
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
