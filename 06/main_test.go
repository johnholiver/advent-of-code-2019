package main

import (
	"github.com/johnholiver/advent-of-code-2019/pkg/graph"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var testInput = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

var testInputDisconnect = `A)B
C)D`

var testInputStar = `A)B
A)C
A)D
A)E`

var testInputLine = `A)B
B)C
C)D
D)E`

var testInputInvLine = `D)E
C)D
B)C
A)B`

var testInputCycle = `A)B
B)C
C)A`

func graphBuilder(input string) *graph.Graph {
	graph := graph.NewGraph()
	file := strings.Split(input, "\n")
	for _, input := range file {
		inputSplit := strings.Split(input, ")")
		src := inputSplit[0]
		dst := inputSplit[1]
		graph.BuildVector(dst, src)
	}
	return graph
}

func Test_orbit(t *testing.T) {
	graph := graphBuilder(testInput)

	assert.Equal(t, 0, orbit(graph.FindNode("COM"), nil))
	assert.Equal(t, 3, orbit(graph.FindNode("D"), nil))
	assert.Equal(t, 7, orbit(graph.FindNode("L"), nil))

	orbitSum := 0
	for _, node := range graph.NodeMap {
		orbitSum += orbit(node, nil)
	}
	assert.Equal(t, 42, orbitSum)
}

func Test_orbit2_disconnected(t *testing.T) {
	graph := graphBuilder(testInputDisconnect)
	assert.Equal(t, 0, orbit(graph.FindNode("A"), nil))
	assert.Equal(t, 0, orbit(graph.FindNode("C"), nil))
	assert.Equal(t, 1, orbit(graph.FindNode("B"), nil))
	assert.Equal(t, 1, orbit(graph.FindNode("D"), nil))
}

func Test_orbit3_star(t *testing.T) {
	graph := graphBuilder(testInputStar)

	assert.Equal(t, 0, orbit(graph.FindNode("A"), nil))
	assert.Equal(t, 1, orbit(graph.FindNode("B"), nil))
	assert.Equal(t, 1, orbit(graph.FindNode("C"), nil))
	assert.Equal(t, 1, orbit(graph.FindNode("D"), nil))
	assert.Equal(t, 1, orbit(graph.FindNode("E"), nil))
}

func Test_orbit4_line(t *testing.T) {
	graph := graphBuilder(testInputLine)

	assert.Equal(t, 0, orbit(graph.FindNode("A"), nil))
	assert.Equal(t, 1, orbit(graph.FindNode("B"), nil))
	assert.Equal(t, 2, orbit(graph.FindNode("C"), nil))
	assert.Equal(t, 3, orbit(graph.FindNode("D"), nil))
	assert.Equal(t, 4, orbit(graph.FindNode("E"), nil))
}

func Test_orbit5_invLine(t *testing.T) {
	graph := graphBuilder(testInputInvLine)

	assert.Equal(t, 0, orbit(graph.FindNode("A"), nil))
	assert.Equal(t, 1, orbit(graph.FindNode("B"), nil))
	assert.Equal(t, 2, orbit(graph.FindNode("C"), nil))
	assert.Equal(t, 3, orbit(graph.FindNode("D"), nil))
	assert.Equal(t, 4, orbit(graph.FindNode("E"), nil))
}

var testInputPart2 = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

func Test_part2(t *testing.T) {
	graph := graphBuilder(testInputPart2)

	assert.Equal(t, 4, orbitBetween(graph, "YOU", "SAN"))
}
