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
