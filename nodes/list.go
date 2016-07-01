package nodes

import "github.com/stampzilla/gozwave/functions"

type Node interface {
	Id() byte
}

type List interface {
	Get(byte) Node
}

type list map[byte]Node

func New(address byte, nodeinfo *functions.FuncGetNodeProtocolInfo) *node {
	return &node{id: address}
}

func NewList() *list {
	l := make(list)
	return &l
}

func (l list) Add(node Node) {
	l[node.Id()] = node
}

func (l list) Get(id byte) Node {
	return l[id]
}

type node struct {
	id byte
}

func (n *node) Id() byte {
	return n.id
}
