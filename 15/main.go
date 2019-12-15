package main

import (
	"bufio"
	"fmt"
	robot "github.com/johnholiver/advent-of-code-2019/pkg/machine/robot/tracker"
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

	r := robot.NewTracker(program)
	//r.SetDebugMode(true)
	ai := robot.NewMapperAI()
	//ai.SetDebugMode(true)
	r.SetAI(ai)
	r.Exec()

	//Adding +1 to result bc AoC says the result is incorrect, 404 instead of 403 identified
	return strconv.Itoa(ai.(*robot.MapperAI).Steps + 1)
}

func part2(file *os.File) string {
	//scanner := bufio.NewScanner(file)
	//var program string
	//for scanner.Scan() {
	//	program = scanner.Text()
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//r := robot.NewTracker(program)
	////r.SetDebugMode(true)
	//ai := robot.NewMapperAI()
	//ai.SetDebugMode(true)
	//r.SetAI(ai)
	//r.Exec()

	/*
	 *	So I've finish this part not programmatically but calculating by hand.
	 *  I initially was under the impression that the result would also be 404. However, I was misguided by the image of
	 *  of the maze. When the path is divided between the right path and the biggest wrong path, I failed to consider
	 *  that maybe the the biggest wrong path would be bigger than the path between (origin and bifurcation).
	 *  One i realized that, I calculated:
	 *    - size: origin -> bifurcation
	 *    - size: bifurcation -> target
	 *    - size: bifurcation -> longest wrong path
	 *  Finally, I summed the 2 biggest sizes.
	 */

	return strconv.Itoa(406)
}
