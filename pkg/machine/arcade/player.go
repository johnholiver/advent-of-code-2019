package arcade

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
)

type Player interface {
	machine.AI
	UpdateBall(p grid.Point)
	UpdatePaddle(p grid.Point)
}

type ArcadeAI struct {
	ballTracker          grid.Point
	paddleTracker        grid.Point
	ballTrackerUpdated   bool
	paddleTrackerUpdated bool
	start                bool
	debugMode            bool
}

func NewArcadeAI() *ArcadeAI {
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

func (ai *ArcadeAI) UpdateBall(p grid.Point) {
	ai.ballTracker = p
	ai.ballTrackerUpdated = true
	if ai.debugMode {
		fmt.Println("Ball Updated:", ai.ballTracker, ai.ballTrackerUpdated, ai.paddleTrackerUpdated)
	}
}
func (ai *ArcadeAI) UpdatePaddle(p grid.Point) {
	ai.paddleTracker = p
	ai.paddleTrackerUpdated = true
	if ai.debugMode {
		fmt.Println("Paddle Updated:", ai.paddleTracker, ai.ballTrackerUpdated, ai.paddleTrackerUpdated)
	}
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
