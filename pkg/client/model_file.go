/*
 * IRS API
 *
 * Package github.com/moov-io/irs implements a file reader and writer written in Go along with a HTTP API and CLI for creating, parsing, validating, and transforming IRS electronic Filing Information Returns Electronically (FIRE). FIRE operates on a byte(ASCII) level making it difficult to interface with JSON and CSV/TEXT file formats. | Input      | Output     | |------------|------------| | JSON       | JSON       | | ASCII FIRE | ASCII FIRE | |            | PDF Form   | |            | SQL        | 
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// File struct for File
type File struct {
	Transmitter TRecord `json:"transmitter"`
	PaymentPersons []PaymentPerson `json:"payment_persons,omitempty"`
	EndTransmitter FRecord `json:"end_transmitter"`
}
