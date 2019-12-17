package screen

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
)

type Programmer struct {
	debugMode bool
	Map       *grid.Grid
	botPos    *grid.Point

	program []int
	pc      int
}

func NewProgrammerAI(program string) machine.AI {
	mapCell := grid.NewGrid(49, 47)

	ai := &Programmer{
		false,
		mapCell,
		grid.NewPoint(0, 0),
		nil,
		-1,
	}
	ai.program = ai.programToInput(program)
	mapCell.SetFormatter(ai.asciiFormatter)

	return ai
}

func (ai *Programmer) SetDebugMode(d bool) {
	ai.debugMode = d
}

func (ai *Programmer) asciiFormatter(e int) string {
	return string(rune(e))
}

func (ai *Programmer) programToInput(command string) []int {
	runes := []rune(command)
	input := make([]int, len(runes))
	for i, c := range runes {
		input[i] = int(c)
	}
	return input
}

func (ai *Programmer) GetNextInput() *int {
	if ai.pc >= len(ai.program)-1 {
		return nil
	}
	ai.pc++
	return &ai.program[ai.pc]
}

func (ai *Programmer) LastOutput(output []int) {
	cellValue := output[0]
	if ai.debugMode {
		fmt.Println("Last output:", ai.botPos, cellValue)
	}
	if cellValue == 10 {
		ai.botPos.X = -1
		ai.botPos.Y++
	} else {
		ai.Map.Get(ai.botPos.X, ai.botPos.Y).Value = cellValue
	}
	ai.botPos.X++

	if ai.debugMode {
		fmt.Println(ai.Map.Print())
	}
}
