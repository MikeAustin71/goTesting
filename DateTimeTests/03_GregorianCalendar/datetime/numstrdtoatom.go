package datetime

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
)

type numStrDtoAtom struct {

	lock *sync.Mutex
}

// compareNumStrDtoAbsoluteValues - Compares the absolute numeric values
// of two NumStrDto input parameters, 'n1Dto' and 'n2Dto'. The
// number signs (+ or -) of the two compared numeric values are
// ignored. Only the absolute numeric values are compared.
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
// ------------------------------------------------------------------------
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
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be prefixed to the beginning of the returned
//       error message.
//
func (nStrDtoAtom *numStrDtoAtom) compareNumStrDtoAbsoluteValues(
	n1Dto *NumStrDto,
	n2Dto *NumStrDto,
	ePrefix string) (
	compareResult int,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.compareNumStrDtoAbsoluteValues() "
	compareResult = -99
	err = nil

	if n1Dto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'n1Dto' is INVALID!\n" +
			"n1Dto = nil pointer!\n")
		return compareResult, err
	}

	if n2Dto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"n1Dto = nil pointer!\n")
		return compareResult, err
	}

	nStrDtoElectron := numStrDtoElectron{}
	_,
		err = nStrDtoElectron.testNumStrDtoValidity(
		n1Dto,
		ePrefix + "Testing Validity of 'n1Dto' ")


	if err != nil {
		return compareResult, err
	}

	_,
		err = nStrDtoElectron.testNumStrDtoValidity(
		n2Dto,
		ePrefix + "Testing Validity of 'n2Dto' ")


	if err != nil {
		return compareResult, err
	}

	var n1DtoAbsFracRunes, n2DtoAbsFracRunes,
	n1DtoAbsIntRunes, n2DtoAbsIntRunes []rune


	n1DtoAbsFracRunes,
	err =
		nStrDtoElectron.getAbsFracRunes(n1Dto, ePrefix)

	if err != nil {
		return compareResult, err
	}

	n2DtoAbsFracRunes,
	err =
		nStrDtoElectron.getAbsFracRunes(n2Dto, ePrefix)

	if err != nil {
		return compareResult, err
	}

	n1DtoAbsIntRunes,
	err =
		nStrDtoElectron.getAbsIntRunes(n1Dto, ePrefix)

	if err != nil {
		return compareResult, err
	}

	n2DtoAbsIntRunes,
	err =
		nStrDtoElectron.getAbsIntRunes(n2Dto, ePrefix)

	if err != nil {
		return compareResult, err
	}

	lenN1IntRunes := len(n1DtoAbsIntRunes)
	lenN2IntRunes := len(n2DtoAbsIntRunes)

	var isN1Zero, isN2Zero bool

	isN1Zero,
	err =
		nStrDtoElectron.isNumStrZeroValue(
		n1Dto,
		ePrefix)

	if err != nil {
		return compareResult, err
	}

	isN2Zero,
	err =
		nStrDtoElectron.isNumStrZeroValue(
		n2Dto,
		ePrefix)

	if err != nil {
		return compareResult, err
	}

	if !isN1Zero && isN2Zero {
		compareResult = 1
		return compareResult, err
	}

	if isN1Zero && !isN2Zero {
		compareResult = -1
		return compareResult, err
	}

	if isN1Zero && isN2Zero {
		compareResult = 0
		return compareResult, err
	}

	if lenN1IntRunes > lenN2IntRunes {
		compareResult = 1
		return compareResult, err
	}

	if lenN1IntRunes < lenN2IntRunes {
		compareResult = -1
		return compareResult, err
	}

	// lenN1IntRunes Must Be Equal to lenN2IntRunes

	for i := 0; i < lenN1IntRunes; i++ {
		n1 := n1DtoAbsIntRunes[i] - 48
		n2 := n2DtoAbsIntRunes[i] - 48

		if n1 > n2 {
			compareResult = 1
			return compareResult, err
		}

		if n1 < n2 {
			compareResult = -1
			return compareResult, err
		}
	}

	// All the integers are equal
	lenN1FracRunes := len(n1DtoAbsFracRunes)
	lenN2FracRunes := len(n2DtoAbsFracRunes)

	lenFracRunesToTest := lenN1FracRunes

	if lenN2FracRunes < lenN1FracRunes {
		lenFracRunesToTest = lenN2FracRunes
	}

	for j := 0; j < lenFracRunesToTest; j++ {
		n1 := n1DtoAbsFracRunes[j] - 48
		n2 := n2DtoAbsFracRunes[j] - 48
		if n1 > n2 {
			compareResult = 1
			return compareResult, err
		}

		if n1 < n2 {
			compareResult = -1
			return compareResult, err
		}

	}

	if lenN1FracRunes > lenN2FracRunes {
		compareResult = 1
		return compareResult, err
	}

	if lenN1FracRunes < lenN2FracRunes {
		compareResult = -1
		return compareResult, err
	}

	compareResult = 0

	return compareResult, err
}

