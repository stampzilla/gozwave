package nodes

type List map[byte]*Node

func NewList() *List {
	l := make(List)
	return &l
}

func (l List) Add(node *Node) {
	l[node.Id] = node
}

func (l List) All() map[byte]*Node {
	return l
}

func (l List) Get(id byte) *Node {
	return l[id]
}
