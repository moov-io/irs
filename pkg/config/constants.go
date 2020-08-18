// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package config

const (
	// TRecordType indicates type of transmitter “T” record
	TRecordType = "T"
	// ARecordType indicates type of payer “A” record
	ARecordType = "A"
	// BRecordType indicates type of payee “B” record
	BRecordType = "B"
	// CRecordType indicates type of end payer “C” record
	CRecordType = "C"
	// KRecordType indicates type of totals “K” record
	KRecordType = "K"
	// FRecordType indicates name of transmission “F” record
	FRecordType = "F"
	// Sub1097BtcType indicates extension block type of payee “B” record for form 1097-BTC
	Sub1097BtcType = "1097-BTC"
	// Sub1098Type indicates extension block type of payee “B” record for form 1098
	Sub1098Type = "1098"
	// Sub1098CType indicates extension block type of payee “B” record for form 1098-C
	Sub1098CType = "1098-C"
	// Sub1098EType indicates extension block type of payee “B” record for form 1098-E
	Sub1098EType = "1098-E"
	// Sub1098FType indicates extension block type of payee “B” record for form 1098-F
	Sub1098FType = "1098-F"
	// Sub1098QType indicates extension block type of payee “B” record for form 1098-Q
	Sub1098QType = "1098-Q"
	// Sub1098TType indicates extension block type of payee “B” record for form 1098-Q
	Sub1098TType = "1098-T"
	// Sub1099IntType indicates extension block type of payee “B” record for form 1099-INT
	Sub1099IntType = "1099-INT"
	// Sub1099MiscType indicates extension block type of payee “B” record for form 1099-MISC
	Sub1099MiscType = "1099-MISC"
	// Sub1099OidType indicates extension block type of payee “B” record for form 1099-OID
	Sub1099OidType = "1099-OID"
	// Sub1099PatrType indicates extension block type of payee “B” record for form 1099-PATR
	Sub1099PatrType = "1099-PATR"
)

const (
	// RecordLength indicates length of general record
	RecordLength = 750
	// SubRecordLength indicates length of sub record for payee “B” record
	SubRecordLength = 207
)

const (
	// BlankString indicates the empty string
	BlankString = " "
	// ZeroString indicates the zero string
	ZeroString = "0"
	// dateFormat indicates data format like as  YYYYMMDD
	DateFormat = "20060102"
)

const (
	// PriorYearDataIndicator indicates that reporting prior year data
	PriorYearDataIndicator = "P"
	// TestFileIndicator indicates  this is a test file
	TestFileIndicator = "T"
	// ForeignEntityIndicator indicates the transmitter is a foreign entity
	ForeignEntityIndicator = "1"
	// Software was purchased from a vendor or other source
	VendorIndicatorPurchased = "V"
	// Software was produced by in-house programmers
	VendorIndicatorProduced = "I"
	// FSFilingProgramApproved indicates  approved and submitting information as part of the CF/SF Program
	FSFilingProgramApproved = "1"
	// LastFilingIndicator indicates this is the last year this payer name and TIN will file information returns electronically or on paper
	LastFilingIndicator = "1"
	// The entity in the Second Payer Name Line Field is the transfer (or paying) agent
	TransferAgentIndicator = "1"
	// The entity shown is not the transfer agent
	NotTransferAgentIndicator = "0"
	// For a one-transaction correction or the first of a two transaction correction
	CorrectedReturnIndicatorG = "G"
	// For a second transaction of a two-transaction correction
	CorrectedReturnIndicatorC = "C"
	// TinType1 is used to identify an employer identification number (EIN)
	TinType1 = "1"
	// TinType2 is used to identify SSN, ITIN, ATIN
	TinType2 = "2"
	// the address of the payee is in a foreign country
	ForeignCountryIndicator = "1"
	// SecondTINNotice indicates notification by the IRS twice within
	// three calendar years that the payee provided an incorrect name and/or TIN combination
	SecondTINNotice = "2"
	// FatcaFilingRequirementIndicator indicates there is a FATCA Filing Requirement
	FatcaFilingRequirementIndicator = "1"
	// DirectSalesIndicator indicates sales of $5,000 or more of
	// consumer products to a person on a buy-sell, depositcommission, or any other commission basis for resale
	// anywhere other than in a permanent retail establishment
	DirectSalesIndicator = "1"

	// Enter “1” (one) if Property Securing Mortgage is the same as payer/borrowers’ address.
	PropertySecuringMortgageIndicator = "1"

	// Enter “1” (one) for general field
	GeneralOneIndicator = "1"

	// Enter “2” (one) for general field
	GeneralTwoIndicator = "2"
)

