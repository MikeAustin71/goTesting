package datetime

import (
	"math/big"
)

// NumStrDto - This Type contains data fields and methods used
// to manage, store and transport number strings.
//
//
// Supporting File hierarchy
// numstrdto.go ->
//  numstrdtoutility.go ->
//    numstrdtohelper.go ->
//     numstrdtomechanics.go ->
//       numstrdtonanobot.go ->
//         numstrdtomolecule.go ->
//           numstrdtoatom.go ->
//             numstrdtoelectron.go
//
type NumStrDto struct {
	signVal int             // An integer value indicating the numeric sign of this number string.
	//                      //   Valid values are +1 or -1
	absAllNumRunes []rune   // An array of runes containing all the numeric digits in a number with
	//                      //   no preceding plus or minus sign character. Example: 123.456 =
	//                      //   []rune{'1','2','3','4','5','6'}
	precision          uint // The number of digits to the right of the decimal point.
	thousandsSeparator rune // Separates thousands in the integer number: '1,000,000,000
	decimalSeparator   rune // Separates integer and fractional elements of a number. '123.456'
	currencySymbol     rune // Currency symbol used in currency string displays
}

// Add - Adds the value of input NumStrDto to the current NumStrDto
// instance.
//
// Input parameter 'ePrefix' is a string consisting of the method chain used to call
// this method. In case of error, this text string is included in the error message.
// Note: Be sure to leave a space at the end of 'ePrefix'.
//
func (nDto *NumStrDto) Add(
	n2Dto NumStrDto,
	ePrefix string) error {

	ePrefix += "NumStrDto.Add() "

	n1Dto := nDto.CopyOut()

	nResult, err := nDto.AddNumStrs(
		n1Dto,
		n2Dto,
		ePrefix)

	if err != nil {
		return err
	}

	nDto.CopyIn(nResult)

	return nil
}

// AddNumStrs - Adds the values represented by two NumStrDto objects and
// returns the result as an NumStrDto.
//
// Input parameter 'ePrefix' is a string consisting of the method chain used to call
// this method. In case of error, this text string is included in the error message.
// Note: Be sure to leave a space at the end of 'ePrefix'.
//
//
// Input parameter 'ePrefix' is a string consisting of the method chain used to call
// this method. In case of error, this text string is included in the error message.
// Note: Be sure to leave a space at the end of 'ePrefix'.
//
func (nDto *NumStrDto) AddNumStrs(
	addend1 NumStrDto,
	addend2 NumStrDto,
	ePrefix string) (
	sum NumStrDto,
	err error) {

	ePrefix += "NumStrDto.AddNumStrs() "

	sum = NumStrDto{}
	err = nil

	var n1DtoSetup, n2DtoSetup NumStrDto

	n1DtoSetup,
	n2DtoSetup, _,
	_,
	err =
		nDto.FormatForMathOps(
			addend1,
			addend2,
			ePrefix)

	if err != nil {
		return sum, err
	}

	newSignVal := n1DtoSetup.signVal

	if n1DtoSetup.signVal != n2DtoSetup.signVal {

		err = n1DtoSetup.SetSignValue(
			1,
			ePrefix + "n1DtoSetup ")

		if err != nil {
			return sum, err
		}

		err = n2DtoSetup.SetSignValue(
			1,
			ePrefix+ "n2DtoSetup ")

		if err != nil {
			return sum, err
		}


		sum,
		err =
			nDto.SubtractNumStrs(
				n1DtoSetup,
				n2DtoSetup,
				ePrefix)

		if err != nil {
			return sum, err
		}

		if nDto.IsNumStrZeroValue(&sum) {
			newSignVal = 1
		}

		err = sum.SetSignValue(
			newSignVal,
			ePrefix + "sum ")

		return sum, err
	}

	precision := n1DtoSetup.precision
	lenN1AllRunes := len(n1DtoSetup.absAllNumRunes)

	n3IntAry := make([]int, lenN1AllRunes+1)
	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0

	for j := lenN1AllRunes - 1; j >= 0; j-- {

		n1 = int(n1DtoSetup.absAllNumRunes[j]) - 48
		n2 = int(n2DtoSetup.absAllNumRunes[j]) - 48

		n3 = n1 + n2 + carry

		carry = 0

		if n3 > 9 {
			n3 = n3 - 10
			carry = 1
		}

		n3IntAry[j+1] = n3

	}

	if carry > 0 {
		n3IntAry[0] = carry
	}

	sum,
	err = nDto.FindIntArraySignificantDigitLimits(
		n3IntAry,
		precision,
		newSignVal,
		ePrefix)

	return sum, err
}

// Compare - compares the signed numeric values of two NumStrDto
// objects. The first object in the comparison is the current
// NumStrDto hereinafter referred to as 'nDto'. The second object
// in the comparison is an instance of NumStrDto passed as an input
// parameter labeled 'nDto2'.  'nDto' is compared to 'n2Dto' and the
// comparison result is returned as an integer value.
//
//     Return Values:
//     -1 = n1Dto is less than n2Dto
//      0 = n1Dto is equal to n2Dto
//     +1 = n1Dto is greater than n2Dto
//
// Examples:
//                          Integer
//                         Comparison
//   n1Dto     n2Dto         Result
//  -------------------------------
//  -9691.23   91.245         -1
//   9691.23   91.245          1
//     -5      82             -1
//      5       5              0
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  n2Dto               *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE data values of internal member variables to
//       achieve the method's objectives.
//
//       This numeric value of 'n2Dto' will be compared to that
//       of the current NumStrDto instance, 'nDto'.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  compareResult       int
//     - This comparison result is generated by comparing the numeric
//       values of input parameters 'n1Dto' and 'n2Dto'. If the method
//       completes successfully, this return parameter will be set to
//       one of three values thereby signaling the results of this
//       comparison.
//
//             'compareResult' Values
//         -1 = n1Dto is less than n2Dto
//          0 = n1Dto is equal to n2Dto
//         +1 = n1Dto is greater than n2Dto
//
//
//  err                error
//     - If this method completes successfully, the returned error Type
//       is set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. Said text will be prefixed
//       to the beginning of the error message.
//
func (nDto *NumStrDto) Compare(
	n2Dto *NumStrDto,
	ePrefix string) (
	compareResult int,
	err error) {

	ePrefix += "NumStrDto.Compare() "
	nStrDtoMolecule := numStrDtoMolecule{}

	compareResult,
	err =
		nStrDtoMolecule.compareNumStrDtoSignedValues(
			nDto,
			n2Dto,
			"")

	return compareResult, err
}

// CompareSignedValues - compares the signed numeric values of two
// NumStrDto objects. The two NumStrDto objects to be compared
// are passed as input parameters 'n1Dto' and 'n2Dto'. 'n1Dto' is
// compared to 'n2Dto' and the comparison result is returned as an
// integer value.
//
//     Return Values:
//     -1 = n1Dto is less than n2Dto
//      0 = n1Dto is equal to n2Dto
//     +1 = n1Dto is greater than n2Dto
//
// Examples:
//                          Integer
//                         Comparison
//   n1Dto     n2Dto         Result
//  -------------------------------
//  -9691.23   91.245         -1
//   9691.23   91.245          1
//     -5      82             -1
//      5       5              0
//
// If either 'n1Dto' or 'n2Dto' are judged to be invalid, this
// method will return an error.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  n1Dto               *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE data values of internal member variables to
//       achieve the method's objectives.
//
//       This numeric value of this NumStrDto instance will be
//       compared to that of the second input parameter, 'n2Dto'.
//
//       If NumStrDto instance 'n1Dto' is judged to be invalid an
//
//
//  n2Dto               *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE data values of internal member variables to
//       achieve the method's objectives.
//
//       This numeric value of this NumStrDto instance will be
//       compared to that of the second input parameter, 'n2Dto'.
//
//       If NumStrDto instance 'n2Dto' is judged to be invalid an
//       error will be returned.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  compareResult       int
//     - This comparison result is generated by comparing the numeric
//       values of input parameters 'n1Dto' and 'n2Dto'. If the method
//       completes successfully, this return parameter will be set to
//       one of three values thereby signaling the results of this
//       comparison.
//
//             'compareResult' Values
//         -1 = n1Dto is less than n2Dto
//          0 = n1Dto is equal to n2Dto
//         +1 = n1Dto is greater than n2Dto
//
//
//  err                error
//     - If this method completes successfully, the returned error Type
//       is set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. Said text will be prefixed
//       to the beginning of the error message.
//
func (nDto *NumStrDto) CompareSignedValues(
	n1Dto *NumStrDto,
	n2Dto *NumStrDto,
	ePrefix string) (
	compareResult int,
	err error) {

	ePrefix += "NumStrDto.CompareSignedValues() "

	nStrDtoMolecule := numStrDtoMolecule{}

	compareResult,
	err =
		nStrDtoMolecule.compareNumStrDtoSignedValues(
			n1Dto,
			n2Dto,
			"")

	return compareResult, err
}

// CompareAbsoluteValues - compares the absolute numeric values
// of two NumStrDto objects. The signs (+ or -) of the two
// compared numeric values are ignored. Only the absolute
// numeric values are compared.
//
// The returned integer value will specify one of the three
// possible outcomes from this comparison:
//
//  -1 = n1Dto absolute value is less than n2Dto absolute value
//   0 = n1Dto absolute value is equal to n2Dto absolute value
//  +1 = n1Dto absolute value is greater than n2Dto absolute value
//
// Examples:
//                                Comparison
//   n1Dto          n2Dto           Result
//  ----------------------------------------
//
//  -9691.23        91.245            1
//   9691.23        91.245            1
//     -5           82               -1
//      5            5                0
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  n1Dto               *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       The absolute numeric value of this NumStrDto instance
//       will be compared to that of input parameter 'n2Dto'.
//
//       If this instance of NumStrDto is judged to be invalid, the method
//       will return an error.
//
//
//  n2Dto               *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       The absolute numeric value of this NumStrDto instance
//       will be compared to that of input parameter 'n2Dto'.
//
//       If this instance of NumStrDto is judged to be invalid, the method
//       will return an error.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  compareResult      int
//     - If this method completes successfully, the value of this returned
//       integer will specify the comparison result obtained from
//       comparing the absolute numeric value of 'n1Dto' to 'n2Dto'.
//       The comparison result will report one of three possible
//       outcomes:
//
//           Comparison
//             Result
//               -1    = n1Dto is less than n2Dto
//                0    = n1Dto is equal to n2Dto
//               +1    = n1Dto is greater than n2Dto
//
func (nDto *NumStrDto) CompareAbsoluteValues(
	n1Dto *NumStrDto,
	n2Dto *NumStrDto,
	ePrefix string) (
	compareResult int,
	err error) {

	ePrefix += "NumStrDto.CompareAbsoluteValues() "

	nStrDtoAtom := numStrDtoAtom{}

	compareResult,
	err = nStrDtoAtom.compareNumStrDtoAbsoluteValues(
		n1Dto,
		n2Dto,
		"")

	if err != nil {
		compareResult = -99
	}

	return compareResult, err
}

// CopyIn - Receives an incoming NumStrDto object
// and copies the information to the current NumStrDto
// data fields.
func (nDto *NumStrDto) CopyIn(nInDto NumStrDto) {

	nStrDtoElectron := numStrDtoElectron{}

	_ = nStrDtoElectron.copyIn(
		nDto,
		&nInDto,
		"")

}

// CopyOut - Creates a copy of the current
// NumStrDto fields and returns a completely
// new instance of NumStrDto
func (nDto *NumStrDto) CopyOut() NumStrDto {

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto,
	err := nStrDtoElectron.copyOut(
		nDto,
		"")

	if err != nil {
		return NumStrDto{}
	}

	return newNumStrDto
}

// Equal - Returns true if the input parameter 'n2Dto' has a numeric
// value instance is equal to the numeric value of the current
// NumStrDto instance, 'nDto'.
//
// If 'nDto' or 'n2Dto' are judged to be invalid, this method will
// return an error.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  n2Dto               *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE data values of internal member variables to
//       achieve the method's objectives.
//
//       This numeric value of this NumStrDto instance will be
//       compared to that of the c, 'n2Dto'.
//
//       If NumStrDto instance 'n2Dto' is judged to be invalid an
//       error will be returned.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  areEqual            bool
//     - After comparing the current NumStrDto instance, 'nDto' to
//       the input parameter 'n2Dto', this returned boolean value
//       will be set to 'true' if the two signed numeric values are
//       equal. Otherwise the two numeric values are unequal and the
//       value is set to 'false'.
//
//
//  err                error
//     - If this method completes successfully, the returned error Type
//       is set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. Said text will be prefixed
//       to the beginning of the error message.
//
func (nDto *NumStrDto) Equal(
	n2Dto *NumStrDto,
	ePrefix string) (
	areEqual bool,
	err error) {

	ePrefix += "NumStrDto.Equal() "

	nStrDtoMolecule := numStrDtoMolecule{}

	var compareResult int

	compareResult,
	err = nStrDtoMolecule.compareNumStrDtoSignedValues(
		nDto,
		n2Dto,
		"")

	if err != nil {
		areEqual = false
	} else if compareResult != 0 {
		areEqual = false
	} else {
		areEqual = true
	}

	return areEqual, err
}

// Empty - Sets all the fields in the NumStrDto
// to their initial or zero state.
func (nDto *NumStrDto) Empty() {

	nStrDtoElectron := numStrDtoElectron{}

	_ =	nStrDtoElectron.empty(
		nDto,
		"")

	return
}

