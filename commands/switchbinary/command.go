package switchbinary

type command struct {
	node  byte
	value bool
}

func New() *command {
	return &command{}

}

func (c *command) SetValue(v bool) *command {
	c.value = v
	return c
}
func (c *command) SetNode(n byte) *command {
	c.node = n
	return c
}

func (c *command) Encode() []byte {
	v := byte(0x00)
	if c.value {
		v = 0xff
	}
	return []byte{
		0x13,   // SEND ZW
		c.node, // NOD ID
		0x03,   // binary switch report
		0x25,   // command class id
		0x01,   // set
		v,      // value
		0x25,   // transmit options
		0x25,   // transmit options
	}
}
