package subrecords

// General subrecord interface
type SubRecord interface {
	Type() string
	Parse([]byte) error
	Ascii() []byte
	Validate() error
}
