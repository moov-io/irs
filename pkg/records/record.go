package records

// General record interface
type Record interface {
	Type() string
	SequenceNumber() int
	SetSequenceNumber(int)
	Parse([]byte) error
	Ascii() []byte
	Validate() error
}
