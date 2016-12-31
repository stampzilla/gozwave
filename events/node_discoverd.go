package events

import "github.com/stampzilla/gozwave/serialapi"

type NodeDiscoverd struct {
	Address int

	serialapi.FuncGetNodeProtocolInfo
}
