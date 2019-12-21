package screen

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
)

type Screen struct {
	p                   *computer.Processor
	ai                  machine.AI
	overrideProcessFunc machine.ProcessingStepFunc
	debugMode           bool
	lastOutput          int
}

func NewScreen(program string) *Screen {
	buildComputer := func(program string) *computer.Processor {
		i := computer_io.NewTape()
		p := computer.NewProcessor(i, nil, nil)
		m := computer_mem.NewRelative(p, program)
		p.Memory = m
		o := computer_io.NewInterruptingTape(p)
		p.Output = o
		return p
	}

	return &Screen{
		buildComputer(program),
		nil,
		nil,
		false,
		-1,
	}
}

func (r *Screen) SetDebugMode(d bool) {
	r.debugMode = d
}

func (r *Screen) SetAI(ai machine.AI) {
	r.ai = ai
}

func (r *Screen) Exec() {
	cnt := 0
	for r.p.IsHalted == false {
		r.ExecOneStep()
		cnt++
	}
}

func (r *Screen) ExecOneStep() {
	var stepInput *int

	if r.ai == nil {
		if r.debugMode {
			fmt.Printf("Robot must have an AI")
		}
		return
	}

	stepInput = r.ai.GetNextInput()
	if r.debugMode && stepInput != nil {
		fmt.Printf("Step: %v => ", *stepInput)
	}

	processingFunc := r.ProcessOne
	if r.overrideProcessFunc != nil {
		processingFunc = r.overrideProcessFunc
	}

	output, done := processingFunc(stepInput)
	if done {
		if r.debugMode {
			fmt.Println("done")
		}
		return
	}

	r.ai.LastOutput(output)
	r.lastOutput = output[0]
}

func (r *Screen) ProcessOne(input *int) ([]int, bool) {
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

func (r *Screen) StartShow(input []int) {
	if len(input) > 0 {
		for _, in := range input {
			r.p.Input.Append(in)
		}
		r.p.Memory.(*computer_mem.RelativeMemory).Variables[0] = 2
	} else {
		r.p.Memory.(*computer_mem.RelativeMemory).Variables[0] = 1
	}
}

func (r *Screen) GetLastOutput() int {
	return r.lastOutput
}