// State Abbreviation Codes
var StateAbbreviationCodes = map[string]string{
	"AL": "Alabama",
	"AK": "Alaska",
	"AS": "American Samoa",
	"AZ": "Arizona",
	"AR": "Arkansas",
	"CA": "California",
	"CO": "Colorado",
	"CT": "Connecticut",
	"DE": "Delaware",
	"DC": "District of Columbia",
	"FL": "Florida",
	"GA": "Georgia",
	"GU": "Guam",
	"HI": "Hawaii",
	"ID": "Idaho",
	"IL": "Illinois",
	"IN": "Indiana",
	"IA": "IA",
	"KS": "KS",
	"KY": "Kentucky",
	"LA": "Louisiana",
	"ME": "Maine",
	"MD": "Maryland",
	"MA": "Massachusetts",
	"MI": "Michigan",
	"MN": "Minnesota",
	"MS": "Mississippi",
	"MO": "Missouri",
	"MT": "Montana",
	"NE": "Nebraska",
	"NV": "Nevada",
	"NH": "New Hampshire",
	"NJ": "New Jersey",
	"NM": "New Mexico",
	"NY": "NY",
	"NC": "North Carolina",
	"ND": "North Dakota",
	"MP": "No. Mariana Islands",
	"OH": "Ohio",
	"OK": "Oklahoma",
	"OR": "Oregon",
	"PA": "Pennsylvania",
	"PR": "Puerto Rico",
	"RI": "Rhode Island",
	"SC": "South Carolina",
	"SD": "South Dakota",
	"TN": "Tennessee",
	"TX": "Texas",
	"UT": "Utah",
	"VT": "Vermont",
	"VA": "Virginia",
	"VI": "U.S. Virgin Islands",
	"WA": "Washington",
	"WV": "West Virginia",
	"WI": "Wisconsin",
	"WY": "Wyoming",
}

// Codes for participating states in the CF/SF Program
var ParticipateStateCodes = map[int]string{
	1:  "Alabama",
	4:  "Arizona",
	5:  "Arkansas",
	6:  "California",
	7:  "Colorado",
	8:  "Connecticut",
	10: "Delaware",
	13: "Georgia",
	15: "Hawaii",
	16: "Idaho",
	18: "Indiana",
	20: "Kansas",
	22: "Louisiana",
	23: "Maine",
	24: "Maryland",
	25: "Massachusetts",
	26: "Michigan",
	27: "Minnesota",
	28: "Mississippi",
	29: "Missouri",
	30: "Montana",
	31: "Nebraska",
	34: "New Jersey",
	35: "New Mexico",
	37: "North Carolina",
	38: "North Dakota",
	39: "Ohio",
	40: "Ohio",
	45: "South Carolina",
	55: "Wisconsin",
}

// Codes for type of return
var TypeOfReturns = map[string]string{
	"BT": "1097-BTC",
	"3":  "1098",
	"X":  "1098-C",
	"2":  "1098-E",
	"FP": "1098-F",
	"QL": "1098-Q",
	"8":  "1098-T",
	"4":  "1099-A",
	"B":  "1099-B",
	"5":  "1099-C",
	"P":  "1099-CAP",
	"1":  "1099-DIV",
	"F":  "1099-G",
	"J":  "1099-H",
	"6":  "1099-INT",
	"MC": "1099-K",
	"LC": "1099-LS",
	"T":  "1099-LTC",
	"A":  "1099-MISC",
	"D":  "1099-OID",
	"7":  "1099-PATR",
	"Q":  "1099-Q",
	"9":  "1099-R",
	"S":  "1099-S",
	"M":  "1099-SA",
	"SB": "1099-SB",
	"N":  "3921",
	"Z":  "3922",
	"L":  "5498",
	"V":  "5498-ESA",
	"K":  "5498-SA",
	"W":  "W-2G",
}

// Available issuer indicators for 1097-BTC
var BtcIssuerIndicator = map[string]string{
	"1": "Issuer of bond",
	"2": "An entity that received a 2018 Form",
}

// Available codes for 1097-BTC
var BtcCode = map[string]string{
	"A": "Account number",
	"C": "CUSIP number",
	"O": "Unique identification number",
}

// Available bond types for 1097-BTC
var BtcBondType = map[string]string{
	"101": "Clean Renewable Energy Bond",
	"199": "Other",
}

