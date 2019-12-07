package main

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_paramModeOpcode(t *testing.T) {
	m := computer.NewMemory("1002,4,3,4,33")
	p := computer.NewProcessor(nil, nil, m)
	p.Process()
	assert.Equal(t, "1002,4,3,4,99", m.String())
}

func Test_negativeParams(t *testing.T) {
	m := computer.NewMemory("1101,100,-1,4,0")
	p := computer.NewProcessor(nil, nil, m)
	p.Process()
	assert.Equal(t, "1101,100,-1,4,99", m.String())
}

func Test_io(t *testing.T) {
	program := "3,0,4,0,99"
	target := 347
	i := computer_io.NewTape()
	i.Set([]int{target})
	o := computer_io.NewTape()
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, fmt.Sprintf("%v,0,4,0,99", target), m.String())
	assert.Equal(t, target, p.Output.Read())
}

func Test_equals_position(t *testing.T) {
	program := "3,9,8,9,10,9,4,9,99,-1,8"
	i := computer_io.NewTape()
	i.Set([]int{8})
	o := computer_io.NewTape()
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.Read())

	i = computer_io.NewTape()
	i.Set([]int{0})
	o = computer_io.NewTape()
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.Read())
}

func Test_lessthen_position(t *testing.T) {
	program := "3,9,7,9,10,9,4,9,99,-1,8"
	i := computer_io.NewTape()
	i.Set([]int{6})
	o := computer_io.NewTape()
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.Read())

	i = computer_io.NewTape()
	i.Set([]int{9})
	o = computer_io.NewTape()
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.Read())
}

func Test_equals_immediate(t *testing.T) {
	program := "3,3,1108,-1,8,3,4,3,99"
	i := computer_io.NewTape()
	i.Set([]int{8})
	o := computer_io.NewTape()
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.Read())

	i = computer_io.NewTape()
	i.Set([]int{0})
	o = computer_io.NewTape()
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.Read())
}

func Test_lessthen_immediate(t *testing.T) {
	program := "3,3,1107,-1,8,3,4,3,99"
	i := computer_io.NewTape()
	i.Set([]int{6})
	o := computer_io.NewTape()
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.Read())

	i = computer_io.NewTape()
	i.Set([]int{9})
	o = computer_io.NewTape()
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.Read())
}

func Test_jump_position(t *testing.T) {
	program := "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"
	i := computer_io.NewTape()
	i.Set([]int{0})
	o := computer_io.NewTape()
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.Read())

	i = computer_io.NewTape()
	i.Set([]int{1})
	o = computer_io.NewTape()
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.Read())
}

func Test_jump_immediate(t *testing.T) {
	program := "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"
	i := computer_io.NewTape()
	i.Set([]int{0})
	o := computer_io.NewTape()
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 0, p.Output.Read())

	i = computer_io.NewTape()
	i.Set([]int{1})
	o = computer_io.NewTape()
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1, p.Output.Read())
}

func Test_large(t *testing.T) {
	program := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	i := computer_io.NewTape()
	i.Set([]int{6})
	o := computer_io.NewTape()
	m := computer.NewMemory(program)
	p := computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 999, p.Output.Read())

	i = computer_io.NewTape()
	i.Set([]int{8})
	o = computer_io.NewTape()
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1000, p.Output.Read())

	i = computer_io.NewTape()
	i.Set([]int{11})
	o = computer_io.NewTape()
	m = computer.NewMemory(program)
	p = computer.NewProcessor(i, o, m)
	p.Process()
	assert.Equal(t, 1001, p.Output.Read())
}
