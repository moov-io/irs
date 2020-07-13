package records

type FRecord struct {
	// Required. Enter “F.”
	RecordType string `json:"record_type" validate:"required"`

	// Enter the total number of Payer “A” Records in the entire file.
	// Right justify the information and fill unused positions with
	// zeros or enter all zeros.
	NumberPayerRecords int `json:"number_of_payer_records"`

	// If this total was entered in the “T” Record, this field may be
	// blank filled. Enter the total number of Payee “B” Records
	// reported in the file. Right justify the information and fill
	// unused positions with zeros.
	TotalNumberPayees int `json:"total_number_of_payees"`

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
	// Record, “00000004” and so on until the final record of the
	// file, the “F” Record.
	RecordSequenceNumber int `json:"record_sequence_number" validate:"required"`
}
