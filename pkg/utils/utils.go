package utils

import (
	"bufio"
	"fmt"
	"os"
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
	case Alphanumeric, Email, Numeric, TelephoneNumber:
		return fmt.Sprintf("%-"+sizeStr+"s", data)
	case ZeroNumeric:
		return fmt.Sprintf("%0"+sizeStr+"d", data)
	case DateYear:
		return fmt.Sprintf("%-"+sizeStr+"d", data)
	}

	return fillString(elm)
}

// to validate fields of record
func Validate(r interface{}, spec map[string]SpecField) error {
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return ErrValidField
		}

		if spec, ok := spec[fieldName]; ok {
			if spec.Required == Required {
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

// File Read
func ReadFile(f *os.File) []byte {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return []byte(strings.Join(lines, ""))
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
	data = strings.TrimRight(data, BlankString)
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
	data = strings.TrimRight(data, BlankString)
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
		return nil
	case ZeroNumeric, DateYear:
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
