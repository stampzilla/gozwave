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

func (m *MultiChannelEncap) Encode() []byte {
	msg := m.msg

	ret := []byte{
		msg[0],     // serialapi function
		msg[1],     // node id
		msg[2] + 4, // length
		MultiInstance,
		0x0d,             // MultiChannelCmd_Encap
		0x01,             // Length
		byte(m.endpoint), // Endpoint
	}

	ret = append(ret, msg[3:]...)

	return ret
}
