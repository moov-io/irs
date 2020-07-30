package file

import (
	"encoding/json"

	"github.com/moov-io/irs/pkg/records"
)

// General file interface
type File interface {
	Parse([]byte) error
	Ascii() []byte
	Validate() error
	SetTCC(string) error
	TCC() (*string, error)
}

// NewFile constructs a file template.
func NewFile() File {
	return &fileInstance{
		Transmitter:    records.NewTRecord(),
		EndTransmitter: records.NewFRecord(),
	}
}

// CreateFile attempts to parse raw metro2 file contents
func CreateFile(buf []byte) (File, error) {
	var err error
	f := NewFile()
	if json.Valid(buf) {
		err = json.Unmarshal(buf, f)
	} else {
		err = f.Parse(buf)
	}
	return f, err
}

func readJsonWithRecord(record records.Record, data interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, record)
	if err != nil {
		return err
	}
	return nil
}

func readJsonWithPerson(person *paymentPerson, data interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, person)
	if err != nil {
		return err
	}
	return nil
}
