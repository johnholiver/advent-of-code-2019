package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"github.com/johnholiver/advent-of-code-2019/pkg/machine/arcade"
)

func main() {
	file, err := input.Load("2019", "13")
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

	arc := arcade.New(program)

	blockTiles := 0
	for {
		output, done := arc.ProcessOne(nil)
		if done {
			break
		}
		tile := output[2]
		if tile == 2 {
			blockTiles++
		}
	}

	return fmt.Sprintf("%v", blockTiles)
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

	arc := arcade.New(program)
	arc.SetDebugMode(true)
	player := arcade.NewArcadeAI()

	arc.PutCoin(player)
	arc.Exec()

	return fmt.Sprintf("%v", arc.Score)
}
