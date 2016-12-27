package commands

// Raw binary data type
type Raw struct {
	data []byte
	node int
}

// NewRaw creates a new Raw
func NewRaw(b []byte) *Raw {
	return &Raw{data: b}
}

// Encode make sure we implement the Encodable interface
func (r Raw) Encode() []byte {

	return append([]byte{
		0x13,                  // SendData // TODO: use functions.SendData somehow without import cycle
		byte(r.node),          // Node id
		byte(len(r.data) + 1), // Length
	}, r.data...)
}

func (r *Raw) SetNode(n int) {
	r.node = n
}
