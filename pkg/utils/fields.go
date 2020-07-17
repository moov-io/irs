package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/moov-io/irs/pkg/config"
)

var (
	alphanumericRegex    = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	numericRegex         = regexp.MustCompile(`^[0-9]+$`)
	yearRegex            = regexp.MustCompile(`((19|20)\d\d)`)
	emailRegex           = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	minPhoneNumberLength = 10
)

// parse field with string
func ParseValue(fields reflect.Value, spec map[string]config.SpecField, record string) error {
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
func ToString(elm config.SpecField, data reflect.Value) string {
	if !data.IsValid() {
		return fillString(elm)
	}

	sizeStr := strconv.Itoa(elm.Length)
	switch elm.Type {
	case config.Alphanumeric, config.Email, config.Numeric, config.TelephoneNumber:
		return fmt.Sprintf("%-"+sizeStr+"s", data)
	case config.ZeroNumeric:
		return fmt.Sprintf("%0"+sizeStr+"d", data)
	case config.DateYear:
		return fmt.Sprintf("%-"+sizeStr+"d", data)
	}

	return fillString(elm)
}

// to validate fields of record
func Validate(r interface{}, spec map[string]config.SpecField) error {
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return ErrValidField
		}

		if spec, ok := spec[fieldName]; ok {
			if spec.Required == config.Required {
				fieldValue := fields.FieldByName(fieldName)
				if fieldValue.IsZero() {
					return ErrFieldRequired
				}
			}
		}

		funcName := validateFuncName(fieldName)
		method := reflect.ValueOf(r).MethodByName(funcName)
		if method.IsValid() {
			response := method.Call(nil)
			if len(response) == 0 {
				continue
			}

			err := method.Call(nil)[0]
			if !err.IsNil() {
				return err.Interface().(error)
			}
		}
	}

	return nil
}

func isValidType(elm config.SpecField, data string) error {
	if elm.Required == config.Required {
		if isBlank(data) {
			return ErrFieldRequired
		}
	}

	// for field with blank
	if isBlank(data) {
		return nil
	}

	switch elm.Type {
	case config.Alphanumeric:
		return isAlphanumeric(data)
	case config.Numeric, config.ZeroNumeric:
		return isNumeric(data)
	case config.TelephoneNumber:
		if len(data) < minPhoneNumberLength {
			break
		}
		return isNumeric(data)
	case config.Email:
		return isEmail(data)
	case config.DateYear:
		return isDateYear(data)
	}

	return ErrValidField
}

func isBlank(data string) bool {
	if len(data) == 0 {
		return true
	}
	return strings.Count(data, config.BlankString) == len(data)
}

func isNumeric(data string) error {
	data = strings.TrimRight(data, config.BlankString)
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
	if !yearRegex.MatchString(data) {
		return ErrValidDate
	}
	return nil
}

func isEmail(data string) error {
	data = strings.TrimRight(data, config.BlankString)
	if !emailRegex.MatchString(data) {
		return ErrEmail
	}
	return nil
}

func fillString(elm config.SpecField) string {
	if elm.Type == config.ZeroNumeric {
		return strings.Repeat(config.ZeroString, elm.Length)
	}
	return strings.Repeat(config.BlankString, elm.Length)
}

func parseValue(elm config.SpecField, field reflect.Value, data string) error {
	switch elm.Type {
	case config.Alphanumeric, config.Email, config.Numeric, config.TelephoneNumber:
		data = strings.TrimRight(data, config.BlankString)
		field.SetString(data)
		return nil
	case config.ZeroNumeric, config.DateYear:
		value, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(value)
		return nil
	}
	return ErrValidField
}

func validateFuncName(name string) string {
	return "Validate" + name
}