// GetAbsAllNumRunes - Returns an array of runes representing all of the
// integer and fractional digits included in the input parameter 'numStrDto'.
// The rune array returned will consist of numeric digits with no sign
// value prefixed. This effectively returns the absolute value of all
// integer and fractional digits combined in one rune array (there is no
// decimal point).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       If this NumStrDto instance is judged to be invalid, an error
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
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  absAllRunes         []rune
//     - If this method is successful, an array of runes (a.k.a characters)
//       will be returned. This 'rune' array consists of all the numeric
//       digits contained in and extracted from input parameter 'numStrDto'.
//       The returned series of numeric character digits therefore represents
//       an absolute or positive value. It never presents with a negative
//       value.
//
//
//  error
//     - If this method completes successfully the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
//
func (nStrDtoAtom *numStrDtoAtom) getAbsoluteValueAllNumRunes(
	numStrDto *NumStrDto,
	ePrefix string) (
	absAllRunes [] rune,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.getAbsoluteValueAllNumRunes() "

	absAllRunes = make([]rune, 0, 50)
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'numStrDto' is a 'nil' pointer!\n")

		return absAllRunes, err
	}

	nStrDtoElectron := numStrDtoElectron{}

	_,
		err = nStrDtoElectron.testNumStrDtoValidity(
		numStrDto,
		ePrefix)

	if err != nil {
		return absAllRunes, err
	}

	lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

	if lenAbsAllNumRunes == 0 {
		return absAllRunes, err
	}

	absAllRunes = make([]rune, lenAbsAllNumRunes, lenAbsAllNumRunes+50)

	for i := 0; i < lenAbsAllNumRunes; i++ {
		absAllRunes[i] = numStrDto.absAllNumRunes[i]
	}

	return absAllRunes, err
}

// getAbsoluteBigInt - Returns the absolute value of all numeric/ digits
// in the number string (nDto.absAllNumRunes). As such, Fractional digits
// to the right of the decimal are included in the consolidate integer
// number. All of the numeric digits in the number string are therefore
// returned as a type *big.Int
//
// This method will fail and return an error if input parameter 'numStrDto'
// is judged to be invalid.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       If this NumStrDto instance is judged to be invalid, an error
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
func (nStrDtoAtom *numStrDtoAtom) getAbsoluteBigInt(
	numStrDto *NumStrDto,
	ePrefix string) (
	bigIntNum *big.Int,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.getAbsoluteBigInt() "

	bigIntNum = big.NewInt(0)
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'numStrDto' is a 'nil' pointer!\n")

		return bigIntNum, err
	}

	nStrDtoElectron := numStrDtoElectron{}

	_,
		err = nStrDtoElectron.testNumStrDtoValidity(
		numStrDto,
		ePrefix)

	if err != nil {
		return bigIntNum, err
	}

	lenAllNumRunes := len(numStrDto.absAllNumRunes)

	base10 := big.NewInt(int64(10))

	for i := 0; i < lenAllNumRunes; i++ {

		bigIntNum =
			big.NewInt(0).Mul(
				bigIntNum,
				base10)

		bigIntNum =
			big.NewInt(0).Add(bigIntNum,
				big.NewInt(int64(numStrDto.absAllNumRunes[i]-48)))

	}

	return bigIntNum, err
}

