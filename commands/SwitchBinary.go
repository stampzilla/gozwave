package commands

type CmdSwitchBinary struct {
	node  byte
	value bool
}

func NewSwitchBinary() *CmdSwitchBinary {
	return &CmdSwitchBinary{}

}

func (c *CmdSwitchBinary) SetValue(v bool) *CmdSwitchBinary {
	c.value = v
	return c
}
func (c *CmdSwitchBinary) SetNode(n int) *CmdSwitchBinary {
	c.node = byte(n)
	return c
}

func (c *CmdSwitchBinary) Encode() []byte {
	v := byte(0x00)
	if c.value {
		v = 0xff
	}
	return []byte{
		0x13,         // SEND ZW
		c.node,       // NOD ID
		0x03,         // length
		SwitchBinary, // command class id
		0x01,         // set
		v,            // value
		0x25,         // transmit options
		//0x08,         // transmit options
	}
}
