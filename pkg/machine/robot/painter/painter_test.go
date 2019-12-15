package painter

import (
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine"
	"github.com/stretchr/testify/assert"
	"testing"
)

var painterProgram = `3,8,1005,8,320,1106,0,11,0,0,0,104,1,104,0,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,1,10,4,10,102,1,8,29,2,1005,1,10,1006,0,11,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,102,1,8,57,1,8,15,10,1006,0,79,1,6,3,10,3,8,102,-1,8,10,101,1,10,10,4,10,108,0,8,10,4,10,101,0,8,90,2,103,18,10,1006,0,3,2,105,14,10,3,8,102,-1,8,10,1001,10,1,10,4,10,108,0,8,10,4,10,101,0,8,123,2,9,2,10,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,1001,8,0,150,1,2,2,10,2,1009,6,10,1,1006,12,10,1006,0,81,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,102,1,8,187,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,0,10,4,10,101,0,8,209,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,231,1,1008,11,10,1,1001,4,10,2,1104,18,10,3,8,102,-1,8,10,1001,10,1,10,4,10,108,1,8,10,4,10,1001,8,0,264,1,8,14,10,1006,0,36,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,0,8,10,4,10,101,0,8,293,1006,0,80,1006,0,68,101,1,9,9,1007,9,960,10,1005,10,15,99,109,642,104,0,104,1,21102,1,846914232732,1,21102,1,337,0,1105,1,441,21102,1,387512115980,1,21101,348,0,0,1106,0,441,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21102,209533824219,1,1,21102,1,395,0,1106,0,441,21101,0,21477985303,1,21102,406,1,0,1106,0,441,3,10,104,0,104,0,3,10,104,0,104,0,21101,868494234468,0,1,21101,429,0,0,1106,0,441,21102,838429471080,1,1,21102,1,440,0,1106,0,441,99,109,2,21201,-1,0,1,21101,0,40,2,21102,472,1,3,21101,0,462,0,1106,0,505,109,-2,2106,0,0,0,1,0,0,1,109,2,3,10,204,-1,1001,467,468,483,4,0,1001,467,1,467,108,4,467,10,1006,10,499,1102,1,0,467,109,-2,2106,0,0,0,109,4,2101,0,-1,504,1207,-3,0,10,1006,10,522,21101,0,0,-3,21202,-3,1,1,22101,0,-2,2,21102,1,1,3,21102,541,1,0,1106,0,546,109,-4,2105,1,0,109,5,1207,-3,1,10,1006,10,569,2207,-4,-2,10,1006,10,569,22102,1,-4,-4,1105,1,637,22102,1,-4,1,21201,-3,-1,2,21202,-2,2,3,21102,588,1,0,1105,1,546,22101,0,1,-4,21102,1,1,-1,2207,-4,-2,10,1006,10,607,21101,0,0,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,629,21201,-1,0,1,21102,629,1,0,105,1,504,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2105,1,0`

func Test_Paint(t *testing.T) {
	robot := NewPainter(painterProgram)

	color := 0
	robot.Paint(color)
	assert.Equal(t, color, robot.Path[len(robot.Path)-1].Value)

	color = 1
	robot.Paint(color)
	assert.Equal(t, color, robot.Path[len(robot.Path)-1].Value)
}

func Test_Move(t *testing.T) {
	robot := NewPainter(painterProgram)

	var newPos *grid.ValuedPoint
	dir := 0
	oP := robot.Path[len(robot.Path)-1].Point
	newPos = robot.Move(dir)
	assert.Equal(t, oP.X-1, newPos.Point.X)
	assert.Equal(t, oP.Y, newPos.Point.Y)
	newPos = robot.Move(dir)
	assert.Equal(t, oP.X-1, newPos.Point.X)
	assert.Equal(t, oP.Y-1, newPos.Point.Y)
	newPos = robot.Move(dir)
	assert.Equal(t, oP.X, newPos.Point.X)
	assert.Equal(t, oP.Y-1, newPos.Point.Y)
	newPos = robot.Move(dir)
	assert.Equal(t, oP.X, newPos.Point.X)
	assert.Equal(t, oP.Y, newPos.Point.Y)

	assert.Equal(t, oP, newPos.Point)

	dir = 1
	newPos = robot.Move(dir)
	assert.Equal(t, oP.X+1, newPos.Point.X)
	assert.Equal(t, oP.Y, newPos.Point.Y)
	newPos = robot.Move(dir)
	assert.Equal(t, oP.X+1, newPos.Point.X)
	assert.Equal(t, oP.Y-1, newPos.Point.Y)
	newPos = robot.Move(dir)
	assert.Equal(t, oP.X, newPos.Point.X)
	assert.Equal(t, oP.Y-1, newPos.Point.Y)
	newPos = robot.Move(dir)
	assert.Equal(t, oP.X, newPos.Point.X)
	assert.Equal(t, oP.Y, newPos.Point.Y)

	assert.Equal(t, oP, newPos.Point)
}