// FindIntArraySignificantDigitLimits - Receives an array of integers and converts them
// to a number string consisting of significant digits. Leading and trailing zeros are
// eliminated.
//
// See related method: NumStrDto{}.FindNumStrSignificantDigitLimits()
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  intArray            []int
//     - An array of integer representing single numeric digits in a number
//       string. If any member values are less than zero ('0'), an error will
//       be returned. Likewise, if any member values are greater than nine
//       ('9'), an error will be returned.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place.
//
//
//  signVal         int
//     - Valid values for this parameter are plus one (+1) or minus
//       one (-1). This number sign value will determine the number
//       sign of the new NumStrDto instance returned by this method.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - If this method completes successfully, the result of the
//       significant digits operation performed by this method will
//       be returned in the form of a new 'NumStrDto' instance.
//
//       The new NumStrDto instance will be configured with default
//       USA numeric separators. To change the numeric separators
//       call one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) FindIntArraySignificantDigitLimits(
	intArray []int,
	precision uint,
	signVal int,
	ePrefix string) (NumStrDto, error) {

	ePrefix += "NumStrDto.FindIntArraySignificantDigitLimits() "

	nStrDtoMech := numStrDtoMechanics{}

	return nStrDtoMech.findIntArraySignificantDigitLimits(
		intArray,
		precision,
		signVal,
		ePrefix)
}

// FindSignificantDigitLimits - Analyzes an array of characters which
// constitute a number string and returns the significant digits in
// the form of a new NumStrDto instance. This operation will effectively
// eliminate leading zeros from the integer value and trailing zeros
// from the fractional value.
//
// See the section below on Example Usage.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  absAllRunes         []rune
//     - An array of characters or runes containing the numeric
//       digits to be evaluated.
//
//
//  precision       uint
//     - The number of numeric digits to the right of the decimal
//       point in the returned new instance of NumStrDto. If the
//       fractional numeric digits to right of the decimal point
//       contain trailing zeros, those trailing zeros will be
//       deleted.
//
//
//  signVal         int
//     - Valid values for this parameter are plus one (+1) or minus
//       one (-1). This number sign value will determine the number
//       sign of the new NumStrDto instance returned by this method.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDto
//     - If this method completes successfully, the result of the
//       significant digits operation performed by this method will
//       be returned in the form of a new 'NumStrDto' instance.
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Example Usage
//
//   <--- Input Parameters -------->   <-- Output -->
//                                       newNumStrDto
//   absAllRunes  precision  signVal        Result
//   -----------------------------------------------
//
//   001236700        4         1          123.67
//   000006700        4         1            0.67
//   001230000        4         1          123.0
//
func (nDto *NumStrDto) FindNumStrSignificantDigitLimits(
	absAllRunes []rune,
	precision uint,
	signVal int,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	ePrefix += "NumStrDto.FindNumStrSignificantDigitLimits() "


	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		nDto,
		ePrefix + "nDto ")

	if err != nil {
		return newNumStrDto, err
	}

	nStrDtoNanobot := numStrDtoNanobot{}

	newNumStrDto,
	err = nStrDtoNanobot.findSignificantDigitLimits(
		numSepsDto,
		absAllRunes,
		precision,
		signVal,
		ePrefix)

	return newNumStrDto, err
}

// FormatForMathOps - receives two NumStrDto objects and converts their strings
// such that both have the same number of integer and fractional digits. This will
// facilitate the performance of string based math operations such as addition and
// subtraction.
//
// The return values represent the formatted NumStrDto objects. The first NumStrDto
// returned always contains the larger absolute value. The second NumStrDto always
// contains the absolute numeric value which is less than or equal to the first
// NumStrDto object returned.
//
// The third parameter returned by this method is an int which will always be set to
// 1 or 0. 1 indicates that the absolute value of the first NumStrDto returned by
// this method is greater than the second NumStrDto returned by this method. If
// the int value returned is zero, it signals that the absolute values
// (not the signed values) of both returned NumStrDto objects are equal.
//
// Input parameter 'ePrefix' is a string consisting of the method chain used to call
// this method. In case of error, this text string is included in the error message.
// Note: Be sure to leave a space at the end of 'ePrefix'.
//
func (nDto *NumStrDto) FormatForMathOps(
	n1Dto,
	n2Dto NumStrDto,
	ePrefix string) (
	n1DtoOut NumStrDto,
	n2DtoOut NumStrDto,
	compare int,
	isOrderReversed bool,
	err error) {

	ePrefix += "NumStrDto.FormatForMathOps() "

	nStrDtoMolecule := numStrDtoMolecule{}

	n1DtoOut,
		n2DtoOut,
		compare,
		isOrderReversed,
		err = nStrDtoMolecule.formatForMathOps(
			&n1Dto,
			&n2Dto,
			ePrefix)

	return n1DtoOut, n2DtoOut, compare, isOrderReversed, err
}

// FormatCurrencyStr - Formats the current NumStrDto numeric value as a currency string.
//
// If the Currency Symbol was not previously set for this NumStrDto, the currency symbol
// is defaulted to the USA standard dollar sign, ('$').
//
// If the Decimal Separator was not previously set for this NumStrDto, the Decimal Separator
// is defaulted to the USA standard period ('.').
//
// If the Thousands Separator was not previously set for this NumStrDto, the Thousands
// Separator is defaulted to the USA standard comma (',').
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  negValMode     NegativeValueFmtMode
//     - Specifies the display mode for negative values:
//
//       LEADMINUSNEGVALFMTMODE   - Negative values formatted with a leading minus sign.
//                                  Example: -$123,456.78
//
//        PARENTHESESNEGVALFMTMODE - Negative values formatted with surrounding parentheses.
//                                    Example: ($123,456.78)
//
//        NumStrDto constants are located in source file:
//               datetime/numstrdtoconstants.go
//
//
//  ePrefix        string
//     - Input parameter 'ePrefix' is a string consisting of the method chain used to call
//       this method. In case of error, this text string is included in the error message.
//       Note: Be sure to leave a space at the end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - If this method completes successfully, this string will be populated
//       with the numeric value extracted from input parameter 'numStrDto'
//       formatted as a currency string incorporating the appropriate currency
//       symbol.
//
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) FormatCurrencyStr(
	negValMode NegativeValueFmtMode,
	ePrefix string) (
	string, error) {

	ePrefix += "NumStrDto.FormatCurrencyStr() "

	nStrDtoAtom := numStrDtoAtom{}

	return nStrDtoAtom.formatCurrencyStr(
		nDto,
		negValMode,
		ePrefix)
}

// FormatNumStr - Formats the numeric value of the current NumStrDto
// as number string consisting of integer digits to the left of the
// decimal point plus fractional digits to the right of the decimal
// point, if such fractional digits exist. The resulting number string
// will NOT contain a currency symbol or thousands separators.
//
// Example: 123456.789
//
// Negative values will be formatted in accordance with input parameter
// 'negValMode'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  negValMode          NegativeValueFmtMode
//     - Specifies the display mode for negative values:
//
//       LEADMINUSNEGVALFMTMODE   - Negative values formatted with
//                                  a leading minus sign.
//                                  Example: -123456.78
//
//       PARENTHESESNEGVALFMTMODE - Negative values formatted with
//                                  surrounding parentheses.
//                                  Example: (123456.78)
//
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  numStr              string
//     - A formatted number string.
//
//
//  err                 error
//     - If this method completes successfully the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) FormatNumStr(
	negValMode NegativeValueFmtMode,
	ePrefix string) (
	numStr string,
	err error) {

	ePrefix += "NumStrDto.FormatNumStr() "

	nStrDtoAtom := numStrDtoAtom{}

	numStr,
	err = nStrDtoAtom.formatNumStr(
		nDto,
		negValMode,
		ePrefix)

	return numStr, err
}

// FormatThousandsStr - Returns the number string delimited with the
// nDto.thousandsSeparator character plus the Decimal Separator if
// applicable.
//
// If the Currency Symbol was not previously set for this NumStrDto,
// the currency symbol is defaulted to the USA standard dollar sign,
// ('$').
//
// If the Decimal Separator was not previously set for this NumStrDto,
// the Decimal Separator is defaulted to the USA standard period ('.').
//
// If the Thousands Separator was not previously set for this NumStrDto, the Thousands
// Separator is defaulted to the USA standard comma (',').
//
// Example:
// thousandsStr = 1000000.234 converted to 1,000,000.234
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  negValMode         NegativeValueFmtMode
//     - Specifies the display mode for negative values:
//
//       LEADMINUSNEGVALFMTMODE   - Negative values formatted with
//                                  a leading minus sign.
//                                  Example: -123,456.78
//
//       PARENTHESESNEGVALFMTMODE - Negative values formatted with
//                                  surrounding parentheses.
//                                  Example: (123,456.78)
//
//        NumStrDto constants are located in source file:
//               datetime/numstrdtoconstants.go
//
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - If this method completes successfully, this string will contain
//       the designated numeric value properly formatted with the correct
//       thousands separator.
//        Example:
//        thousandsStr = 1000000.234 converted to 1,000,000.234
//
//
//  error
//     - If this method completes successfully the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) FormatThousandsStr(
	negValMode NegativeValueFmtMode,
	ePrefix string) (
	string,
	error) {

	ePrefix += "NumStrDto.FormatThousandsStr() "

	nStrDtoAtom := numStrDtoAtom{}

	return nStrDtoAtom.formatThousandsStr(
		nDto,
		negValMode,
		ePrefix)
}

// GetAbsoluteBigInt - Returns the absolute value of all numeric
// digits in the number string (nDto.absAllNumRunes). As such,
// Fractional digits to the right of the decimal are included
// in the consolidate integer number. All of the numeric digits
// in the number string are therefore returned as a *big.Int
// This method will fail if the NumStrDto has not been properly
// initialized with a valid number string.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *big.Int
//     - If this method completes successfully, the current NumStrDto instance
//       will be converted to and returned as a type *big.Int numeric value.
//
//
//  error
//     - If this method completes successfully the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) GetAbsoluteBigInt(
	ePrefix string) (
	*big.Int,
	error) {

	ePrefix += "NumStrDto.GetAbsoluteBigInt() "

	nStrDtoAtom := numStrDtoAtom{}

	return nStrDtoAtom.getAbsoluteBigInt(
		nDto,
		ePrefix)
}

// GetAbsAllNumRunes - Returns an array of runes representing
// all of the integer and fractional digits included in the
// current NumStrDto instance. The rune array returned will
// consist of numeric digits with no sign value prefixed. This
// effectively returns the absolute value of all integer and
// fractional digits combined in one rune array (there is no
// decimal point).
//
// If this NumStrDto instance is judged to be invalid, a zero
// length rune array will be returned.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  absAllRunes         []rune
//     - This method will an array of runes (a.k.a characters). This 'rune'
//       array consists of all the numeric digits contained in and extracted
//       from the current NumStrDto instance.
//
//       The returned series of numeric character digits therefore represents
//       an absolute or positive value. It never presents with a negative
//       value.
//
func (nDto *NumStrDto) GetAbsAllNumRunes() (
	absAllRunes []rune) {

	nStrDtoAtom := numStrDtoAtom{}

	absAllRunes,
	_ = nStrDtoAtom.getAbsoluteValueAllNumRunes(
		nDto,
		"")

	return absAllRunes
}

// GetAbsFracRunes - Returns all of the fractional digits
// to the right of the decimal place in the current NumStrDto
// instance as an array of runes. The rune array is not signed;
// that is, the rune array does not contain a '+' or '-' character
// in the first array position. The rune array is therefore said
// to represent the absolute value of the fractional digits in the
// current NumStrDto numeric value.
//
func (nDto *NumStrDto) GetAbsFracRunes() (
	absFracRunes []rune) {

	nStrDtoElectron := numStrDtoElectron{}

	absFracRunes,
	_ = nStrDtoElectron.getAbsFracRunes(
		nDto,
		"")

	return absFracRunes
}

// GetAbsFracRunesLength - Returns the length of the
// fractional digits in the number string.
//
func (nDto *NumStrDto) GetAbsFracRunesLength() int {

	nStrDtoElectron := numStrDtoElectron{}

	return nStrDtoElectron.getAbsFracRunesLength(nDto)
}

// GetAbsIntRunes - Returns all of the integer digits included
// in the current NumStrDto numeric value as an array of runes.
// The returned rune array does not contain a sign value in the
// first position and therefore represents the absolute or positive
// value of all the integer digits. The integer digits of a NumStrDto
// numeric includes all of the digits to the left of the decimal point.
//
// If the current NumStrDto consists of zero integers and fractional
// digits (Example: '0.123456'), this method will return a rune array
// consisting one array element with a '0' value.
//
func (nDto *NumStrDto) GetAbsIntRunes() []rune {

	nStrDtoElectron := numStrDtoElectron{}

	absIntRunes,
	_:=
		nStrDtoElectron.getAbsIntRunes(
		nDto,
		"")

	return absIntRunes
}

// GetAbsIntRunesLength - Returns the length of the
// integer portion of the number string.
//
func (nDto *NumStrDto) GetAbsIntRunesLength() int {

	nStrDtoElectron := numStrDtoElectron{}

	return nStrDtoElectron.getAbsIntRunesLength(nDto)
}

