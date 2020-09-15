// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/moov-io/irs/pkg/config"
	PDF "github.com/moov-io/irs/pkg/pdf_generator"
	"github.com/moov-io/irs/pkg/records"
	"github.com/moov-io/irs/pkg/utils"
)

//  paymentPerson identifies the person making payments
type paymentPerson struct {
	Payer    records.Record   `json:"payer"`
	Payees   []records.Record `json:"payees"`
	EndPayer records.Record   `json:"end_payer"`
	States   []records.Record `json:"states,omitempty"`
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

// Ascii returns pdf buffer of “Person” record
func (p *paymentPerson) Pdf() ([]byte, error) {
	if p.Payer == nil {
		return nil, utils.ErrNonExistPayer
	}

	payer, ok := p.Payer.(*records.ARecord)
	if !ok {
		return nil, utils.ErrNonExistPayer
	}

	returnType := config.TypeOfReturns[payer.TypeOfReturn]
	pdfType := PDF.PdfMscCopyB
	switch returnType {
	case config.Sub1099MiscType:
		if payer.CombinedFSFilingProgram != config.FSFilingProgramApproved {
			pdfType = PDF.PdfMscCopyC
		}
	case config.Sub1099NecType:
		pdfType = PDF.PdfNecCopyB
		if payer.CombinedFSFilingProgram != config.FSFilingProgramApproved {
			pdfType = PDF.PdfNecCopyC
		}
	default:
		return nil, utils.ErrUnsupportedPdf
	}

	pdf := PDF.Pdf1099Misc{Type: pdfType}
	err := p.fillingPdfInfoMisc(&pdf)
	if err != nil {
		fmt.Println("error 4")
		return nil, err
	}

	return PDF.GeneratePdf(&pdf)
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

		newPayee, err := records.NewBRecord(typeOfReturn)
		if err != nil {
			return readPtr, err
		}
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
				newRecord, err := records.NewBRecord(typeOfReturn)
				if err != nil {
					return err
				}
				err = readJsonWithRecord(newRecord, data)
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
	existedIndicator := ""
	for _, payee := range p.Payees {
		bRecord, ok := payee.(*records.BRecord)
		if !ok {
			return utils.NewErrUnexpectedRecord("payee", payee)
		}
		indicator := bRecord.CorrectedReturnIndicator
		if len(indicator) == 0 {
			indicator = "N"
		}

		if existedIndicator == "" {
			existedIndicator = indicator
		}

		if existedIndicator != indicator {
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

	// 4. verify  CF/SF code
	err = p.validateFSCodes()
	if err != nil {
		return err
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

	if p.States != nil && len(p.States) > 0 {
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

func (p *paymentPerson) validateFSCodes() error {
	existed := make(map[string]interface{})
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

	aRecord, _, err := p.getRecords()
	if err != nil {
		return err
	}

	if aRecord.CombinedFSFilingProgram != config.FSFilingProgramApproved {
		return nil
	}
	if len(p.Payees) == 0 || len(p.States) == 0 {
		return utils.ErrCFSFProgram
	}

	payeeCodes := make(map[string]bool)
	for _, payee := range p.Payees {
		bRecord, ok := payee.(*records.BRecord)
		if !ok {
			return utils.NewErrUnexpectedRecord("payee", payee)
		}
		code, exited := config.ParticipateStateCodes[bRecord.FederalState()]
		if !exited {
			return utils.NewErrValidValue("combined federal state code")
		}
		payeeCodes[code] = true
	}

	stateCodes := make(map[string]bool)
	for _, state := range p.States {
		kRecord, ok := state.(*records.KRecord)
		if !ok {
			return utils.NewErrUnexpectedRecord("state", state)
		}
		code, exited := config.StateAbbreviationCodes[kRecord.CombinedFederalStateCode]
		if !exited {
			return utils.NewErrValidValue("combined federal state code")
		}
		stateCodes[code] = true
	}

	if !eq(payeeCodes, stateCodes) {
		return utils.ErrCFSFState
	}

	return nil
}

func (p *paymentPerson) fillingPdfInfoMisc(pdf *PDF.Pdf1099Misc) error {
	payer, cRecord, err := p.getRecords()
	if err != nil {
		return err
	}
	payee, ok := p.Payees[0].(*records.BRecord)
	if !ok {
		return utils.ErrNonExistPayee
	}

	pdf.Corrected = true
	err = fillPayer(pdf, payer, cRecord)
	if err != nil {
		return err
	}

	return fillRecipient(pdf, payee)
}

func eq(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}

	return true
}

func fillAmounts(amountCodes []string, pdf *PDF.Pdf1099Misc, cRecord *records.CRecord) error {
	for _, amountCode := range amountCodes {
		control, err := cRecord.ControlTotal(amountCode)
		if err != nil {
			return err
		}
		switch amountCode {
		case "1":
			pdf.Rents = control
			pdf.Nonemployee = control
		case "2":
			pdf.Royalties = control
		case "3":
			pdf.Other = control
		case "4":
			pdf.Federal = control
		case "5":
			pdf.Fishing = control
		case "6":
			pdf.Medical = control
		case "8":
			pdf.Substitute = control
		case "A":
			pdf.Crop = control
		case "B":
			pdf.Excess = control
		case "C":
			pdf.Gross = control
		case "D":
			pdf.Section = control
		case "E":
			pdf.Nonqualified = control
		}
	}
	return nil
}

func fillPayer(pdf *PDF.Pdf1099Misc, payer *records.ARecord, cRecord *records.CRecord) error {
	pdf.PayerTin = payer.TIN
	info := make([]string, 0)
	name := make([]string, 0)
	if len(payer.FirstPayerNameLine) > 0 {
		name = append(name, payer.FirstPayerNameLine)
	}
	if len(payer.SecondPayerNameLine) > 0 {
		name = append(name, payer.SecondPayerNameLine)
	}
	if len(name) > 0 {
		info = append(info, strings.Join(name, " "))
	}
	if len(payer.PayerShippingAddress) > 0 {
		info = append(info, payer.PayerShippingAddress)
	}
	if len(payer.PayerCity) > 0 {
		info = append(info, payer.PayerCity)
	}
	if len(payer.PayerState) > 0 {
		info = append(info, payer.PayerState)
	}
	if len(payer.PayerZipCode) > 0 {
		info = append(info, payer.PayerZipCode)
	}
	if len(payer.PayerTelephoneNumber) > 0 {
		info = append(info, payer.PayerTelephoneNumber)
	}
	pdf.PayerInfo = strings.Join(info, "\r")

	amountCodes := strings.Split(payer.AmountCodes, "")
	err := fillAmounts(amountCodes, pdf, cRecord)
	if err != nil {
		return err
	}
	return nil
}

func fillRecipient(pdf *PDF.Pdf1099Misc, payee *records.BRecord) error {
	pdf.AccountNumber = payee.PayerAccountNumber
	pdf.Street = payee.PayeeMailingAddress
	pdf.RecipientTin = payee.TIN
	info := make([]string, 0)
	if len(payee.PayeeCity) > 0 {
		info = append(info, payee.PayeeCity)
	}
	if len(payee.PayeeState) > 0 {
		info = append(info, payee.PayeeState)
	}
	if len(payee.PayeeZipCode) > 0 {
		info = append(info, payee.PayeeZipCode)
	}
	pdf.City = strings.Join(info, ",")
	name := make([]string, 0)
	if len(payee.FirstPayeeNameLine) > 0 {
		name = append(name, payee.FirstPayeeNameLine)
	}
	if len(payee.SecondPayeeNameLine) > 0 {
		name = append(name, payee.SecondPayeeNameLine)
	}
	pdf.RecipientName = strings.Join(name, " ")

	fatca, err := payee.Fatca()
	if err != nil {
		return err
	}
	if fatca != nil && *fatca == config.FatcaFilingRequirementIndicator {
		pdf.Fatca = true
	}
	tin, err := payee.SecondTIN()
	if err != nil {
		return err
	}
	if tin != nil && *tin == config.SecondTINNotice {
		pdf.SecondTin = true
	}
	sale, err := payee.SecondTIN()
	if err != nil {
		return err
	}
	if sale != nil && *sale == config.DirectSalesIndicator {
		pdf.DirectSale = true
	}

	pdf.StateTax1, pdf.StateIncome1, err = payee.IncomeTax()
	if err != nil {
		return err
	}
	if payee.FederalState() > 0 {
		pdf.StateNo1 = fmt.Sprintf("%02d", payee.FederalState())
	}
	return nil
}
