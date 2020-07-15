package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	. "github.com/moov-io/irs/pkg/config"
)

var (
	alphanumericRegex    = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	numericRegex         = regexp.MustCompile(`^[0-9]+$`)
	yearRegex            = regexp.MustCompile(`((19|20)\d\d)`)
	emailRegex           = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	minPhoneNumberLength = 10
)

// parse field with string
func ParseValue(fields reflect.Value, spec map[string]SpecField, record string) error {
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}

		field := fields.FieldByName(fieldName)
		spec, ok := spec[fieldName]
		if !ok || !field.IsValid() || !field.CanSet() {
			return ErrValidField
		}

		if len(record) < spec.Start+spec.Length {
			return ErrShortRecord
		}

		data := record[spec.Start : spec.Start+spec.Length]
		if err := isValidType(spec, data); err != nil {
			return err
		}

		if err := parseValue(spec, field, data); err != nil {
			return err
		}
	}
	return nil
}

// to string from field
func ToString(elm SpecField, data reflect.Value) string {
	if !data.IsValid() {
		return fillString(elm)
	}

	sizeStr := strconv.Itoa(elm.Length)
	switch elm.Type {
	case Alphanumeric, Email:
		return fmt.Sprintf("%-"+sizeStr+"s", data)
	case ZeroNumeric:
		return fmt.Sprintf("%0"+sizeStr+"d", data)
	case TelephoneNumber, DateYear, Numeric:
		return fmt.Sprintf("%-"+sizeStr+"d", data)
	}

	return fillString(elm)
}

func isValidType(elm SpecField, data string) error {
	if elm.Required == Required {
		if isBlank(data) {
			return ErrFieldRequired
		}
	}

	// for field with blank
	if isBlank(data) {
		return nil
	}

	switch elm.Type {
	case Alphanumeric:
		return isAlphanumeric(data)
	case Numeric, ZeroNumeric:
		return isNumeric(data)
	case TelephoneNumber:
		if len(data) < minPhoneNumberLength {
			break
		}
		return isNumeric(data)
	case Email:
		return isEmail(data)
	case DateYear:
		return isDateYear(data)
	}

	return ErrValidField
}

func isBlank(data string) bool {
	if len(data) == 0 {
		return true
	}
	return false
}

func isNumeric(data string) error {
	if !numericRegex.MatchString(data) {
		return ErrNumeric
	}
	return nil
}

func isAlphanumeric(data string) error {
	if alphanumericRegex.MatchString(data) {
		return ErrNonAlphanumeric
	}
	return nil
}

func isDateYear(data string) error {
	if yearRegex.MatchString(data) {
		return ErrValidDate
	}
	return nil
}

func isEmail(data string) error {
	if !emailRegex.MatchString(data) {
		return ErrEmail
	}
	return nil
}

func fillString(elm SpecField) string {
	if elm.Type == ZeroNumeric {
		return strings.Repeat(ZeroString, elm.Length)
	}
	return strings.Repeat(BlankString, elm.Length)
}

func parseValue(elm SpecField, field reflect.Value, data string) error {
	switch elm.Type {
	case Alphanumeric, Email, Numeric, TelephoneNumber:
		data = strings.TrimRight(data, BlankString)
		field.SetString(data)
	case ZeroNumeric, DateYear:
		value, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(value)
	}
	return ErrValidField
}