// GetAbsNumStr - Returns all digits in the current NumStrDto numeric
// value as a pure, unsigned number string. If fractional digits exists
// they are included and NOT separated by a decimal separator.
//
//     Examples:
//
//      numstr      result
//      ------      ------
//      123.45      12345
//     -123.45      12345
//
//
// IMPORTANT
//
// This method does NOT validate the current NumStrDto instance
// before returning the value. To run a validity check on the
// current NumStrDto instance first call one of the two following
// methods:
//
//  NumStrDto.IsValidInstanceError() bool
//                OR
//  NumStrDto.IsValidInstanceError(ePrefix string) error
//
func (nDto *NumStrDto) GetAbsNumStr() (
	absValNumStr string) {

	nStrDtoElectron := numStrDtoElectron{}

	var err error
	absValNumStr = ""

	absValNumStr,
	err = nStrDtoElectron.getAbsoluteValueNumStr(
		nDto,
		"NumStrDto.GetAbsNumStr() ")

	if err != nil {
		absValNumStr = err.Error()
	}

	return absValNumStr
}

// GetBigInt - returns a integer of type *big.Int representing
// the signed integer value of the current NumStrDto instance.
// Decimal numbers like '-123.456' will be returned as signed
// integer values, '-123456'.
//
// This method will fail if the NumStrDto has not been properly initialized
// with a valid number string.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bigIntNum           *big.Int
//     - If this method completes successfully, the numeric value of the
//       current NumStrDto instance will converted to and returned as a
//       type *big.Int integer value.
//
//
//  err                 error
//     - If this method completes successfully the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) GetBigInt(
	ePrefix string) (
	bigIntNum *big.Int,
	err error) {

	ePrefix += "NumStrDto.GetBigInt()"

	nStrDtoMolecule := numStrDtoMolecule{}

	bigIntNum,
	err = nStrDtoMolecule.getSignedBigIntNum(
		nDto,
		ePrefix)

	return bigIntNum, err
}

// GetCurrencySymbol - Returns the character currently designated
// as the currency symbol for this number string.
//
// In the USA, the currency symbol is the dollar sign ('$').
//
// For a list of Major Currency Unicode Symbols, see constants
// located in:
//   datetime/numstrdtoconstants.go
//
// Example: $123.45
//
func (nDto *NumStrDto) GetCurrencySymbol() rune {

	nStrDtoElectron := numStrDtoElectron{}

	currencySymbol,
	_ := nStrDtoElectron.getCurrencySymbol(
		nDto,
		"")

	return currencySymbol

}

// GetCurrencyParen - Returns the number string delimited with the
// nDto.thousandsSeparator character and the currency symbol. If
// the value is negative, the number will be surrounded in parentheses.
//
// Example:
// numstr = 1000000.23
// GetCurrencyParen(ePrefix) = $1,000,000.23
//
// numstr = -1000000.23
// GetCurrencyParen(ePrefix) = ($1,000,000.23)
//
// Note: If the current NumStrDto is invalid, this method
// returns an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  currencyParenStr    string
//     - If this method completes successfully, the numeric value of the
//       current NumStrDto instance will be formatted with a currency
//       symbol and thousands separators. Additionally, if the numeric
//       value is negative, the text will be surrounded with parentheses.
//       Example:
//         numstr = 1000000.23
//         GetCurrencyParen(ePrefix) = $1,000,000.23
//
//         numstr = -1000000.23
//         GetCurrencyParen(ePrefix) = ($1,000,000.23)
//
//
//  err                 error
//     - If this method completes successfully the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) GetCurrencyParen(
	ePrefix string) (
	currencyParenStr string,
	err error) {

	ePrefix += "NumStrDto.GetCurrencyParen() "

	nStrDtoAtom := numStrDtoAtom{}

	currencyParenStr,
	err = nStrDtoAtom.formatCurrencyStr(
		nDto,
		PARENTHESESNEGVALFMTMODE,
		ePrefix)

	return currencyParenStr, err
}

// GetCurrencyStr - Returns the number string delimited with the
// nDto.thousandsSeparator character and the currency symbol.
// If the value is negative, a leading minus sign will be prefixed
// to the currency display.
//
// Example:
//  numstr = 1000000.23
//  GetCurrencyStr(ePrefix) = $1,000,000.23
//
//  numstr = -1000000.23
//  GetCurrencyStr(ePrefix) = -$1,000,000.23
//
// Note: If the current NumStrDto is invalid, this method
// returns an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  currStr             string
//     - If this method completes successfully, the numeric value encapsulated
//       by this NumStrDto instance will be returned and formatted with a
//       currency symbol and thousands separators. Additionally, if the
//       numeric value is negative, this text string will be prefixed with a
//       leading minus sign.
//
//       Example:
//        numstr = 1000000.23
//        GetCurrencyStr(ePrefix) = $1,000,000.23
//
//        numstr = -1000000.23
//        GetCurrencyStr(ePrefix) = -$1,000,000.23
//
//
//  err                 error
//     - If this method completes successfully the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) GetCurrencyStr(
	ePrefix string) (
	currStr string,
	err error) {

	ePrefix += "NumStrDto.GetCurrencyStr() "

	nStrDtoAtom := numStrDtoAtom{}

	currStr,
	err = nStrDtoAtom.formatCurrencyStr(
		nDto,
		LEADMINUSNEGVALFMTMODE,
		ePrefix)

	return currStr, err
}

// GetDecimalSeparator - returns the character designated
// as the decimal separator for the current NumStrDto instance.
//
// In the USA, the decimal separator is the period character ('.').
//
// Example:  123.456
//
func (nDto *NumStrDto) GetDecimalSeparator() rune {

	nStrDtoElectron := numStrDtoElectron{}

	decimalSeparator,
	_ :=  nStrDtoElectron.getDecimalSeparator(nDto, "")

	return decimalSeparator
}

// GetNumericSeparatorsDto - Returns a structure containing the
// character or rune values for decimal point separator, thousands
// separator and currency symbol.
//
func (nDto *NumStrDto) GetNumericSeparatorsDto() NumericSeparatorDto {

	nStrDtoAtom := numStrDtoAtom{}

	numSeps,
	_ := nStrDtoAtom.getNumericSeparatorsDto(
		nDto,
		"")

	return numSeps
}

// GetNumParen - Returns the numeric value of the current NumStrDto
// instance as a signed number string. The resulting number string
// will NOT contain a currency symbol or thousands separators. It
// will contain a decimal separator and fractional digits if such
// fractional digits exist.
//
// Note: If the current NumStrDto is invalid, this method will return
// an error string.
//
// If the sign of the numeric value is negative, the resulting number
// string will be surrounded in parentheses.
//
// Examples:
//  numeric value     result
//   123456.78        123456.78
//  -123456.78       (123456.78)
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  numParenStr         string
//     - If this method completes successfully and the numeric value
//       encapsulated by the current NumStrDto instance is negative, that
//       value will be formatted with surrounding parentheses.
//
//       Examples:
//        numeric value     result
//         123456.78        123456.78
//        -123456.78       (123456.78)
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) GetNumParen(
	ePrefix string) (
	numParenStr string,
	err error) {

	ePrefix += "NumStrDto.GetNumParen() "

	nStrDtoAtom := numStrDtoAtom{}

	numParenStr, err = nStrDtoAtom.formatNumStr(
		nDto,
		PARENTHESESNEGVALFMTMODE,
		ePrefix)

	return numParenStr, err
}

// GetNumStr - returns the numeric value of the current NumStrDto
// instance as a signed number string. The resulting number string
// will NOT contain a currency symbol or thousands separators. It
// will contain a decimal separator and fractional digits if such
// fractional digits exist.
//
// Note: If the current NumStrDto is invalid, this method will return
// an error.
//
// Examples:
//   123456.78
//  -123456.78
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  numStr              string
//     - If this method completes successfully, the numeric value of the
//       current NumStrDto instance will be returned as a plain number
//       string with no thousands separator or currency symbol. Negative
//       values will be formatted with a leading minus sign.
//
//       Examples:
//         123456.78
//        -123456.78
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) GetNumStr(
	ePrefix string) (
	numStr string,
	err error) {

	ePrefix += "NumStrDto.GetNumStr() "

	nStrDtoMolecule := numStrDtoMolecule{}

	numStr,
	err = nStrDtoMolecule.getNumStr(
		nDto,
		ePrefix )

	return numStr, err
}

// GetNumStrDto - Returns a deep copy of the current NumStrDto
// instance.
//
// The returned NumStrDto instance will contain numeric
// separators (decimal separator, thousands separator
// and currency symbol) copied from the current NumStrDto
// instance.
//
// Before returning the NumStrDto result, this method
// performs a validity test on the current NumStrDto instance.
//
// This method is necessary in order to fulfill the requirements
// of the INumMgr interface.
//
func (nDto *NumStrDto) GetNumStrDto() (
	newNumStrDto NumStrDto,
	err error) {

	ePrefix := "NumStrDto.GetNumStrDto() "

	nStrDtoElectron := numStrDtoElectron{}

	_,
	err = nStrDtoElectron.testNumStrDtoValidity(
		nDto,
		ePrefix)

	if err != nil {
		return newNumStrDto, err
	}

	newNumStrDto,
	err = nStrDtoElectron.copyOut(
		nDto,
		ePrefix)

	return newNumStrDto, err
}

// GetPrecisionInt - Returns the precision of the current NumStrDto
// instance.
//
// 'precision', as defined here, specifies the number of numeric
// digits to the right of the decimal place. To compute the
// location of the decimal point in a string of numeric digits, go
// to the right most digit in the number string and count left
// 'precision' digits.
//
// The value of 'precision' returned by this method will always be
// greater than or equal to zero '0' ( precision >= zero ).
//
// Internally, type NumStrDto stores 'precision' as an unsigned
// integer value. This method converts and returns 'precision' as
// an integer value for purposes of convenience.
//
//
// Example:
//   1.234     GetPrecisionInt() = 3
//   5         GetPrecisionInt() = 0
//   0.12345   GetPrecisionInt() = 5
//
//    Number                    Fractional
//    String     precision        Number
//    123456         3            123.456
//
// IMPORTANT
//
// This method does NOT validate the current NumStrDto instance
// before returning the value. To run a validity check on the
// current NumStrDto instance, first call one of the two following
// methods:
//
//  NumStrDto.IsValidInstanceError() bool
//                OR
//  NumStrDto.IsValidInstanceError(ePrefix string) error
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  precision           int
//     - This method returns the 'precision' value associated with
//       the current NumStrDto instance as a type int.
//
func (nDto *NumStrDto) GetPrecisionInt() (
	precision int) {

	var precisionUint uint
	var err error

	nStrDtoElectron := numStrDtoElectron{}

	precisionUint,
	err = nStrDtoElectron.getPrecisionUint(
		nDto,
		"")

	if err != nil {
		precision = 0
	} else {
		precision = int(precisionUint)
	}

	return precision
}

// GetPrecisionUint - Returns the precision of the
// current NumStrDto Instance as an unsigned integer
// (uint). precision represents the number of fractional
// digits to the right of the decimal point.
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal point in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
//  1.234     GetPrecisionInt() = 3
//  5         GetPrecisionInt() = 0
//  0.12345   GetPrecisionInt() = 5
//
//        Number                 Fractional
//        String   precision       Number
//        123456       3          123.456
//
// IMPORTANT
//
// This method does NOT validate the current NumStrDto instance
// before returning the value. To run a validity check on the
// current NumStrDto instance, first call one of the two following
// methods:
//
//  NumStrDto.IsValidInstanceError() bool
//                OR
//  NumStrDto.IsValidInstanceError(ePrefix string) error
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  precision           uint
//     - This method returns the 'precision' value associated with
//       the current NumStrDto instance as a type uint.
//
func (nDto *NumStrDto) GetPrecisionUint() (
	precision uint) {

	nStrDtoElectron := numStrDtoElectron{}

	var err error

	precision,
		err = nStrDtoElectron.getPrecisionUint(
		nDto,
		"")

	if err != nil {
		precision = 0
	}

	return precision
}

// GetRationalNumber - Returns the numeric value of the current
// NumStrDto instance expressed as a type *big.Rat, or Big Rational
// Number. This rational number is a signed value reflecting signed
// numeric value of the the current NumStrDto instance.
//
// For more information on the *big.Rat, rational number type,
// reference:
//   https://golang.org/pkg/math/big/
//
// This method will return an error if the NumStrDto instance is
// judged to be invalid.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bigRatNum           *big.Rat
//     - If this method completes successfully, the signed numerical
//       value of the current NumStrDto instance will be converted and
//       returned in this parameter as a signed type *big.Rat, or Big
//       Rational Number.
//
//       For more information on the *big.Rat, rational number type,
//       reference:
//           https://golang.org/pkg/math/big/
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be prefixed to the beginning of the
//       returned error message.
//
func (nDto *NumStrDto) GetRationalNumber(
	ePrefix string) (
	bigRatNum *big.Rat,
	err error) {

	ePrefix += "NumStrDto.GetRationalNumber() "

	nStrDtoAtom := numStrDtoAtom{}

	bigRatNum,
	err = nStrDtoAtom.getBigRationalNum(
		nDto,
		ePrefix)

	return bigRatNum, err
}

