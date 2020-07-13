package records

// General record interface
type Record interface {
	Type() string
	TCC() string
	SetTCC(string) string
	TIN() string
	Parse(string) error
	String() string
	Validate() error
}