// getBigRationalNum - Receives an instance of of NumStrDto and
// converts the associated signed numerical value to a type
// *big.Rat, big rational number.
//
// For more information on the *big.Rat, rational number type,
// reference:
//   https://golang.org/pkg/math/big/
//
// This method will return an error if input parameter 'numStrDto'
// is judged to be an invalid instance of NumStrDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT change the values of internal member variables to achieve
//       the method's objectives. If this NumStrDto instance is judged
//       to be invalid, an error will be returned.
//
//       The signed numerical value encapsulated by this NumStrDto
//       instance will be converted and returned as a type *big.Rat,
//       Big Rational number.
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
//  bigRatNum           *big.Rat
//     - If this method completes successfully, the signed numerical
//       value of the input parameter 'numStrDto' will be converted
//       and returned in this parameter as a signed type *big.Rat,
//       or Big Rational Number.
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
func (nStrDtoAtom *numStrDtoAtom) getBigRationalNum(
	numStrDto *NumStrDto,
	ePrefix string) (
	bigRatNum *big.Rat,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.getBigRationalNum() "

	err = nil
	bigRatNum = big.NewRat(1, 1)

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return bigRatNum, err
	}

	nStrDtoElectron := numStrDtoElectron{}

	err = nStrDtoElectron.setNumericSeparatorsToDefaultIfEmpty(
		numStrDto,
		ePrefix + "numStrDto ")

	if err != nil {
		return bigRatNum, err
	}

	absInt, isOk :=
		big.NewInt(0).SetString(
			string(numStrDto.absAllNumRunes),
			10)

	if !isOk {
		err = fmt.Errorf(ePrefix + "\n" +
			"Conversion of nDto.absAllNumRunes to big.Int Failed! "+
			"nDto.absIntRunes= '%v'", numStrDto.absAllNumRunes)
		return bigRatNum, err
	}

	base10 := big.NewInt(10)

	bigPrecision := big.NewInt(int64(numStrDto.precision))

	scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)

	absRatNum := big.NewRat(1, 1).SetFrac(absInt, scaleFactor)

	// int value = +1 or -1
	numSign := numStrDto.signVal

	if numSign == -1 {
		bigRatNum = big.NewRat(1,1).
			Set(
			big.NewRat(1,1).Neg(absRatNum))
	} else {
		bigRatNum = big.NewRat(1,1).
			Set(absRatNum)
	}

	return bigRatNum, err
}


// getNumericSeparatorsDto - Returns a structure containing the
// character or rune values for decimal point separator, thousands
// separator and currency symbol. These numeric separators are
// extracted from input parameter
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       Numeric separators extracted from this NumStrDto instance are
//       used to populate the returned instance of NumericSeparatorDto.
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
//  numSepsDto          NumericSeparatorDto
//     - If this method completes successfully, an instance of
//       NumericSeparatorDto will be returned populated with the
//       numeric separators associated with input parameter,
//       'numStrDto'. The NumericSeparatorDto data structure is
//       shown below:
//
//          type NumericSeparatorDto struct {
//            DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
//            ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
//            CurrencySymbol     rune // Currency Symbol
//          }
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoAtom *numStrDtoAtom) getNumericSeparatorsDto(
	numStrDto *NumStrDto,
	ePrefix string) (
	numSepsDto NumericSeparatorDto,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.getNumericSeparatorsDto() "

	numSepsDto = NumericSeparatorDto{}

	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return numSepsDto, err
	}

	nStrDtoElectron := numStrDtoElectron{}

	numSepsDto.DecimalSeparator,
	err = nStrDtoElectron.getDecimalSeparator(
		numStrDto,
		ePrefix)

	if err != nil {
		return numSepsDto, err
	}

	numSepsDto.ThousandsSeparator,
	err = nStrDtoElectron.getThousandsSeparator(
		numStrDto,
		ePrefix)

	if err != nil {
		return numSepsDto, err
	}

	numSepsDto.CurrencySymbol,
	err = nStrDtoElectron.getCurrencySymbol(
		numStrDto,
		ePrefix)

	if err != nil {
		return numSepsDto, err
	}

	return numSepsDto, err
}

