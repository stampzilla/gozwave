package events

import "github.com/stampzilla/gozwave/functions"

type NodeDiscoverd struct {
	Address int

	functions.FuncGetNodeProtocolInfo
}
