package main

import (
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var program = `3,8,1005,8,320,1106,0,11,0,0,0,104,1,104,0,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,1,10,4,10,102,1,8,29,2,1005,1,10,1006,0,11,3,8,1002,8,-1,10,101,1,10,10,4,10,108,0,8,10,4,10,102,1,8,57,1,8,15,10,1006,0,79,1,6,3,10,3,8,102,-1,8,10,101,1,10,10,4,10,108,0,8,10,4,10,101,0,8,90,2,103,18,10,1006,0,3,2,105,14,10,3,8,102,-1,8,10,1001,10,1,10,4,10,108,0,8,10,4,10,101,0,8,123,2,9,2,10,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,1001,8,0,150,1,2,2,10,2,1009,6,10,1,1006,12,10,1006,0,81,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,1,10,4,10,102,1,8,187,3,8,102,-1,8,10,1001,10,1,10,4,10,1008,8,0,10,4,10,101,0,8,209,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,101,0,8,231,1,1008,11,10,1,1001,4,10,2,1104,18,10,3,8,102,-1,8,10,1001,10,1,10,4,10,108,1,8,10,4,10,1001,8,0,264,1,8,14,10,1006,0,36,3,8,1002,8,-1,10,1001,10,1,10,4,10,108,0,8,10,4,10,101,0,8,293,1006,0,80,1006,0,68,101,1,9,9,1007,9,960,10,1005,10,15,99,109,642,104,0,104,1,21102,1,846914232732,1,21102,1,337,0,1105,1,441,21102,1,387512115980,1,21101,348,0,0,1106,0,441,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21102,209533824219,1,1,21102,1,395,0,1106,0,441,21101,0,21477985303,1,21102,406,1,0,1106,0,441,3,10,104,0,104,0,3,10,104,0,104,0,21101,868494234468,0,1,21101,429,0,0,1106,0,441,21102,838429471080,1,1,21102,1,440,0,1106,0,441,99,109,2,21201,-1,0,1,21101,0,40,2,21102,472,1,3,21101,0,462,0,1106,0,505,109,-2,2106,0,0,0,1,0,0,1,109,2,3,10,204,-1,1001,467,468,483,4,0,1001,467,1,467,108,4,467,10,1006,10,499,1102,1,0,467,109,-2,2106,0,0,0,109,4,2101,0,-1,504,1207,-3,0,10,1006,10,522,21101,0,0,-3,21202,-3,1,1,22101,0,-2,2,21102,1,1,3,21102,541,1,0,1106,0,546,109,-4,2105,1,0,109,5,1207,-3,1,10,1006,10,569,2207,-4,-2,10,1006,10,569,22102,1,-4,-4,1105,1,637,22102,1,-4,1,21201,-3,-1,2,21202,-2,2,3,21102,588,1,0,1105,1,546,22101,0,1,-4,21102,1,1,-1,2207,-4,-2,10,1006,10,607,21101,0,0,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,629,21201,-1,0,1,21102,629,1,0,105,1,504,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2105,1,0`

func Test_Paint(t *testing.T) {
	robot := NewRobot(program)

	color := 0
	robot.Paint(color)
	assert.Equal(t, color, robot.Path[len(robot.Path)-1].Value)

	color = 1
	robot.Paint(color)
	assert.Equal(t, color, robot.Path[len(robot.Path)-1].Value)
}

func Test_Move(t *testing.T) {
	robot := NewRobot(program)

	dir := 0
	oP := robot.Path[len(robot.Path)-1].Point
	robot.Move(dir)
	assert.Equal(t, oP.X-1, robot.Path[len(robot.Path)-1].Point.X)
	assert.Equal(t, oP.Y, robot.Path[len(robot.Path)-1].Point.Y)
	robot.Move(dir)
	assert.Equal(t, oP.X-1, robot.Path[len(robot.Path)-1].Point.X)
	assert.Equal(t, oP.Y-1, robot.Path[len(robot.Path)-1].Point.Y)
	robot.Move(dir)
	assert.Equal(t, oP.X, robot.Path[len(robot.Path)-1].Point.X)
	assert.Equal(t, oP.Y-1, robot.Path[len(robot.Path)-1].Point.Y)
	robot.Move(dir)
	assert.Equal(t, oP.X, robot.Path[len(robot.Path)-1].Point.X)
	assert.Equal(t, oP.Y, robot.Path[len(robot.Path)-1].Point.Y)

	assert.Equal(t, oP, robot.Path[len(robot.Path)-1].Point)

	dir = 1
	robot.Move(dir)
	assert.Equal(t, oP.X+1, robot.Path[len(robot.Path)-1].Point.X)
	assert.Equal(t, oP.Y, robot.Path[len(robot.Path)-1].Point.Y)
	robot.Move(dir)
	assert.Equal(t, oP.X+1, robot.Path[len(robot.Path)-1].Point.X)
	assert.Equal(t, oP.Y-1, robot.Path[len(robot.Path)-1].Point.Y)
	robot.Move(dir)
	assert.Equal(t, oP.X, robot.Path[len(robot.Path)-1].Point.X)
	assert.Equal(t, oP.Y-1, robot.Path[len(robot.Path)-1].Point.Y)
	robot.Move(dir)
	assert.Equal(t, oP.X, robot.Path[len(robot.Path)-1].Point.X)
	assert.Equal(t, oP.Y, robot.Path[len(robot.Path)-1].Point.Y)

	assert.Equal(t, oP, robot.Path[len(robot.Path)-1].Point)
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

func Test_Path(t *testing.T) {
	robot := NewRobot(program)

	assert.Equal(t, 0, len(robot.Path[:len(robot.Path)-1]))
	uP := uniquePointsInPath(robot.Path[:len(robot.Path)-1])
	assert.Equal(t, 0, len(uP))

	repeat := 6
	for repeat > 0 {
		robot.ExecOneStep(robot.processOne)
		repeat--
	}
	assert.Equal(t, 6, len(robot.Path[:len(robot.Path)-1]))
	uP = uniquePointsInPath(robot.Path[:len(robot.Path)-1])
	assert.Equal(t, 5, len(uP))

	robot.ExecOneStep(robot.processOne)
	assert.Equal(t, 7, len(robot.Path[:len(robot.Path)-1]))
	uP = uniquePointsInPath(robot.Path[:len(robot.Path)-1])
	assert.Equal(t, 5, len(uP))
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

func mockProcessOne(color, dir int) func(int) (int, int, bool) {
	return func(int) (int, int, bool) {
		return color, dir, false
	}
}

func Test_ExecOneStepPaint(t *testing.T) {
	robot := NewRobot(program)

	colorInput := 0
	dirInput := 0
	robot.ExecOneStep(mockProcessOne(colorInput, dirInput))
	lP := robot.Path[len(robot.Path)-2]
	cP := robot.Path[len(robot.Path)-1]

	assert.Equal(t, colorInput, lP.Value)
	assert.Equal(t, 0, cP.Value)

	colorInput = 1
	dirInput = 0
	robot.ExecOneStep(mockProcessOne(colorInput, dirInput))
	lP = robot.Path[len(robot.Path)-2]
	cP = robot.Path[len(robot.Path)-1]

	assert.Equal(t, colorInput, lP.Value)
	assert.Equal(t, 0, cP.Value)
}

func Test_ExecOneStepMove(t *testing.T) {
	robot := NewRobot(program)

	assert.Equal(t, "U", robot.DirectionFacing)

	//Turn left
	colorInput := 0
	dirInput := 0
	robot.ExecOneStep(mockProcessOne(colorInput, dirInput))
	assert.Equal(t, "L", robot.DirectionFacing)
	lP := robot.Path[len(robot.Path)-2]
	cP := robot.Path[len(robot.Path)-1]

	assert.Equal(t, lP.Point.X-1, cP.Point.X)
	assert.Equal(t, lP.Point.Y, cP.Point.Y)

	//Turn left
	colorInput = 0
	dirInput = 0
	robot.ExecOneStep(mockProcessOne(colorInput, dirInput))
	assert.Equal(t, "D", robot.DirectionFacing)
	lP = robot.Path[len(robot.Path)-2]
	cP = robot.Path[len(robot.Path)-1]

	assert.Equal(t, lP.Point.X, cP.Point.X)
	assert.Equal(t, lP.Point.Y-1, cP.Point.Y)

	//Turn right
	colorInput = 0
	dirInput = 1
	robot.ExecOneStep(mockProcessOne(colorInput, dirInput))
	assert.Equal(t, "L", robot.DirectionFacing)
	lP = robot.Path[len(robot.Path)-2]
	cP = robot.Path[len(robot.Path)-1]

	assert.Equal(t, lP.Point.X-1, cP.Point.X)
	assert.Equal(t, lP.Point.Y, cP.Point.Y)

	//Turn right
	colorInput = 0
	dirInput = 1
	robot.ExecOneStep(mockProcessOne(colorInput, dirInput))
	assert.Equal(t, "U", robot.DirectionFacing)
	lP = robot.Path[len(robot.Path)-2]
	cP = robot.Path[len(robot.Path)-1]

	assert.Equal(t, lP.Point.X, cP.Point.X)
	assert.Equal(t, lP.Point.Y+1, cP.Point.Y)
}

func Test_part1(t *testing.T) {
	robot := NewRobot(program)

	robot.Exec()

	//fmt.Printf("Robot Path size: %v\n", len(robot.Path))
	//fmt.Printf("RobotPath:%v\n",robot.Path)
	uniquePoints := uniquePointsInPath(robot.Path[:len(robot.Path)-1])
	assert.Equal(t, 2469, len(uniquePoints))

}
