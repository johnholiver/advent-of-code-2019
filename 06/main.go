package main

import (
	"bufio"
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg/graph"
	"github.com/johnholiver/advent-of-code-2019/pkg/input"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := input.Load("2019", "6")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Result part1: %v\n", part1(file))

	file.Seek(0, io.SeekStart)
	fmt.Printf("Result part2: %v\n", part2(file))
}

func part1(file *os.File) string {
	graph := graph.NewGraph()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		inputSplit := strings.Split(input, ")")
		src := inputSplit[0]
		dst := inputSplit[1]
		graph.BuildVector(dst, src)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	orbitSum := 0
	for _, node := range graph.NodeMap {
		orbitSum += orbit(node, nil)
	}

	return strconv.Itoa(orbitSum)
}

func orbit(node *graph.Node, target *graph.Node) int {
	stepsToRoot := 0
	for {
		if node.Parent == target {
			break

		}
		stepsToRoot += 1
		node = node.Parent
	}
	return stepsToRoot
}

func part2(file *os.File) string {
	graph := graph.NewGraph()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		inputSplit := strings.Split(input, ")")
		src := inputSplit[0]
		dst := inputSplit[1]
		graph.BuildVector(dst, src)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return strconv.Itoa(orbitBetween(graph, "YOU", "SAN"))
}

func orbitBetween(g *graph.Graph, value1, value2 string) int {
	n1 := g.FindNode(value1)
	path1 := pathNode(n1)
	n2 := g.FindNode(value2)
	path2 := pathNode(n2)

	var turningNode *graph.Node
	for _, nPath1 := range path1 {
		for _, nPath2 := range path2 {
			if nPath1.Value == nPath2.Value {
				turningNode = nPath1
				break
			}
		}
		if turningNode != nil {
			break
		}
	}

	n1ToTN := orbit(n1, turningNode)
	n2ToTN := orbit(n2, turningNode)

	return n1ToTN + n2ToTN
}

func pathNode(node *graph.Node) []*graph.Node {
	path := make([]*graph.Node, 0)

	for {
		if node.Parent == nil {
			break

		}
		path = append(path, node.Parent)
		node = node.Parent
	}

	return path
}
