package config

import "sort"

type SpecField struct {
	Start    int
	Length   int
	Type     int
	Required string
}

type SpecRecord struct {
	Key   int
	Name  string
	Field SpecField
}

// field properties
const (
	Nullable   = ""
	Required   = "Y"
	Applicable = "A"
	Expandable = "E"
)

// field types
const (
	Alphanumeric = 1 << iota
	AlphanumericRightAlign
	Numeric
	ZeroNumeric
	TelephoneNumber
	Email
	DateYear
)

var (
	// Transmitter “T” Record
	TRecordLayout = map[string]SpecField{
		"RecordType":                   {0, 1, Alphanumeric, Required},
		"PaymentYear":                  {1, 4, DateYear, Required},
		"PriorYearDataIndicator":       {5, 1, Alphanumeric, Required},
		"TIN":                          {6, 9, Numeric, Required},
		"TCC":                          {15, 5, Alphanumeric, Required},
		"Blank1":                       {20, 7, Alphanumeric, Nullable},
		"TestFileIndicator":            {27, 1, Alphanumeric, Applicable},
		"ForeignEntityIndicator":       {28, 1, Alphanumeric, Applicable},
		"TransmitterName":              {29, 40, Alphanumeric, Required},
		"TransmitterNameContinuation":  {69, 40, Alphanumeric, Applicable},
		"CompanyName":                  {109, 40, Alphanumeric, Required},
		"CompanyNameContinuation":      {149, 40, Alphanumeric, Applicable},
		"CompanyMailingAddress":        {189, 40, Alphanumeric, Required},
		"CompanyCity":                  {229, 40, Alphanumeric, Required},
		"CompanyState":                 {269, 2, Alphanumeric, Required},
		"CompanyZipCode":               {271, 9, Numeric, Required},
		"Blank2":                       {280, 15, Alphanumeric, Nullable},
		"TotalNumberPayees":            {295, 8, ZeroNumeric, Applicable},
		"ContactName":                  {303, 40, Alphanumeric, Required},
		"ContactTelephoneNumber":       {343, 15, TelephoneNumber, Required},
		"ContactEmailAddress":          {358, 50, Email, Applicable},
		"Blank3":                       {408, 91, Alphanumeric, Nullable},
		"RecordSequenceNumber":         {499, 8, ZeroNumeric, Required},
		"Blank4":                       {507, 10, Alphanumeric, Nullable},
		"VendorIndicator":              {517, 1, Alphanumeric, Required},
		"VendorName":                   {518, 40, Alphanumeric, Required},
		"VendorMailingAddress":         {558, 40, Alphanumeric, Required},
		"VendorCity":                   {598, 40, Alphanumeric, Required},
		"VendorState":                  {638, 2, Alphanumeric, Required},
		"VendorZipCode":                {640, 9, Numeric, Required},
		"VendorContactName":            {649, 40, Alphanumeric, Required},
		"VendorContactTelephoneNumber": {689, 15, TelephoneNumber, Required},
		"Blank5":                       {704, 35, Alphanumeric, Nullable},
		"VendorForeignEntityIndicator": {739, 1, Alphanumeric, Applicable},
		"Blank6":                       {740, 8, Alphanumeric, Nullable},
		"Blank7":                       {748, 2, Alphanumeric, Nullable},
	}
	// Payer “A” Record
	ARecordLayout = map[string]SpecField{
		"RecordType":              {0, 1, Alphanumeric, Required},
		"PaymentYear":             {1, 4, DateYear, Required},
		"CombinedFSFilingProgram": {5, 1, Alphanumeric, Applicable},
		"Blank1":                  {6, 5, Alphanumeric, Nullable},
		"TIN":                     {11, 9, Numeric, Required},
		"PayerNameControl":        {20, 4, Alphanumeric, Applicable},
		"LastFilingIndicator":     {24, 1, Alphanumeric, Applicable},
		"TypeOfReturn":            {25, 2, Alphanumeric, Required},
		"AmountCodes":             {27, 16, Alphanumeric, Required},
		"Blank2":                  {43, 8, Alphanumeric, Nullable},
		"ForeignEntityIndicator":  {51, 1, Alphanumeric, Applicable},
		"FirstPayerNameLine":      {52, 40, Alphanumeric, Required},
		"SecondPayerNameLine":     {92, 40, Alphanumeric, Applicable},
		"TransferAgentIndicator":  {132, 1, Alphanumeric, Required},
		"PayerShippingAddress":    {133, 40, Alphanumeric, Required},
		"PayerCity":               {173, 40, Alphanumeric, Required},
		"PayerState":              {213, 2, Alphanumeric, Required},
		"PayerZipCode":            {215, 9, Numeric, Required},
		"PayerTelephoneNumber":    {224, 15, TelephoneNumber, Required},
		"Blank3":                  {239, 260, Alphanumeric, Nullable},
		"RecordSequenceNumber":    {499, 8, ZeroNumeric, Required},
		"Blank4":                  {507, 241, Alphanumeric, Nullable},
		"Blank5":                  {748, 2, Alphanumeric, Nullable},
	}
	// Payee “B” Record
	BRecordLayout = map[string]SpecField{
		"RecordType":               {0, 1, Alphanumeric, Required},
		"PaymentYear":              {1, 4, DateYear, Required},
		"CorrectedReturnIndicator": {5, 1, Alphanumeric, Applicable},
		"NameControl":              {6, 4, Alphanumeric, Applicable},
		"TypeOfTIN":                {10, 1, Numeric, Applicable},
		"TIN":                      {11, 9, Numeric, Required},
		"PayerAccountNumber":       {20, 20, Alphanumeric, Applicable},
		"PayerOfficeCode":          {40, 4, Alphanumeric, Applicable},
		"Blank1":                   {44, 10, Alphanumeric, Nullable},
		"PaymentAmount1":           {54, 12, ZeroNumeric, Required},
		"PaymentAmount2":           {66, 12, ZeroNumeric, Required},
		"PaymentAmount3":           {78, 12, ZeroNumeric, Required},
		"PaymentAmount4":           {90, 12, ZeroNumeric, Required},
		"PaymentAmount5":           {102, 12, ZeroNumeric, Required},
		"PaymentAmount6":           {114, 12, ZeroNumeric, Required},
		"PaymentAmount7":           {126, 12, ZeroNumeric, Required},
		"PaymentAmount8":           {138, 12, ZeroNumeric, Required},
		"PaymentAmount9":           {150, 12, ZeroNumeric, Required},
		"PaymentAmountA":           {162, 12, ZeroNumeric, Required},
		"PaymentAmountB":           {174, 12, ZeroNumeric, Required},
		"PaymentAmountC":           {186, 12, ZeroNumeric, Required},
		"PaymentAmountD":           {198, 12, ZeroNumeric, Required},
		"PaymentAmountE":           {210, 12, ZeroNumeric, Required},
		"PaymentAmountF":           {222, 12, ZeroNumeric, Required},
		"PaymentAmountG":           {234, 12, ZeroNumeric, Required},
		"ForeignCountryIndicator":  {246, 1, Alphanumeric, Applicable},
		"FirstPayeeNameLine":       {247, 40, Alphanumeric, Required},
		"SecondPayeeNameLine":      {287, 40, Alphanumeric, Applicable},
		"Blank2":                   {327, 40, Alphanumeric, Nullable},
		"PayeeMailingAddress":      {367, 40, Alphanumeric, Required},
		"Blank3":                   {407, 40, Alphanumeric, Nullable},
		"PayeeCity":                {447, 40, Alphanumeric, Required},
		"PayeeState":               {487, 2, Alphanumeric, Required},
		"PayeeZipCode":             {489, 9, Numeric, Required},
		"Blank4":                   {498, 1, Alphanumeric, Nullable},
		"RecordSequenceNumber":     {499, 8, ZeroNumeric, Required},
		"Blank5":                   {507, 36, Alphanumeric, Nullable},
		"Reserved":                 {543, 207, Alphanumeric, Expandable},
	}
	// End of Payer “C” Record
	CRecordLayout = map[string]SpecField{
		"RecordType":           {0, 1, Alphanumeric, Required},
		"NumberPayees":         {1, 8, ZeroNumeric, Required},
		"Blank1":               {9, 6, Alphanumeric, Nullable},
		"ControlTotal1":        {15, 18, ZeroNumeric, Required},
		"ControlTotal2":        {33, 18, ZeroNumeric, Required},
		"ControlTotal3":        {51, 18, ZeroNumeric, Required},
		"ControlTotal4":        {69, 18, ZeroNumeric, Required},
		"ControlTotal5":        {87, 18, ZeroNumeric, Required},
		"ControlTotal6":        {105, 18, ZeroNumeric, Required},
		"ControlTotal7":        {123, 18, ZeroNumeric, Required},
		"ControlTotal8":        {141, 18, ZeroNumeric, Required},
		"ControlTotal9":        {159, 18, ZeroNumeric, Required},
		"ControlTotalA":        {177, 18, ZeroNumeric, Required},
		"ControlTotalB":        {195, 18, ZeroNumeric, Required},
		"ControlTotalC":        {213, 18, ZeroNumeric, Required},
		"ControlTotalD":        {231, 18, ZeroNumeric, Required},
		"ControlTotalE":        {249, 18, ZeroNumeric, Required},
		"ControlTotalF":        {267, 18, ZeroNumeric, Required},
		"ControlTotalG":        {285, 18, ZeroNumeric, Required},
		"Blank2":               {303, 196, Alphanumeric, Nullable},
		"RecordSequenceNumber": {499, 8, ZeroNumeric, Required},
		"Blank3":               {507, 241, Alphanumeric, Nullable},
		"Blank4":               {748, 2, Alphanumeric, Nullable},
	}
	// State Totals “K” Record
	KRecordLayout = map[string]SpecField{
		"RecordType":                  {0, 1, Alphanumeric, Required},
		"NumberPayees":                {1, 8, ZeroNumeric, Required},
		"Blank1":                      {9, 6, Alphanumeric, Nullable},
		"ControlTotal1":               {15, 18, ZeroNumeric, Required},
		"ControlTotal2":               {33, 18, ZeroNumeric, Required},
		"ControlTotal3":               {51, 18, ZeroNumeric, Required},
		"ControlTotal4":               {69, 18, ZeroNumeric, Required},
		"ControlTotal5":               {87, 18, ZeroNumeric, Required},
		"ControlTotal6":               {105, 18, ZeroNumeric, Required},
		"ControlTotal7":               {123, 18, ZeroNumeric, Required},
		"ControlTotal8":               {141, 18, ZeroNumeric, Required},
		"ControlTotal9":               {159, 18, ZeroNumeric, Required},
		"ControlTotalA":               {177, 18, ZeroNumeric, Required},
		"ControlTotalB":               {195, 18, ZeroNumeric, Required},
		"ControlTotalC":               {213, 18, ZeroNumeric, Required},
		"ControlTotalD":               {231, 18, ZeroNumeric, Required},
		"ControlTotalE":               {249, 18, ZeroNumeric, Required},
		"ControlTotalF":               {267, 18, ZeroNumeric, Required},
		"ControlTotalG":               {285, 18, ZeroNumeric, Required},
		"Blank2":                      {303, 196, Alphanumeric, Nullable},
		"RecordSequenceNumber":        {499, 8, ZeroNumeric, Required},
		"Blank3":                      {507, 199, Alphanumeric, Nullable},
		"StateIncomeTaxWithheldTotal": {706, 18, Numeric, Applicable},
		"LocalIncomeTaxWithheldTotal": {724, 18, Numeric, Applicable},
		"Blank4":                      {742, 4, Alphanumeric, Nullable},
		"CombinedFSCode":              {746, 2, Alphanumeric, Required},
		"Blank5":                      {748, 2, Alphanumeric, Nullable},
	}
	// End of Transmission “F” Record
	FRecordLayout = map[string]SpecField{
		"RecordType":           {0, 1, Alphanumeric, Required},
		"NumberPayerRecords":   {1, 8, ZeroNumeric, Required},
		"Zero":                 {9, 21, ZeroNumeric, Applicable},
		"Blank2":               {30, 19, Alphanumeric, Nullable},
		"TotalNumberPayees":    {49, 8, ZeroNumeric, Applicable},
		"Blank3":               {57, 442, Alphanumeric, Nullable},
		"RecordSequenceNumber": {499, 8, ZeroNumeric, Required},
		"Blank4":               {507, 241, Alphanumeric, Nullable},
		"Blank5":               {748, 2, Alphanumeric, Nullable},
	}
)