// Amount codes for the type of return being reported.
var AmountCodes = map[string]map[string]string{
	"1097-BTC": {
		"1": "Total Aggregate",
		"2": "January payments",
		"3": "February payments",
		"4": "March payments",
		"5": "April payments",
		"6": "May payments",
		"7": "June payments",
		"8": "July payments",
		"9": "August payments",
		"A": "September payments",
		"B": "October payments",
		"C": "November payments",
		"D": "December payments",
	},
	"1098": {
		"1": "Mortgage",
		"2": "Points",
		"3": "Refund",
		"4": "Mortgage Insurance Premium",
		"5": "Blank",
		"6": "Outstanding Mortgage Principal",
	},
	"1098-C": {
		"4": "Gross proceeds from sales",
		"6": "Value of goods or services",
	},
	"1098-E": {
		"1": "Student loan interest received by the lender",
	},
	"1098-F": {
		"1": "Total amount required to be paid",
		"2": "Restitution/remediation amount",
		"3": "Compliance amount",
	},
	"1098-Q": {
		"1": "January payments",
		"2": "February payments",
		"3": "March payments",
		"4": "April payments",
		"5": "May payments",
		"6": "June payments",
		"7": "July payments",
		"8": "August payments",
		"9": "September payments",
		"A": "October payments",
		"B": "November payments",
		"C": "December payments",
		"D": "Total premiums",
		"E": "Annuity amount on start date",
		"F": "FMV of QLAC",
	},
	"1098-T": {
		"1": "Payments",
		"3": "Payments",
		"4": "Scholarships",
		"5": "Adjustments",
		"7": "Reimbursements",
	},
	"1099-A": {
		"2": "Balance of principal outstanding",
		"4": "Fair market value of the property",
	},
	"1099-B": {
		"2": "Proceeds",
		"3": "Cost",
		"4": "Federal income tax withheld",
		"5": "Wash Sale Loss Disallowed",
		"7": "Bartering",
		"9": "Profit",
		"A": "Unrealized profit",
		"B": "Unrealized profit",
		"C": "Aggregate profit",
		"D": "Accrued Market Discount",
	},
	"1099-C": {
		"2": "Amount of debt discharged",
		"3": "Interest included in Amount Code 2",
		"7": "Fair market value of property",
	},
	"1099-CAP": {
		"2": "Aggregate amount received",
	},
	"1099-DIV": {
		"1": "Total ordinary dividends",
		"2": "Qualified dividends",
		"3": "Total capital gain distribution",
		"5": "Section 199A Dividends",
		"6": "Unrecaptured Section 1250 gain",
		"7": "Section 1202 gain",
		"8": "Collectibles (28%) rate gain",
		"9": "Nondividend distributions",
		"A": "Federal income tax withheld",
		"B": "Investment expenses",
		"C": "Foreign tax paid",
		"D": "Cash liquidation distributions",
		"E": "Non-cash liquidation distributions",
		"F": "Exempt interest dividends",
		"G": "Specified private activity bond interest dividends",
	},
	"1099-G": {
		"1": "Unemployment",
		"2": "State or local income tax",
		"4": "Federal income tax withheld",
		"5": "Reemployment Trade Adjustment ",
		"6": "Taxable grants",
		"7": "Agriculture payments",
		"9": "Market gain",
	},
	"1099-H": {
		"1": "Gross amount of health insurance advance payments",
		"2": "Gross amount of health insurance payments for January",
		"3": "Gross amount of health insurance payments for February",
		"4": "Gross amount of health insurance payments for March",
		"5": "Gross amount of health insurance payments for April",
		"6": "Gross amount of health insurance payments for May",
		"7": "Gross amount of health insurance payments for June",
		"8": "Gross amount of health insurance payments for July",
		"9": "Gross amount of health insurance payments for August",
		"A": "Gross amount of health insurance payments for September",
		"B": "Gross amount of health insurance payments for October",
		"C": "Gross amount of health insurance payments for November",
		"D": "Gross amount of health insurance payments for December",
	},
	"1099-INT": {
		"1": "Interest income not included in Amount Code 3",
		"2": "Early withdrawal penalty",
		"3": "Interest on U.S. Savings Bonds and Treasury obligations",
		"4": "Federal income tax withheld (backup withholding)",
		"5": "Investment expenses",
		"6": "Foreign tax paid",
		"8": "Tax-exempt interest",
		"9": "Specified private activity bond",
		"A": "Market discount",
		"B": "Bond premium",
		"D": "Bond premium on tax exempt bond",
		"E": "Bond premium on Treasury obligation",
	},
	"1099-K": {
		"1": "Gross amount of payment card/third party network transactions",
		"2": "Card not present transactions",
		"4": "Federal Income tax withheld",
		"5": "January payments",
		"6": "February payments",
		"7": "March payments",
		"8": "April payments",
		"9": "May payments",
		"A": "June payments",
		"B": "July payments",
		"C": "August payments",
		"D": "September payments",
		"E": "October payments",
		"F": "November payments",
		"G": "December payments",
	},
	"1099-LS": {
		"1": "Amount paid to payment recipient",
	},
	"1099-LTC": {
		"1": "Gross long-term care benefits paid",
		"2": "Accelerated death benefits paid",
	},
	"1099-MISC": {
		"1": "Rents",
		"2": "Royalties",
		"3": "Other income",
		"4": "Federal income tax withheld",
		"5": "Fishing boat proceeds",
		"6": "Medical and health care payments",
		"7": "Nonemployee compensation (NEC)",
		"8": "Substitute payments in lieu of dividends or interest",
		"A": "Crop insurance proceeds",
		"B": "Excess golden parachute payment",
		"C": "Gross proceeds paid to an attorney in connection with legal services",
		"D": "Section 409A deferrals",
		"E": "Section 409A income",
	},
	"1099-OID": {
		"1": "Original issue discount for 2019",
		"2": "Other periodic interest",
		"3": "Early withdrawal penalty",
		"4": "Federal income tax withheld",
		"5": "Bond premium",
		"6": "Original issue discount on U.S.",
		"7": "Investment expenses",
		"A": "Market discount",
		"B": "Acquisition premium",
		"C": "Tax-Exempt OID",
	},
	"1099-PATR": {
		"1": "Patronage dividends",
		"2": "Nonpatronage distributions",
		"3": "Per-unit retain allocations",
		"4": "Federal income tax withheld",
		"5": "Redemption of nonqualified notices and retain allocations",
		"6": "Deduction for domestic production activities income",
		"B": "Qualified Payments",
		"7": "Investment credit",
		"8": "Work opportunity credit",
		"9": "Patron’s alternative minimum tax (AMT) adjustment",
		"A": "For filer’s use for pass-through credits and deduction",
	},
	"1099-Q": {
		"1": "Gross distribution",
		"2": "Earnings (or loss)",
		"3": "Basis",
	},
	"1099-R": {
		"1": "Gross distribution",
		"2": "Taxable amount",
		"3": "Capital gain",
		"4": "Federal income tax withheld",
		"5": "Employee contributions/designated Roth contributions or insurance premiums",
		"6": "Net unrealized appreciation in employer’s securities",
		"8": "Other",
		"9": "Total employee contributions",
		"A": "Traditional IRA/SEP/SIMPLE",
		"B": "Amount allocable to IRR",
	},
	"1099-S": {
		"2": "Gross proceeds",
		"5": "Buyer’s part of real estate tax",
	},
	"1099-SA": {
		"1": "Gross distribution",
		"2": "Earnings on excess contributions",
		"4": "Fair market value of the account on the date of death",
	},
	"1099-SB": {
		"1": "Investment in contract",
		"2": "Surrender amount",
	},
	"3921": {
		"3": "Exercise price per share",
		"4": "Fair market value of share on exercise date",
	},
	"3922": {
		"3": "Fair market value per share on grant date",
		"4": "Fair market value on exercise date",
		"5": "Exercise price per share",
		"8": "Exercise price per share determined  as if the option was exercised on the date the option was granted",
	},
	"5498": {
		"1": "IRA contributions",
		"2": "Rollover contributions",
		"3": "Roth conversion amount",
		"4": "Recharacterized contributions",
		"5": "Fair market value of account",
		"6": "Life insurance cost included in Amount Code 1",
		"7": "FMV of certain specified assets",
		"8": "SEP contributions",
		"9": "SIMPLE contributions",
		"A": "Roth IRA contributions",
		"B": "RMD amount",
		"C": "Postponed Contribution",
		"D": "Repayments",
	},
	"5498-ESA": {
		"1": "Coverdell ESA contributions",
		"2": "Rollover contributions",
	},
	"5498-SA": {
		"1": "Employee",
		"2": "Total contributions made in 2019",
		"3": "Total HSA",
		"4": "Rollover contributions",
		"5": "Fair market value of HSA",
	},
	"W-2G": {
		"1": "Reportable winnings",
		"2": "Federal income tax withheld",
		"7": "Winnings from identical wagers",
	},
}

// Amount codes for Positions 544-750 for Form 1098-F.
var PaymentCodes1098F = map[string]string{
	"B": "Multiple payers/defendants",
	"C": "Multiple payees",
	"D": "Property included in settlement",
	"E": "Settlement payments to nongovernmental entities, i.e., charities",
	"F": "Settlement paid in full as of time of filing",
	"G": "No payment received as of time of filing",
	"H": "Deferred prosecution agreement",
}

const (
	OutputJsonFormat = "json"
	OutputIrsFormat  = "irs"
)
