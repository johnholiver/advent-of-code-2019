package graph

import "github.com/johnholiver/advent-of-code-2019/pkg"

type Graph struct {
	NodeMap   map[interface{}]*Node
	Roots     map[interface{}]*Node
	formatter pkg.InterfaceFormatter
}

func NewGraph() *Graph {
	return &Graph{
		make(map[interface{}]*Node, 0),
		make(map[interface{}]*Node, 0),
		defaultFormatter,
	}
}

func (g *Graph) SetFormatter(formatter pkg.InterfaceFormatter) *Graph {
	g.formatter = formatter
	return g
}

func defaultFormatter(e interface{}) string {
	return e.(string)
}

func (g *Graph) String() string {
	s := ""
	for _, rootNode := range g.Roots {
		s += rootNode.String()
	}
	return s
}

func (g *Graph) FindNode(value interface{}) *Node {
	return g.NodeMap[value]
}

func (g *Graph) BuildVector(value interface{}, parentValue interface{}) *Node {
	var parentNode *Node
	if parentValue != nil {
		parentNode = g.FindNode(parentValue)
		if parentNode == nil {
			parentNode = g.BuildVector(parentValue, nil)
			g.Roots[parentValue] = parentNode
			if _, ok := g.Roots[value]; ok {
				g.Roots[value] = nil
			}
		}
	}

	currentNode, ok := g.NodeMap[value]
	if !ok {
		currentNode = NewNode(value)
		currentNode.SetFormatter(g.formatter)
		g.NodeMap[value] = currentNode
	}

	if parentNode != nil {
		parentNode.AddChild(currentNode)
	}

	return currentNode
}

func (g *Graph) BuildBranch(value interface{}, parentNode *Node) *Node {
	currentNode := NewNode(value)
	currentNode.SetFormatter(g.formatter)

	if parentNode == nil {
		g.Roots[value] = currentNode
	}
	if parentNode != nil {
		parentNode.AddChild(currentNode)
	}

	return currentNode
}
