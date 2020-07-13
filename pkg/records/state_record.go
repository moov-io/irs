package records

type KRecord struct {
	// Required. Enter “K.”
	RecordType string `json:"record_type" validate:"required"`

	// Required. Enter the total number of “B” Records being
	// coded for this state. Right justify the information and fill
	// unused positions with zeros.
	NumberPayees int `json:"number_of_payees" validate:"required"`

	// Required. Accumulate totals of any payment amount fields
	// in the “B” Records for each state being reported into the
	// appropriate control total fields of the appropriate “K” Record.
	// Each payment amount must contain U.S. dollars and cents.
	// The right-most two positions represent cents in the payment
	// amount fields. Control totals must be right justified and fill
	// unused positions with zeros. All control total fields are
	// eighteen positions in length. Do not enter dollar signs,
	// commas, decimal points, or negative payments, except those
	// items that reflect a loss on Form 1099-B or 1099-OID.
	// Positive and negative amounts are indicated by placing a “+”
	// (plus) or “-” (minus) sign in the left-most position of the
	// payment amount field.
	ControlTotal1 int `json:"control_total_1" validate:"required"`
	ControlTotal2 int `json:"control_total_2" validate:"required"`
	ControlTotal3 int `json:"control_total_3" validate:"required"`
	ControlTotal4 int `json:"control_total_4" validate:"required"`
	ControlTotal5 int `json:"control_total_5" validate:"required"`
	ControlTotal6 int `json:"control_total_6" validate:"required"`
	ControlTotal7 int `json:"control_total_7" validate:"required"`
	ControlTotal8 int `json:"control_total_8" validate:"required"`
	ControlTotal9 int `json:"control_total_9" validate:"required"`
	ControlTotalA int `json:"control_total_A" validate:"required"`
	ControlTotalB int `json:"control_total_B" validate:"required"`
	ControlTotalC int `json:"control_total_C" validate:"required"`
	ControlTotalD int `json:"control_total_D" validate:"required"`
	ControlTotalE int `json:"control_total_E" validate:"required"`
	ControlTotalF int `json:"control_total_F" validate:"required"`
	ControlTotalG int `json:"control_total_G" validate:"required"`

	// Required. Enter the number of the record as it appears
	// within the file. The record sequence number for the “T”
	// Record will always be “1” (one), since it is the first record on
	// the file and the file can have only one “T” Record in a file.
	// Each record, thereafter, must be increased by one in
	// ascending numerical sequence, that is, 2, 3, 4, etc. Right
	// justify numbers with leading zeros in the field. For example,
	// the “T” Record sequence number would appear as
	// “00000001” in the field, the first “A” Record would be
	// “00000002,” the first “B” Record, “00000003,” the second “B”
	// Record, “00000004” and so on through the final record of the
	// file, the “F” Record.
	RecordSequenceNumber int `json:"record_sequence_number" validate:"required"`

	// Aggregate totals of the state income tax withheld field in the
	// Payee “B” Records. Otherwise, enter blanks. (This field is for
	// the convenience of filers.)
	StateIncomeTaxWithheldTotal int `json:"state_income_tax_withheld_total"`

	// Aggregate totals of the local income tax withheld field in the
	// Payee “B” Records. Otherwise, enter blanks. (This field is for
	// the convenience of filers.)
	LocalIncomeTaxWithheldTotal int `json:"local_income_tax_withheld_total"`

	// Required. Enter the CF/SF code assigned to the state which
	// is to receive the information.
	CombinedFSCode string `json:"combined_federal_state_code" validate:"required"`
}
