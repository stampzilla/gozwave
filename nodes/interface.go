package nodes

type Node interface {
	Id() byte

	Identify()
}

type List interface {
	Add(node Node)
	Get(byte) Node
}