// getScaleFactor - Returns the scale factor for the NumStrDto
// instance passed as input parameter 'numStrDto'.
//
// Scale factor is defined by 10 raised to the power of the
// precision value associated with input parameter, 'numStrDto'
// (numStrDto.precision).
//
// 'precision', as defined in connection with type NumStrDto, is
// the number of digits to the right of the decimal point.
//
// For example, if the 'precision' value of 'numStrDto' is '5', the
// resulting scale factor would be computed as follows:
//
//     precision = 5
//     scale factor = 10^5 = 100,000
//
// The computed 'Scale Factor' is returned as a type *big.Int.
// For more information on type big.Int, reference:
//           https://golang.org/pkg/math/big/
//
// Be advised that this method will return an error if input
// parameter 'numStrDto' is judged to be an invalid instance of
// NumStrDto.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT change the values of internal member variables to achieve
//       the method's objectives. If this NumStrDto instance is judged
//       to be invalid, an error will be returned.
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
//  scaleFactor         *big.Int
//     - Scale factor is defined by 10 raised to the power of nDto.precision.
//       NumStrDto.precision is the number of digits to the right of the
//       decimal point. 'scaleFactor' is returned as a type *big.Int integer
//       value.
//
//       Example:
//        'numStrDto' numerical value = 5.123
//        Precision = 3
//        Scale Factor = 10^3 = 1,000
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nStrDtoAtom *numStrDtoAtom) getScaleFactor(
	numStrDto *NumStrDto,
	ePrefix string) (
	scaleFactor *big.Int,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.getScaleFactor() "

	err = nil
	scaleFactor = big.NewInt(0)

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto has a 'nil' pointer!\n")
		return scaleFactor, err
	}

	nStrDtoElectron := numStrDtoElectron{}

	err = nStrDtoElectron.setNumericSeparatorsToDefaultIfEmpty(
		numStrDto,
		ePrefix + "numStrDto ")

	if err != nil {
		return scaleFactor, err
	}

	if len(numStrDto.absAllNumRunes) == 0 {
		err = errors.New(ePrefix + "\n" +
			"Input parameter 'numStrDto' has a zero length number string.\n" +
			"len(numStrDto.absAllNumRunes) == 0\n")

		return scaleFactor, err

	}

	if numStrDto.precision == 0 {
		scaleFactor = big.NewInt(1)
		return scaleFactor, err
	}

	base10 := big.NewInt(0).SetInt64(int64(10))

	bigPrecision := big.NewInt(0).
		SetInt64(int64(numStrDto.precision))

	scaleFactor =
		big.NewInt(0).
			Exp(base10, bigPrecision, nil)

	return scaleFactor, err
}

