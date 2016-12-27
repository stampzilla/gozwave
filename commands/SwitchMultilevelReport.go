package commands

type SwitchMultilevelReport struct {
	*report
	node  byte
	Level byte
	data  []byte
}

func NewSwitchMultilevelReport(data []byte) *SwitchMultilevelReport {
	if len(data) < 1 {
		data = []byte{0x00}
	}

	return &SwitchMultilevelReport{
		Level: data[0],
		data:  data,
	}
}