// GetScaleFactor - Returns the scale factor associated with the
// current NumStrDto instance.
//
// Scale factor is defined by 10 raised to the power of the
// 'precision' value associated with the current NumStrDto
// instance (NumStrDto.precision).
//
// 'precision', as defined in connection with type NumStrDto, is
// the number of digits to the right of the decimal point.
//
// For example, if the 'precision' value of the current NumStrDto
// is '5', the resulting scale factor would be computed as follows:
//
//     precision = 5
//     scale factor = 10^5 = 100,000
//
//
// The computed 'Scale Factor' is returned as a type *big.Int. For
// more information on type big.Int, reference:
//           https://golang.org/pkg/math/big/
//
// Be advised that this method will return an error if the current
// NumStrDto instance is judged to be an invalid.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  scaleFactor         *big.Int
//     - Scale factor is defined by 10 raised to the power of nDto.precision.
//       NumStrDto.precision is the number of digits to the right of the
//       decimal point. 'scaleFactor' is returned as a type *big.Int integer
//       value.
//
//       Example:
//       Current NumStrDto instance numerical value = 5.123
//       Precision = 3
//       Scale Factor = 10^3 = 1,000
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) GetScaleFactor(
	ePrefix string) (
	scaleFactor *big.Int,
	err error) {

	ePrefix += "NumStrDto.GetScaleFactor() "

	nStrDtoAtom := numStrDtoAtom{}

	scaleFactor,
	err = nStrDtoAtom.getScaleFactor(
		nDto,
		ePrefix)

	return scaleFactor, err
}

// GetSign - Returns the numeric sign value for the current
// NumStrDto instance. If no processing errors are encountered,
// th return values will be either +1 or -1.
//
//         +1 = a positive number greater than or equal to zero ('0')
//                       - OR -
//         -1 = a negative number less than zero ('0')
//
// If a processing error is encountered, this method will return
// an int value of -99.
//
// IMPORTANT
//
// This method does NOT validate the current NumStrDto instance
// before returning the value. To run a validity check on the
// current NumStrDto instance first call one of the two following
// methods:
//
//  NumStrDto.IsValidInstanceError() bool
//                OR
//  NumStrDto.IsValidInstanceError(ePrefix string) error
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  signValue           int
//     - If this method completes successfully, it will return the
//       numeric sign value associated with the numeric value
//       encapsulated by input parameter 'numStrDto'. The numeric
//       sign value associated with this NumStrDto instance will
//       always be one of three values:
//
//         +1 = a positive number greater than or equal to zero ('0')
//         -1 = a negative number less than zero ('0')
//        -99 = Processing Error
//
//
func (nDto *NumStrDto) GetSign() (
	signValue int) {

	nStrDtoElectron := numStrDtoElectron{}

	var err error

	signValue,
	err = nStrDtoElectron.getNumericSignValue(
		nDto,
		"")

	if err != nil {
		signValue = -99
	}

	return signValue
}

// GetThouParen - Returns the number string delimited with the
// nDto.thousandsSeparator character. Negative values are
// surrounded in parentheses.
//
// Example:
// numstr = 1000000.234
// GetThouParen() = 1,000,000.234
//
// numstr = -1000000.234
// GetThouParen() = (1,000,000.234)
//
// Note: If the current NumStrDto is invalid, this method
// returns an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  thousandsParenStr   string
//     - If this method completes successfully, the returned string will contain
//       the numeric value of the current NumStrDto instance formatted with
//       thousands separators. In addition, negative values will be surrounded with
//       parentheses.
//
//       Example:
//       numstr = 1000000.234
//       GetThouParen() = 1,000,000.234
//
//       numstr = -1000000.234
//       GetThouParen() = (1,000,000.234)
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) GetThouParen(
	ePrefix string) (
	thousandsParenStr string,
	err error) {

	ePrefix += "NumStrDto.GetThouParen() "

	nStrDtoAtom := numStrDtoAtom{}

	return nStrDtoAtom.formatThousandsStr(
		nDto,
		PARENTHESESNEGVALFMTMODE,
		ePrefix)
}

// GetThouStr - Returns the number string delimited with the
// nDto.thousandsSeparator character plus the Decimal Separator
// if applicable.
//
// Example:
// numstr = 1000000.234
// GetThouStr() = 1,000,000.234
//
// numstr = -1000000.234
// GetThouStr() = -1,000,000.234
//
// Note: If the current NumStrDto is invalid, this method
// returns an error.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  thousandsStr        string
//     - If this method completes successfully, the returned string will contain
//       the numeric value of the current NumStrDto instance formatted with
//       thousands separators. In addition, negative values will be prefixed
//       with a leading minus sign ('-').
//
//       Example:
//        numstr = 1000000.234
//        GetThouStr() = 1,000,000.234
//
//        numstr = -1000000.234
//        GetThouStr() = -1,000,000.234
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) GetThouStr(
	ePrefix string) (
	thousandsStr string,
	err error) {

	ePrefix += "NumStrDto.GetThouStr() "

	nStrDtoAtom := numStrDtoAtom{}

	return nStrDtoAtom.formatThousandsStr(
		nDto,
		LEADMINUSNEGVALFMTMODE,
		ePrefix)
}

// GetThousandsSeparator - returns a rune which represents
// the character currently used to separate thousands in
// the display of the current NumStrDto number string.
//
// In the USA, the thousands separator is a comma character.
//
// Example: 1,000,000,000
//
func (nDto *NumStrDto) GetThousandsSeparator() rune {

	nStrDtoElectron := numStrDtoElectron{}

	thousandsSeparator,
	_ := nStrDtoElectron.getThousandsSeparator(
		nDto,
		"")

	return thousandsSeparator
}

// GetZeroNumStrDto - returns a new NumStrDto initialized
// to zero value. If the parameter numFracDigits is set
// to a value greater than zero, then an equal number of
// zero characters will be added to the right of the
// decimal point.
//
// Examples:
//
// Input Parameter     newNumStrDto
//  numFracDigits        Result
//     0                  "0"
//     2                  "0.00"
//
func (nDto *NumStrDto) GetZeroNumStrDto(
	numFracDigits uint) (
	newNumStrDto NumStrDto) {

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto,
	_ = nStrDtoElectron.newZeroNumStrDto(
		nDto,
		numFracDigits,
		"")

	return newNumStrDto
}

// HasNumericDigits - returns 'false' if the number
// string for the current NumStrDto instance is uninitialized
// and contains no numeric digits. In this case, the length
// of the number string is zero characters.
//
// If this method returns 'true' it signals that there is at
// least one numeric digit in the number string, even if that
// digit is zero.
func (nDto *NumStrDto) HasNumericDigits() bool {

	nStrDtoElectron := numStrDtoElectron{}

	_,
	err := nStrDtoElectron.testNumStrDtoValidity(
		nDto,
		"NumStrDto.HasNumericDigits() ")

	if err != nil {
		return false
	}

	return true
}

// IsNumStrZeroValue - Returns 'true' if all the digits in the number
// string for the current NumStrDto instance are zero.
func (nDto *NumStrDto) IsNumStrZeroValue(
	numDto *NumStrDto) (isZeroValue bool) {

	nStrDtoElectron := numStrDtoElectron{}

	isZeroValue,
	_ = nStrDtoElectron.isNumStrZeroValue(
		numDto,
		"")

	return isZeroValue
}

// IsFractionalValue - Returns 'true' if the numeric value of the
// current NumStrDto object includes a fractional value; that is,
// the number has fractional digits to the right of the decimal
// point.
func (nDto *NumStrDto) IsFractionalValue(
	) (isFractionalValue bool) {

	nStrDtoElectron := numStrDtoElectron{}

	isFractionalValue, _ = nStrDtoElectron.isFractionalValue(nDto,"")

	return isFractionalValue
}

