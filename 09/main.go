package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	computer_io "github.com/johnholiver/advent-of-code-2019/pkg/computer/io"
	computer_mem "github.com/johnholiver/advent-of-code-2019/pkg/computer/memory"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := input.Load("2019", "9")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Result part1: %v\n", part1(file))

	file.Seek(0, io.SeekStart)
	fmt.Printf("Result part2: %v\n", part2(file))
}

func part1(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	i := computer_io.NewTape()
	i.Set([]int{1})
	o := computer_io.NewTape()

	p := computer.NewProcessor(i, o, nil)
	m := computer_mem.NewRelative(p, program)
	p.Memory = m
	p.Process()

	output := p.Output.Read()

	return strconv.Itoa(output)
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var program string
	for scanner.Scan() {
		program = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	i := computer_io.NewTape()
	i.Set([]int{2})
	o := computer_io.NewTape()

	p := computer.NewProcessor(i, o, nil)
	m := computer_mem.NewRelative(p, program)
	p.Memory = m
	p.Process()

	output := p.Output.Read()

	return strconv.Itoa(output)
}
