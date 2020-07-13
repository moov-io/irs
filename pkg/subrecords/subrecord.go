package subrecords

// General subrecord interface
type SubRecord interface {
	Type() string
	Parse(string) error
	String() string
	Validate() error
}
