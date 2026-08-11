package file

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateFile_EmptyNoPanic(t *testing.T) {
	require.NotPanics(t, func() {
		f, err := CreateFile(nil)
		_ = err
		if f != nil {
			_ = f.Ascii()
			_ = f.Validate()
		}
	})
	require.NotPanics(t, func() {
		f, err := CreateFile([]byte{})
		_ = err
		if f != nil {
			_ = f.Ascii()
		}
	})
	require.NotPanics(t, func() {
		// JSON with missing person fields
		f, err := CreateFile([]byte(`{"PaymentPersons":[{}]}`))
		_ = err
		if f != nil {
			_ = f.Ascii()
		}
	})
}

// Truncated file ending mid C-record previously panicked in paymentPerson.Parse
// (slice bounds out of range). See corpus test/fuzz/testdata/fuzz/FuzzCreateFile/292f0f96f8173139
func TestCreateFile_TruncatedCRecordNoPanic(t *testing.T) {
	// T + A + two full B records + single trailing "C" byte (no full 750-byte C record)
	rec := func(typ byte) []byte {
		b := make([]byte, 750)
		b[0] = typ
		copy(b[1:], []byte("2017"))
		return b
	}
	input := make([]byte, 0, 3001)
	input = append(input, rec('T')...)
	input = append(input, rec('A')...)
	input = append(input, rec('B')...)
	input = append(input, rec('B')...)
	input = append(input, 'C') // truncated C

	require.NotPanics(t, func() {
		f, err := CreateFile(input)
		require.Error(t, err)
		_ = f
	})
}
