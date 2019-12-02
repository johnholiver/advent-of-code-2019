package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
//2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
//2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
//1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.

func Test_part1(t *testing.T) {
	m := initMemory("1,0,0,0,99")
	processMemory(m)
	assert.Equal(t, "2,0,0,0,99", m.String())

	m = initMemory("2,3,0,3,99")
	processMemory(m)
	assert.Equal(t, "2,3,0,6,99", m.String())

	m = initMemory("2,4,4,5,99,0")
	processMemory(m)
	assert.Equal(t, "2,4,4,5,99,9801", m.String())

	m = initMemory("1,1,1,4,99,5,6,0,99")
	processMemory(m)
	assert.Equal(t, "30,1,1,4,2,5,6,0,99", m.String())
}
