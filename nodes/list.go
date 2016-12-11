package nodes

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/stampzilla/gozwave/serialapi"
)

type List struct {
	nodes      map[int]*Node
	connection *serialapi.Connection

	sync.RWMutex
}

func NewList() *List {
	l := List{
		nodes: make(map[int]*Node),
	}
	return &l
}

func (l *List) Add(node *Node) {
	l.Lock()
	if l.nodes == nil {
		l.nodes = make(map[int]*Node)
	}

	node.connection = l.connection
	for k, _ := range node.Endpoints {
		node.Endpoints[k].node = node
	}
	l.nodes[int(node.Id)] = node
	l.Unlock()
}

func (l List) All() map[int]*Node {
	l.RLock()
	defer l.RUnlock()

	return l.nodes
}

func (l List) Get(id int) *Node {
	l.RLock()
	defer l.RUnlock()

	return l.nodes[id]
}

func (l *List) SetConnection(connection *serialapi.Connection) {
	l.connection = connection
}

func (l List) MarshalJSON() ([]byte, error) {
	s := make(map[string]*Node)

	for k, v := range l.All() {
		s[strconv.Itoa(k)] = v
	}

	return json.Marshal(s)
}

func (l *List) UnmarshalJSON(data []byte) error {
	s := make(map[string]*Node)

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	for _, v := range s {
		l.Add(v)
	}

	return nil
}
