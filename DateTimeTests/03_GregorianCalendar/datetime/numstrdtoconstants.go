package datetime

/*
 Number String Dto Constants
 ===========================

 Overview and General Usage
 ==========================
 This source file is used to store constants used by
 by type 'NumStrDto' and related number string formatting
 utilities.

*/

// Source Currency Info
// https://gist.github.com/bzerangue/5484121
// http://symbologic.info/currency.htm
// http://www.xe.com/symbols.php
//
// NumStrCurrencySymbols - an array containing
// most of the world's major currency symbols
// stored as type 'rune'.
//
var NumStrCurrencySymbols = []rune{
	'\U00000024', // Australia Dollar 								 0
	'\U00008236', // Brazil Real											 1
	'\U00000024', // Canada Dollar 										 2
	'\U000000a5', // China Yuan												 3
	'\U00000024', // Colombia Peso										 4
	'\U0004b10d', // Czech Republic Koruna						 5
	'\U000000a3', // Egypt Pound											 6
	'\U000020ac', // Euro €  													 7
	'\U00070116', // Hungary Forint										 8
	'\U00107114', // Iceland Krona										 9
	'\U00082112', // Indonesia Rupiah									10
	'\U000020aa', // Israel Shekel  									11
	'\U000000a5', // Japan Yen  											12
	'\U000020a9', // Korea Won  											13
	'\U0000524d', // Malaysia Ringgit									14
	'\U00000024', // Mexico Peso  										15
	'\U00006b72', // Norway Krone											16
	'\U00000192', // Netherlands Antilles Guilder			17
	'\U000020a8', // Pakistan Rupee 									18
	'\U000020bd', // Russian Ruble  									19
	'\U0000fdfc', // Saudi Arabia Riyal 							20
	'\U00000082', // South Africa Rand								21
	'\U00006b72', // Sweden Krona											22
	'\U000020a3', // Switzerland Franc								23
	'\U00000024', // Taiwan NewBigIntNum Dollar								24
	'\U000020ba', // TURKISH LIRA											25
	'\U00066115', // Venezuela Bolivar								26
	'\U00008363', // Viet Nam Dong										27
	'\U00000024', // United States Dollar  						28
	'\U000000a3', // United Kingdom Pound (£)					29
	'\U000020a3', // French Franc  						        30
	'\U000020a4', // Italy Lira  						          31
	'\U000020bf', // Bitcoin  						            32
	'\U000000a2'} // United States Cent		            33

// PrecisionScaleMode - Specifies the scaling mode used in
// setting precision or the number of digits to the right
// of the decimal place.
type PrecisionScaleMode int

func (scaleMode PrecisionScaleMode) String() string {

	return PrecisionScaleModeLabels[scaleMode]
}

const (
	SCALEPRECISIONRIGHT PrecisionScaleMode = iota

	SCALEPRECISIONLEFT
)

var PrecisionScaleModeLabels = [...]string{"ScalePrecisionRight", "ScalePrecisionLeft"}

// NumStrFmtMode - Designates the type of number string formatting
// applied when converting a number to a string.
type NumStrFmtMode int

func (nstrFmtMode NumStrFmtMode) String() string {
	return NumStrFmtModeLabels[nstrFmtMode]
}

const (

	// PUREINTEGERFMT - Specifies a pure number string with no decimal
	// point, no thousands separators and no currency symbol.
	// Example: 123456789
	//
	PUREINTEGERFMT NumStrFmtMode = iota

	// INTSTRDECIMALFMT - Specifies an integer string, decimal point and
	// fractional digits. No currency symbol or thousands separator are
	// injected in to the final number string.
	// Example: 12345.678
	//
	INTSTRDECIMALFMT

	// THOUSANDSNUMSTRFMT - Specifies that the output number string will
	// have a decimal point separating fractional digits. Integer numbers
	// to the left of the decimal point will have a thousands separator
	// character injected after every third character.
	// Example: 123,456,789.23
	//
	THOUSANDSNUMSTRFMT

	// CURRENCYNUMSTRFMT - Specifies a Currency String. The output number string
	// will include a currency symbol, thousands separators and a decimal point.
	CURRENCYNUMSTRFMT
)

var NumStrFmtModeLabels = [...]string{"PureIntegerString", "IntegerDecimalString", "ThousandsNumString", "CurrencyNumString"}

type NegativeValueFmtMode int

func (negValFmtMode NegativeValueFmtMode) String() string {
	return NegativeValueFmtModeLabels[negValFmtMode]
}

const (

// LEADMINUSNEGVALFMTMODE - Negative values formatted with
//                          a leading minus sign.
//                          Example: -123456.78
//
	LEADMINUSNEGVALFMTMODE NegativeValueFmtMode = iota

// PARENTHESESNEGVALFMTMODE - Negative values formatted with
//                            surrounding parentheses.
//                            Example: (123456.78)
//
	PARENTHESESNEGVALFMTMODE

// ABSOLUTEPURENUMSTRFMTMODE - Formats a pure number string with
//                              absolute (positive) integer value
//                              and no decimal point separator.
//                             Example: (12345678)
	ABSOLUTEPURENUMSTRFMTMODE
)

var NegativeValueFmtModeLabels = [...]string{"LeadingMinusSign", "SurroundingParentheses", "AbsolutePureNumberString"}
