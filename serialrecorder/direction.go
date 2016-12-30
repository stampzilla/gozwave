package serialrecorder

// Direction indicates if its a read or write row
type Direction bool

// DirectionRead is a read
const DirectionRead Direction = false

// DirectionWrite is a write
const DirectionWrite Direction = true

// IsRead check if direction was read
func (d Direction) IsRead() bool {
	return !bool(d)
}

// IsWrite check if direction was write
func (d Direction) IsWrite() bool {
	return bool(d)
}

func (d Direction) String() string {
	if d.IsRead() {
		return "read"
	}

	return "write"
}
