package file

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/records"
	"github.com/moov-io/irs/pkg/utils"
)

//  paymentPerson identifies the person making payments
type paymentPerson struct {
	Payer    records.Record   `json:"payer"`
	Payees   []records.Record `json:"payees"`
	EndPayer records.Record   `json:"end_payer"`
	States   []records.Record `json:"states"`
}

// Type returns type of “Person” record
func (p *paymentPerson) Type() string {
	return "Person"
}

// Ascii returns fire ascii of “Person” record
func (p *paymentPerson) Ascii() []byte {
	var buf bytes.Buffer

	buf.Grow(config.RecordLength)
	buf.Write(p.Payer.Ascii())

	for _, payee := range p.Payees {
		buf.Grow(config.RecordLength)
		buf.Write(payee.Ascii())
	}

	buf.Grow(config.RecordLength)
	buf.Write(p.EndPayer.Ascii())

	for _, state := range p.States {
		buf.Grow(config.RecordLength)
		buf.Write(state.Ascii())
	}

	return buf.Bytes()
}

// Validate performs some checks on the record and returns an error if not Validated
func (p *paymentPerson) Validate() error {
	var err error
	if err = p.validateRecords(); err != nil {
		return err
	}

	if err = p.Payer.Validate(); err != nil {
		return err
	}

	for _, payee := range p.Payees {
		if err = payee.Validate(); err != nil {
			return err
		}
	}

	if err = p.EndPayer.Validate(); err != nil {
		return err
	}

	for _, state := range p.States {
		if err = state.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// SequenceNumber returns sequence number of the record
func (p *paymentPerson) SequenceNumber() int {
	if p.Payer == nil {
		return 0
	}
	return p.Payer.SequenceNumber()
}

// SequenceNumber set sequence number of the record
func (p *paymentPerson) SetSequenceNumber(int) {}

// Parse attempts to parse with raw data.
func (p *paymentPerson) Parse(buf []byte) (int, error) {
	bufSize := len(buf)
	readPtr := 0

	if string(buf[readPtr]) != config.ARecordType || bufSize < config.RecordLength {
		return readPtr, utils.ErrInvalidAscii
	}

	if p.Payer == nil {
		p.Payer = records.NewARecord()
	}
	err := p.Payer.Parse(buf[readPtr : readPtr+config.RecordLength])
	if err != nil {
		return readPtr, err
	}
	readPtr += config.RecordLength

	typeOfReturn, err := p.getTypeOfReturn()
	if err != nil {
		return readPtr, err
	}

	p.Payees = []records.Record{}
	for bufSize > readPtr && string(buf[readPtr]) == config.BRecordType {
		if bufSize < readPtr+config.RecordLength {
			return readPtr, utils.ErrInvalidAscii
		}

		newPayee := records.NewBRecord(typeOfReturn)
		if err = newPayee.Parse(buf[readPtr : readPtr+config.RecordLength]); err != nil {
			return readPtr, err
		}

		readPtr += config.RecordLength
		p.Payees = append(p.Payees, newPayee)
	}

	if bufSize <= readPtr {
		return readPtr, utils.ErrInvalidAscii
	}

	if string(buf[readPtr]) == config.CRecordType || bufSize < readPtr+config.RecordLength {
		if p.EndPayer == nil {
			p.EndPayer = records.NewCRecord()
		}
		err := p.EndPayer.Parse(buf[readPtr : readPtr+config.RecordLength])
		if err != nil {
			return readPtr, err
		}
		readPtr += config.RecordLength
	}

	p.States = []records.Record{}
	for bufSize > readPtr && string(buf[readPtr]) == config.KRecordType {
		if bufSize < readPtr+config.RecordLength {
			return readPtr, utils.ErrInvalidAscii
		}

		newState := records.NewKRecord()
		if err = newState.Parse(buf[readPtr : readPtr+config.RecordLength]); err != nil {
			return readPtr, err
		}

		readPtr += config.RecordLength
		p.States = append(p.States, newState)
	}

	return readPtr, nil
}

// UnmarshalJSON parses a JSON blob
func (p *paymentPerson) UnmarshalJSON(data []byte) error {
	dummy := make(map[string]interface{})
	err := json.Unmarshal(data, &dummy)
	if err != nil {
		return err
	}

	for name, record := range dummy {
		switch name {
		case "payer":
			p.Payer = records.NewARecord()
			err := readJsonWithRecord(p.Payer, record)
			if err != nil {
				return err
			}
		case "end_payer":
			p.EndPayer = records.NewCRecord()
			err := readJsonWithRecord(p.EndPayer, record)
			if err != nil {
				return err
			}
		}
	}

	typeOfReturn, err := p.getTypeOfReturn()
	if err != nil {
		return err
	}

	for name, record := range dummy {
		buf, err := json.Marshal(record)
		if err != nil {
			return err
		}

		switch name {
		case "payees":
			var list []interface{}
			err = json.Unmarshal(buf, &list)
			if err != nil {
				return nil
			}

			p.Payees = make([]records.Record, 0)
			for _, data := range list {
				newRecord := records.NewBRecord(typeOfReturn)
				err := readJsonWithRecord(newRecord, data)
				if err != nil {
					return err
				}
				p.Payees = append(p.Payees, newRecord)
			}
		case "states":
			var list []interface{}
			err = json.Unmarshal(buf, &list)
			if err != nil {
				return nil
			}
			p.States = make([]records.Record, 0)
			for _, data := range list {
				newRecord := records.NewKRecord()
				err := readJsonWithRecord(newRecord, data)
				if err != nil {
					return err
				}
				p.States = append(p.States, newRecord)
			}
		}
	}

	// sort by record sequence number
	if p.Payees != nil {
		sort.SliceStable(p.Payees, func(i, j int) bool {
			return p.Payees[i].SequenceNumber() < p.Payees[j].SequenceNumber()
		})
	}
	if p.States != nil {
		sort.SliceStable(p.States, func(i, j int) bool {
			return p.States[i].SequenceNumber() < p.States[j].SequenceNumber()
		})
	}

	return nil
}

func (p *paymentPerson) integrationCheck() error {
	if err := p.validateRecords(); err != nil {
		return err
	}
	_, cRecord, err := p.getRecords()
	if err != nil {
		return err
	}

	// 1. verify corrected return indicator
	existed := make(map[string]interface{})
	for _, payee := range p.Payees {
		bRecord, ok := payee.(*records.BRecord)
		if !ok {
			return utils.NewErrUnexpectedRecord("payee", payee)
		}
		indicator := bRecord.CorrectedReturnIndicator
		if len(indicator) == 0 {
			indicator = "N"
		}
		existed[indicator] = bRecord
		if len(existed) > 1 {
			return utils.ErrIncorrectReturnIndicator
		}
	}

	// 2. verify number of payees
	if len(p.Payees) != cRecord.NumberPayees {
		return utils.ErrInvalidNumberPayees
	}

	// 3. verify payment codes
	err = p.validatePaymentCodes()
	if err != nil {
		return err
	}

	// 4. verify combined federal/state code in K records
	existed = make(map[string]interface{})
	for _, state := range p.States {
		kRecord, ok := state.(*records.KRecord)
		if !ok {
			return utils.NewErrUnexpectedRecord("state", state)
		}
		_, ok = existed[kRecord.CombinedFederalStateCode]
		if ok {
			return utils.ErrDuplicatedFSCode
		}
		existed[kRecord.CombinedFederalStateCode] = kRecord
	}

	return nil
}

func (p *paymentPerson) validatePaymentCodes() error {
	aRecord, cRecord, err := p.getRecords()
	if err != nil {
		return err
	}
	amountCodes := strings.Split(aRecord.AmountCodes, "")

	// check amount codes between a record and b records
	existedCodes := make(map[string]bool)
	for _, payee := range p.Payees {
		bRecord, ok := payee.(*records.BRecord)
		if !ok {
			return utils.NewErrUnexpectedRecord("payee", payee)
		}
		for _, existed := range strings.Split(bRecord.PaymentCodes(), "") {
			existedCodes[existed] = true
		}
	}
	if len(existedCodes) != len(amountCodes) {
		return utils.ErrUnexpectedPaymentAmount
	}

	// check amount codes between a record and c record
	if cRecord.TotalCodes() != aRecord.AmountCodes {
		return utils.ErrUnexpectedTotalAmount
	}

	// check amounts between c record and b records
	for _, amountCode := range amountCodes {
		control, err := cRecord.ControlTotal(amountCode)
		if err != nil {
			return err
		}

		for _, payee := range p.Payees {
			bRecord, _ := payee.(*records.BRecord)
			amount, err := bRecord.PaymentAmount(amountCode)
			if err != nil {
				return err
			}
			control -= amount
		}

		if control != 0 {
			return utils.ErrInvalidTotalAmounts
		}
	}

	if p.States != nil {
		// check amount codes between a record and k records
		existedCodes = make(map[string]bool)
		for _, state := range p.States {
			kRecord, ok := state.(*records.KRecord)
			if !ok {
				return utils.NewErrUnexpectedRecord("state", state)
			}
			for _, existed := range strings.Split(kRecord.PaymentCodes(), "") {
				existedCodes[existed] = true
			}
		}
		if len(existedCodes) != len(amountCodes) {
			return utils.ErrUnexpectedTotalAmount
		}

		// check amounts between c record and k records
		for _, amountCode := range amountCodes {
			control, err := cRecord.ControlTotal(amountCode)
			if err != nil {
				return err
			}

			for _, state := range p.States {
				kRecord, _ := state.(*records.KRecord)
				amount, err := kRecord.ControlTotal(amountCode)
				if err != nil {
					return err
				}
				control -= amount
			}

			if control != 0 {
				return utils.ErrInvalidTotalAmounts
			}
		}
	}

	return nil
}

func (p *paymentPerson) validateRecords() error {
	if p.Payer == nil {
		return utils.ErrNonExistPayer
	}
	if p.EndPayer == nil {
		return utils.ErrNonExistEndPayer
	}
	if p.Payees == nil {
		return utils.ErrNonExistPayee
	}
	return nil
}

func (p *paymentPerson) getRecords() (*records.ARecord, *records.CRecord, error) {
	aRecord, ok := p.Payer.(*records.ARecord)
	if !ok {
		return nil, nil, utils.NewErrUnexpectedRecord("payer", p.Payer)
	}

	cRecord, ok := p.EndPayer.(*records.CRecord)
	if !ok {
		return nil, nil, utils.NewErrUnexpectedRecord("end of payer", p.EndPayer)
	}

	return aRecord, cRecord, nil
}

func (p *paymentPerson) getTypeOfReturn() (string, error) {
	typeOfReturn := ""
	if p.Payer == nil {
		return typeOfReturn, utils.ErrNonExistPayer
	}
	aRecord, ok := p.Payer.(*records.ARecord)
	if !ok {
		return typeOfReturn, utils.NewErrUnexpectedRecord("payer", p.Payer)
	}
	typeOfReturn, ok = config.TypeOfReturns[aRecord.TypeOfReturn]
	if !ok {
		return typeOfReturn, utils.ErrInvalidTypeOfReturn
	}
	return typeOfReturn, nil
}