func Test_transformDir(t *testing.T) {
	assert.Equal(t, "L", transformDir("U", 0))
	assert.Equal(t, "D", transformDir("L", 0))
	assert.Equal(t, "R", transformDir("D", 0))
	assert.Equal(t, "U", transformDir("R", 0))

	assert.Equal(t, "R", transformDir("U", 1))
	assert.Equal(t, "D", transformDir("R", 1))
	assert.Equal(t, "L", transformDir("D", 1))
	assert.Equal(t, "U", transformDir("L", 1))
}

func Test_lastPoint(t *testing.T) {
	var path []*grid.ValuedPoint

	path = []*grid.ValuedPoint{
		grid.NewValuedPoint(0, 0, 0), //0>1>0
		grid.NewValuedPoint(0, 1, 0), //0>0>0
		grid.NewValuedPoint(1, 1, 0), //0>1>1
		grid.NewValuedPoint(1, 0, 0), //0>0>1
		grid.NewValuedPoint(0, 0, 1),
		grid.NewValuedPoint(0, 1, 0),
		grid.NewValuedPoint(1, 1, 1),
		grid.NewValuedPoint(1, 0, 0),
		grid.NewValuedPoint(0, 0, 0),
		grid.NewValuedPoint(0, 1, 0),
		grid.NewValuedPoint(1, 1, 1),
		grid.NewValuedPoint(1, 0, 1),
	}

	assert.Nil(t, lastPoint(path, *grid.NewPoint(2, 2)))
	assert.Equal(t, path[8], lastPoint(path, *grid.NewPoint(0, 0)))
	assert.Equal(t, path[9], lastPoint(path, *grid.NewPoint(0, 1)))
	assert.Equal(t, path[10], lastPoint(path, *grid.NewPoint(1, 1)))
	assert.Equal(t, path[11], lastPoint(path, *grid.NewPoint(1, 0)))

}

func mockProcessOne(color, dir int) machine.ProcessingStepFunc {
	return func(*int) ([]int, bool) {
		output := make([]int, 2)
		output[0] = color
		output[1] = dir
		return output, false
	}
}

func Test_ExecOneStepPaint(t *testing.T) {
	robot := NewPainter(painterProgram)

	colorInput := 0
	dirInput := 0
	robot.overrideProcessFunc = mockProcessOne(colorInput, dirInput)
	robot.ExecOneStep()
	lP := robot.Path[len(robot.Path)-2]
	cP := robot.Path[len(robot.Path)-1]

	assert.Equal(t, colorInput, lP.Value)
	assert.Equal(t, 0, cP.Value)

	colorInput = 1
	dirInput = 0
	robot.overrideProcessFunc = mockProcessOne(colorInput, dirInput)
	robot.ExecOneStep()
	lP = robot.Path[len(robot.Path)-2]
	cP = robot.Path[len(robot.Path)-1]

	assert.Equal(t, colorInput, lP.Value)
	assert.Equal(t, 0, cP.Value)
}

func Test_ExecOneStepMove(t *testing.T) {
	robot := NewPainter(painterProgram)

	assert.Equal(t, "U", robot.DirectionFacing)

	//Turn left
	colorInput := 0
	dirInput := 0
	robot.overrideProcessFunc = mockProcessOne(colorInput, dirInput)
	robot.ExecOneStep()
	assert.Equal(t, "L", robot.DirectionFacing)
	lP := robot.Path[len(robot.Path)-2]
	cP := robot.Path[len(robot.Path)-1]

	assert.Equal(t, lP.Point.X-1, cP.Point.X)
	assert.Equal(t, lP.Point.Y, cP.Point.Y)

	//Turn left
	colorInput = 0
	dirInput = 0
	robot.overrideProcessFunc = mockProcessOne(colorInput, dirInput)
	robot.ExecOneStep()
	assert.Equal(t, "D", robot.DirectionFacing)
	lP = robot.Path[len(robot.Path)-2]
	cP = robot.Path[len(robot.Path)-1]

	assert.Equal(t, lP.Point.X, cP.Point.X)
	assert.Equal(t, lP.Point.Y-1, cP.Point.Y)

	//Turn right
	colorInput = 0
	dirInput = 1
	robot.overrideProcessFunc = mockProcessOne(colorInput, dirInput)
	robot.ExecOneStep()
	assert.Equal(t, "L", robot.DirectionFacing)
	lP = robot.Path[len(robot.Path)-2]
	cP = robot.Path[len(robot.Path)-1]

	assert.Equal(t, lP.Point.X-1, cP.Point.X)
	assert.Equal(t, lP.Point.Y, cP.Point.Y)

	//Turn right
	colorInput = 0
	dirInput = 1
	robot.overrideProcessFunc = mockProcessOne(colorInput, dirInput)
	robot.ExecOneStep()
	assert.Equal(t, "U", robot.DirectionFacing)
	lP = robot.Path[len(robot.Path)-2]
	cP = robot.Path[len(robot.Path)-1]

	assert.Equal(t, lP.Point.X, cP.Point.X)
	assert.Equal(t, lP.Point.Y+1, cP.Point.Y)
}
