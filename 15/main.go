package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine/robot"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/johnholiver/advent-of-code-2019/pkg/input"
)

func main() {
	file, err := input.Load("2019", "15")
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

	////This robot maps the area
	//robot := robot.NewMapper(program)
	//
	//robot.Exec()
	//
	////This robot A-star
	robot := robot.NewTracker(program)

	robot.Exec()

	return strconv.Itoa(len(robot.Path))
}

func part2(file *os.File) string {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//TODO: Massage the input, line by line
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return "implement me"
}
