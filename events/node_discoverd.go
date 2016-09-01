package events

import "github.com/stampzilla/gozwave/functions"

type NodeDiscoverd struct {
	Address byte

	functions.FuncGetNodeProtocolInfo
}
