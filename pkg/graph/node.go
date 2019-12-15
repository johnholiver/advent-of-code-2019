package graph

import (
	"fmt"
	"github.com/johnholiver/advent-of-code-2019/pkg"
)

type Node struct {
	Value     interface{}
	Parent    *Node
	Children  []*Node
	formatter pkg.InterfaceFormatter
}

func NewNode(value interface{}) *Node {
	return &Node{
		value,
		nil,
		make([]*Node, 0),
		defaultFormatter,
	}
}

func (n *Node) SetFormatter(formatter pkg.InterfaceFormatter) *Node {
	n.formatter = formatter
	return n
}

func (n *Node) AddChild(c *Node) {
	c.Parent = n
	n.Children = append(n.Children, c)
}

func (n *Node) Print() {
	fmt.Printf("%v", n.formatter(n.Value))
	for _, c := range n.Children {
		c.Print()
	}
}
