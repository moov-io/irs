/*
 * IRS API
 *
 * Package github.com/moov-io/irs implements a file reader and writer written in Go along with a HTTP API and CLI for creating, parsing, validating, and transforming IRS electronic Filing Information Returns Electronically (FIRE). FIRE operates on a byte(ASCII) level making it difficult to interface with JSON and CSV/TEXT file formats.  | Input      | Output     |  |------------|------------|  | JSON       | JSON       |  | ASCII FIRE | ASCII FIRE |  |            | PDF Form   |  |            | SQL        |
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"time"
)

// BRecordWith1098C struct for BRecordWith1098C
type BRecordWith1098C struct {
	RecordType                           string    `json:"record_type"`
	PaymentYear                          int32     `json:"payment_year"`
	CorrectedReturnIndicator             string    `json:"corrected_return_indicator,omitempty"`
	PayeesNameControl                    string    `json:"payees_name_control,omitempty"`
	TypeOfTin                            string    `json:"type_of_tin,omitempty"`
	PayeesTin                            string    `json:"payees_tin"`
	PayersAccountNumberForPayee          string    `json:"payers_account_number_for_payee,omitempty"`
	PayersOfficeCode                     string    `json:"payers_office_code,omitempty"`
	PaymentAmount1                       int32     `json:"payment_amount_1,omitempty"`
	PaymentAmount2                       int32     `json:"payment_amount_2,omitempty"`
	PaymentAmount3                       int32     `json:"payment_amount_3,omitempty"`
	PaymentAmount4                       int32     `json:"payment_amount_4,omitempty"`
	PaymentAmount5                       int32     `json:"payment_amount_5,omitempty"`
	PaymentAmount6                       int32     `json:"payment_amount_6,omitempty"`
	PaymentAmount7                       int32     `json:"payment_amount_7,omitempty"`
	PaymentAmount8                       int32     `json:"payment_amount_8,omitempty"`
	PaymentAmount9                       int32     `json:"payment_amount_9,omitempty"`
	PaymentAmountA                       int32     `json:"payment_amount_A,omitempty"`
	PaymentAmountB                       int32     `json:"payment_amount_B,omitempty"`
	PaymentAmountC                       int32     `json:"payment_amount_C,omitempty"`
	PaymentAmountD                       int32     `json:"payment_amount_D,omitempty"`
	PaymentAmountE                       int32     `json:"payment_amount_E,omitempty"`
	PaymentAmountF                       int32     `json:"payment_amount_F,omitempty"`
	PaymentAmountG                       int32     `json:"payment_amount_G,omitempty"`
	PaymentAmountH                       int32     `json:"payment_amount_H,omitempty"`
	PaymentAmountJ                       int32     `json:"payment_amount_J,omitempty"`
	ForeignCountryIndicator              string    `json:"foreign_country_indicator,omitempty"`
	FirstPayeeNameLine                   string    `json:"first_payee_name_line"`
	SecondPayeeNameLine                  string    `json:"second_payee_name_line,omitempty"`
	PayeeMailingAddress                  string    `json:"payee_mailing_address"`
	PayeeCity                            string    `json:"payee_city"`
	PayeeState                           string    `json:"payee_state"`
	PayeeZipCode                         string    `json:"payee_zip_code"`
	RecordSequenceNumber                 int32     `json:"record_sequence_number"`
	TransactionIndicator                 string    `json:"transaction_indicator,omitempty"`
	TransferAfterImprovementsIndicator   string    `json:"transfer_after_improvements_indicator,omitempty"`
	TransferMarketValueIndicator         string    `json:"transfer_market_value_indicator,omitempty"`
	Year                                 int32     `json:"year,omitempty"`
	Make                                 string    `json:"make,omitempty"`
	Model                                string    `json:"model,omitempty"`
	VehicleIdentificationNumber          string    `json:"vehicle_identification_number,omitempty"`
	VehicleDescription                   string    `json:"vehicle_description,omitempty"`
	DateContribution                     time.Time `json:"date_contribution,omitempty"`
	DoneeIndicator                       string    `json:"donee_indicator,omitempty"`
	IntangibleReligiousBenefitsIndicator string    `json:"intangible_religious_benefits_indicator,omitempty"`
	DeductionLessIndicator               string    `json:"deduction_less_indicator,omitempty"`
	GoodsServices                        string    `json:"goods_services,omitempty"`
	DateSale                             time.Time `json:"date_sale,omitempty"`
	SpecialDataEntries                   string    `json:"special_data_entries,omitempty"`
}