// formatCurrencyStr - Formats the NumStrDto numeric value contained
// in input parameter 'numStrDto' as a currency string.
//
// If the Currency Symbol was not previously set for 'numStrDto',
// the currency symbol is defaulted to the USA standard dollar sign,
// ('$').
//
// If the Decimal Separator was not previously set for 'numStrDto',
// the Decimal Separator is defaulted to the USA standard period ('.').
//
// If the Thousands Separator was not previously set for 'numStrDto',
// the Thousands Separator is defaulted to the USA standard comma (',').
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT change the values of internal member variables to achieve
//       the method's objectives. If this NumStrDto instance is judged
//       to be invalid, an error will be returned.
//
//
//  negValMode          NegativeValueFmtMode
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
//  currencyStr         string
//     - If this method completes successfully, this string will be populated
//       with the numeric value extracted from input parameter 'numStrDto'
//       formatted as a currency string incorporating the appropriate currency
//       symbol.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoAtom *numStrDtoAtom) formatCurrencyStr(
	numStrDto *NumStrDto,
	negValMode NegativeValueFmtMode,
	ePrefix string) (
	currencyStr string,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.formatCurrencyStr() "

	if numStrDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'numStrDto' is a 'nil' pointer!\n")

		return currencyStr, err
	}

	nStrDtoElectron := numStrDtoElectron{}

	err = nStrDtoElectron.setNumericSeparatorsToDefaultIfEmpty(
		numStrDto,
		ePrefix)

	if err != nil {
		return currencyStr, err
	}

	_,
		err = nStrDtoElectron.testNumStrDtoValidity(
		numStrDto,
		ePrefix)

	if err != nil {
		return currencyStr, err
	}

	lenAllNumRunes := len(numStrDto.absAllNumRunes)

	lenOut := lenAllNumRunes

	lenIntRunes := lenAllNumRunes - int(numStrDto.precision)

	seps := lenIntRunes / 3

	mod := lenIntRunes - (seps * 3)

	if mod == 0 {
		seps--
	}

	// adjust for thousands delimiters
	lenOut += seps

	// adjust for negative sign value
	if numStrDto.signVal == -1 {
		if negValMode == LEADMINUSNEGVALFMTMODE {
			lenOut++
		} else {
			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			lenOut += 2
		}
	}

	// adjust for decimal point
	if numStrDto.precision > 0 {
		lenOut++
	}

	// adjust for currency symbol
	lenOut++

	outRunes := make([]rune, lenOut)
	outIdx := lenOut - 1
	allNumsIdx := lenAllNumRunes - 1

	// If negative value and parenthesis formatting
	// specified, format trailing parenthesis.
	if numStrDto.signVal == -1 &&
		negValMode == PARENTHESESNEGVALFMTMODE {
		outRunes[outIdx] = ')'
		outIdx--
	}

	if numStrDto.precision > 0 {

		for i := 0; i < int(numStrDto.precision); i++ {
			outRunes[outIdx] = numStrDto.absAllNumRunes[allNumsIdx]
			outIdx--
			allNumsIdx--
		}

		outRunes[outIdx] = numStrDto.decimalSeparator
		outIdx--
	}

	sepCnt := 0

	for i := 0; i < lenIntRunes; i++ {

		sepCnt++

		if sepCnt == 4 && seps > 0 {
			sepCnt = 1
			seps--
			outRunes[outIdx] = numStrDto.thousandsSeparator
			outIdx--
		}

		outRunes[outIdx] = numStrDto.absAllNumRunes[allNumsIdx]
		outIdx--
		allNumsIdx--

	}

	outRunes[outIdx] = numStrDto.currencySymbol

	// If required, add leading negative
	// value sign
	if numStrDto.signVal == -1 &&
		negValMode == PARENTHESESNEGVALFMTMODE {
		outRunes[0] = '('
	} else if numStrDto.signVal == -1 {
		outRunes[0] = '-'
	}

	currencyStr = string(outRunes)

	return currencyStr, err
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
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. This NumStrDto will supply the
//       numeric value which will be used to create the returned
//       number string.
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
func (nStrDtoAtom *numStrDtoAtom) formatNumStr(
	numStrDto *NumStrDto,
	negValMode NegativeValueFmtMode,
	ePrefix string) (
	numStr string,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.formatNumStr() "

	if numStrDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'numStrDto' is a 'nil' pointer!\n")

		return numStr, err
	}

	nStrDtoElectron := numStrDtoElectron{}

	// Do this to ensure default separator
	// initialized if equal to zero.
	_,
	_ = nStrDtoElectron.getDecimalSeparator(
		numStrDto,
		ePrefix)

	_,
	err = nStrDtoElectron.testNumStrDtoValidity(
		numStrDto,
		ePrefix)

	if err != nil {
		return numStr, err
	}

	lenAllNumRunes := len(numStrDto.absAllNumRunes)

	lenOut := lenAllNumRunes

	lenIntRunes := lenAllNumRunes - int(numStrDto.precision)

	// adjust for negative sign value
	if numStrDto.signVal == -1 {
		if negValMode == LEADMINUSNEGVALFMTMODE {
			lenOut++
		} else {
			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			lenOut += 2
		}
	}

	// adjust for decimal point
	if numStrDto.precision > 0 {
		lenOut++
	}

	outRunes := make([]rune, lenOut)
	outIdx := lenOut - 1

	if numStrDto.signVal == -1 &&
		negValMode == PARENTHESESNEGVALFMTMODE {
		outRunes[outIdx] = ')'
		outIdx--
	}

	allNumsIdx := lenAllNumRunes - 1

	if numStrDto.precision > 0 {

		for i := 0; i < int(numStrDto.precision); i++ {
			outRunes[outIdx] = numStrDto.absAllNumRunes[allNumsIdx]
			outIdx--
			allNumsIdx--
		}

		outRunes[outIdx] = numStrDto.decimalSeparator
		outIdx--
	}

	for i := 0; i < lenIntRunes; i++ {

		outRunes[outIdx] = numStrDto.absAllNumRunes[allNumsIdx]
		outIdx--
		allNumsIdx--
	}

	if numStrDto.signVal == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes[0] = '-'
		} else {
			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			outRunes[0] = '('
		}

	}

	numStr = string(outRunes)

	return numStr, err
}

