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
	file, err := input.Load("2019", "2")
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

	//Restore 1202 state
	m.Variables[1] = 12
	m.Variables[2] = 2

	p := computer.NewProcessor(m)

	err := p.Process()
	if err != nil {
		log.Fatal(err)
	}

	return strconv.Itoa(p.Memory.Variables[0])
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var line string
	var m *computer.Memory

	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Find 19690720
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			m = computer.NewMemory(line)

			m.Variables[1] = noun
			m.Variables[2] = verb

			p := computer.NewProcessor(m)
			err := p.Process()
			if err != nil {
				log.Fatal(err)
			}

			if p.Memory.Variables[0] == 19690720 {
				return strconv.Itoa(100*noun + verb)
			}
		}
	}

	return "failed"
}