// IsValidInstanceError - Performs a diagnostic review of the current NumStrDto
// instance and returns 'nil' if the NumStrDto object is valid in all
// respects.
//
// Input parameter 'ePrefix' is a string consisting of the method chain used to call
// this method. In case of error, this text string is included in the error message.
// Note: Be sure to leave a space at the end of 'ePrefix'.
//
// If the NumStrDto instance is judged invalid, an error message is returned.
func (nDto *NumStrDto) IsValidInstanceError(
	ePrefix string) (
	err error) {

	ePrefix += "NumStrDto.IsValidInstanceError() "

	err = nil

	nStrDtoElectron := numStrDtoElectron{}

	_,
	err = nStrDtoElectron.testNumStrDtoValidity(
		nDto,
		ePrefix)

	return err
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrDto instance and returns 'true' if the NumStrDto object
// is valid in all respects.
//
// If the current NumStrDto instance is judged invalid, this method
// returns 'false'. Otherwise, if the current NumStrDto instance
// qualifies as valid, this method returns 'true'.
//
// For a validity test which returns a detailed error message,
// see method NumStrDto.IsValidInstanceError()
//
func (nDto *NumStrDto) IsValidInstance() (
	isValid bool) {

	nStrDtoElectron := numStrDtoElectron{}
	isValid,
	_ = nStrDtoElectron.testNumStrDtoValidity(
		nDto,
		"")

	return isValid
}

// IsZero - Returns 'true' if the numeric value of the current
// NumStrDto instance is zero.
//
func (nDto *NumStrDto) IsZero(
	) (isZeroValue bool) {

	nStrDtoElectron := numStrDtoElectron{}

	isZeroValue,
	_ = nStrDtoElectron.isNumStrZeroValue(
		nDto,
		"")

	return isZeroValue
}

// Multiply - Multiplies the numeric value of the current NumStrDto
// instance by the numeric value of the input parameter 'multiplier'.
// The result of this multiplication, or product, is then stored
// in the current NumStrDto instance.
//
// Be advised that if the current NumStrDto instance proves invalid,
// an error will be returned.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  multiplier          *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE the values of internal member variables to
//       achieve the method's objectives.
//
//       The numerical value of 'multiplier' will be multiplied
//       by the numerical value of current NumStrDto instance. The
//       result, or product, will then be stored in the current
//       NumStrDto instance.
//
//       If the 'multiplier' instance proves to be invalid, an error
//       will be returned.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err
//     - If this method completes successfully, the returned error Type is
//       set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. The 'ePrefix' text will be
//       prefixed to the beginning of the returned error message.
//
func (nDto *NumStrDto) Multiply(
	multiplier *NumStrDto,
	ePrefix string) error {

	ePrefix += "NumStrDto.Multiply() "

	nStrDtoUtil := numStrDtoUtility{}

	return nStrDtoUtil.multiplyInPlace(
		nDto,
		multiplier,
		ePrefix)
}

// MultiplyNumStrs - Multiplies two NumStrDto instances and returns
// the product as a new NumStrDto instance.
//
// If either of the two input parameters, 'multiplicand' or
// 'multiplier' proves to be invalid, this method will return an
// error.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  multiplicand        *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE the values of internal member variables to
//       achieve the method's objectives.
//
//       The numerical value of 'multiplicand' will be multiplied
//       by input parameter 'multiplier' to generate and return
//       the product as NumStrDto instance.
//
//       If this NumStrDto instance proves to be invalid, this
//       method will return an error.
//
//
//  multiplier          *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT CHANGE the values of internal member variables to
//       achieve the method's objectives.
//
//       The numerical value of 'multiplier' will be multiplied
//       by input parameter 'multiplicand' to generate and return
//       the product as NumStrDto instance.
//
//       If this NumStrDto instance proves to be invalid, this
//       method will return an error.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  product             NumStrDto
//     - If this method completes successfully, the product obtained from
//       multiplying input parameters 'multiplicand' and 'multiplier' will 
//       be returned in a new instance of NumStrDto.
//
//
//  err                error
//     - If this method completes successfully, the returned error Type is
//       set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. The 'ePrefix' text will be
//       prefixed to the beginning of the returned error message.
//
func (nDto *NumStrDto) MultiplyNumStrs(
	multiplicand NumStrDto,
	multiplier NumStrDto,
	ePrefix string) (
	product NumStrDto,
	err error) {

	ePrefix += "NumStrDto.MultiplyNumStrs() "

	nStrDtoHelper := numStrDtoHelper{}

	product,
	err = nStrDtoHelper.multiplyNumStrs(
		&multiplicand,
		&multiplier,
		ePrefix)

	return product, err
}

// New - Used to create and return an empty or zero NumStrDto
// instance. This method initializes the NumStrDto data fields to
// their zero values.
//
// Example Usage
//
// n := NumStrDto{}.New()
// n2, err := n.ParseNumStr("123.456")
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       of zero (0).
//
func (nDto NumStrDto) New() NumStrDto {

	nStrDtoElectron := numStrDtoElectron{}

	return nStrDtoElectron.newBaseZeroNumStrDto(0)
}

// NewBigFloat - Creates a new NumStrDto instance from a Big Float value
// (*big.Float) and a precision specification.
//
// For more information on the *big.Float floating point numeric value,
// reference:
//   https://golang.org/pkg/math/big/
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See the 'Example Usage' section below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  bigFloatNum         *big.Float
//     - A type *big.Float floating point numeric value. For details
//       on type *big.Float, reference:
//         https://golang.org/pkg/math/big/
//
//     This floating point numeric value will be converted to a new
//     instance of type NumStrDto.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place. The final value will be
//       rounded to 'precision' digits after the decimal point.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//     f64Num := float64(123.456)
//     bigFloatNum := big.NewFloat(f64Num)
//     precision := uint(2)
//     ePrefix := "calling method name "
//
//          nDto, err  :=
//              NumStrDto{}.NewBigFloat(
//                bigFloatNum,
//                precision,
//                ePrefix)
//
//           nDto is now equal to 123.46
//
//  Examples:
//  ---------
//                                newNumStrDto
//  bigFloatNum     precision        Result
//  -------------   --------------------------
//
//   12.3456            4               12.3456
//   123456.5           0           123457
//   1234.56            1             1234.6
//
func (nDto NumStrDto) NewBigFloat(
	bigFloatNum *big.Float,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewBigFloat() "


	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newBigFloat(
		numSepsDto,
		bigFloatNum,
		precision,
		ePrefix)
}

// NewBigFloatWithNumSeps - Creates a new NumStrDto instance from a
// Big Float value (*big.Float) and a precision specification.
//
// For more information on the *big.Float floating point numeric value,
// reference:
//   https://golang.org/pkg/math/big/
//
// This method requires numeric separators to be configured in the
// returned 'NumStrDto' instance to be specified through input
// parameter 'numSepsDto'. This allows the use of numeric separators
// from multiple nationalities and cultures.
//
// See the 'Example Usage' section below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numSepsDto          NumericSeparatorDto
//     - An instance of NumericSeparatorDto which will be used to supply
//       the numeric separators for the new NumStrDto instance returned
//       by this method. Numeric separators include the Thousands
//       Separator, Decimal Separator and the Currency Symbol.
//
//       The data fields included in the NumericSeparatorDto are
//       listed as follows:
//
//          type NumericSeparatorDto struct {
//
//            DecimalSeparator   rune // Character used to separate
//                                    //  integer and fractional digits ('.')
//
//            ThousandsSeparator rune // Character used to separate thousands
//                                    //  (1,000,000,000
//
//            CurrencySymbol     rune // Currency Symbol
//          }
//
//       If any of the data fields in this passed structure
//       'customSeparators' are set to zero ('0'), they will
//       be reset to USA default values. USA default numeric
//       separators are listed as follows:
//
//             Currency Symbol: '$'
//         Thousands Separator: ','
//           Decimal Separator: '.'
//
//
//  bigFloatNum         *big.Float
//     - A type *big.Float floating point numeric value. For details
//       on type *big.Float, reference:
//         https://golang.org/pkg/math/big/
//
//     This floating point numeric value will be converted to a new
//     instance of type NumStrDto.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place. The final value will be
//       rounded to 'precision' digits after the decimal point.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will be configured with default
//       numeric separators extracted from input parameter 'numSepsDto'.
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//     f64Num := float64(123.456)
//     bigFloatNum := big.NewFloat(f64Num)
//     precision := uint(2)
//     ePrefix := "calling method name "
//
//          nDto, err  :=
//              NumStrDto{}.NewBigFloat(
//                bigFloatNum,
//                precision,
//                ePrefix)
//
//           nDto is now equal to 123.46
//
//  Examples:
//  ---------
//                                newNumStrDto
//  bigFloatNum     precision        Result
//  -------------   --------------------------
//
//   12.3456            4               12.3456
//   123456.5           0           123457
//   1234.56            1             1234.6
//
func (nDto NumStrDto) NewBigFloatWithNumSeps(
	numSepsDto NumericSeparatorDto,
	bigFloatNum *big.Float,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewBigFloatWithNumSeps() "

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newBigFloat(
		numSepsDto,
		bigFloatNum,
		precision,
		ePrefix)
}

// NewBigInt - Creates and returns a new NumStrDto instance
// extracted from a signed Big Integer number (type *big.Int) and
// a precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place.
//
// Numeric separators include the Currency Symbol, the Decimal
// Delimiter and the Thousands Separator. These separator character
// values vary according to nationality and culture.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  bigIntNum           *big.Int
//     - This numeric value will be converted to a new instance of
//       type NumStrDto. Type 'big.Int' is designed to handle very
//       large integer values. For more information on type 'big.Int',
//       reference:
//          https://golang.org/pkg/math/big/
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place. If 'precision' has a value
//       greater than zero, the returned NumStrDto will be configured
//       as a floating point value.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If the method completes successfully, the returned error Type
//       is set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example #1
//     int64Num := int64(123456)
//     bigIntNum := big.NewInt(int64Num)
//     precision := uint(3)
//     ePrefix := "calling method name "
//
//          nDto, err :=
//            NumStrDto{}.NewBigInt(
//            bigIntNum,
//            precision,
//            ePrefix)
//
//           nDto is now equal to 123.456
//
//  Example #2
//     int64Num := int64(123456)
//     bigIntNum := big.NewInt(int64Num)
//     precision := uint(0)
//     ePrefix := "calling method name "
//
//          nDto, err :=
//            NumStrDto{}.NewBigInt(
//            bigIntNum,
//            precision,
//            ePrefix)
//
//           nDto is now equal to 123456
//
//  Examples:
//  ---------
//
//  <-- Input Parameters -->     <--- Output --->
//                                 newNumStrDto
//  bigIntNum    precision            Result
//  ---------------------------------------------
//
//   123456          4                  12.3456
//   123456          0              123456
//   123456          1               12345.6
//   123456          7                   0.0123456
//
func (nDto NumStrDto) NewBigInt(
	bigIntNum *big.Int,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewBigInt() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
	err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newBigInt(
		numSepsDto,
		bigIntNum,
		precision,
		ePrefix)
}

// NewFloat32 - Creates a new NumStrDto instance from a float32
// and precision specification.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See the Example Usage section below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  f32                float32
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place. The final value will be
//       rounded to 'precision' digits after the decimal point.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  This method is designed to be used in conjunction with the NumStrDto{}
//  syntax thereby allowing NumStrDto type creation and initialization in
//  one step.
//
//           f32 := float32(123.456)
//     precision := uint(2)
//     nDto, err := NumStrDto{}.NewFloat32(
//                  uint64Num,
//                  precision,
//                  "calling method name ")
//
//           nDto is now equal to 123.46
//
//  Examples:
//  ---------
//                                newNumStrDto
//     f32       precision          Result
//  ------------------------------------------
//
//   12.3456         4                  12.3456
//   123456.5        0              123457
//   1234.56         1                1234.6
//
func (nDto NumStrDto) NewFloat32(
	f32 float32,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewFloat32() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newFloat64(
		numSepsDto,
		float64(f32),
		precision,
		ePrefix)
}

// NewFloat64 - Creates a new NumStrDto instance from a float64
// and precision specification.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See the 'Example Usage' section below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  f64                 float64
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place. The final value will be
//       rounded to 'precision' digits after the decimal point.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  This method is designed to be used in conjunction with the NumStrDto{}
//  syntax thereby allowing NumStrDto type creation and initialization in
//  one step.
//
//           f64 := float64(123.456)
//     precision := uint(2)
//     nDto, err := NumStrDto{}.NewFloat64(
//                  uint64Num,
//                  precision,
//                  "calling method name ")
//
//           nDto is now equal to 123.46
//
//  Examples:
//  ---------
//                                newNumStrDto
//     f64       precision          Result
//  ------------------------------------------
//
//   12.3456         4                  12.3456
//   123456.5        0              123457
//   1234.56         1                1234.6
//
func (nDto NumStrDto) NewFloat64(
	f64 float64,
	precision uint,
	ePrefix string) (NumStrDto, error) {

	ePrefix += "NumStrDto.NewFloat64() "


	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newFloat64(
		numSepsDto,
		f64,
		precision,
		ePrefix)
}

// Creates a new NumStrDto from an int and a precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place.
//
// Example: NumStrDto{}.NewInt(123456, 3) yields a new NumStrDto
// instance with a numeric value of 123.456.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
// --------------------------------------------------------------------------------------------------
//
// Input Parameters
//
//  intNum              int
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If the method completes successfully, the returned error Type
//       is set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     intNum := 123456
//     precision := uint(3)
//
//          nDto, err :=
//            NumStrDto{}.NewInt(
//            intNum,
//            precision,
//            "calling method name ")
//
//           nDto is now equal to 123.456
//
//  Examples:
//  ---------
//
//  <-- Input Parameters -->     <--- Output --->
//                                 newNumStrDto
//  intNum      precision             Result
//  ---------------------------------------------
//
//   123456          4                  12.3456
//   123456          0              123456
//   123456          1               12345.6
//   123456          7                   0.0123456
//
func (nDto NumStrDto) NewInt(
	intNum int,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewInt()"

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newInt64(
		numSepsDto,
		int64(intNum),
		precision,
		ePrefix)
}

// NewIntExponent - Returns a new NumStrDto instance. The numeric
// value is set using an integer multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//       numeric value = integer X 10^exponent
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See 'Example Usage' below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  intNum              int
//     - This numeric value will be multiplied by 10^exponent and
//       converted to a new instance of type NumStrDto.
//
//  exponent            int
//     - 10^exponent is multiplied by input parameter 'int64Num' to
//       generate a new instance of type NumStrDto.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  nDto := NumStrDto{}.NewIntExponent(123456, -3)
//  nDto is now equal to "123.456", precision = 3
//
//  nDto := NumStrDto{}.NewIntExponent(123456, 3)
//  decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//                                   NumStrDto
//   intNum        exponent            Result
//   --------------------------------------
//
//   123456           -3                123.456
//   123456            3              123456.000
//   123456            0              123456
//
func (nDto NumStrDto) NewIntExponent(
	intNum int,
	exponent int,
	ePrefix string) (
	NumStrDto,
	error ) {

	ePrefix += "NumStrDto.NewIntExponent() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newInt64Exponent(
		numSepsDto,
		int64(intNum),
		exponent,
		ePrefix)
}

// NewInt32 - Creates a new NumStrDto instance from an int32 and a
// precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place.
//
// The 'NewInt32' method is designed to used in conjunction with
// NumStrDto{} syntax thereby allowing NumStrDto type creation and
// initialization in one step.
//
// Example: NumStrDto{}.NewInt32(123456, 3) yields a new NumStrDto
// instance with a numeric value of 123.456.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  int32Num            int32
//     - An 'int32' integer value which will be converted into a new
//       instance of type NumStrDto.
//
//
//  precision           uint
//     - The number of digits to the right of the decimal place. This
//       will ensure that input parameter 'int32Num' is rounded to the
//       correct 'precision'.
//
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNDto             NumStrDto
//     - If this method completes successfully, a new instance of NumStrDto
//       will be returned which encapsulates the numeric value specified by
//       input parameters 'int32Num' and 'precision'.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nDto NumStrDto) NewInt32(
	int32Num int32,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewInt32() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newInt64(
		numSepsDto,
		int64(int32Num),
		precision,
		ePrefix)
}

// NewInt32Exponent - Returns a new NumStrDto instance. The numeric
// value is set using the 'int32Num' parameter multiplied by 10 raised
// to the power of the 'exponent' parameter.
//
//         numeric value = int32 X 10^exponent
//
// For example, if exponent is -3, precision is set equal to 'int32Num'
// divided by 10^+3. Example:
//
//                          'newNumStrDto'
//   int32Num    exponent       Result
//    123456        -3          123.456
//
// If exponent is +3, int32Num is multiplied by 10 raised to the
// power of exponent and precision is set equal to exponent.
//
//                          'newNumStrDto'
//   int32Num    exponent       Result
//    123456        +3        123456.000
//
// Numeric separators include the Currency Symbol, the Decimal
// Delimiter and the Thousands Separator. These separator character
// values vary according to nationality and culture.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  int32Num            int32
//     - The base numeric value to be adjusted and converted to a
//       new instance of NumStrDto.
//
//  exponent            int
//     - The exponent value to be applied to input parameter 'int32Num'.
//
//
//  ePrefix             string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//     int32Num := int32(123456)
//     exponent = -3
//
//  nDto,err := NumStrDto{}.NewInt32Exponent(
//          int32Num,
//          exponent,
//          "calling method name ")
//   -- nDto is now equal to "123.456", precision = 3
//
//
//     int32Num := int32(123456)
//     exponent = 3
//
//  nDto,err := NumStrDto{}.NewInt32Exponent(
//          int32Num,
//          exponent,
//          "calling method name ")
//
//  -- nDto is now equal to "123456.000", precision = 3
//
//  Examples:
//  ---------
//
//  <---- Input Parameters ---->       <-- Output -->
//                                       newNumStrDto
//   int32Num          exponent            Result
//   123456               -3               123.456
//   123456                3            123456.000
//   123456                0            123456
//
func (nDto NumStrDto) NewInt32Exponent(
	int32Num int32,
	exponent int,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += " NumStrDto.NewInt32Exponent() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newInt64Exponent(
		numSepsDto,
		int64(int32Num),
		exponent,
		ePrefix)
}

// NewInt64 - Creates a new NumStrDto instance from an int64 and a
// precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place.
//
// The 'NewInt64' method is designed to used in conjunction with
// NumStrDto{} syntax thereby allowing NumStrDto type creation and
// initialization in one step.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See the 'Example Usage' section below.
//
//
// --------------------------------------------------------------------------------------------------
//
// Input Parameters
//
//  int64Num            int64
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If the method completes successfully, the returned error Type
//       is set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     int64Num := int64(123456)
//     precision := uint(3)
//
//          nDto, err :=
//            NumStrDto{}.NewUint64(
//            int64Num,
//            precision,
//            "calling method name ")
//
//           nDto is now equal to 123.456
//
//  Examples:
//  ---------
//
//  <-- Input Parameters -->     <--- Output --->
//                                 newNumStrDto
//  int64Num     precision            Result
//  ---------------------------------------------
//
//   123456          4                  12.3456
//   123456          0              123456
//   123456          1               12345.6
//   123456          7                   0.0123456
//
func (nDto NumStrDto) NewInt64(
	int64Num int64,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error)  {

	ePrefix += "NumStrDto.NewInt64() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newInt64(
		numSepsDto,
		int64Num,
		precision,
		ePrefix)
}

