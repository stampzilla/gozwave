package nodes

type Node interface {
	Id() byte

	Identify()
}

type List interface {
	Add(node Node)
	All() map[byte]Node
	Get(byte) Node
}
