package datetime

import "fmt"

// Numeric Separator Data Transfer Object
// Used to transmit symbols used for decimal point,
// thousands separator and currency symbol.
//
type NumericSeparatorDto struct {
	DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
	ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
	CurrencySymbol     rune // Currency Symbol
}

// Equal - Compares two NumericSeparatorDto's and returns 'true' if they
// are equivalent.
//
func (numSep *NumericSeparatorDto) Equal(numSep2 NumericSeparatorDto) bool {

	if numSep.DecimalSeparator != numSep2.DecimalSeparator {
		return false
	}

	if numSep.ThousandsSeparator != numSep2.ThousandsSeparator {
		return false
	}

	if numSep.CurrencySymbol != numSep2.CurrencySymbol {
		return false
	}

	return true
}

// New - Returns a new instance of NumericSeparatorDto. The
// rune values are automatically set to USA defaults.
//
func (numSep NumericSeparatorDto) New() NumericSeparatorDto {

	n2 := NumericSeparatorDto{}

	n2.DecimalSeparator = '.'

	n2.ThousandsSeparator = ','

	n2.CurrencySymbol = '$'

	return n2
}

// New - Returns a new instance of NumericSeparatorDto. The
// rune values are automatically set to USA defaults.
//
func (numSep NumericSeparatorDto) NewFromNumStrDto(
	numStrDto NumStrDto) NumericSeparatorDto {

	n2 := NumericSeparatorDto{}

	if numStrDto.decimalSeparator == 0 {
		n2.DecimalSeparator = '.'
	} else {
		n2.DecimalSeparator =
			numStrDto.decimalSeparator
	}

	if numStrDto.thousandsSeparator == 0 {
		n2.ThousandsSeparator = ','
	} else {
		n2.ThousandsSeparator =
			numStrDto.thousandsSeparator
	}

	if numStrDto.currencySymbol == 0 {
		n2.CurrencySymbol = '$'
	} else {
		n2.CurrencySymbol = numStrDto.currencySymbol
	}

	return n2
}

// SetFromNumStrDto - Sets the current instance of NumericSeparatorDto
// to the values provided by passed NumStrDto instance 'numStrDto'.
//
func (numSep *NumericSeparatorDto) SetFromNumStrDto(
	numStrDto *NumStrDto) {


	if numStrDto.decimalSeparator == 0 {
		numSep.DecimalSeparator = '.'
	} else {
		numSep.DecimalSeparator =
			numStrDto.decimalSeparator
	}

	if numStrDto.thousandsSeparator == 0 {
		numSep.ThousandsSeparator = ','
	} else {
		numSep.ThousandsSeparator =
			numStrDto.thousandsSeparator
	}

	if numStrDto.currencySymbol == 0 {
		numSep.CurrencySymbol = '$'
	} else {
		numSep.CurrencySymbol = numStrDto.currencySymbol
	}

}

// SetToUSADefaultsIfEmpty - If any of the NumericSeparatorDTo rune values
// are zero, this method will set those elements to USA default values.
//
func (numSep *NumericSeparatorDto) SetToUSADefaultsIfEmpty() {

	if numSep.DecimalSeparator == 0 {
		numSep.DecimalSeparator = '.'
	}

	if numSep.ThousandsSeparator == 0 {
		numSep.ThousandsSeparator = ','
	}

	if numSep.CurrencySymbol == 0 {
		numSep.CurrencySymbol = '$'
	}
}

// SetToUSADefaults - Sets all of the numeric separator
// member variables to default USA values.
//
// USA default numeric separators are listed as follows:
//
//       Currency Symbol: '$'
//   Thousands Separator: ','
//     Decimal Separator: '.'
//
func (numSep *NumericSeparatorDto) SetToUSADefaults() {

	numSep.DecimalSeparator = '.'

	numSep.ThousandsSeparator = ','

	numSep.CurrencySymbol = '$'

}

// String - Provides a formatted listing of the contents from the current
// NumericSeparatorDto instance.
//
func (numSep *NumericSeparatorDto) String() string {
	return fmt.Sprintf("Decimal Separator: %q  Thousands Separator: %q  Currency Symbol: %q",
		numSep.DecimalSeparator, numSep.ThousandsSeparator, numSep.CurrencySymbol)
}

