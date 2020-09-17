// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pdf_generator

import (
	"os"
	"strings"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&PdfTest{})

// Pdf test
type PdfTest struct{}

func (t *PdfTest) SetUpSuite(c *check.C) {}

func (t *PdfTest) SetUpTest(c *check.C) {
	pdfConverter = "pdftk"
	convertParam1 = "fill_form"
	convertParam3 = "cat"
}

func (t *PdfTest) TestPdfWithMscCopyB(c *check.C) {
	pdf := Pdf1099Misc{Type: PdfMscCopyB}
	f, err := GeneratePdf(&pdf)
	c.Assert(err, check.IsNil)
	files := make([][]byte, 0)
	_, err = MergePdfs(files)
	c.Assert(err, check.NotNil)
	files = append(files, f)
	_, err = MergePdfs(files)
	c.Assert(err, check.IsNil)
	files = append(files, f)
	_, err = MergePdfs(files)
	c.Assert(err, check.IsNil)

	templateFdf, err := pdf.getTemplateFdf()
	c.Assert(err, check.IsNil)
	newFdf, err := pdf.generateFDF("")
	c.Assert(err, check.IsNil)
	c.Assert(strings.ReplaceAll(string(templateFdf), "\r", ""),
		check.Equals, strings.ReplaceAll(string(newFdf), "\r", ""))
}

func (t *PdfTest) TestPdfWithMscCopyC(c *check.C) {
	pdf := Pdf1099Misc{Type: PdfMscCopyC}
	_, err := GeneratePdf(&pdf)
	c.Assert(err, check.IsNil)
	templateFdf, err := pdf.getTemplateFdf()
	c.Assert(err, check.IsNil)
	newFdf, err := pdf.generateFDF("")
	c.Assert(err, check.IsNil)
	c.Assert(strings.ReplaceAll(string(templateFdf), "\r", ""),
		check.Equals, strings.ReplaceAll(string(newFdf), "\r", ""))
}

func (t *PdfTest) TestPdfWithNecCopyB(c *check.C) {
	pdf := Pdf1099Misc{Type: PdfNecCopyB}
	_, err := GeneratePdf(&pdf)
	c.Assert(err, check.IsNil)
	templateFdf, err := pdf.getTemplateFdf()
	c.Assert(err, check.IsNil)
	newFdf, err := pdf.generateFDF("")
	c.Assert(err, check.IsNil)
	c.Assert(strings.ReplaceAll(string(templateFdf), "\r", ""),
		check.Equals, strings.ReplaceAll(string(newFdf), "\r", ""))
}

func (t *PdfTest) TestPdfWithNecCopyC(c *check.C) {
	pdf := Pdf1099Misc{Type: PdfNecCopyC}
	templateFdf, err := pdf.getTemplateFdf()
	c.Assert(err, check.IsNil)
	newFdf, err := pdf.generateFDF("")
	c.Assert(err, check.IsNil)
	c.Assert(strings.ReplaceAll(string(templateFdf), "\r", ""),
		check.Equals, strings.ReplaceAll(string(newFdf), "\r", ""))
	pdf = Pdf1099Misc{Type: PdfNecCopyC, VoID: true, Corrected: true, Fatca: true, Section: 1000}
	_, err = GeneratePdf(&pdf)
	c.Assert(err, check.IsNil)
}

func (t *PdfTest) TestPdfWithUnknownTemplate(c *check.C) {
	_, err := GeneratePdf(nil)
	c.Assert(err, check.NotNil)
	pdf := Pdf1099Misc{Type: "Unknown"}
	_, err = GeneratePdf(&pdf)
	c.Assert(err, check.NotNil)
	_, err = pdf.getSpecFdf()
	c.Assert(err, check.NotNil)
	_, err = pdf.getTemplateFdf()
	c.Assert(err, check.NotNil)
	_, err = pdf.getTemplateFile()
	c.Assert(err, check.NotNil)
	_, err = pdf.generateFDF("")
	c.Assert(err, check.NotNil)
	_, err = GeneratePdf(&pdf)
	c.Assert(err, check.NotNil)
	os.RemoveAll(pdf.tempDir)
}

func (t *PdfTest) TestPdfWithErrorParam(c *check.C) {
	pdf := Pdf1099Misc{Type: PdfNecCopyC}
	convertParam1 = "Unknown"
	_, err := pdf.generatePDF("")
	c.Assert(err, check.NotNil)
	pdfConverter = "Unknown"
	_, err = pdf.generatePDF("")
	c.Assert(err, check.NotNil)
	os.RemoveAll(pdf.tempDir)
}

func (t *PdfTest) TestPdfMergeWithUnknownParam(c *check.C) {
	pdf := Pdf1099Misc{Type: PdfMscCopyB}
	f, err := GeneratePdf(&pdf)
	c.Assert(err, check.IsNil)
	files := make([][]byte, 0)
	_, err = MergePdfs(files)
	c.Assert(err, check.NotNil)
	files = append(files, f)
	_, err = MergePdfs(files)
	c.Assert(err, check.IsNil)
	convertParam3 = "fill_form"
	files = append(files, f)
	_, err = MergePdfs(files)
	c.Assert(err, check.NotNil)
	pdfConverter = "unknown"
	_, err = MergePdfs(files)
	c.Assert(err, check.NotNil)
}
