package main

import (
	"github.com/johnholiver/advent-of-code-2019/pkg/grid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var input1 = `.#..#
.....
#####
....#
...##`

var output1 = `.7..7
.....
67775
....7
...87`

var input2 = `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`

var input3 = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`

var input4 = `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`

var input5 = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

var input6 = `....#.....#.#...##..........#.......#......
.....#...####..##...#......#.........#.....
.#.#...#..........#.....#.##.......#...#..#
.#..#...........#..#..#.#.......####.....#.
##..#.................#...#..........##.##.
#..##.#...#.....##.#..#...#..#..#....#....#
##...#.............#.#..........#...#.....#
#.#..##.#.#..#.#...#.....#.#.............#.
...#..##....#........#.....................
##....###..#.#.......#...#..........#..#..#
....#.#....##...###......#......#...#......
.........#.#.....#..#........#..#..##..#...
....##...#..##...#.....##.#..#....#........
............#....######......##......#...#.
#...........##...#.#......#....#....#......
......#.....#.#....#...##.###.....#...#.#..
..#.....##..........#..........#...........
..#.#..#......#......#.....#...##.......##.
.#..#....##......#.............#...........
..##.#.....#.........#....###.........#..#.
...#....#...#.#.......#...#.#.....#........
...####........#...#....#....#........##..#
.#...........#.................#...#...#..#
#................#......#..#...........#..#
..#.#.......#...........#.#......#.........
....#............#.............#.####.#.#..
.....##....#..#...........###........#...#.
.#.....#...#.#...#..#..........#..#.#......
.#.##...#........#..#...##...#...#...#.#.#.
#.......#...#...###..#....#..#...#.........
.....#...##...#.###.#...##..........##.###.
..#.....#.##..#.....#..#.....#....#....#..#
.....#.....#..............####.#.........#.
..#..#.#..#.....#..........#..#....#....#..
#.....#.#......##.....#...#...#.......#.#..
..##.##...........#..........#.............
...#..##....#...##..##......#........#....#
.....#..........##.#.##..#....##..#........
.#...#...#......#..#.##.....#...#.....##...
...##.#....#...........####.#....#.#....#..
...#....#.#..#.........#.......#..#...##...
...##..............#......#................
........................#....##..#........#`

func Test_buildGrid(t *testing.T) {
	g := buildGrid(input6)
	assert.Equal(t, 43, g.Width)
	assert.Equal(t, 43, g.Height)
}

func Test_pointInLine(t *testing.T) {
	assert.False(t, pointInLine(
		*grid.NewPoint(4, 0),
		*grid.NewPoint(0, 3),
		*grid.NewPoint(1, 2),
	))

	assert.True(t, pointInLine(
		*grid.NewPoint(4, 0),
		*grid.NewPoint(4, 3),
		*grid.NewPoint(4, 2),
	))

	assert.False(t, pointInLine(
		*grid.NewPoint(1, 2),
		*grid.NewPoint(2, 2),
		*grid.NewPoint(4, 0),
	))

	assert.True(t, pointInLine(
		*grid.NewPoint(1, 0),
		*grid.NewPoint(3, 4),
		*grid.NewPoint(2, 2),
	))
}

func Test_pointLos(t *testing.T) {
	g := buildGrid(input1)

	//Inner range
	assert.Equal(t, 7, pointLos2(g, 1, 2))
	assert.Equal(t, 7, pointLos2(g, 2, 2))
	assert.Equal(t, 7, pointLos2(g, 3, 2))
	//Border
	assert.Equal(t, 7, pointLos2(g, 1, 0))
	assert.Equal(t, 7, pointLos2(g, 4, 0)) //fails
	assert.Equal(t, 5, pointLos2(g, 4, 2))
	assert.Equal(t, 7, pointLos2(g, 4, 3)) //fails
	assert.Equal(t, 7, pointLos2(g, 4, 4)) //fails
	assert.Equal(t, 8, pointLos2(g, 3, 4))
	assert.Equal(t, 6, pointLos2(g, 0, 2))
}

func Test_part1(t *testing.T) {
	valuedPoint1 := biggestLineOfSight(lineOfSight(buildGrid(input1)))
	assert.Equal(t, grid.NewValuedPoint(3, 4, 8), valuedPoint1)

	valuedPoint2 := biggestLineOfSight(lineOfSight(buildGrid(input2)))
	assert.Equal(t, grid.NewValuedPoint(5, 8, 33), valuedPoint2)

	valuedPoint3 := biggestLineOfSight(lineOfSight(buildGrid(input3)))
	assert.Equal(t, grid.NewValuedPoint(1, 2, 35), valuedPoint3)

	valuedPoint4 := biggestLineOfSight(lineOfSight(buildGrid(input4)))
	assert.Equal(t, grid.NewValuedPoint(6, 3, 41), valuedPoint4)

	valuedPoint5 := biggestLineOfSight(lineOfSight(buildGrid(input5)))
	assert.Equal(t, grid.NewValuedPoint(11, 13, 210), valuedPoint5)
}

func Test_transform_and_mirror(t *testing.T) {
	g := buildGrid(input1)
	g.Transform(-2, -2)
	g.MirrorY()
	assert.Equal(t, g.Get(0, 0).Point.X, -2)
	assert.Equal(t, g.Get(0, 0).Point.Y, 2)
}

func Test_polar(t *testing.T) {
	pTop := grid.NewPoint(0, 1)
	pRig := grid.NewPoint(1, 0)
	pBot := grid.NewPoint(0, -1)
	pLef := grid.NewPoint(-1, 0)
	pTRi := grid.NewPoint(1, 1)

	ppTop := grid.NewPolarPoint(pTop)
	ppRig := grid.NewPolarPoint(pRig)
	ppBot := grid.NewPolarPoint(pBot)
	ppLef := grid.NewPolarPoint(pLef)
	ppTRi := grid.NewPolarPoint(pTRi)

	ppTop.Rotate(float64(-90))
	ppRig.Rotate(float64(-90))
	ppBot.Rotate(float64(-90))
	ppLef.Rotate(float64(-90))
	ppTRi.Rotate(float64(-90))

	assert.Equal(t, float64(360), ppTop.Angle)
	assert.Equal(t, float64(270), ppRig.Angle)
	assert.Equal(t, float64(180), ppBot.Angle)
	assert.Equal(t, float64(90), ppLef.Angle)
	assert.Equal(t, float64(315), ppTRi.Angle)
}

var input1_part2 = `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`

func Test_part2(t *testing.T) {
	g := buildGrid(input1)
	pp := part2(g, grid.NewPoint(3, 4), 6)
	assert.Equal(t, grid.NewPoint(0, 2), pp)
}

func Test_part2_input6(t *testing.T) {
	g := buildGrid(input5)
	elfBet := part2(g, grid.NewPoint(11, 13), 200)
	assert.Equal(t, grid.NewPoint(8, 2), elfBet)
}

func Test_part2_newInput1(t *testing.T) {
	g := buildGrid(input1_part2)

	elfBet := part2(g, grid.NewPoint(8, 3), 1)
	assert.Equal(t, grid.NewPoint(8, 1), elfBet)
}

func Test_part2_newInput2(t *testing.T) {
	g := buildGrid(input1_part2)

	elfBet := part2(g, grid.NewPoint(8, 3), 15)
	assert.Equal(t, grid.NewPoint(16, 4), elfBet)
}
