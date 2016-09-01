package nodes

type Node interface {
	Id() byte

	Identify(chan struct{})
}

type ListInterface interface {
	Add(node Node)
	All() map[byte]Node
	Get(byte) Node
}
