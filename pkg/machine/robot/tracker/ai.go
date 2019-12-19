package tracker

import (
	"fmt"

	"github.com/johnholiver/advent-of-code-2019/pkg/graph"
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
)

type MapperAI struct {
	debugMode bool
	Map       *grid.Grid
	botPos    *grid.Point
	lastInput int
	lastMove  int
	moved     bool

	pathGraph    *graph.Graph
	Steps        int
	HighestSteps int
	totalSteps   int
}

func NewMapperAI() machine.AI {
	mapCell := grid.NewGrid(41, 41)
	pathGraph := graph.NewGraph()
	botPos := grid.NewPoint(21, 19)
	mapCell.Get(botPos.X, botPos.Y).Value = 1
	pathGraph.BuildVector(botPos.String(), nil)

	ai := &MapperAI{
		false,
		mapCell,
		botPos,
		0,
		0,
		false,
		pathGraph,
		0,
		0,
		0,
	}
	mapCell.SetFormatter(ai.mapperFormatter)
	//pathGraph.SetFormatter(ai.pointFormatter)

	return ai
}

func (ai *MapperAI) SetDebugMode(d bool) {
	ai.debugMode = d
}

func (ai *MapperAI) mapperFormatter(e interface{}) string {
	cast := e.(int)
	switch cast {
	case 0:
		return "?"
	case 1:
		return "."
	case 2:
		return "X"
	case 3:
		return "#"
	case 4:
		return "o"
	}
	return "e"
}

func (ai *MapperAI) pointFormatter(point interface{}) string {
	p := point.(*grid.Point)
	return p.String()
}

func contraryDirection(dir int) int {
	switch dir {
	case 1:
		return 2
	case 2:
		return 1
	case 3:
		return 4
	case 4:
		return 3
	}
	return 0
}

func (ai *MapperAI) wouldMoveTo() *grid.Point {
	direction := ""
	switch ai.lastInput {
	case 1:
		direction = "U"
	case 2:
		direction = "D"
	case 3:
		direction = "L"
	case 4:
		direction = "R"
	}

	newPoint := grid.NewPoint(ai.botPos.X, ai.botPos.Y)
	walker := grid.NewWalker(newPoint, grid.NewVector(direction, 1))
	walker.WalkOne()

	return newPoint
}

func (ai *MapperAI) directionFromTo(from, to *grid.Point) int {
	if from.Y == to.Y+1 {
		return 1 //up
	}
	if from.Y == to.Y-1 {
		return 2 //down
	}
	if from.X == to.X-1 {
		return 3 //left
	}
	if from.X == to.X+1 {
		return 4 //right
	}

	panic("From -> To are not adjacent")
}

//3nd try (similar to 1st, but "keeping left hand on the wall of the maze")
func (ai *MapperAI) GetNextInput() *int {
	if ai.moved {
		ai.moved = false
		ai.lastInput = contraryDirection(ai.lastInput)
	}

	var nextFromLast func(dir int) int
	nextFromLast = nextFromLastClockwise

	ai.lastInput = nextFromLast(ai.lastInput)

	return &ai.lastInput
}

func nextFromLastClockwise(dir int) int {
	switch dir {
	case 0:
		return 1
	case 1:
		return 4
	case 2:
		return 3
	case 3:
		return 1
	case 4:
		return 2
	case 5:
		return 1
	}
	panic("WTF")
}

func nextFromLastCounterClockwise(dir int) int {
	switch dir {
	case 0:
		return 1
	case 1:
		return 3
	case 2:
		return 4
	case 3:
		return 2
	case 4:
		return 1
	case 5:
		return 1
	}
	panic("WTF")
}

func (ai *MapperAI) LastOutput(output []int) {
	nextPos := ai.wouldMoveTo()
	status := output[0]

	mapCell := ai.Map.Get(nextPos.X, nextPos.Y)
	mapCell.Value = status
	if status == 0 {
		mapCell.Value = 3
	}

	switch status {
	case 1:
		if ai.pathGraph.FindNode(nextPos.String()) == nil {
			//New node
			ai.pathGraph.BuildVector(nextPos.String(), ai.botPos.String())
			ai.Steps++
			if ai.Steps > ai.HighestSteps {
				ai.HighestSteps = ai.Steps
			}
		} else {
			robotCell := ai.Map.Get(ai.botPos.X, ai.botPos.Y)
			robotCell.Value = 4
			//backtrack
			ai.Steps--
		}

		ai.totalSteps++
		ai.lastMove = ai.lastInput
		ai.botPos = nextPos
		ai.moved = true
	}

	if ai.debugMode {
		fmt.Print(ai.Map)
		fmt.Printf("Steps: [%v][%v][%v]\n", ai.Steps, ai.HighestSteps, ai.totalSteps)
		if status == 2 {
			fmt.Printf("X: %v %v/%v\n", nextPos, ai.Steps, ai.HighestSteps)
		}
	}

}

