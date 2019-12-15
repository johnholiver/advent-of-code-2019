package arcade

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
)

type ArcadeAI struct {
	ballTracker          grid.Point
	paddleTracker        grid.Point
	ballTrackerUpdated   bool
	paddleTrackerUpdated bool
	start                bool
	debugMode            bool
}

func NewArcadeAI() machine.AI {
	return &ArcadeAI{
		*grid.NewPoint(0, 0),
		*grid.NewPoint(0, 0),
		false,
		false,
		true,
		false,
	}
}

func (a *ArcadeAI) SetDebugMode(d bool) {
	a.debugMode = d
}

func (ai *ArcadeAI) GetNextInput() *int {
	if !((ai.start && (ai.ballTrackerUpdated && ai.paddleTrackerUpdated)) || (!ai.start && ai.ballTrackerUpdated)) {
		return nil
	}

	ai.start = false
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

func (ai *ArcadeAI) LastOutput(output []int) {
	x := output[0]
	y := output[1]
	tile := output[2]

	switch tile {
	case 3:
		ai.updatePaddle(*grid.NewPoint(x, y))
	case 4:
		ai.updateBall(*grid.NewPoint(x, y))
	}
}

func (ai *ArcadeAI) updateBall(p grid.Point) {
	ai.ballTracker = p
	ai.ballTrackerUpdated = true
	if ai.debugMode {
		fmt.Println("Ball Updated:", ai.ballTracker, ai.ballTrackerUpdated, ai.paddleTrackerUpdated)
	}
}
func (ai *ArcadeAI) updatePaddle(p grid.Point) {
	ai.paddleTracker = p
	ai.paddleTrackerUpdated = true
	if ai.debugMode {
		fmt.Println("Paddle Updated:", ai.paddleTracker, ai.ballTrackerUpdated, ai.paddleTrackerUpdated)
	}
}