// formatThousandsStr - Returns the number string delimited with the
// nDto.thousandsSeparator character plus the Decimal Separator if
// applicable. The number string is extracted from input parameter
// 'numStrDto'.
//
// If the Currency Symbol was not previously set for 'numStrDto',
// the currency symbol is defaulted to the USA standard dollar sign,
// ('$').
//
// If the Decimal Separator was not previously set for 'numStrDto',
// the Decimal Separator is defaulted to the USA standard period ('.').
//
// If the Thousands Separator was not previously set for 'numStrDto',
// the Thousands Separator is defaulted to the USA standard comma (',').
//
//
// Example:
// thousandsStr = 1000000.234 converted to 1,000,000.234
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. If this NumStrDto instance is judged
//       to be invalid, an error will be returned.
//
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
//  thousandsStr        string
//     - If this method completes successfully, this string will contain
//       the designated numeric value properly formatted with the correct
//       thousands separator.
//        Example:
//        thousandsStr = 1000000.234 converted to 1,000,000.234
//
//
//  err                 error
//     - If this method completes successfully the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'.
//
func (nStrDtoAtom *numStrDtoAtom) formatThousandsStr(
	numStrDto *NumStrDto,
	negValMode NegativeValueFmtMode,
	ePrefix string) (
	thousandsStr string,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.formatThousandsStr() "

	if numStrDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'numStrDto' is a 'nil' pointer!\n")

		return thousandsStr, err
	}

	nStrDtoElectron := numStrDtoElectron{}

	// Do this to ensure default separator
	// initialized if equal to zero.
	_,
		_ = nStrDtoElectron.getDecimalSeparator(
		numStrDto,
		ePrefix)

	_,
		err = nStrDtoElectron.testNumStrDtoValidity(
		numStrDto,
		ePrefix)

	if err != nil {
		return thousandsStr, err
	}

	lenAllNumRunes := len(numStrDto.absAllNumRunes)

	lenOut := lenAllNumRunes

	lenIntRunes := lenAllNumRunes - int(numStrDto.precision)

	seps := lenIntRunes / 3

	mod := lenIntRunes - (seps * 3)

	if mod == 0 {
		seps--
	}

	// adjust for thousands delimiters
	lenOut += seps

	// adjust for negative sign value
	if numStrDto.signVal == -1 {
		if negValMode == LEADMINUSNEGVALFMTMODE {
			lenOut++
		} else {
			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			lenOut += 2
		}

	}

	// adjust for decimal point
	if numStrDto.precision > 0 {
		lenOut++
	}

	outRunes := make([]rune, lenOut)
	outIdx := lenOut - 1

	if numStrDto.signVal == -1 &&
		negValMode == PARENTHESESNEGVALFMTMODE {
		outRunes[outIdx] = ')'
		outIdx--
	}

	allNumsIdx := lenAllNumRunes - 1

	if numStrDto.precision > 0 {

		for i := 0; i < int(numStrDto.precision); i++ {
			outRunes[outIdx] = numStrDto.absAllNumRunes[allNumsIdx]
			outIdx--
			allNumsIdx--
		}

		outRunes[outIdx] = numStrDto.decimalSeparator
		outIdx--
	}

	sepCnt := 0

	for i := 0; i < lenIntRunes; i++ {

		sepCnt++

		if sepCnt == 4 && seps > 0 {
			sepCnt = 1
			seps--
			outRunes[outIdx] = numStrDto.thousandsSeparator
			outIdx--
		}

		outRunes[outIdx] = numStrDto.absAllNumRunes[allNumsIdx]
		outIdx--
		allNumsIdx--

	}

	if numStrDto.signVal == -1 {
		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes[0] = '-'
		} else {
			outRunes[0] = '('
		}

	}

	thousandsStr = string(outRunes)

	return thousandsStr, err
}

