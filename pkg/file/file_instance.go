package file

import (
	"bytes"
	"encoding/json"
	"sort"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/records"
	"github.com/moov-io/irs/pkg/utils"
)

// File contains the structures of irs file.
type fileInstance struct {
	Transmitter    records.Record   `json:"transmitter"`
	PaymentPersons []*paymentPerson `json:"payment_persons"`
	EndTransmitter records.Record   `json:"end_transmitter"`
}

// SetTCC set transmitter control code
func (f *fileInstance) SetTCC(code string) error {
	tRecord, _, err := f.getRecords()
	if err != nil {
		return err
	}
	if len(code) != 5 {
		return utils.ErrInvalidTCC
	}

	tRecord.TCC = code
	return nil
}

// TCC returns transmitter control code
func (f *fileInstance) TCC() (*string, error) {
	tRecord, _, err := f.getRecords()
	if err != nil {
		return nil, err
	}
	return &tRecord.TCC, nil
}

// Validate performs some checks on the file and returns an error if not Validated
func (f *fileInstance) Validate() error {
	err := f.validateRecords()
	if err != nil {
		return err
	}

	err = f.validateRecordSequenceNumber()
	if err != nil {
		return err
	}

	err = f.integrationCheck()
	if err != nil {
		return err
	}

	return nil
}

// Parse attempts to initialize a *File object assuming the input is valid raw data.
func (f *fileInstance) Parse(buf []byte) error {
	bufSize := len(buf)
	readPtr := 0
	if string(buf[readPtr]) != config.TRecordType || bufSize < config.RecordLength {
		return utils.ErrInvalidAscii
	}

	if f.Transmitter == nil {
		f.Transmitter = records.NewTRecord()
	}
	err := f.Transmitter.Parse(buf[readPtr : readPtr+config.RecordLength])
	if err != nil {
		return err
	}
	readPtr += config.RecordLength

	f.PaymentPersons = []*paymentPerson{}
	for string(buf[readPtr]) == config.ARecordType {
		currentPerson := &paymentPerson{}
		readSize, err := currentPerson.Parse(buf[readPtr:])
		if err != nil {
			return err
		}
		readPtr += readSize
		f.PaymentPersons = append(f.PaymentPersons, currentPerson)
	}

	if string(buf[readPtr]) != config.FRecordType || bufSize < readPtr+config.RecordLength {
		return utils.ErrInvalidAscii
	}

	if f.EndTransmitter == nil {
		f.EndTransmitter = records.NewFRecord()
	}
	err = f.EndTransmitter.Parse(buf[readPtr : readPtr+config.RecordLength])
	if err != nil {
		return err
	}

	return nil
}

// String writes the File struct to raw string.
func (f *fileInstance) Ascii() []byte {
	var buf bytes.Buffer

	if f.Transmitter != nil {
		buf.Grow(config.RecordLength)
		buf.Write(f.Transmitter.Ascii())
	}

	for _, person := range f.PaymentPersons {
		ascii := person.Ascii()
		buf.Grow(len(ascii))
		buf.Write(ascii)
	}

	if f.EndTransmitter != nil {
		buf.Grow(config.RecordLength)
		buf.Write(f.EndTransmitter.Ascii())
	}

	return buf.Bytes()
}

