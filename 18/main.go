package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"log"
	"os"
)

func main() {
	file, err := input.Load("2019", "18")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1(file)
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//TODO: Massage the input, line by line
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
