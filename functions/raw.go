package functions

// Raw binary data type
type Raw []byte

// NewRaw creates a new Raw
func NewRaw(b []byte) Raw {
	return Raw(b)
}

// Encode make sure we implement the Encodable interface
func (r Raw) Encode() []byte {
	return r
}
