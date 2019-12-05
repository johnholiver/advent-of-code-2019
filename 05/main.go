package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/computer"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := input.Load("2019", "5")
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
	var m *computer.Memory

	for scanner.Scan() {
		line := scanner.Text()
		m = computer.NewMemory(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	i := computer.NewIO([]int{1})
	o := computer.NewIO(make([]int, 1))

	p := computer.NewProcessor(i, o, m)

	err := p.Process()
	if err != nil {
		log.Fatal(err)
	}

	p.Output.Reset()
	return strconv.Itoa(p.Output.Read())
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var m *computer.Memory

	for scanner.Scan() {
		line := scanner.Text()
		m = computer.NewMemory(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	i := computer.NewIO([]int{5})
	o := computer.NewIO(make([]int, 1))

	p := computer.NewProcessor(i, o, m)

	err := p.Process()
	if err != nil {
		log.Fatal(err)
	}

	p.Output.Reset()
	return strconv.Itoa(p.Output.Read())
}