// UnmarshalJSON parses a JSON blob
func (f *fileInstance) UnmarshalJSON(data []byte) error {
	dummy := make(map[string]interface{})
	err := json.Unmarshal(data, &dummy)
	if err != nil {
		return err
	}

	for name, record := range dummy {
		buf, err := json.Marshal(record)
		if err != nil {
			return err
		}

		switch name {
		case "payment_persons":
			var list []interface{}
			err = json.Unmarshal(buf, &list)
			if err != nil {
				return nil
			}
			f.PaymentPersons = make([]*paymentPerson, 0)
			for _, data := range list {
				newRecord := &paymentPerson{}
				err := readJsonWithPerson(newRecord, data)
				if err != nil {
					return err
				}
				f.PaymentPersons = append(f.PaymentPersons, newRecord)
			}
		case "end_transmitter":
			if f.EndTransmitter == nil {
				f.EndTransmitter = records.NewFRecord()
			}
			err := readJsonWithRecord(f.EndTransmitter, record)
			if err != nil {
				return err
			}
		case "transmitter":
			if f.Transmitter == nil {
				f.Transmitter = records.NewTRecord()
			}
			err := readJsonWithRecord(f.Transmitter, record)
			if err != nil {
				return err
			}
		}
	}

	// sort by record sequence number
	if f.PaymentPersons != nil {
		sort.SliceStable(f.PaymentPersons, func(i, j int) bool {
			return f.PaymentPersons[i].SequenceNumber() < f.PaymentPersons[j].SequenceNumber()
		})
	}

	return nil
}

func (f *fileInstance) validateRecords() error {
	if f.Transmitter == nil || f.EndTransmitter == nil {
		return utils.ErrInvalidFile
	}

	err := f.Transmitter.Validate()
	if err != nil {
		return err
	}

	for _, person := range f.PaymentPersons {
		err = person.Validate()
		if err != nil {
			return err
		}
	}

	err = f.EndTransmitter.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (f *fileInstance) validateRecordSequenceNumber() error {
	sequenceNumber := 1
	if sequenceNumber != f.Transmitter.SequenceNumber() {
		return utils.NewErrRecordSequenceNumber(f.Transmitter.Type())
	}
	sequenceNumber++

	for _, person := range f.PaymentPersons {
		if sequenceNumber != person.SequenceNumber() {
			return utils.NewErrRecordSequenceNumber(person.Type())
		}
		sequenceNumber++

		for _, payee := range person.Payees {
			if sequenceNumber != payee.SequenceNumber() {
				return utils.NewErrRecordSequenceNumber(payee.Type())
			}
			sequenceNumber++
		}

		if sequenceNumber != person.EndPayer.SequenceNumber() {
			return utils.NewErrRecordSequenceNumber(person.EndPayer.Type())
		}
		sequenceNumber++

		for _, state := range person.States {
			if sequenceNumber != state.SequenceNumber() {
				return utils.NewErrRecordSequenceNumber(state.Type())
			}
			sequenceNumber++
		}
	}

	if sequenceNumber != f.EndTransmitter.SequenceNumber() {
		return utils.NewErrRecordSequenceNumber(f.EndTransmitter.Type())
	}

	return nil
}

func (f *fileInstance) integrationCheck() error {
	for _, person := range f.PaymentPersons {
		if err := person.integrationCheck(); err != nil {
			return err
		}
	}

	tRecord, fRecord, err := f.getRecords()
	if err != nil {
		return err
	}

	if fRecord.NumberPayerRecords != len(f.PaymentPersons) {
		return utils.ErrInvalidNumberPayers
	}

	numberPayees := 0
	if tRecord.TotalNumberPayees > 0 {
		numberPayees = tRecord.TotalNumberPayees
	} else if fRecord.NumberPayerRecords > 0 {
		numberPayees = fRecord.TotalNumberPayees
	}

	if numberPayees != f.getNumberPayees() {
		return utils.ErrInvalidNumberPayees
	}

	return nil
}

func (f *fileInstance) getRecords() (*records.TRecord, *records.FRecord, error) {
	if f.Transmitter == nil || f.EndTransmitter == nil {
		return nil, nil, utils.ErrInvalidFile
	}

	tRecord, ok := f.Transmitter.(*records.TRecord)
	if !ok {
		return nil, nil, utils.NewErrUnexpectedRecord("transmitter", f.Transmitter)
	}

	fRecord, ok := f.EndTransmitter.(*records.FRecord)
	if !ok {
		return nil, nil, utils.NewErrUnexpectedRecord("end of transmitter", f.EndTransmitter)
	}

	return tRecord, fRecord, nil
}

func (f *fileInstance) getNumberPayees() int {
	number := 0
	for _, person := range f.PaymentPersons {
		number += len(person.Payees)
	}
	return number
}