//===== Idea Graveyard ===

/*
 *	1nd try (very manual!)
 *	I was trying to run non programmatically, I thought the puzzle wouldn't be so big... eventually I realized that this
 *	would take more time than building some heuristic
 */

//func (ai *MapperAI) GetNextInput_1try() *int {
//	if ai.moved {
//		ai.moved = false
//		ai.lastInput = 0
//	}
//
//	fmt.Println("AI Step:",ai.Steps)
//	var nextFromLast func(dir int) int
//	switch {
//	case ai.Steps < 70:
//		nextFromLast = nextFromLastDownRight
//	case ai.Steps >= 70 && ai.Steps < 110:
//		nextFromLast = nextFromLastUpLeft
//	case ai.Steps >= 110 && ai.Steps < 180:
//		nextFromLast = nextFromLastLeftUp
//	case ai.Steps >= 170:
//		nextFromLast = nextFromLastDownRight
//	}
//
//
//
//	ai.lastInput = nextFromLast(ai.lastInput)
//	avoid := contraryDirection(ai.lastMove)
//	if ai.lastInput == avoid {
//		ai.lastInput = nextFromLast(ai.lastInput)
//	}
//	//mod
//	if ai.lastInput == 5 {
//		ai.lastInput = nextFromLast(ai.lastInput)
//	}
//
//	return &ai.lastInput
//}
//
//func nextFromLastRightDown (dir int) int{
//	switch dir {
//	case 0:
//		return 4
//	case 1:
//		return 4
//	case 2:
//		return 1
//	case 3:
//		return 2
//	case 4:
//		return 3
//	case 5:
//		return 4
//	}
//	panic("WTF")
//}
//
//func nextFromLastLeftUp (dir int) int{
//	switch dir {
//	case 0:
//		return 3
//	case 1:
//		return 2
//	case 2:
//		return 3
//	case 3:
//		return 4
//	case 4:
//		return 1
//	case 5:
//		return 3
//	}
//	panic("WTF")
//}
//
//func nextFromLastUpLeft (dir int) int{
//	switch dir {
//	case 0:
//		return dir+1
//	case 1:
//		return dir+1
//	case 2:
//		return dir+1
//	case 3:
//		return dir+1
//	case 4:
//		return dir+1
//	case 5:
//		return 1
//	}
//	panic("WTF")
//}
//
//func nextFromLastDownRight (dir int) int{
//	switch dir {
//	case 0:
//		return 2
//	case 1:
//		return 4
//	case 2:
//		return 1
//	case 3:
//		return 2
//	case 4:
//		return 3
//	case 5:
//		return 2
//	}
//	panic("WTF")
//}

/*
 *	2nd try (incomplete)
 *	My initial idea is that I will build a tree with the possible paths, then I can traverse the graph from the target
 *  node til the root. However, the fact that the robot directly move with the input made "looking" if there is a path
 *  much harder than necessary. I ended up dropping this idea in the middle in favor of the maze walk with hand in the
 * 	wall.
 */

//func (ai *MapperAI) GetNextInput_2try() *int {
//	if ai.mode == Scanning {
//		ai.lastInput = ai.scanningNextInput()
//	}
//	if ai.mode ==  MovingForward {
//		ai.lastInput = 1
//	}
//
//	return &ai.lastInput
//}
//
//
//var lastScan = 0
//func (ai *MapperAI) scanningNextInput() int {
//	//if ai.moved {
//	//	// Add to tree path
//	//	ai.moved = false
//	//	return contraryDirection(ai.revertMoveDuringScan)
//	//}
//
//	lastScan++
//	if lastScan == contraryDirection(ai.lastMove) {
//		lastScan++
//	}
//
//	if lastScan == 5 {
//		lastScan = 0
//		ai.mode = MovingForward
//	}
//
//	return lastScan
//}
