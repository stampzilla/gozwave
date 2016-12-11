package commands

import "github.com/stampzilla/gozwave/interfaces"

type MultiChannelEncap struct {
	msg      interfaces.Encodable
	endpoint int
}

func NewMultiChannelEncap(msg interfaces.Encodable, endpoint int) *MultiChannelEncap {
	return &MultiChannelEncap{
		msg:      msg,
		endpoint: endpoint,
	}
}

func (self *MultiChannelEncap) Encode() []byte {
	msg := self.msg.Encode()

	ret := []byte{
		msg[0],     // serialapi function
		msg[1],     // node id
		msg[2] + 4, // length
		MultiInstance,
		0x0d,                // MultiChannelCmd_Encap
		0x01,                // Length
		byte(self.endpoint), // Endpoint
	}

	ret = append(ret, msg[3:]...)

	return ret
}