// NewInt64Exponent - Returns a new NumStrDto instance. The numeric
// value is set using an int64 value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//    numeric value = int64 X 10^exponent
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See 'Example Usage' below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  int64Num            int64
//     - This numeric value will be multiplied by 10^exponent and
//       converted to a new instance of type NumStrDto.
//
//  exponent            int
//     - 10^exponent is multiplied by input parameter 'int64Num' to
//       generate a new instance of type NumStrDto.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//
//     int64Num := int64(123456)
//     exponent := -3
//
//  nDto,err := NumStrDto{}.NewInt64Exponent(
//              int64Num,
//              exponent)
//
//  -- nDto is now equal to "123.456", precision = 3
//
//
//     int64Num := int64(123456)
//     exponent := 3
//
//  nDto,err := NumStrDto{}.NewInt64Exponent(
//              int64Num,
//              exponent)
//
//  -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//                                Decimal
//   int64Num    exponent          Result
//    123456        -3              123.456
//    123456         3           123456.000
//    123456         0           123456
//
func (nDto NumStrDto) NewInt64Exponent(
	int64Num int64,
	exponent int,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewInt64Exponent() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newInt64Exponent(
		numSepsDto,
		int64Num,
		exponent,
		ePrefix)
}

// NewNumStr - Used to create and return a populated NumStrDto instance
// using a valid number string as an input parameter.
//
// A valid number string 'may' be prefixed with numeric sign value of
// plus ('+') or minus ('-'). The absence of a leading numeric sign
// character will default the numeric value to plus or a positive
// numeric value. A valid number string 'may' also include a decimal
// delimiter such as a decimal point to separate integer and fractional
// digits in the number string. With these two exceptions, all other
// characters in a valid number string must be text characters between
// '0' and '9'.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
// See the 'Example Usage' section below.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStr              string
//     - A valid number string. A valid number string 'may' be
//       prefixed with a numeric sign value of plus ('+') or
//       minus ('-'). The absence of a leading numeric sign
//       character will default the numeric value to plus or a
//       positive numeric value. A valid number string 'may'
//       also include a decimal delimiter such as a decimal
//       point to separate integer and fractional digits
//       within the number string. With these two exceptions,
//       all other characters in a valid number string must be
//       numeric values represented by text characters between
//       '0' and '9'.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//    numStr := "123.456"
//    ePrefix := "calling method name "
//
//    newNumStrDto, err := NumStrDto{}.NewNumStr(
//                numstr,
//                ePrefix)
//
//     - newNumStrDto is now equal to 123.456
//
//
//    numStr := "123456"
//    ePrefix := "calling method name "
//
//    newNumStrDto, err := NumStrDto{}.NewNumStr(
//                numstr,
//                ePrefix)
//
//     - newNumStrDto is now equal to 123456
//
func (nDto NumStrDto) NewNumStr(
	numStr string,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewNumStr() "


	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newNumStr(
		numSepsDto,
		numStr,
		ePrefix)
}

// NewNumStrWithNumSeps - Receives a number string as input and
// returns a new NumStrDto instance. The input parameter
// 'numSepsDto' contains numeric separators (decimal separator,
// thousands separator and currency symbol) which will be used to
// parse the number string.
//
// In addition, the numeric separators contained in input parameter
// 'numSepsDto' will be copied to the returned NumStrDto instance.
//
// This method is identical to method NumStrDto.NewNumStr() with
// the sole exception that this method allows the caller to specify
// numeric separators for nationalities other than the USA.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStr              string
//     - A valid number string. A valid number string 'may' be
//       prefixed with a numeric sign value of plus ('+') or
//       minus ('-'). The absence of a leading numeric sign
//       character will default the numeric value to plus or a
//       positive numeric value. A valid number string 'may'
//       also include a decimal delimiter such as a decimal
//       point to separate integer and fractional digits
//       within the number string. With these two exceptions,
//       all other characters in a valid number string must be
//       numeric values represented by text characters between
//       '0' and '9'.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//
//  numSepsDto          NumericSeparatorDto
//     - An instance of NumericSeparatorDto which will be used to supply
//       the numeric separators for the new NumStrDto instance returned
//       by this method. Numeric separators include the Thousands
//       Separator, Decimal Separator and the Currency Symbol.
//
//       The data fields included in the NumericSeparatorDto are
//       listed as follows:
//
//          type NumericSeparatorDto struct {
//
//            DecimalSeparator   rune // Character used to separate
//                                    //  integer and fractional digits ('.')
//
//            ThousandsSeparator rune // Character used to separate thousands
//                                    //  (1,000,000,000
//
//            CurrencySymbol     rune // Currency Symbol
//          }
//
//       If any of the data fields in this passed structure
//       'customSeparators' are set to zero ('0'), they will
//       be reset to USA default values. USA default numeric
//       separators are listed as follows:
//
//             Currency Symbol: '$'
//         Thousands Separator: ','
//           Decimal Separator: '.'
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//    numStr := "123.456"
//    ePrefix := "calling method name "
//    numSepsDto := NewNumStrWithNumSeps{}
//    numSepsDto.SetToUSADefaults()
//
//    newNumStrDto, err := NumStrDto{}.NewNumStrWithNumSeps(
//                numstr,
//                numSepsDto,
//                ePrefix)
//
//     - newNumStrDto is now equal to 123.456
//
//
//    numStr := "123456"
//    ePrefix := "calling method name "
//    numSepsDto := NewNumStrWithNumSeps{}
//    numSepsDto.SetToUSADefaults()
//
//    newNumStrDto, err := NumStrDto{}.NewNumStrWithNumSeps(
//                numstr,
//                numSepsDto,
//                ePrefix)
//
//     - newNumStrDto is now equal to 123456
//
func (nDto NumStrDto) NewNumStrWithNumSeps(
	numStr string,
	numSepsDto NumericSeparatorDto,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewNumStrWithNumSeps() "


	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newNumStr(
		numSepsDto,
		numStr,
		ePrefix)
}

// NewRational - Creates a new NumStrDto instance from a rational
// number and a precision specification.
//
// For information on Big Rational Numbers (*big.Rat), reference:
//    https://golang.org/pkg/math/big/
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  bigRatNum           *big.Rat
//     - This 'big' Rational Number will be converted into a
//       a returned instance of NumStrDto. The numeric value
//       of the big Rational Number will be represented as
//       a fractional or floating point number with a 'precision'
//       number of digits after the decimal point.
//
//       For more information on type *big.Rat, reference:
//         https://golang.org/pkg/math/big/
//
//
//  precision           uint
//     - The number of digits which will be placed to the right
//       of the decimal point in the returned new instance of
//       NumStrDto. This fractional floating point value will be
//       rounded to 'precision' digits.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (nDto NumStrDto) NewRational(
	bigRat *big.Rat,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewRational() "


	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newRational(
		numSepsDto,
		bigRat,
		precision,
		ePrefix)
}

// NewUint - Creates a new NumStrDto instance from an uint and a
// precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// -----------------------------------------------------------------------
//
// Input Parameters
//
//  uintNum             uint
//     - The numeric value which will be converted and encapsulated
//       in the returned instance of NumStrDto.
//
//
//  precision           uint
//     - The 'precision' input parameter specifies the number of digits
//       to be formatted to the right of the decimal place.
//
//
//  ePrefix             string
//     - Input parameter 'ePrefix' is a string consisting of the method
//       chain used to call this method. In case of error, this text
//       string is included in the error message. Note: Be sure to
//       leave a space at the end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDto
//     - If this method completes successfully, a new instance of NumStrDto
//       will be returned encapsulating a numeric value calculated from the
//       input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//      uintNum := uint(123456)
//      precision := uint(3)
//      nDto := NumStrDto{}.NewUint(uintNum, precision)
//             nDto is now equal to 123.456
//
//
//                                    NumStrDto
//   uintNum       precision           Result
//   -------------------------------------------
//   123456           4                 12.3456
//   123456           0             123456
//   123456           1              12345.6
//
func (nDto NumStrDto) NewUint(
	uintNum uint,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewUint() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newUint64(
		numSepsDto,
		uint64(uintNum),
		precision,
		ePrefix)
}

// NewUintExponent - Returns a new NumStrDto instance. The numeric value
// is set using an uint value multiplied by 10 raised to the power of the
// 'exponent' input parameter.
//
//   numeric value = 'uintNum' X 10^'exponent'
//
// Numeric separators include the Currency Symbol, the Decimal
// Delimiter and the Thousands Separator. These separator character
// values vary according to nationality and culture.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See 'Example Usage' section below.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  uintNum             uint
//     - This uint64 is multiplied by 10^exponent (input parameter exponent)
//       to calculate the final numeric value which is returned in a new
//       instance of NumStrDto.
//
//
//  exponent            int
//     - This is the exponent value. Input parameter uint64Num is multiplied
//       by 10 raised to the power of this 'exponent' parameter in order to
//       calculate the numeric value contained in the new instance of
//       NumStrDto returned by this method.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//    - If this method completes successfully, it will return a new
//       instance of NumStrDto. The numeric value contained in this new
//       instance is calculated by multiplying input parameter
//       'uint64Num' times 10 raised to the power of input parameter
//       'exponent'.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//     uintNum  := uint(123456)
//     exponent := -3
//
//  nDto,err := NumStrDto{}.NewUintExponent(
//          uintNum,
//          exponent,
//          "calling method name ")
//   -- nDto is now equal to "123.456", precision = 3
//
//
//     numSeps.SetToUSADefaults()
//     uintNum := uint(123456)
//     exponent := 3
//
//  nDto,err := NumStrDto{}.NewUintExponent(
///         uintNum,
//          exponent,
//          "calling method name ")
//  -- nDto is now equal to "123456.000", precision = 3
//
//  Examples:
//  ---------
//                                         NumStrDto
//   uintNum           exponent             Result
//   123456               -3               123.456
//   123456                3            123456.000
//   123456                0            123456
//
func (nDto NumStrDto) NewUintExponent(
	uintNum uint,
	exponent int,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewUintExponent() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newUint64Exponent(
			numSepsDto,
			uint64(uintNum),
			exponent,
			ePrefix)
}

// NewUint32 - Creates a new NumStrDto instance from an uint32 and a
// precision specification.
//
// Numeric separators include the Currency Symbol, the Decimal
// Delimiter and the Thousands Separator. These separator character
// values vary according to nationality and culture.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See the Example Usage section below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  uint32Num           uint32
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     uint32Num := uint32(123456)
//     precision := uint(3)
//
//          nDto, err :=
//            NumStrDto{}.NewUint32(
//            uint32Num,
//            precision,
//            "calling method name ")
//
//           nDto is now equal to 123.456
//
//  Examples:
//  ---------

//  <-- Input Parameters -->     <-- Output -->
//                                newNumStrDto
//  uint64Num     precision          Result
//  ------------------------------------------
//
//   123456          4                  12.3456
//   123456          0              123456
//   123456          1               12345.6
//
func (nDto NumStrDto) NewUint32(
	uint32Num uint32,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewUint32() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newUint64(
		numSepsDto,
		uint64(uint32Num),
		precision,
		ePrefix)
}

// NewUint32Exponent - Returns a new NumStrDto instance. The numeric
// value is set using an uint32 value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//      numeric value = uint32Num X 10^exponent
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See the Example Usage section below.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  uint32Num           uint32
//     - This uint32 is multiplied by 10^exponent (input parameter exponent)
//       to calculate the final numeric value which is returned in a new
//       instance of NumStrDto.
//
//
//  exponent            int
//     - This is the exponent value. Input parameter uint64Num is multiplied
//       by 10 raised to the power of this 'exponent' parameter in order to
//       calculate the numeric value contained in the new instance of
//       NumStrDto returned by this method.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDto
//    - If this method completes successfully, it will return a new
//       instance of NumStrDto. The numeric value contained in this new
//       instance is calculated by multiplying input parameter
//       'uint64Num' times 10 raised to the power of input parameter
//       'exponent'.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//     uint32Num := uint32(123456)
//     exponent = -3
//
//  nDto,err := NumStrDto{}.NewUint32Exponent(
//               uint32Num,
//               exponent,
//              "calling method name ")
//   -- nDto is now equal to "123.456", precision = 3
//
//
//     uint32Num := uint32(123456)
//     exponent = 3
//
//  nDto,err := numStrDtoUtility.NewUint32Exponent(
//          uint32Num,
//          exponent,
//          "calling method name ")
//
//  -- nDto is now equal to "123456.000", precision = 3
//
//  Examples:
//  ---------
//
//  <---- Input Parameters ---->       <-- Output -->
//                                       newNumStrDto
//   uint32Num          exponent            Result
//   123456               -3               123.456
//   123456                3            123456.000
//   123456                0            123456
//
func (nDto NumStrDto) NewUint32Exponent(
	uint32Num uint32,
	exponent int,
	ePrefix string) (
	NumStrDto,
	error ) {

	ePrefix += "NumStrDto.NewUint32Exponent() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newUint64Exponent(
		numSepsDto,
		uint64(uint32Num),
		exponent,
		ePrefix)
}

// NewUint64 - Creates a new NumStrDto instance from an uint64 and a
// precision specification.
//
// Numeric separators include the Currency Symbol, the Decimal
// Delimiter and the Thousands Separator. These separator character
// values vary according to nationality and culture.
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
// See the Example Usage section below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  uint64Num           uint64
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
//
//
//  precision           uint
//     - 'precision' specifies the number of digits to be formatted
//       to the right of the decimal place.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  This method is designed to be used in conjunction with the NumStrDto{}
//  syntax thereby allowing NumStrDto type creation and initialization in
//  one step.
//
//     uint64Num := uint64(123456)
//     precision := uint(3)
//     nDto, err := NumStrDto{}.NewUint64(
//                  uint64Num,
//                  precision,
//                  "calling method name ")
//
//           nDto is now equal to 123.456
//
//  Examples:
//  ---------
//
//  <-- Input Parameters -->      <- Output ->
//                                newNumStrDto
//  uint64Num     precision          Result
//  ------------------------------------------
//
//   123456          4                  12.3456
//   123456          0              123456
//   123456          1               12345.6
//
func (nDto NumStrDto) NewUint64(
	uint64Num uint64,
	precision uint,
	ePrefix string) (
	NumStrDto,
	error) {

	ePrefix += "NumStrDto.NewUint64() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newUint64(
		numSepsDto,
		uint64Num,
		precision,
		ePrefix)
}

