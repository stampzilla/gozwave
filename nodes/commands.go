package nodes

import (
	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/interfaces"
)

func (n *Node) On() {
	var send interfaces.Encodable

	switch {
	case n.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinary()
		cmd.SetValue(true)
		cmd.SetNode(n.Id)

		send = cmd
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(100)
		cmd.SetNode(n.Id)

		send = cmd
	default:
		return
	}

	n.connection.Write(send)
}

func (n *Node) Off() {
	var send interfaces.Encodable

	switch {
	case n.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinary()
		cmd.SetValue(false)
		cmd.SetNode(n.Id)

		send = cmd
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(0)
		cmd.SetNode(n.Id)

		send = cmd
	default:
		return
	}

	n.connection.Write(send)
}

func (n *Node) Level(value float64) {
	var send interfaces.Encodable

	switch {
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(value)
		cmd.SetNode(n.Id)

		send = cmd
	default:
		return
	}

	n.connection.Write(send)
}
