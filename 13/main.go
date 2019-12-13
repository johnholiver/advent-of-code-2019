package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
)

func main() {
	file, err := input.Load("2019", "13")
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
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	arcade := NewArcade(program)

	blockTiles := 0
	for arcade.p.IsHalted == false {
		_, _, tile, _ := arcade.ExecOneStep(nil)
		if tile == 2 {
			blockTiles++
		}
	}

	return fmt.Sprintf("%v", blockTiles)
}

type Arcade struct {
	originalProgram string
	p               *computer.Processor
	g               *grid.Grid
	Coins           int
	Score           int
}

func NewArcade(program string) *Arcade {
	a := &Arcade{
		program,
		nil,
		grid.NewGrid(36, 21),
		0,
		0,
	}
	a.Reset()
	return a
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

func (a *Arcade) Exec(ai *ArcadeAI) {
	for a.p.IsHalted == false {
		var stepInput *int
		if a.Coins > 0 {
			stepInput = ai.GetNextInput()
		}

		x, y, tile, _ := a.ExecOneStep(stepInput)

		if x == -1 && y == 0 {
			a.Score = tile
		} else {
			switch tile {
			case 3:
				ai.UpdatePaddle(*grid.NewPoint(x, y))
			case 4:
				ai.UpdateBall(*grid.NewPoint(x, y))
			}

			a.g.Get(x, y).Value = tile
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

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Training AI
	arcade := NewArcade(program)
	ai := NewArcadeAI()
	arcade.Exec(ai)

	fmt.Println(arcade.g.Print(arcadeFormatter))
	fmt.Println(ai.ballTracker)
	fmt.Println(ai.paddleTracker)
	//Playing
	arcade.PutCoin()
	arcade.Exec(ai)

	fmt.Println(arcade.g.Print(arcadeFormatter))

	return fmt.Sprintf("%v", arcade.Score)
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

type ArcadeAI struct {
	ballTracker          grid.Point
	paddleTracker        grid.Point
	ballTrackerUpdated   bool
	paddleTrackerUpdated bool
}

func NewArcadeAI() *ArcadeAI {
	return &ArcadeAI{
		*grid.NewPoint(0, 0),
		*grid.NewPoint(0, 0),
		false,
		false,
	}
}

func (ai *ArcadeAI) UpdateBall(p grid.Point) {
	ai.ballTracker = p
	ai.ballTrackerUpdated = true
}
func (ai *ArcadeAI) UpdatePaddle(p grid.Point) {
	ai.paddleTracker = p
	ai.paddleTrackerUpdated = true
}
func (ai *ArcadeAI) GetNextInput() *int {
	if !ai.ballTrackerUpdated && !ai.paddleTrackerUpdated {
		return nil
	}

	ai.ballTrackerUpdated = false
	ai.paddleTrackerUpdated = false

	input := 0
	if ai.ballTracker.X > ai.paddleTracker.X {
		input = 1
	}
	if ai.ballTracker.X < ai.paddleTracker.X {
		input = -1
	}

	return &input
}
