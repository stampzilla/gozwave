package commands

type Report interface {
	SetNode(byte)
}

type report struct {
	node byte
}

func (r *report) SetNode(n byte) {
	if r == nil {
		r = &report{}
		//r = nr
	}

	r.node = n
}