var (
	// Record Layout Positions 544-750 for Form 1097-BTC
	Sub1097BTCLayout = map[string]SpecField{
		"Blank1":             {0, 3, Alphanumeric, Nullable},
		"IssuerIndicator":    {3, 1, Numeric, Required},
		"Blank2":             {4, 8, Alphanumeric, Nullable},
		"Code":               {12, 1, Alphanumeric, Required},
		"Blank3":             {13, 3, Alphanumeric, Nullable},
		"UniqueIdentifier":   {16, 39, AlphanumericRightAlign, Applicable},
		"BondType":           {55, 3, Alphanumeric, Required},
		"Blank4":             {58, 61, Alphanumeric, Nullable},
		"SpecialDataEntries": {119, 60, Alphanumeric, Applicable},
		"Blank5":             {179, 26, Alphanumeric, Nullable},
		"Blank6":             {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1099-INT
	Sub1099INTLayout = map[string]SpecField{
		"SecondTinNotice":        {0, 1, Alphanumeric, Applicable},
		"Blank1":                 {1, 2, Alphanumeric, Nullable},
		"ForeignCountry":         {3, 40, Alphanumeric, Applicable},
		"CUSIP":                  {43, 13, AlphanumericRightAlign, Applicable},
		"FATCA":                  {56, 1, Alphanumeric, Applicable},
		"Blank2":                 {57, 62, Alphanumeric, Nullable},
		"SpecialDataEntries":     {119, 60, Alphanumeric, Applicable},
		"StateIncomeTaxWithheld": {179, 12, ZeroNumeric, Applicable},
		"LocalIncomeTaxWithheld": {191, 12, ZeroNumeric, Applicable},
		"CombinedFSCode":         {203, 2, Alphanumeric, Required},
		"Blank3":                 {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1099-MISC
	Sub1099MISCLayout = map[string]SpecField{
		"SecondTinNotice":        {0, 1, Alphanumeric, Applicable},
		"Blank1":                 {1, 2, Alphanumeric, Nullable},
		"DirectSalesIndicator":   {3, 1, Alphanumeric, Applicable},
		"FATCA":                  {4, 1, Alphanumeric, Applicable},
		"Blank2":                 {5, 114, Alphanumeric, Nullable},
		"SpecialDataEntries":     {119, 60, Alphanumeric, Applicable},
		"StateIncomeTaxWithheld": {179, 12, ZeroNumeric, Applicable},
		"LocalIncomeTaxWithheld": {191, 12, ZeroNumeric, Applicable},
		"CombinedFSCode":         {203, 2, Alphanumeric, Required},
		"Blank3":                 {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1099-OID
	Sub1099OIDLayout = map[string]SpecField{
		"SecondTinNotice":        {0, 1, Alphanumeric, Applicable},
		"Blank1":                 {1, 2, Alphanumeric, Nullable},
		"Description":            {3, 39, Alphanumeric, Applicable},
		"FATCA":                  {42, 1, Alphanumeric, Applicable},
		"Blank2":                 {43, 76, Alphanumeric, Nullable},
		"SpecialDataEntries":     {119, 60, Alphanumeric, Applicable},
		"StateIncomeTaxWithheld": {179, 12, ZeroNumeric, Applicable},
		"LocalIncomeTaxWithheld": {191, 12, ZeroNumeric, Applicable},
		"CombinedFSCode":         {203, 2, Alphanumeric, Required},
		"Blank3":                 {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1099-PATR
	Sub1099PATRLayout = map[string]SpecField{
		"SecondTinNotice":        {0, 1, Alphanumeric, Applicable},
		"Blank1":                 {1, 118, Alphanumeric, Nullable},
		"SpecialDataEntries":     {119, 60, Alphanumeric, Applicable},
		"StateIncomeTaxWithheld": {179, 12, ZeroNumeric, Applicable},
		"LocalIncomeTaxWithheld": {191, 12, ZeroNumeric, Applicable},
		"CombinedFSCode":         {203, 2, Alphanumeric, Required},
		"Blank3":                 {205, 2, Alphanumeric, Nullable},
	}
)

func ToSpecifications(fieldsFormat map[string]SpecField) []SpecRecord {
	var records []SpecRecord
	for key, field := range fieldsFormat {
		records = append(records, SpecRecord{field.Start, key, field})
	}
	sort.Slice(records, func(i, j int) bool {
		if records[i].Key == records[j].Key {
			return records[i].Name < records[j].Name
		}
		return records[i].Key < records[j].Key
	})
	return records
}
