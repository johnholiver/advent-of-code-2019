package main

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_puzzle2(t *testing.T) {
	i := computer.NewIO(nil)
	o := computer.NewIO(nil)
	m := computer.NewMemory("1,0,0,0,99")
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, "2,0,0,0,99", m.String())

	m = computer.NewMemory("2,3,0,3,99")
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, "2,3,0,6,99", m.String())

	m = computer.NewMemory("2,4,4,5,99,0")
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, "2,4,4,5,99,9801", m.String())

	m = computer.NewMemory("1,1,1,4,99,5,6,0,99")
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, "30,1,1,4,2,5,6,0,99", m.String())
}

func Test_paramModeOpcode(t *testing.T) {
	i := computer.NewIO(nil)
	o := computer.NewIO(nil)
	m := computer.NewMemory("1002,4,3,4,33")
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, "1002,4,3,4,99", m.String())
}

func Test_negativeParams(t *testing.T) {
	i := computer.NewIO(nil)
	o := computer.NewIO(nil)
	m := computer.NewMemory("1101,100,-1,4,0")
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, "1101,100,-1,4,99", m.String())
}

func Test_io(t *testing.T) {
	x := 347
	i := computer.NewIO([]int{x})
	o := computer.NewIO(make([]int, 1))
	m := computer.NewMemory("3,0,4,0,99")
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, fmt.Sprintf("%v,0,4,0,99", x), m.String())
	assert.Equal(t, x, p.Output.ReadAt(0))
}

func Test_equals_position(t *testing.T) {
	program := "3,9,8,9,10,9,4,9,99,-1,8"
	i := computer.NewIO([]int{8})
	o := computer.NewIO(make([]int, 1))
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.ReadAt(0))

	i = computer.NewIO([]int{0})
	o = computer.NewIO(make([]int, 1))
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.ReadAt(0))
}

func Test_lessthen_position(t *testing.T) {
	program := "3,9,7,9,10,9,4,9,99,-1,8"
	i := computer.NewIO([]int{6})
	o := computer.NewIO(make([]int, 1))
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.ReadAt(0))

	i = computer.NewIO([]int{9})
	o = computer.NewIO(make([]int, 1))
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.ReadAt(0))
}

func Test_equals_immediate(t *testing.T) {
	program := "3,3,1108,-1,8,3,4,3,99"
	i := computer.NewIO([]int{8})
	o := computer.NewIO(make([]int, 1))
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.ReadAt(0))

	i = computer.NewIO([]int{0})
	o = computer.NewIO(make([]int, 1))
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.ReadAt(0))
}

func Test_lessthen_immediate(t *testing.T) {
	program := "3,3,1107,-1,8,3,4,3,99"
	i := computer.NewIO([]int{6})
	o := computer.NewIO(make([]int, 1))
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.ReadAt(0))

	i = computer.NewIO([]int{9})
	o = computer.NewIO(make([]int, 1))
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.ReadAt(0))
}

func Test_jump_position(t *testing.T) {
	program := "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"
	i := computer.NewIO([]int{0})
	o := computer.NewIO(make([]int, 1))
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.ReadAt(0))

	i = computer.NewIO([]int{1})
	o = computer.NewIO(make([]int, 1))
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.ReadAt(0))
}

func Test_jump_immediate(t *testing.T) {
	program := "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"
	i := computer.NewIO([]int{0})
	o := computer.NewIO(make([]int, 1))
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.ReadAt(0))

	i = computer.NewIO([]int{1})
	o = computer.NewIO(make([]int, 1))
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.ReadAt(0))
}

func Test_large(t *testing.T) {
	program := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	i := computer.NewIO([]int{6})
	o := computer.NewIO(make([]int, 1))
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 999, p.Output.ReadAt(0))

	i = computer.NewIO([]int{8})
	o = computer.NewIO(make([]int, 1))
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1000, p.Output.ReadAt(0))

	i = computer.NewIO([]int{11})
	o = computer.NewIO(make([]int, 1))
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1001, p.Output.ReadAt(0))
}