// NewUint64Exponent - Returns a new NumStrDto instance. The numeric
// value is set using an uint64 value multiplied by 10 raised to the
// power of the 'exponent' input parameter.
//
//   numeric value = 'uint64Num' X 10^'exponent'
//
// Numeric separators used by this method to configure the returned
// NumStrDto instance will be taken from the current NumStrDto instance.
// Unless this current NumStrDto instance was specifically configured
// for numeric separators, default USA numeric separators will be
// be applied to the returned NumStrDto instance. Default USA numeric
// separators are defined as:
//
//        decimal separator = '.'
//        thousands separator = ','
//        currency symbol = '$'
//
// To accommodate other nationalities and cultures, one can change
// the numeric separators by calling one of the following methods
// on the returned NumStrDto instance :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  uint64Num           uint64
//     - This uint64 is multiplied by 10^exponent (input parameter exponent)
//       to calculate the final numeric value which is returned in a new
//       instance of NumStrDto.
//
//
//  exponent            int
//     - This is the exponent value. Input parameter uint64Num is multiplied
//       by 10 raised to the power of this 'exponent' parameter in order to
//       calculate the numeric value contained in the new instance of
//       NumStrDto returned by this method.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//    - If this method completes successfully, it will return a new
//       instance of NumStrDto. The numeric value contained in this new
//       instance is calculated by multiplying input parameter
//       'uint64Num' times 10 raised to the power of input parameter
//       'exponent'.
//
//       The new NumStrDto instance will likely be configured with
//       default USA numeric separators. To accommodate other
//       nationalities and cultures, change the numeric separators
//       by calling one of the following methods on 'newNumStrDto' :
//
//         NumStrDto.SetNumericSeparatorsDto()
//         NumStrDto.SetCurrencySymbol()
//         NumStrDto.SetDecimalSeparator()
//         NumStrDto.SetThousandsSeparator()
//
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//     uint64Num := uint64(123456)
//     exponent := -3
//
//  nDto,err := NumStrDto{}.NewUint64Exponent(
//          uint64Num,
//          exponent,
//          "calling method name ")
//   -- nDto is now equal to "123.456", precision = 3
//
//
//     numSeps.SetToUSADefaults()
//     uint64Num := uint64(123456)
//     exponent := 3
//
//  nDto,err := NumStrDto{}.NewUint64Exponent(
///          uint64Num,
//          exponent,
//          "calling method name ")
//  -- nDto is now equal to "123456.000", precision = 3
//
//  Examples:
//  ---------
//                                         NumStrDto
//   uint64Num          exponent            Result
//   123456               -3               123.456
//   123456                3            123456.000
//   123456                0            123456
//
func (nDto NumStrDto) NewUint64Exponent(
	uint64Num uint64,
	exponent int,
	ePrefix string) (
	NumStrDto,
	error){

	ePrefix += "NumStrDto.NewUint64Exponent() "


	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		&nDto,
		ePrefix)

	if err != nil {
		return NumStrDto{}, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.newUint64Exponent(
		numSepsDto,
		uint64Num,
		exponent,
		ePrefix)
}

// NewZero - Creates a new NumStrDto with a value of zero and a
// precision specified by input parameter 'precision'.
//
// This method will return a zero NumStrDto instance set to
// default USA separators (Currency Symbol, Thousands Separator,
// Decimal Separator).
//
// Example Usage:
//
// Input Parameter      NumStrDto
//  precision            Result
//     0                  "0"
//     2                  "0.00"
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  precision       uint
//     - The number of zeros which will be placed to the right of the
//       decimal point in the returned new instance of NumStrDto.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrDto
//     - This method will return a new instance of NumStrDto populated
//       with zero values. The separator values will be set to USA
//       defaults (Currency Symbol, Thousands Separator,
//       Decimal Separator).
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     Input Parameter   newNumStrDto
//       precision          Result
//          0                 "0"
//          2                 "0.00"
//
func (nDto NumStrDto) NewZero(
	precision uint) NumStrDto {

	nStrDtoElectron := numStrDtoElectron{}

	return nStrDtoElectron.newBaseZeroNumStrDto(
		precision)
}

// ParseNumStr - receives a raw string and converts to a properly
// formatted number string. The string is returned via a NumStrDto type.
// Returned number strings may consist of a leading negative sign ('-')
// numeric digits and may include a decimal separator ('.'). The NumStrDto
// breaks the string down into sign, Integer and Fractional components.
//
// The numeric separators (decimal separator, thousands separator and
// currency symbol) taken from the current NumStrDto instance will be
// copied to the NumStrDto instance returned by this method.
//
// Input parameter 'ePrefix' is a string consisting of the method chain
// used to call this method. In case of error, this text string is
// included in the error message. Note: Be sure to leave a space at the
// end of 'ePrefix'.
//
func (nDto *NumStrDto) ParseNumStr(
	numStr string,
	ePrefix string) (
	outputNDto NumStrDto,
	err error) {

	ePrefix += "NumStrDto.ParseNumStr() "

	nStrDtoElectron := numStrDtoElectron{}

	outputNDto = nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto

	numSepsDto,
	err = nStrDtoAtom.getNumericSeparatorsDto(
		nDto,
		ePrefix + "nDto ")

	if err != nil {
		return outputNDto, err
	}

	outputNDto,
	err = nStrDtoAtom.parseNumStr(
		numStr,
		numSepsDto,
		ePrefix)

	return outputNDto, err
}

// ScaleNumStr - This method receives a signed number string
// (signedNumStr) and proceeds to create a new NumStrDto instance.
// The decimal point within the signed number string will be
// left or right depending on the value of input parameter
// 'scaleMode'.
//
// The current instance of NumStrDto will not be altered in any
// way and an no member variable data values will be changed.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  signedNumStr        string
//     - A valid number string. The leading digit may optionally
//       be a '+' or '-' indicating numeric sign value. If '+'
//       or '-' characters are not present in the first character
//       position, the number is assumed to represent a positive
//       numeric value ('+'). In addition to leading plus or minus
//       characters, the number string may contain a decimal point
//       separating integer and fractional digits. All other
//       characters in this number string must be numeric digits.
//
//
//  shiftPrecision      uint
//     - The number of positions which the decimal point will be
//       shifted. If 'shiftPrecision is Equal to zero, no action
//       will be taken, no error will be issued and the original
//       signedNumStr will be converted to a NumStrDto instance
//       and returned to the caller.
//
//
//  scaleMode           PrecisionScaleMode
//     - A constant with one of two Scale Mode values. These
//       constant values are located in source code file:
//             datetime/numstrdtoconstants.go
//
//       SCALEPRECISIONLEFT - Shifts the decimal point
//                            from its current position to the left.
//
//       SCALEPRECISIONRIGHT - Shifts the decimal point from its current
//                             position to the right.
//
//       Note: See Methods numStrDtoMolecule.shiftPrecisionRight() and
//             numStrDtoMolecule.shiftPrecisionLeft()
//             for additional information.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDto
//     - If successful, the method returns the result of the Shift Left
//       precision operation in the form of a new 'NumStrDto' instance.
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (nDto *NumStrDto) ScaleNumStr(
	signedNumStr string,
	shiftPrecision uint,
	scaleMode PrecisionScaleMode,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	ePrefix += "NumStrDto.ScaleNumStr() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto

	numSepsDto,
	err = nStrDtoAtom.getNumericSeparatorsDto(
		nDto,
		ePrefix + "nDto ")

	if err != nil {
		return newNumStrDto, err
	}

	nStrDtoNanobot := numStrDtoNanobot{}

	newNumStrDto,
	err = nStrDtoNanobot.scaleNumStr(
		numSepsDto,
		signedNumStr,
		shiftPrecision,
		scaleMode,
		ePrefix)

	return newNumStrDto, err
}

// SetCurrencySymbol - assigns the input parameter rune as the
// currency symbol to be used by the NumStrDto.
//
// In the USA, the currency symbol is the dollar sign ('$').
//
// Note: If a zero value is submitted as input, Currency Symbol
// will default to the USA dollar sign ('$').
//
// For a list of Major Currency Unicode Symbols, see constants
// located in: datetime/numstrdtoconstants.go
//
// Example: $123.45
//
func (nDto *NumStrDto) SetCurrencySymbol(currencySymbol rune) {

	if currencySymbol == 0 {
		currencySymbol = '$'
	}

	nDto.currencySymbol = currencySymbol
}

// SetDecimalSeparator - Assigns a rune or character to the internal
// data field, 'decimalSeparator'. The Decimal Separator is used to
// separate the integer and fractional elements of a number string.
//
// In the USA, the Decimal Separator is a period character ('.').
//
// Note: If a zero value is submitted as input, the Decimal Separator
// will default to the USA standard period character ('.').
//
// Example: 123.456
//
func (nDto *NumStrDto) SetDecimalSeparator(decimalSeparator rune) {

	if decimalSeparator == 0 {
		decimalSeparator = '.'
	}

	nDto.decimalSeparator = decimalSeparator
}

// SetThousandsSeparator - Sets the value of the character which will be
// used to separate thousands in the display of the NumStrDto number
// string. In the USA the typical thousands separator is the comma.
//
// If if a zero value is submitted, the Thousands Separator will default
// to the comma character.
//
// Example:
// 1,000,000
//
func (nDto *NumStrDto) SetThousandsSeparator(thousandsSeparator rune) {

	if thousandsSeparator == 0 {
		thousandsSeparator = ','
	}

	nDto.thousandsSeparator = thousandsSeparator

}

// ShiftPrecisionLeft - Shifts the relative position of a decimal point within a number
// string. The position of the decimal point is shifted 'shiftPrecision' positions to
// the left of the current decimal point position.
//
// This is equivalent to: result = signedNumStr / 10^precision or signedNumStr divided
// by 10 raised to the power of precision.
//
// See the Example Usage section below.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  signedNumStr        string
//     - A valid number string. The leading digit may optionally
//       be a '+' or '-' indicating numeric sign value. If '+'
//       or '-' characters are not present in the first character
//       position, the number is assumed to represent a positive
//       numeric value ('+'). In addition to leading plus or minus
//       characters, the number string may contain a decimal point
//       separating integer and fractional digits. All other
//       characters in this number string must be numeric digits.
//
//
//  shiftPrecision      uint
//     - The number of digits by which the current decimal point
//       point position in the number string, 'signedNumStr' will
//       be shifted to the left.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDto
//     - If successful, the method returns the result of the Shift Left
//       precision operation in the form of a 'NumStrDto' instance.
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//                               Shift-Left
//  signedNumStr    precision     Result
//  "123456.789"       3        "123.456789"
//  "123456.789"       2       "1234.56789"
//  "123456.789"       6          "0.123456789"
//  "123456789"        6        "123.456789"
//  "123"              5          "0.00123"
//   "0"               3          "0.000"
//   "0.000"           2          "0.00000"
//  "123456.789"       0     "123456.789"      - zero 'shiftPrecision' has no effect on
//                                               original number string
// "-123456.789"       0       "-123.456789"
// "-123456.789"       3       "-123.456789"
// "-123456789"        6       "-123.456789"
//
func (nDto *NumStrDto) ShiftPrecisionLeft(
	signedNumStr string,
	shiftLeftPrecision uint,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	ePrefix += "NumStrDto.ShiftPrecisionLeft() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSeparators  NumericSeparatorDto

	numSeparators,
	err = nStrDtoAtom.getNumericSeparatorsDto(
		nDto,
		ePrefix)

	if err != nil {
		return newNumStrDto, err
	}

	nStrDtoMolecule := numStrDtoMolecule{}

	newNumStrDto,
	err = nStrDtoMolecule.shiftPrecisionLeft(
		numSeparators,
		signedNumStr,
		shiftLeftPrecision,
		ePrefix)

	return newNumStrDto, err
}

