package nodes

type list map[byte]Node

func NewList() *list {
	l := make(list)
	return &l
}

func (l list) Add(node Node) {
	l[node.Id()] = node

	go node.Identify()
}

func (l list) All() map[byte]Node {
	return l
}

func (l list) Get(id byte) Node {
	return l[id]
}
