package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//R75,D30,R83,U83,L12,D49,R71,U7,L72
//U62,R66,U55,R34,D71,R55,D58,R83 = distance 159

//R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
//U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135

func Test_part1(t *testing.T) {
	wire1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	wire2 := "U62,R66,U55,R34,D71,R55,D58,R83"
	assert.Equal(t, 159, smallestManhattan(wire1, wire2))

	wire1 = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	wire2 = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	assert.Equal(t, 135, smallestManhattan(wire1, wire2))
}