// parseNumStr - receives a raw string and converts to a properly
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
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. This NumStrDto will supply the
//       numeric value which will be used to create the returned
//       number string.
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
func (nStrDtoAtom *numStrDtoAtom) parseNumStr(
	numStr string,
	numericSeparators NumericSeparatorDto,
	ePrefix string)(
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoAtom.lock == nil {
		nStrDtoAtom.lock = new(sync.Mutex)
	}

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	ePrefix += "numStrDtoAtom.parseNumStr() "

	err = nil

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto = nStrDtoElectron.newBaseZeroNumStrDto(0)

	if len(numStr) == 0 {
		err = errors.New(ePrefix + "\n" +
			"Input parameter 'numStr' is invalid!\n" +
			"'numStr' is an empty string.\n")

		return newNumStrDto, err
	}

	numericSeparators.SetToUSADefaultsIfEmpty()

	n2Dto := nStrDtoElectron.newBaseZeroNumStrDto(0)

	n2Dto.signVal = 1


	err =
		nStrDtoElectron.setNumericSeparatorsDto(
			&n2Dto,
			numericSeparators,
			ePrefix + "n2Dto ")

	if err != nil {
		return newNumStrDto, err
	}


	baseRunes := []rune(numStr)
	lBaseRunes := len(baseRunes)
	isStartRunes := false
	isEndRunes := false
	isMinusSignFound := false
	isFractionalValue := false

	absFracRunes := make([]rune, 0, 20)
	absIntRunes := make([]rune, 0, 20)

	for i := 0; i < lBaseRunes && isEndRunes == false; i++ {

		if baseRunes[i] != '-' &&
			baseRunes[i] != n2Dto.decimalSeparator &&
			(baseRunes[i] < '0' || baseRunes[i] > '9') {

			continue

		} else if baseRunes[i] == '-' &&
			isMinusSignFound == false &&
			isStartRunes == false && isEndRunes == false &&
			i+1 < lBaseRunes &&
			((baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9') ||
				baseRunes[i+1] == n2Dto.decimalSeparator) {

			isMinusSignFound = true
			n2Dto.signVal = -1
			isStartRunes = true
			continue

		} else if isEndRunes == false &&
			baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, baseRunes[i])
			isStartRunes = true

			if isFractionalValue {
				absFracRunes = append(absFracRunes, baseRunes[i])
			} else {
				absIntRunes = append(absIntRunes, baseRunes[i])
			}

		} else if isEndRunes == false &&
			i+1 < lBaseRunes &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' &&
			baseRunes[i] == n2Dto.decimalSeparator {

			isFractionalValue = true
			continue

		}

		if i == lBaseRunes-1 {
			isEndRunes = true
		}

	}

	lenAbsAllNumRunes := len(n2Dto.absAllNumRunes)

	if lenAbsAllNumRunes == 0 {
		newNumStrDto = nStrDtoElectron.newBaseZeroNumStrDto(0)
		return newNumStrDto, err
	}

	lenAbsIntNumRunes := len(absIntRunes)
	if lenAbsIntNumRunes == 0 {
		absIntRunes = append(absIntRunes, '0')
	}

	lenAbsAllNumRunes = len(n2Dto.absAllNumRunes)
	lenAbsIntNumRunes = len(absIntRunes)
	lenAbsFracNumRunes := len(absFracRunes)

	isZeroVal := true

	for i := 0; i < lenAbsAllNumRunes; i++ {
		if n2Dto.absAllNumRunes[i] != '0' {
			isZeroVal = false
		}
	}

	if isZeroVal {
		newNumStrDto = nStrDtoElectron.newBaseZeroNumStrDto(uint(lenAbsFracNumRunes))
		return newNumStrDto, err
	}

	if isFractionalValue {
		n2Dto.precision = uint(len(absFracRunes))
	}

	if lenAbsAllNumRunes != lenAbsIntNumRunes+lenAbsFracNumRunes {
		n2Dto.absAllNumRunes = []rune{}
		newLenAbsAllNumRunes := lenAbsIntNumRunes + lenAbsFracNumRunes

		for i := 0; i < newLenAbsAllNumRunes; i++ {
			if i < lenAbsIntNumRunes {

				n2Dto.absAllNumRunes =
					append(n2Dto.absAllNumRunes, absIntRunes[i])
			} else {
				n2Dto.absAllNumRunes =
					append(n2Dto.absAllNumRunes, absFracRunes[i-lenAbsIntNumRunes])
			}
		}

		lenAbsAllNumRunes = len(n2Dto.absAllNumRunes)
	}

	newNumStrDto,
		err = nStrDtoElectron.copyOut(
		&n2Dto,
		ePrefix+ "n2Dto->newNumStrDto ")

	if err !=nil {
		return newNumStrDto, err
	}

	_,
		err = nStrDtoElectron.testNumStrDtoValidity(
		&newNumStrDto,
		ePrefix + "Final newNumStrDto ")

	return newNumStrDto, err
}



