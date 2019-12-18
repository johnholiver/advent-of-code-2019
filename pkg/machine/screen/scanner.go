package screen

import (
	"fmt"

	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
)

type ScannerAI struct {
	debugMode bool
	Map       *grid.Grid
	botPos    *grid.Point
	X         int
}

func NewScannerAI() machine.AI {
	mapCell := grid.NewGrid(49, 47)

	ai := &ScannerAI{
		false,
		mapCell,
		grid.NewPoint(0, 0),
		-1,
	}
	mapCell.SetFormatter(ai.asciiFormatter)

	return ai
}

func (ScannerAI) asciiFormatter(e interface{}) string {
	cast := e.(int)
	return string(rune(cast))
}

func (ai *ScannerAI) SetDebugMode(d bool) {
	ai.debugMode = d
}

func (ai *ScannerAI) GetNextInput() *int {
	return nil
}

func (ai *ScannerAI) LastOutput(output []int) {
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