// ShiftPrecisionRight - Shifts the relative precision of a decimal
// point with a number string. The position of the decimal point is
// shifted 'shiftRightPrecision' positions to the right of the
// original decimal point position.
//
// This is equivalent to: result = signedNumStr X 10^shiftRightPrecision
// or signedNumStr Multiplied by 10 raised to the power of
// 'shiftRightPrecision'.
//
// See the Example Usage section below.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  signedNumStr        string
//     - A valid number string. The leading digit may optionally
//       be a '+' or '-' indicating numeric sign value. If '+'
//       or '-' characters are not present in the first character
//       position, the number is assumed to represent a positive
//       numeric value ('+'). In addition to leading plus or minus
//       characters, the number string may contain a decimal point
//       separating integer and fractional digits. All other
//       characters in this number string must be numeric digits.
//
//
//  shiftRightPrecision uint
//     - The number of digits by which the current decimal point
//       point position in the number string, 'signedNumStr' will
//       be shifted to the right.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDto
//     - If successful, the method returns the result of the Shift Left
//       precision operation in the form of a new 'NumStrDto' instance.
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
// Notice that Zero 'shiftRightPrecision' has no effect on the original
// number string.
//
//   -------------------------------------------------------
//   signedNumStr   shiftRightPrecision       Result
//   -------------------------------------------------------
//   "123456.789"            3              "123456789"
//   "123456.789"            2              "12345678.9"
//   "123456.789"            6              "123456789000"
//   "123456789"             6              "123456789000000"
//   "123"                   5              "12300000"
//   "0"                     3              "0"
//   "123456.789"            0              "123456.789"
//  "-123456.789"            0             "-123456.789"
//  "-123456.789"            3             "-123456789"
//  "-123456789"             6             "-123456789000000"
//
func (nDto *NumStrDto) ShiftPrecisionRight(
	signedNumStr string,
	shiftRightPrecision uint) (
	newNumStrDto NumStrDto,
	err error) {

	ePrefix := "NumStrDto.ShiftPrecisionRight() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSeparators  NumericSeparatorDto

	numSeparators,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		nDto,
		ePrefix)

	if err != nil {
		return newNumStrDto, err
	}

	nStrDtoMolecule := numStrDtoMolecule{}

	newNumStrDto,
		err = nStrDtoMolecule.shiftPrecisionRight(
		numSeparators,
		signedNumStr,
		shiftRightPrecision,
		ePrefix)

	return newNumStrDto, err
}

// SetNumericSeparators - Used to assign values for the Decimal and Thousands separators as well
// as the Currency Symbol to be used in displaying the current number string.
//
// Note: If zero values are submitted as input, the values will default to USA standards.
//
// USA Examples:
//
// Decimal Separator period ('.') 		= 123.456
// Thousands Separator comma (',') 		= 1,000,000,000
// Currency Symbol dollar sign ('$')	= $123
//
func (nDto *NumStrDto) SetNumericSeparators(
	decimalSeparator,
	thousandsSeparator,
	currencySymbol rune) {

	if decimalSeparator == 0 {
		decimalSeparator = '.'
	}

	if thousandsSeparator == 0 {
		thousandsSeparator = ','
	}

	if currencySymbol == 0 {
		currencySymbol = '$'
	}

	nDto.decimalSeparator = decimalSeparator
	nDto.thousandsSeparator = thousandsSeparator
	nDto.currencySymbol = currencySymbol
}

// SetNumericSeparatorsDto - Sets the values of numeric separators:
//
//                           decimal point separator
//                           thousands separator
//                           currency symbol
//
// These separators are set based on values transmitted through input
// parameter 'customSeparators'. Separator values from 'customSeparators'
// are copied to the current NumStrDto instance.
//
// If any of the values contained in input parameter 'customSeparators'
// is set to zero or nil, an error will be returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  customSeparators    NumericSeparatorDto
//     - If any of the data fields in this passed structure 'customSeparators'
//       are set to zero ('0'), an this method will return an error.
//
//       The separator values contained in this input parameter will be
//       copied to current NumStrDto instance. The data fields included
//       in the NumericSeparatorDto are listed as follows:
//
//          type NumericSeparatorDto struct {
//            DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
//            ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
//            CurrencySymbol     rune // Currency Symbol
//          }
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) SetNumericSeparatorsDto(
	customSeparators NumericSeparatorDto,
	ePrefix string) error {

	ePrefix += "NumStrDto.SetNumericSeparatorsDto() "

	nStrDtoElectron := numStrDtoElectron{}

	return nStrDtoElectron.setNumericSeparatorsDto(
		nDto,
		customSeparators,
		ePrefix)
}

// SetNumericSeparatorsToDefaultIfEmpty - If numeric separators are
// set to zero or nil, this method will set those numeric
// separators to the USA defaults. This means that the
// Decimal separator is set to ('.'), the Thousands separator
// is set to (',') and the currency symbol is set to '$'.
//
// If the numeric separators were previously set to a value
// other than zero or nil, that value is not altered by this
// method.
//
// Effectively, this method ensures that numeric separators
// are set to valid values.
//
func (nDto *NumStrDto) SetNumericSeparatorsToDefaultIfEmpty() {

	nStrDtoElectron := numStrDtoElectron{}

	_ = nStrDtoElectron.setNumericSeparatorsToDefaultIfEmpty(
			nDto,
			"")

	return
}

// SetNumericSeparatorsToUSADefault - Sets Numeric separators:
//   Decimal Point Separator
//   Thousands Separator
//   Currency Symbol
//
// to United States of America (USA) defaults.
//
// Call specific methods to set numeric separators for other countries or
// cultures:
// NumStrDto.SetDecimalSeparator()
// NumStrDto.SetThousandsSeparator()
// NumStrDto.SetCurrencySymbol()
//
func (nDto *NumStrDto) SetNumericSeparatorsToUSADefault() {

	nStrDtoElectron := numStrDtoElectron{}

	_ = nStrDtoElectron.setNumericSeparatorsToUSADefault(
		nDto,
		"")

	return
}

// SetNumStr - Sets the value of the current NumStrDto instance
// to the number string received as input.
func (nDto *NumStrDto) SetNumStr(numStr string) error {

	ePrefix := "NumStrDto.SetNumStr() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto
	var err error

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		nDto,
		ePrefix)

	if err != nil {
		return err
	}

	nStrDtoUtil := numStrDtoUtility{}

	return nStrDtoUtil.setNumStr(
		nDto,
		numSepsDto,
		numStr,
		ePrefix)
}

// SetPrecision - parses the incoming number string and applies the designated 'precision'.
// 'precision' determines the number of digits to the right of the decimal place. The boolean
// parameter 'roundResult' is used to apply rounding in those cases where 'precision' dictates
// a reduction in the number of digits to the right of the decimal place. See 'Examples' below.
//
//
// --------------------------------------------------------------------------------------------------
//
// Input Parameters
//
//  signedNumStr        string
//     - A valid number string
//
//  precision           uint
//     - The 'precision' values designates the number of places to the right of the
//       decimal point which will be realized upon completion of this operation.
//
//  roundResult         bool
//     - If the 'precision' value is less than the current number of places to the
//       right of the decimal point, this method will truncate the existing fractional
//       digits. If 'roundResult' is set to true, this truncation operation will
//       include rounding the last digit.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//       ------------ Input Parameters -----------      ---- Result --------
// Test
//  No   signedNumStr      precision   roundResult
//  ---- ------------------------------------------------------------------
//  1    "123456789"          7           false         "123456789.0000000"
//  2    "123456789"          7           true          "123456789.0000000"
//  3    "-123456789"         7           false        "-123456789.0000000"
//  4    "-123456789"         7           true         "-123456789.0000000"
//  5    "123456.789"         2           true             "123456.79"
//  6    "123456.789"         2           false            "123456.78"
//  7    "123456.789"         5           false            "123456.78900"
//  8    "123.456789"         1           false               "123.4"
//  9    "123.456789"         1           true                "123.5"
// 10    "-123.456789"        1           false              "-123.4"
// 11    "-123.456789"        1           true               "-123.5"
// 12    "123456.789"         0           true             "123457"
// 13    "-123456.789"        0           true            "-123457"
// 14    "123456.789"         0           false            "123456"
// 15    "-123456.789"        0           false           "-123456"
// 16    "123457"             1           false            "123457.0"
// 17    "123457"             1           true             "123457.0"
// 18    "-123457"            1           false           "-123457.0"
// 19    "-123457"            1           true            "-123457.0"
//
func (nDto *NumStrDto) SetPrecision(
	signedNumStr string,
	precision uint,
	roundResult bool,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	ePrefix += "NumStrDto.SetPrecision() "

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto

	numSepsDto,
		err = nStrDtoAtom.getNumericSeparatorsDto(
		nDto,
		ePrefix + "nDto ")

	if err != nil {
		return newNumStrDto, err
	}

	nStrDtoMolecule := numStrDtoMolecule{}

	newNumStrDto,
	err = nStrDtoMolecule.setPrecision(
		numSepsDto,
		signedNumStr,
		precision,
		roundResult,
		ePrefix)

	return newNumStrDto, err
}

// SetSignValue - Sets the sign of the numeric value
// for the current NumStrDto.
//
// Input parameter 'newSignVal' must be set to one of two possible values:
//
//  +1 = positive numeric values greater than or equal to zero
//  -1 = negative numeric values less than zero
//
// If 'newSignVal' is not set to one of the two valid numeric sign values
// listed above, this method will return an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Member variables will be tested for
//       validity.
//
//
//  newSignVal          int
//     - This input parameter must be set to one of two possible values:
//
//       +1 for positive numeric values greater than or equal to zero
//
//       -1 for negative numeric values less than zero
//
//       If 'newSignVal' is not set to one of the two valid numeric sign
//       values listed above, this method will return an error.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) SetSignValue(
	newSignVal int,
	ePrefix string) error {

	ePrefix += "NumStrDto.SetSignValue() "

	nStrDtoElectron := numStrDtoElectron{}

	return nStrDtoElectron.setSignValue(
		nDto,
		newSignVal,
		ePrefix)
}

// SetThisPrecision - Sets precision for the current NumStrDto instance.
// 'precision' identifies the number of decimal places to the right of the
// decimal point.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  precision           uint
//     - The number of numeric digits to the right of the decimal place
//       which will be configured in the numeric value encapsulated within
//       the current NumStrDto instance.
//
//
//  roundResult         bool
//     - If the 'precision' value is less than the current number of places
//       to the	right of the decimal point, this method will truncate the
//       existing fractional digits. If 'roundResult' is set to true, this
//       truncation operation will include rounding the last digit.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                error
//     - If this method completes successfully, the returned error Type is
//       set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'.
//
func (nDto *NumStrDto) SetThisPrecision(
	precision uint,
	roundResult bool,
	ePrefix string) error {

	ePrefix += "NumStrDto.SetThisPrecision() "

	nStrDtoNanobot := numStrDtoNanobot{}

	return nStrDtoNanobot.setNumStrDtoPrecision(
		nDto,
		precision,
		roundResult,
		ePrefix)
}

// Subtract - Subtracts the value of an input NumStrDto from the
// current NumStrDto instance.
func (nDto *NumStrDto) Subtract(
	n2Dto NumStrDto,
	ePrefix string) (
	err error)  {

	ePrefix += "NumStrDto.Subtract() "

	err = nil

	n1Dto := nDto.CopyOut()

	var difference NumStrDto

	difference,
	err = nDto.SubtractNumStrs(
		n1Dto,
		n2Dto,
		ePrefix)

	if err != nil {
		return err
	}

	nDto.CopyIn(difference)

	return err
}

// SubtractNumStrs - Subtracts the numeric values represented by two NumStrDto
// objects.
//
// Input parameter 'ePrefix' is a string consisting of the method chain used to
// call this method. In case of error, this text string is included in the error
// message. Note: Be sure to leave a space at the end of 'ePrefix'.
//
func (nDto *NumStrDto) SubtractNumStrs(
	minuend NumStrDto,
	subtrahend NumStrDto,
	ePrefix string) (
	difference NumStrDto,
	err error) {

	ePrefix += "NumStrDto.SubtractNumStrs() "

	difference = NumStrDto{}
	err = nil

	var n1NumDto, n2NumDto NumStrDto
	var compare int
	var isReversed bool

	n1NumDto,
	n2NumDto,
	compare,
	isReversed,
	err = nDto.FormatForMathOps(
		minuend,
		subtrahend, ePrefix)

	if err != nil {
		return difference, err
	}

	if compare == 0 {
		difference = nDto.GetZeroNumStrDto(n1NumDto.precision)
		return difference, err
	}

	newSignVal := n1NumDto.signVal
	precision := n1NumDto.precision

	if n1NumDto.signVal != n2NumDto.signVal {

		err = n1NumDto.SetSignValue(1,
			ePrefix + "n1NumDto ")

		if err != nil {
			return difference, err
		}

		err = n2NumDto.SetSignValue(
			1,
			ePrefix + "n2NumDto")

		if err != nil {
			return difference, err
		}

		difference,
		err = nDto.AddNumStrs(
			n1NumDto,
			n2NumDto,
			ePrefix)

		if err != nil {
			difference = NumStrDto{}
			return difference, err
		}

		err = difference.SetSignValue(
			newSignVal,
			ePrefix + "difference ")

		if err != nil {
			difference = NumStrDto{}
		}

		return difference, err
	}

	// Change sign for subtraction
	newSignVal = n1NumDto.signVal

	if isReversed {
		newSignVal = newSignVal * -1
	}

	lenN1AllRunes := len(n1NumDto.absAllNumRunes)

	n1IntAry := make([]int, lenN1AllRunes)
	n2IntAry := make([]int, lenN1AllRunes)
	n3IntAry := make([]int, lenN1AllRunes)

	for i := 0; i < lenN1AllRunes; i++ {

		n1IntAry[i] = int(n1NumDto.absAllNumRunes[i]) - 48
		n2IntAry[i] = int(n2NumDto.absAllNumRunes[i]) - 48

	}

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0
	// Main Subtraction Routine
	for j := lenN1AllRunes - 1; j >= 0; j-- {

		n1 = n1IntAry[j]
		n2 = n2IntAry[j]
		n3 = 0

		if n1-carry-n2 < 0 {
			n1 += 10
			n3 = n1 - n2 - carry
			carry = 1
		} else {
			n3 = n1 - n2 - carry
			carry = 0
		}

		n3IntAry[j] = n3

	}

	difference,
	err =
		nDto.FindIntArraySignificantDigitLimits(
			n3IntAry,
			precision,
			newSignVal,
			ePrefix)

	return difference, err
}
