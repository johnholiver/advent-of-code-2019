package graph

type Graph struct {
	NodeMap map[string]*Node
	Roots   map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		make(map[string]*Node, 0),
		make(map[string]*Node, 0),
	}
}

func (g *Graph) BuildVector(value string, parentValue *string) *Node {
	var parentNode *Node
	if parentValue != nil {
		parentNode = g.FindNode(*parentValue)
		if parentNode == nil {
			parentNode = g.BuildVector(*parentValue, nil)
			g.Roots[*parentValue] = parentNode
			if _, ok := g.Roots[value]; ok {
				g.Roots[value] = nil
			}
		}
	}

	currentNode, ok := g.NodeMap[value]
	if !ok {
		currentNode = NewNode(value)
		g.NodeMap[value] = currentNode
	}

	if parentNode != nil {
		parentNode.AddChild(currentNode)
	}

	return currentNode
}

func (g *Graph) FindNode(value string) *Node {
	return g.NodeMap[value]
}

func (g *Graph) Print() {
	for _, rootNode := range g.Roots {
		rootNode.Print()
	}
}
