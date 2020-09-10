// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/moov-io/irs/pkg/config"
)

var (
	upperAlphanumericRegex = regexp.MustCompile(`[^ A-Z0-9!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	numericRegex           = regexp.MustCompile(`^[0-9]+$`)
	dateRegex              = regexp.MustCompile(`^(19|20)\d\d(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])`)
	yearRegex              = regexp.MustCompile(`((19|20)\d\d)`)
	emailRegex             = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	minPhoneNumberLength   = 10
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
		if err := isValidType(fieldName, spec, data); err != nil {
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
	if elm.Required == config.Expandable {
		return ""
	}

	if !data.IsValid() {
		return fillString(elm)
	}

	sizeStr := strconv.Itoa(elm.Length)

	switch elm.Type {
	case config.Alphanumeric, config.Email, config.Numeric, config.TelephoneNumber:
		return fmt.Sprintf("%-"+sizeStr+"s", data)
	case config.AlphanumericRightAlign:
		return fmt.Sprintf("%"+sizeStr+"s", data)
	case config.ZeroNumeric:
		if elm.Required == config.Omitted && data.Interface().(int) == 0 {
			return fmt.Sprintf("%"+sizeStr+"s", config.BlankString)
		}
		return fmt.Sprintf("%0"+sizeStr+"d", data)
	case config.Percent:
		if data.Interface().(int) == 100 {
			return fmt.Sprintf("%"+sizeStr+"s", config.BlankString)
		}
		return fmt.Sprintf("%0"+sizeStr+"d", data)
	case config.DateYear:
		if elm.Required == config.Omitted && data.Interface().(int) == 0 {
			return fmt.Sprintf("%"+sizeStr+"s", config.BlankString)
		}
		return fmt.Sprintf("%-"+sizeStr+"d", data)
	case config.Date:
		if datetime, ok := data.Interface().(time.Time); ok && !datetime.IsZero() {
			return datetime.Format(config.DateFormat)
		}
		return fmt.Sprintf("%"+sizeStr+"s", config.BlankString)
	}

	return fillString(elm)
}

// to validate fields of record
func Validate(r interface{}, spec map[string]config.SpecField, rType string) error {
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
					return NewErrFieldRequired(fieldName)
				}
				if fieldName == "RecordType" {
					if rType != fieldValue.String() {
						return NewErrRecordType(rType)
					}
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

// to get field
func GetField(from interface{}, fieldName string) (reflect.Value, error) {
	fields := reflect.ValueOf(from).Elem()
	if !fields.IsValid() {
		return fields, ErrValidField
	}
	field := fields.FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() {
		return field, ErrValidField
	}
	return field, nil
}

// to copy fields between struct instances
func CopyStruct(from interface{}, to interface{}) {
	fromFields := reflect.ValueOf(from).Elem()
	toFields := reflect.ValueOf(to).Elem()
	for i := 0; i < fromFields.NumField(); i++ {
		fieldName := fromFields.Type().Field(i).Name
		// skip local variable
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}
		fromField := fromFields.FieldByName(fieldName)
		toField := toFields.FieldByName(fieldName)
		if fromField.IsValid() && toField.CanSet() {
			toField.Set(fromField)
		}
	}
}

// to check amount codes
func CheckAvailableCodes(codes string, codeMap map[string]string) bool {
	codes = strings.TrimRight(codes, config.BlankString)
	codeList := strings.Split(codes, "")
	sort.Strings(codeList)
	if strings.Join(codeList, "") != codes {
		return false
	}

	repeated := map[string]int{}
	for i := 0; i < len(codeList); i++ {
		repeated[codeList[i]]++
	}

	for code, v := range repeated {
		if v > 1 {
			return false
		}
		if _, ok := codeMap[code]; !ok {
			return false
		}
	}

	return true
}

func isValidType(fieldName string, elm config.SpecField, data string) error {
	if elm.Required == config.Required {
		if isBlank(data) {
			return NewErrFieldRequired(fieldName)
		}
	}

	// for field with blank
	if isBlank(data) {
		return nil
	}

	switch elm.Type {
	case config.Alphanumeric, config.AlphanumericRightAlign:
		return isAlphanumeric(data)
	case config.Numeric, config.ZeroNumeric, config.Percent:
		return IsNumeric(data)
	case config.TelephoneNumber:
		if len(data) < minPhoneNumberLength {
			break
		}
		return IsNumeric(data)
	case config.Email:
		return isEmail(data)
	case config.DateYear:
		return isDateYear(data)
	case config.Date:
		if isBlank(data) {
			return nil
		}
		return isDate(data)
	}

	return NewErrValidValue(fieldName)
}

func isBlank(data string) bool {
	if len(data) == 0 {
		return true
	}
	return strings.Count(data, config.BlankString) == len(data)
}

func IsNumeric(data string) error {
	data = strings.TrimRight(data, config.BlankString)
	if !numericRegex.MatchString(data) {
		return ErrNumeric
	}
	return nil
}

func isAlphanumeric(data string) error {
	if upperAlphanumericRegex.MatchString(data) {
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

func isDate(data string) error {
	if !dateRegex.MatchString(data) {
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
	case config.Alphanumeric, config.AlphanumericRightAlign, config.Email, config.Numeric, config.TelephoneNumber:
		data = strings.TrimRight(data, config.BlankString)
		field.SetString(data)
		return nil
	case config.ZeroNumeric, config.DateYear:
		data = strings.Trim(data, config.BlankString)
		if len(data) == 0 {
			field.SetInt(0)
			return nil
		}
		value, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(value)
		return nil
	case config.Percent:
		data = strings.Trim(data, config.BlankString)
		if len(data) == 0 {
			field.SetInt(100)
			return nil
		}
		value, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(value)
		return nil
	case config.Date:
		if isBlank(data) {
			return nil
		}
		_date, err := time.Parse(config.DateFormat, data)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(_date))
		return nil
	}
	return ErrValidField
}

func validateFuncName(name string) string {
	return "Validate" + name
}

func ValidateCombinedFSCode(code int) error {
	if _, ok := config.ParticipateStateCodes[code]; !ok {
		return NewErrValidValue("combined federal state code")
	}
	return nil
}
