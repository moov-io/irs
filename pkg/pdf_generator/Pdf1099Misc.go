package pdf_generator

import (
	"fmt"
	"github.com/moov-io/irs/pkg/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"
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
	StateTax1     string
	StateTax2     string
	StateNo1      string
	StateNo2      string
	StateIncome1  string
	StateIncome2  string
	tempDir       string
}

var fdf1099MiscPatterns = map[string]string{
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

const (
	pdfConverter = "pdftk"
	specFDF      = "spec.fdf"
	templateFDF  = "template.fdf"
	templatePDF  = "template.pdf"
	applyFDF     = "apply.fdf"
	resultPDF    = "result.pdf"
	timeFormat   = "20060102150405"
)

func (p *Pdf1099Misc) getSpecFdf() ([]byte, error) {
	switch p.Type {
	case PdfMscCopyB, PdfMscCopyC, PdfNecCopyB, PdfNecCopyC:
	default:
		return nil, utils.ErrUnknownPdfTemplate
	}
	return ioutil.ReadFile(filepath.Join(p.Type, specFDF))
}

func (p *Pdf1099Misc) getTemplateFdf() ([]byte, error) {
	switch p.Type {
	case PdfMscCopyB, PdfMscCopyC, PdfNecCopyB, PdfNecCopyC:
	default:
		return nil, utils.ErrUnknownPdfTemplate
	}
	return ioutil.ReadFile(filepath.Join(p.Type, templateFDF))
}

func (p *Pdf1099Misc) getTemplateFile() (*string, error) {
	switch p.Type {
	case PdfMscCopyB, PdfMscCopyC, PdfNecCopyB, PdfNecCopyC:
	default:
		return nil, utils.ErrUnknownPdfTemplate
	}
	filePath := filepath.Join(p.Type, templatePDF)
	return &filePath, nil
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

		if pattern, ok := fdf1099MiscPatterns[fieldName]; ok {
			field := fields.FieldByName(fieldName)
			switch field.Type().String() {
			case "string":
				newFdf = strings.ReplaceAll(newFdf, pattern, field.String())
			case "bool":
				if field.Bool() {
					switch fieldName {
					case "VoID":
						newFdf = strings.ReplaceAll(newFdf, pattern, strings.ReplaceAll(pattern, "Off", "1"))
					case "Corrected":
						newFdf = strings.ReplaceAll(newFdf, pattern, strings.ReplaceAll(pattern, "Off", "2"))
					default:
						newFdf = strings.ReplaceAll(newFdf, pattern, strings.ReplaceAll(pattern, "Off", "Yes"))
					}
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

	err = ioutil.WriteFile(fileName, []byte(newFdf), os.ModePerm)
	if err != nil {
		return nil, err
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
	cmd := exec.Command(execFile, *template, "fill_form", fdfFile, "output", result)
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(result)
}

// Generate pdf file form Pdf1099Misc struct using pdftk
func (p *Pdf1099Misc) Generate() ([]byte, error) {
	t := time.Now()
	p.tempDir = "." + t.Format(timeFormat)
	err := os.Mkdir(p.tempDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	fdfFile := filepath.Join(p.tempDir, templateFDF)
	_, err = p.generateFDF(fdfFile)
	if err != nil {
		return nil, err
	}

	buf, err := p.generatePDF(fdfFile)
	if err != nil {
		return nil, err
	}

	err = os.RemoveAll(p.tempDir)
	return buf, err
}
