package commands

type SwitchBinaryReport struct {
	*report
	node   byte
	Status bool
	data   []byte
}

func NewSwitchBinaryReport(data []byte) *SwitchBinaryReport {
	if len(data) < 1 {
		data = []byte{0x00}
	}

	return &SwitchBinaryReport{
		Status: data[0] == 0xFF,
		data:   data,
	}
}
