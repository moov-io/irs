// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pdf_generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/moov-io/irs/pkg/utils"
)

const (
	// 1099-MISC Copy B
	PdfMscCopyB = "1099msc_copy_b"
	// 1099-MISC Copy C
	PdfMscCopyC = "1099msc_copy_c"
	// 1099-MISC(NEC) Copy B
	PdfNecCopyB = "1099nec_copy_b"
	// 1099-MISC(NEC) Copy C
	PdfNecCopyC = "1099nec_copy_c"
)

// Pdf struct for 1099-MISC
type Pdf1099Misc struct {
	Type          string
	VoID          bool
	Corrected     bool
	Fatca         bool
	SecondTin     bool
	DirectSale    bool
	PayerInfo     string
	PayerTin      string
	RecipientTin  string
	RecipientName string
	Street        string
	City          string
	AccountNumber string
	Rents         int
	Royalties     int
	Other         int
	Federal       int
	Fishing       int
	Medical       int
	Substitute    int
	Crop          int
	Gross         int
	Section       int
	Excess        int
	Nonqualified  int
	Nonemployee   int
	StateTax1     int
	StateTax2     int
	StateIncome1  int
	StateIncome2  int
	StateNo1      string
	StateNo2      string
	tempDir       string
}

var fdf1099MiscPatternsCopyC = map[string]string{
	"VoID":          "/Off#?#/T (c2_1[0])",
	"Corrected":     "/Off#?#/T (c2_1[1])",
	"Fatca":         "/Off#?#/T (c2_2[0])",
	"SecondTin":     "/Off#?#/T (c2_3[0])",
	"DirectSale":    "/Off#?#/T (c2_4[0])",
	"PayerInfo":     "PAYER Information",
	"PayerTin":      "PAYER TIN",
	"RecipientTin":  "RECIP TIN",
	"RecipientName": "RECIPIENT Name",
	"Street":        "Street Address",
	"City":          "ZIP, Postal Code",
	"AccountNumber": "Account Number",
	"Rents":         "Rents",
	"Royalties":     "Royalties",
	"Other":         "Other Income",
	"Federal":       "Federal Income",
	"Fishing":       "Fishing",
	"Medical":       "Medical Health",
	"Substitute":    "Substitute",
	"Crop":          "Crop",
	"Gross":         "Gross",
	"Section":       "Section 409A",
	"Excess":        "Excess Golden",
	"Nonqualified":  "Nonqualified",
	"Nonemployee":   "Nonemployee",
	"StateTax1":     "State tax1",
	"StateTax2":     "State tax2",
	"StateNo1":      "State no1",
	"StateNo2":      "State no2",
	"StateIncome1":  "State income1",
	"StateIncome2":  "State income2",
}

var fdf1099MiscPatternsCopyB = map[string]string{
	"Corrected":     "/Off#?#/T (c2_1[0])",
	"Fatca":         "/Off#?#/T (c2_2[0])",
	"DirectSale":    "/Off#?#/T (c2_4[0])",
	"PayerInfo":     "PAYER Information",
	"PayerTin":      "PAYER TIN",
	"RecipientTin":  "RECIP TIN",
	"RecipientName": "RECIPIENT Name",
	"Street":        "Street Address",
	"City":          "ZIP, Postal Code",
	"AccountNumber": "Account Number",
	"Rents":         "Rents",
	"Royalties":     "Royalties",
	"Other":         "Other Income",
	"Federal":       "Federal Income",
	"Fishing":       "Fishing",
	"Medical":       "Medical Health",
	"Substitute":    "Substitute",
	"Crop":          "Crop",
	"Gross":         "Gross",
	"Section":       "Section 409A",
	"Excess":        "Excess Golden",
	"Nonqualified":  "Nonqualified",
	"Nonemployee":   "Nonemployee",
	"StateTax1":     "State tax1",
	"StateTax2":     "State tax2",
	"StateNo1":      "State no1",
	"StateNo2":      "State no2",
	"StateIncome1":  "State income1",
	"StateIncome2":  "State income2",
}

