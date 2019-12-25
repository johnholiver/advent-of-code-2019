package main

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/life"
	"github.com/stretchr/testify/assert"
	"testing"
)

var input0 = `....#
#..#.
#..##
..#..
#....
`

var input1 = `#..#.
####.
###.#
##.##
.##..
`

var input2 = `#####
....#
....#
...#.
#.###
`

var input3 = `#....
####.
...##
#.##.
.##.#
`

var input4 = `####.
....#
##..#
.....
##...
`

var inputBio = `.....
.....
.....
#....
.#...
`

var inputEmpty = `.....
.....
.....
.....
.....
`

func Test_NewEmptyWorld(t *testing.T) {
	w := life.NewWorld(5, 5, 0)
	w.FillEmpty(life.NewMonoverseCell)
	assert.Equal(t, inputEmpty, w.String())
}
func Test_WorldString(t *testing.T) {
	w0 := buildWorld(input0, life.NewMonoverseCell)
	assert.Equal(t, input0, w0.String())
}

func Test_WorldTick(t *testing.T) {
	w0 := buildWorld(input0, life.NewMonoverseCell)
	assert.Equal(t, input0, w0.String())
	w1 := w0.Tick()
	assert.Equal(t, input1, w1.String())
	w2 := w1.Tick()
	assert.Equal(t, input2, w2.String())
	w3 := w2.Tick()
	assert.Equal(t, input3, w3.String())
	w4 := w3.Tick()
	assert.Equal(t, input4, w4.String())
}

func Test_WorldBiodiversityRating(t *testing.T) {
	w := buildWorld(inputBio, life.NewMonoverseCell)
	assert.Equal(t, 2129920, w.BiodiversityRating())
}

func Test_Multiverse(t *testing.T) {
	w := buildWorld(input0, life.NewMultiverseCell)
	mv := life.NewMultiverse(w)
	mv.OutterDepth = -2
	mv.InnerDepth = 2
	fmt.Println(mv)
	mv.Tick()
	fmt.Println(mv)
	mv.Tick()
	fmt.Println(mv)
}

func Test_MultiverseAoC(t *testing.T) {
	w := buildWorld(input0, life.NewMultiverseCell)
	mv := life.NewMultiverse(w)
	mv.OutterDepth = -5
	mv.InnerDepth = 5

	for i := 1; i <= 10; i++ {
		mv.Tick()
	}
	assert.Equal(t, 99, mv.CountBugs())
}
