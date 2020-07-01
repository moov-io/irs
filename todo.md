# todo

Future work: 

Allow PDF's to be filled from the data in the system for sumbidsion to payees. 

Golang: 
https://github.com/desertbit/fillpdf/network
utilizes [PDFtk](https://www.pdflabs.com/tools/pdftk-server/)

Finding all the PDF field properties from the command line: 

```text

wade-mbp:Desktop wadearnold$ pdftk 1099-misc-2019.pdf dump_data_fields
---
FieldType: Button
FieldName: topmostSubform[0].CopyA[0].CopyAHeader[0].c1_1[0]
FieldFlags: 1
FieldValue: Off
FieldJustification: Left
FieldStateOption: 1
FieldStateOption: Off
--
FieldType: Button
FieldName: topmostSubform[0].CopyA[0].CopyAHeader[0].c1_1[1]
FieldFlags: 1
FieldValue: Off
FieldJustification: Left
FieldStateOption: 2
FieldStateOption: Off
---
```

Free website for finding what each of the fields are for a PDF