var (
	pdfConverter  = "pdftk"
	specFDF       = "spec.fdf"
	templateFDF   = "template.fdf"
	templatePDF   = "template.pdf"
	resultPDF     = "result.pdf"
	timeFormat    = "20060102150405"
	convertParam1 = "fill_form"
	convertParam2 = "output"
	convertParam3 = "cat"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func (p *Pdf1099Misc) getSpecFdf() ([]byte, error) {
	switch p.Type {
	case PdfMscCopyB, PdfMscCopyC, PdfNecCopyB, PdfNecCopyC:
	default:
		return nil, utils.ErrUnknownPdfTemplate
	}
	return ioutil.ReadFile(filepath.Join(basePath, p.Type, specFDF))
}

func (p *Pdf1099Misc) getTemplateFdf() ([]byte, error) {
	switch p.Type {
	case PdfMscCopyB, PdfMscCopyC, PdfNecCopyB, PdfNecCopyC:
	default:
		return nil, utils.ErrUnknownPdfTemplate
	}
	return ioutil.ReadFile(filepath.Join(basePath, p.Type, templateFDF))
}

func (p *Pdf1099Misc) getTemplateFile() (*string, error) {
	switch p.Type {
	case PdfMscCopyB, PdfMscCopyC, PdfNecCopyB, PdfNecCopyC:
	default:
		return nil, utils.ErrUnknownPdfTemplate
	}
	filePath := filepath.Join(basePath, p.Type, templatePDF)
	return &filePath, nil
}

func (p *Pdf1099Misc) getPattern() (map[string]string, error) {
	switch p.Type {
	case PdfMscCopyB, PdfNecCopyB:
		return fdf1099MiscPatternsCopyB, nil
	case PdfMscCopyC, PdfNecCopyC:
		return fdf1099MiscPatternsCopyC, nil
	}
	return nil, utils.ErrUnknownPdfTemplate
}

func (p *Pdf1099Misc) generateFDF(fileName string) ([]byte, error) {
	buf, err := p.getSpecFdf()
	if err != nil {
		return nil, err
	}
	spec := string(buf)
	newFdf := spec

	fields := reflect.ValueOf(p).Elem()
	if !fields.IsValid() {
		return nil, utils.ErrFdfGenerate
	}

	for i := 0; i < fields.NumField(); i++ {
		fieldName := fields.Type().Field(i).Name
		if !fields.IsValid() {
			return nil, utils.ErrFdfGenerate
		}

		pattern, err := p.getPattern()
		if err != nil {
			return nil, err
		}

		if pattern, ok := pattern[fieldName]; ok {
			field := fields.FieldByName(fieldName)
			switch field.Type().String() {
			case "string":
				newFdf = strings.ReplaceAll(newFdf, pattern, field.String())
			case "bool":
				if field.Bool() {
					newPattern := strings.ReplaceAll(pattern, "Off", "Yes")
					switch fieldName {
					case "VoID":
						newPattern = strings.ReplaceAll(pattern, "Off", "1")
					case "Corrected":
						newPattern = strings.ReplaceAll(pattern, "Off", "2")
					}
					if p.Type == PdfNecCopyB || p.Type == PdfNecCopyC {
						switch fieldName {
						case "Fatca", "SecondTin":
							newPattern = strings.ReplaceAll(pattern, "Off", "1")
						}
					}

					newFdf = strings.ReplaceAll(newFdf, pattern, newPattern)
				}
			case "int":
				if field.Int() > 0 {
					value := fmt.Sprintf("%.2f", float32(field.Int()/100))
					newFdf = strings.ReplaceAll(newFdf, pattern, value)
				} else {
					newFdf = strings.ReplaceAll(newFdf, pattern, "")
				}
			}
		}
	}
	newFdf = strings.ReplaceAll(newFdf, "#?#", "\n")

	if fileName != "" {
		err = ioutil.WriteFile(fileName, []byte(newFdf), os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return []byte(newFdf), nil
}

func (p *Pdf1099Misc) generatePDF(fdfFile string) ([]byte, error) {
	execFile, err := exec.LookPath(pdfConverter)
	if err != nil {
		return nil, err
	}

	template, err := p.getTemplateFile()
	if err != nil {
		return nil, err
	}

	result := filepath.Join(p.tempDir, resultPDF)
	cmd := exec.Command(execFile, *template, convertParam1, fdfFile, convertParam2, result)
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(result)
}

// Generate pdf file form Pdf1099Misc struct using pdftk
func GeneratePdf(p *Pdf1099Misc) ([]byte, error) {
	if p == nil {
		return nil, utils.ErrInvalidFile
	}

	randStr, err := utils.RandAlphanumericString(40)
	if err != nil {
		return nil, err
	}

	p.tempDir = filepath.Join(basePath, "."+randStr)
	err = os.Mkdir(p.tempDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	fdfFile := filepath.Join(p.tempDir, templateFDF)
	_, err = p.generateFDF(fdfFile)
	if err != nil {
		return returnWithRemoveTmp(p.tempDir, err)
	}

	buf, err := p.generatePDF(fdfFile)
	if err != nil {
		return returnWithRemoveTmp(p.tempDir, err)
	}

	err = os.RemoveAll(p.tempDir)
	return buf, err
}

// Generate pdf file form Pdf1099Misc struct using pdftk
func MergePdfs(files [][]byte) ([]byte, error) {
	randStr, err := utils.RandAlphanumericString(40)
	if err != nil {
		return nil, err
	}

	if len(files) < 1 {
		return nil, utils.ErrPdfMerge
	}

	if len(files) == 1 {
		return files[0], nil
	}

	execFile, err := exec.LookPath(pdfConverter)
	if err != nil {
		return nil, err
	}

	tempDir := filepath.Join(basePath, "."+randStr)
	err = os.Mkdir(tempDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	params := make([]string, 0)
	for index, f := range files {
		newFile := filepath.Join(tempDir, resultPDF+fmt.Sprintf("%v", index))
		err := ioutil.WriteFile(newFile, f, os.ModePerm)
		if err != nil {
			return returnWithRemoveTmp(tempDir, err)
		}
		params = append(params, newFile)
	}

	result := filepath.Join(tempDir, resultPDF)
	params = append(params, convertParam3)
	params = append(params, convertParam2)
	params = append(params, result)
	cmd := exec.Command(execFile, params...)
	err = cmd.Run()
	if err != nil {
		return returnWithRemoveTmp(tempDir, err)
	}

	buf, err := ioutil.ReadFile(result)
	if err != nil {
		return returnWithRemoveTmp(tempDir, err)
	}

	err = os.RemoveAll(tempDir)
	return buf, err
}

func returnWithRemoveTmp(tempDir string, err error) ([]byte, error){
	os.RemoveAll(tempDir)
	return nil, err
}
