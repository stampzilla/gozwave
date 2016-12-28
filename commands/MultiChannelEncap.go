package commands

type MultiChannelEncap struct {
	msg      []byte
	endpoint int
}

func NewMultiChannelEncap(msg []byte, endpoint int) *MultiChannelEncap {
	return &MultiChannelEncap{
		msg:      msg,
		endpoint: endpoint,
	}
}

func (self *MultiChannelEncap) Encode() []byte {
	msg := self.msg

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
