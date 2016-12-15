package commands

type CmdSwitchMultilevel struct {
	node  byte
	level byte
}

func NewSwitchMultilevel() *CmdSwitchMultilevel {

	return &CmdSwitchMultilevel{}
}

func (c *CmdSwitchMultilevel) SetValue(v float64) *CmdSwitchMultilevel {
	if v < 1 {
		v = 0
	}
	if v > 99 {
		v = 99
	}
	// TODO: Round to integer first to fix precision error
	c.level = byte(v)
	return c
}
func (c *CmdSwitchMultilevel) SetNode(n int) *CmdSwitchMultilevel {
	c.node = byte(n)
	return c
}

func (c *CmdSwitchMultilevel) Encode() []byte {
	return []byte{
		0x13,             // SEND ZW
		c.node,           // NOD ID
		0x03,             // length
		SwitchMultilevel, // command class id
		0x01,             // set
		c.level,          // value
		0x25,             // transmit options
	}
}
