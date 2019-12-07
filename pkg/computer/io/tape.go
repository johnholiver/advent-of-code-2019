package io

import "github.com/johnholiver/advent-of-code-2019/pkg/computer"

type Tape struct {
	values []int
	cursor int
}

func NewTape() *Tape {
	return &Tape{
		values: make([]int, 0),
		cursor: 0,
	}
}

var _ computer.IO = &Tape{}

//Cursor operations
func (io *Tape) Reset() {
	io.cursor = 0
}

func (io *Tape) Previous() {
	io.cursor--
}

func (io *Tape) Next() {
	io.cursor++
}

//Read operations
func (io *Tape) Read() int {
	value := io.values[io.cursor]
	io.Next()
	return value
}

func (io *Tape) ReadAt(index int) int {
	return io.values[index]
}

//Write operations
func (io *Tape) Set(values []int) {
	io.values = values
}

func (io *Tape) Write(value int) {
	if io.cursor >= len(io.values) {
		io.values = append(io.values, value)
	} else {
		io.values[io.cursor] = value
	}
	io.Next()
}

func (io *Tape) WriteAt(value int, index int) {
	io.values[index] = value
}

func (io *Tape) Append(value int) {
	io.values = append(io.values, value)
}
