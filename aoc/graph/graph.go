package graph

type Node struct {
	parents  []*Node
	children map[string]*Node
	Value    string
}

type Graph struct {
	Nodes []*Node
}

func New() *Graph {
	return &Graph{}
}

func (g *Graph) MakeNode(value string) *Node {
	newNode := &Node{Value: value, children: make(map[string]*Node)}
	g.Nodes = append(g.Nodes, newNode)
	return newNode
}

func (g *Graph) FindNode(value string) *Node {
	for _, n := range g.Nodes {
		if n.Value == value {
			return n
		}
	}
	return nil
}

func (g *Graph) DependsOn(parent string, child string) {
	p := g.FindNode(parent)
	c := g.FindNode(child)
	if c == nil {
		panic("Could not find child node")
	}
	if p == nil {
		panic("Could not find parent node")
	}

	for _, parent := range c.GetParents() {
		if parent == p {
			panic("Readding parent")
		}
	}

	for _, child := range p.GetChildren() {
		if child == c {
			panic("Readding child")
		}
	}

	p.children[c.Value] = c
	c.parents = append(c.parents, p)
}

func (n *Node) GetChildren() map[string]*Node {
	return n.children
}

func (n *Node) GetParents() []*Node {
	return n.parents
}
