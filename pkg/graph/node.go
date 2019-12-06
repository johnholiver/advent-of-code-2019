package graph

import "fmt"

type Node struct {
	Value    string
	Parent   *Node
	Children []*Node
}

func NewNode(value string) *Node {
	return &Node{
		value,
		nil,
		make([]*Node, 0),
	}
}

func (n *Node) AddChild(c *Node) {
	c.Parent = n
	n.Children = append(n.Children, c)
}

func (n *Node) Print() {
	fmt.Printf("%v", n.Value)
	for _, c := range n.Children {
		c.Print()
	}
}
