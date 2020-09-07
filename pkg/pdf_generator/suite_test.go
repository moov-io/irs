// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pdf_generator

import (
	"gopkg.in/check.v1"
	"strings"
	"testing"
)

func Test(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&PdfTest{})

// Pdf test
type PdfTest struct{}

func (t *PdfTest) SetUpSuite(c *check.C) {}

func (t *PdfTest) SetUpTest(c *check.C) {
	pdfConverter = "pdftk"
	convertParam1 = "fill_form"
}

func (t *PdfTest) TestPdfWithMscCopyB(c *check.C) {
	pdf := Pdf1099Misc{Type: PdfMscCopyB}
	_, err := pdf.Generate()
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
	_, err := pdf.Generate()
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
	_, err := pdf.Generate()
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
	_, err = pdf.Generate()
	c.Assert(err, check.IsNil)
}

func (t *PdfTest) TestPdfWithUnknownTemplate(c *check.C) {
	pdf := Pdf1099Misc{Type: "Unknown"}
	_, err := pdf.Generate()
	c.Assert(err, check.NotNil)
	_, err = pdf.getSpecFdf()
	c.Assert(err, check.NotNil)
	_, err = pdf.getTemplateFdf()
	c.Assert(err, check.NotNil)
	_, err = pdf.getTemplateFile()
	c.Assert(err, check.NotNil)
	_, err = pdf.generateFDF("")
	c.Assert(err, check.NotNil)
	_, err = pdf.Generate()
	c.Assert(err, check.NotNil)
}

func (t *PdfTest) TestPdfWithErrorParam(c *check.C) {
	pdf := Pdf1099Misc{Type: PdfNecCopyC}
	convertParam1 = "Unknown"
	_, err := pdf.generatePDF("")
	c.Assert(err, check.NotNil)
	pdfConverter = "Unknown"
	_, err = pdf.generatePDF("")
	c.Assert(err, check.NotNil)
}
