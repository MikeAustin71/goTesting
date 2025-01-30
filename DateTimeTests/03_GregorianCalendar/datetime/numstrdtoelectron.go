package datetime

import (
	"errors"
	"fmt"
	"sync"
)

type numStrDtoElectron struct {

	lock *sync.Mutex
}

// CopyIn - Receives an incoming NumStrDto object
// and copies the information to the current NumStrDto
// data fields.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       change and overwrite the values of internal member variables to achieve the method's objectives.
//
//       This NumStrDto will receive all the data values contained
//       in input parameter 'nInDto'. When the copy operation is
//       completed, the data values in 'numStrDto' will be identical
//       to those contained in input parameter, 'nInDto'.
//
//
//  nInDto              *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       All data values in this NumStrDto instance will be copied to
//       input parameter 'numStrDto'.
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
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) copyIn(
	numStrDto *NumStrDto,
	nInDto *NumStrDto,
	ePrefix string) (
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.copyIn() "
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return err
	}

	if nInDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'nInDto' is INVALID!\n" +
			"nInDto = nil pointer!\n")
		return err
	}

	numStrDto.signVal = nInDto.signVal
	numStrDto.absAllNumRunes = nInDto.absAllNumRunes
	numStrDto.precision = nInDto.precision
	numStrDto.thousandsSeparator = nInDto.thousandsSeparator
	numStrDto.decimalSeparator = nInDto.decimalSeparator
	numStrDto.currencySymbol = nInDto.currencySymbol

	return err
}

// CopyOut - Creates a deep copy of the data fields contained in
// input parameter 'numStrDto' and returns that data as a new
// instance NumStrDto.
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
//  newNumStrDto        NumStrDt
//     - If this method completes successfully, a deep copy of
//       input parameter 'numStrDto' will be returned.
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
func (nStrDtoElectron *numStrDtoElectron) copyOut(
	numStrDto *NumStrDto,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.copyOut() "

	newNumStrDto = NumStrDto{}
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return newNumStrDto, err
	}

	newNumStrDto.signVal =
		numStrDto.signVal

	newNumStrDto.absAllNumRunes =
		numStrDto.absAllNumRunes

	newNumStrDto.precision =
		numStrDto.precision

	newNumStrDto.thousandsSeparator =
		numStrDto.thousandsSeparator

	newNumStrDto.decimalSeparator =
		numStrDto.decimalSeparator

	newNumStrDto.currencySymbol =
		numStrDto.currencySymbol

	return newNumStrDto, err
}

// empty - Performs an 'empty' operation on an instance
// of NumStrDto. All data values encapsulated by input
// parameter 'numStrDto' are set to their zero or default
// values.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       set ALL the data values of internal member variables
//       to their zero or default values.
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
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) empty(
	numStrDto *NumStrDto,
	ePrefix string) (
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}
	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.empty() "

	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return err
	}

	// Perform an 'empty' operation
	numStrDto.signVal = 0
	numStrDto.absAllNumRunes = []rune{}
	numStrDto.precision = 0
	numStrDto.thousandsSeparator = ','
	numStrDto.decimalSeparator = '.'
	numStrDto.currencySymbol = '$'

	return err
}



// getAbsFracRunes - Returns all of the fractional digits
// to the right of the decimal place in the NumStrDto input
// parameter, 'numStrDto'. The fractional digits will be returned
// as an array of runes. The rune array is not signed; that is,
// the rune array does not contain a '+' or '-' character in the
// first array position. The rune array is therefore said to
// represent the absolute value of the fractional digits in the
// 'numStrDto' numeric value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Any fractional digits contained in
//       the numeric value encapsulated by this NumStrDto instance
//       will be returned as an array of runes.
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
//  absFracRunes        []rune
//     - If this method completes successfully, an array of runes will
//       be returned. This rune array contains the fractional digits to
//       the right of the decimal point extracted from input parameter
//       numStrDto.
//
//       Note that the rune array is not signed; that is, the rune array
//       will not contain a '+' or '-' character in the first array
//       position. The rune array is therefore said to represent the
//       absolute value of the fractional digits in the input parameter
//       'numStrDto' numeric value.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) getAbsFracRunes(
	numStrDto *NumStrDto,
	ePrefix string) (
	absFracRunes []rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.getAbsFracRunes() "

	absFracRunes = make([]rune, 0, 20)

	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return absFracRunes, err
	}

	precision := int(numStrDto.precision)

	lenAllNums := len(numStrDto.absAllNumRunes)

	if lenAllNums == 0 ||
		precision < 0 ||
		precision > lenAllNums {
		return absFracRunes, err
	}

	absFracRunes = make([]rune, precision, precision+50)

	lenIntNums := lenAllNums - int(numStrDto.precision)

	for i := lenIntNums; i < lenAllNums; i++ {
		absFracRunes[i-lenIntNums] = numStrDto.absAllNumRunes[i]
	}

	return absFracRunes, err
}

// getAbsFracRunesLength - Returns the length of the
// fractional digits in the number string.
//
func (nStrDtoElectron *numStrDtoElectron) getAbsFracRunesLength(
	numStrDto *NumStrDto) (
	absFracRunesLen int) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	absFracRunesLen = 0

	if numStrDto == nil {
		return absFracRunesLen
	}

	return int(numStrDto.precision)
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
func (nStrDtoElectron *numStrDtoElectron) getAbsIntRunes(
	numStrDto *NumStrDto,
	ePrefix string) (
	absIntRunes []rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.getAbsIntRunes() "

	absIntRunes = []rune{}

	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return absIntRunes, err
	}

	lenAllNum := len(numStrDto.absAllNumRunes)

	precision := int(numStrDto.precision)

	if lenAllNum == 0 ||
		precision < 0 ||
		precision >= lenAllNum {
		return absIntRunes, err
	}

	lenIntNum := lenAllNum - precision

	absIntRunes = make([]rune, lenIntNum, lenIntNum+50)

	for i := 0; i < lenIntNum; i++ {
		absIntRunes[i] = numStrDto.absAllNumRunes[i]
	}

	return absIntRunes, err
}


// getAbsIntRunesLength - Returns the length of the
// integer portion of the number string.
//
func (nStrDtoElectron *numStrDtoElectron) getAbsIntRunesLength(
	numStrDto *NumStrDto) (
	absIntRunesLen int) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	absIntRunesLen = 0

	if numStrDto == nil {
		return absIntRunesLen
	}

	absIntRunesLen =
		len(numStrDto.absAllNumRunes) -
			int(numStrDto.precision)

	return absIntRunesLen
}

// getAbsoluteValueNumStr - Returns all numeric digits residing in
// the NumStrDto instance passed as input parameter 'numStrDto'. This
// number string is a pure, unsigned number string consisting entirely
// of numeric digits. There is no leading number sign ('+' or '-') and
// this is no decimal delimiter ('.'). If fractional digits exists they
// are included and NOT separated by a decimal separator.
//
//     Examples:
//
//      numstr      result
//      ------      ------
//      123.45      12345
//     -123.45      12345
//
// Effectively, this method will return the member variable 'absAllNumRunes'
// from the input parameter, 'numStrDto'.
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
//       The absolute value rune array from this instance will be
//       returned.
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
// absValNumStr         string
//     - A string of numeric digits representing the absolute numeric value
//       encapsulated in the NumStrDto instance passed as input parameter
//       'numStrDto'.
//
//
//  error
//     - If this method completes successfully the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note that this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. 'ePrefix' will be prefixed to the beginning
//       of the returned error message.
//
//
func (nStrDtoElectron *numStrDtoElectron) getAbsoluteValueNumStr(
	numStrDto *NumStrDto,
	ePrefix string) (
	absValNumStr string,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.getAbsoluteValueNumStr() "

	absValNumStr = ""
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return absValNumStr, err
	}

	absValNumStr = string(numStrDto.absAllNumRunes)

	return absValNumStr, err
}

// getDecimalSeparator - returns the character designated
// as the decimal separator for the current NumStrDto instance.
//
// In the USA, the decimal separator is the period character ('.').
//
// Example:  123.456
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
//  decimalSeparator    rune
//     - If this method completes successfully, this rune value will
//       contain the decimal separator associated with input parameter
//       'numStrDto'.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) getDecimalSeparator(
	numStrDto *NumStrDto,
	ePrefix string) (
	decimalSeparator rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.getDecimalSeparator() "

	decimalSeparator = 0
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return decimalSeparator, err
	}

	if numStrDto.decimalSeparator == 0 {
		numStrDto.decimalSeparator = '.'
	}

	decimalSeparator = numStrDto.decimalSeparator

	return decimalSeparator, err
}

// getCurrencySymbol - Returns the character currently designated
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
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
//  currencySymbol      rune
//     - If this method completes successfully, this rune value will
//       contain the currency symbol associated with input parameter
//       'numStrDto'.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) getCurrencySymbol(
	numStrDto *NumStrDto,
	ePrefix string) (
	currencySymbol rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.getCurrencySymbol() "

	currencySymbol = 0
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return currencySymbol, err
	}

	if numStrDto.currencySymbol == 0 {
		numStrDto.currencySymbol = '$'
	}

	currencySymbol = numStrDto.currencySymbol

	return currencySymbol, err
}

//
// IMPORTANT
//
// This method does NOT validate the input parameter 'numStrDto'
// before returning the value. To run a validity check on the
// 'numStrDto' instance, first call the following method:
//
//  numStrDtoElectron.testNumStrDtoValidity() (bool, error)
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       The 'precision' value returned by this method will be extracted
//       from this NumStrDto instance.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
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
//       always be one of two values:
//
//         +1 = a positive number greater than or equal to zero ('0')
//                       - OR -
//         -1 = a negative number less than zero ('0')
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
func (nStrDtoElectron *numStrDtoElectron) getNumericSignValue(
	numStrDto *NumStrDto,
	ePrefix string) (
	signValue int,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.getPrecisionUint() "

	signValue = 0
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return signValue, err
	}

	signValue = numStrDto.signVal

	return signValue, err
}


// getPrecisionUint - Returns the precision associated with NumStrDto
// instance passed as an input parameter.
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
// integer value.
//
//
// Example:
//   1.234     getPrecisionUint() = 3
//   5         getPrecisionUint() = 0
//   0.12345   getPrecisionUint() = 5
//
//    Number                    Fractional
//    String     precision        Number
//    123456         3            123.456
//
// IMPORTANT
//
// This method does NOT validate the input parameter 'numStrDto'
// before returning the value. To run a validity check on the
// 'numStrDto' instance, first call the following method:
//
//  numStrDtoElectron.testNumStrDtoValidity() (bool, error)
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//       The 'precision' value returned by this method will be extracted
//       from this NumStrDto instance.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  precision           uint
//     - This method returns the 'precision' value associated with
//       the input parameter 'numStrDto' instance as a type uint.
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
func (nStrDtoElectron *numStrDtoElectron) getPrecisionUint(
	numStrDto *NumStrDto,
	ePrefix string) (
	precision uint,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.getPrecisionUint() "

	precision = 0
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return precision, err
	}

	precision = numStrDto.precision

	return precision, err
}

// getThousandsSeparator - returns a rune which represents
// the character currently used to separate thousands in
// the display of the current NumStrDto number string
// designated by input parameter 'numStrDto' .
//
// In the USA, the thousands separator is a comma character.
//
// Example: 1,000,000,000
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
//  thousandsSeparator  rune
//     - If this method completes successfully, this rune value will
//       contain the thousands separator associated with input parameter
//       'numStrDto'.
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
func (nStrDtoElectron *numStrDtoElectron) getThousandsSeparator(
	numStrDto *NumStrDto,
	ePrefix string) (
	thousandsSeparator rune,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.getThousandsSeparator() "


	thousandsSeparator = 0
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return thousandsSeparator, err
	}

	if numStrDto.thousandsSeparator == 0 {
		numStrDto.thousandsSeparator = ','
	}

	thousandsSeparator = numStrDto.thousandsSeparator

	return thousandsSeparator, err
}

// isNumStrZeroValue - Returns 'true' if all the digits in
// the number string for the passed NumStrDto instance are
// zero.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Separator values from this NumStrDto
//       instance will be used to populate the returned new instance of
//       NumStrDto.
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
//  isZeroValue         bool
//     - If this method completes successfully, this returned boolean
//       value will signal whether all the numeric digits in input
//       parameter 'numStrDto' are zero. If 'true', the numeric value
//       is zero and all digits in the number string are are zero.
//       If 'false', one or more digits are non-zero and the numeric
//       value is non-zero.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) isNumStrZeroValue(
	numStrDto *NumStrDto,
	ePrefix string) (
	isZeroValue bool,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.isNumStrZeroValue() "

	isZeroValue = false
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return isZeroValue, err
	}

	lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

	if lenAbsAllNumRunes == 0 {

		isZeroValue = false

	} else {

		isZeroValue = true

		for i := 0; i < lenAbsAllNumRunes; i++ {
			if numStrDto.absAllNumRunes[i] != '0' {
				isZeroValue = false
			}
		}

	}

	return isZeroValue, err
}

// isFractionalValue - Returns 'true' if the numeric value of
// input parameter 'numStrDto' (type NumStrDto) includes a
// fractional value; that is, the number has fractional digits
// to the right of the decimal point.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. The numeric value encapsulated by
//       this instance of NumStrDto will be evaluated to determine if
//       it contains fractional digits to the right of the decimal
//       point.
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
//  isFractionalValue   bool
//     - If this method completes successfully, this returned boolean
//       value will signal whether the numeric value encapsulated by
//       input parameter 'numStrDto' includes fractional digits to the
//       right of the decimal point. If 'true', 'numStrDto' has fractional
//       digits. If 'false', 'numStrDto' is an integer value with no
//       fractional digits to the right of the decimal point.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) isFractionalValue(
	numStrDto *NumStrDto,
	ePrefix string) (
	isFractionalValue bool,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.isFractionalValue() "

	isFractionalValue = false
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return isFractionalValue, err
	}

	if numStrDto.precision > 0 {
		isFractionalValue = true
	} else {
		isFractionalValue = false
	}

	return isFractionalValue, err
}

// newBaseZeroNumStrDto - Returns a new NumStrDto object
// initialized to zero value. If the parameter 'precision'
// is set to a value greater than zero, then an equal number
// of zero digits will be added to the right of the decimal
// point.
//
// This method will return a zero NumStrDto instance set to
// default USA separators (Currency Symbol, Thousands Separator,
// Decimal Separator).
//
// Example Usage:
//
// Input Parameter      newNumStrDto
//    precision            Result
//       0                  "0"
//       2                  "0.00"
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
//  newNumStrDto        NumStrDto
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
func (nStrDtoElectron *numStrDtoElectron) newBaseZeroNumStrDto(
	precision uint) (
	newNumStrDto NumStrDto) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	newNumStrDto = NumStrDto{}

	newNumStrDto.absAllNumRunes = make([]rune, 0, 50)

	// Set USA defaults for thousands separators,
	// decimal separators and currency Symbols
	newNumStrDto.thousandsSeparator = ','
	newNumStrDto.decimalSeparator = '.'
	newNumStrDto.currencySymbol = '$'
	newNumStrDto.signVal = 1
	newNumStrDto.signVal = 1
	newNumStrDto.precision = 0

	newNumStrDto.absAllNumRunes =
		append(newNumStrDto.absAllNumRunes, '0')

	if precision > 0 {

		for i := uint(0); i < precision; i++ {
			newNumStrDto.absAllNumRunes = append(newNumStrDto.absAllNumRunes, '0')
		}

		newNumStrDto.precision = precision
	}

	return newNumStrDto
}

// newZeroNumStrDto - returns a new NumStrDto initialized
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
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Separator values from this NumStrDto
//       instance will be used to populate the returned new instance of
//       NumStrDto.
//
//
//  numFracDigits       uint
//     - The number of zeros which will be placed to the right of the
//       decimal point in the returned new instance of NumStrDto.
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
//     - If this method completes successfully, this returned new instance
//       of NumStrDto will be populated with zero values and will contain
//       the same separator values as those passed through input parameter
//       'newNumStrDto'.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) newZeroNumStrDto(
	numStrDto *NumStrDto,
	numFracDigits uint,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.newZeroNumStrDto() "

	newNumStrDto = NumStrDto{}
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return newNumStrDto, err
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if numStrDto.thousandsSeparator == 0 {
		numStrDto.thousandsSeparator = ','
	}

	if numStrDto.decimalSeparator == 0 {
		numStrDto.decimalSeparator = '.'
	}

	if numStrDto.currencySymbol == 0 {
		numStrDto.currencySymbol = '$'
	}

	newNumStrDto.signVal = 1

	newNumStrDto.thousandsSeparator =
		numStrDto.thousandsSeparator

	newNumStrDto.decimalSeparator =
		numStrDto.decimalSeparator

	newNumStrDto.currencySymbol =
		numStrDto.currencySymbol

	newNumStrDto.signVal = 1
	newNumStrDto.precision = 0

	newNumStrDto.absAllNumRunes = make([] rune, 0, 50)

	newNumStrDto.absAllNumRunes =
		append(newNumStrDto.absAllNumRunes, '0')

	if numFracDigits > 0 {

		for i := uint(0); i < numFracDigits; i++ {
			newNumStrDto.absAllNumRunes = append(newNumStrDto.absAllNumRunes, '0')
		}

		newNumStrDto.precision = numFracDigits
	}

	return newNumStrDto, err
}

// setNumericSeparatorsToDefaultIfEmpty - Sets the numeric
// separators in the passed instance of NumStrDto to their
// their USA defaults IF any of the separators are set to
// zero.
//
// Effectively, this method ensures that numeric separators
// in a NumStrDto instance are set to valid values.
//
// USA default numeric separators are listed as follows:
//
//       Currency Symbol: '$'
//   Thousands Separator: ','
//     Decimal Separator: '.'
//
//
// If the numeric separators were previously set to a value
// other than zero or nil, that value is not altered by this
// method.
//
// Effectively, this method ensures that numeric separators
// are set to valid values.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       ensure that separator values are set to valid values. If
//       any of the current separator values are zero, this method
//       method will set those values to USA defaults.
//
//       USA default numeric separators are listed as follows:
//
//             Currency Symbol: '$'
//         Thousands Separator: ','
//           Decimal Separator: '.'
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
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) setNumericSeparatorsToDefaultIfEmpty(
	numStrDto *NumStrDto,
	ePrefix string) (
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.setNumericSeparatorsToDefaultIfEmpty() "

	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return err
	}

	if numStrDto.decimalSeparator == 0 {
		numStrDto.decimalSeparator = '.'
	}

	if numStrDto.thousandsSeparator == 0 {
		numStrDto.thousandsSeparator =','
	}

	if numStrDto.currencySymbol == 0 {
		numStrDto.currencySymbol = '$'
	}

	return err
}

// setNumericSeparatorsDto - Sets the values of numeric separators:
//
//                           decimal point separator
//                           thousands separator
//                           currency symbol
//
// These separators are set based on values transmitted through input
// parameter 'customSeparators'.
//
// If any of the values contained in input parameter 'customSeparators'
// is set to zero or nil, an error will be returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       populate the separator values for Currency Symbol, Thousands
//       Separator and Decimal Separator with values copied from
//       the second input parameter 'customSeparators'.
//
//
//  customSeparators    NumericSeparatorDto
//     - If any of the data fields in this passed structure 'customSeparators'
//       are set to zero ('0'), an this method will return an error.
//
//       The separator values contained in this input parameter will be
//       copied to input parameter 'numStrDto'. The data fields included
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
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) setNumericSeparatorsDto(
	numStrDto *NumStrDto,
	customSeparators NumericSeparatorDto,
	ePrefix string) (
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.setNumericSeparatorsDto() "

	if numStrDto == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'numStrDto' is a 'nil' pointer!\n")

		return err
	}

	if customSeparators.DecimalSeparator == 0 {
		err = errors.New(ePrefix +
			"Error: Input Parameter customSeparators.DecimalSeparator" +
			" is set to '0'\n" +
			"Invalid rune!\n")
		return err
	}

	if customSeparators.ThousandsSeparator == 0 {
		err = errors.New(ePrefix +
			"Error: Input Parameter customSeparators.ThousandsSeparator" +
			" is set to '0'\n" +
			"Invalid rune!\n")
		return err
	}

	if customSeparators.CurrencySymbol == 0 {
		err = errors.New(ePrefix +
			"Error: Input Parameter customSeparators.CurrencySymbol" +
			" is set to '0'\n" +
			"Invalid rune!\n")
		return err
	}

	numStrDto.decimalSeparator =
		customSeparators.DecimalSeparator

	numStrDto.thousandsSeparator =
		customSeparators.ThousandsSeparator

	numStrDto.currencySymbol =
		customSeparators.CurrencySymbol

	return err
}

// setNumericSeparatorsToUSADefault - Sets the numeric
// separators in the passed instance of NumStrDto to the
// USA default values.
//
// USA default numeric separators are listed as follows:
//
//       Currency Symbol: '$'
//   Thousands Separator: ','
//     Decimal Separator: '.'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method will
//       set the numeric separator values contained in this instance
//       to USA default values.
//
//       USA default numeric separators are listed as follows:
//             Currency Symbol: '$'
//         Thousands Separator: ','
//           Decimal Separator: '.'
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
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) setNumericSeparatorsToUSADefault(
	numStrDto *NumStrDto,
	ePrefix string) (
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.setNumericSeparatorsToUSADefault() "

	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return err
	}

	numStrDto.decimalSeparator = '.'

	numStrDto.decimalSeparator = '.'

	numStrDto.currencySymbol = '$'


	return err
}

// SetSignValue - Sets the sign of the numeric value for the input parameter
// 'numStrDto'.
//
// Input parameter 'newSignVal' must be set to one of two possible values:
//
//  +1 for positive numeric values greater than or equal to zero
//
//  -1 for negative numeric values less than zero
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
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (nStrDtoElectron *numStrDtoElectron) setSignValue(
	numStrDto *NumStrDto,
	newSignVal int,
	ePrefix string) (
	err error) {
	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.setSignValue() "

	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return err
	}

	if newSignVal != -1 && newSignVal != 1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Invalid sign value passed. Sign value must be +1 or -1.\n"+
			"This sign value= %v\n", newSignVal)
		return err
	}

	numStrDto.signVal = newSignVal

	return err
}

// testNumStrDtoValidity - Evaluates the validity of a NumStrDto
// instance. If the instance is valid this method returns 'true'
// and sets the returned error type to nil.
//
// If the instance is invalid, the method returns 'false' plus an
// error message.
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
//  isValid             bool
//     - If this method completes successfully, this boolean value will
//       signal 'false' if 'numStrDto' member variables contains an
//       invalid value. Absent processing errors, if the validity tests
//       show that 'numStrDto' member variables values are valid, this
//       'isValid' flag will be set to 'true'.
//
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//       An error will also be returned if the method determines that
//       one or more of the 'numStrDto' member variables contain
//       invalid values.
//
func (nStrDtoElectron *numStrDtoElectron) testNumStrDtoValidity(
	numStrDto *NumStrDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrDtoElectron.lock == nil {
		nStrDtoElectron.lock = new(sync.Mutex)
	}

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	ePrefix += "numStrDtoElectron.testNumStrDtoValidity() "

	isValid = false
	err = nil

	if numStrDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'numStrDto' is INVALID!\n" +
			"numStrDto = nil pointer!\n")
		return isValid, err
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if numStrDto.thousandsSeparator == 0 {
		numStrDto.thousandsSeparator = ','
	}

	if numStrDto.decimalSeparator == 0 {
		numStrDto.decimalSeparator = '.'
	}

	if numStrDto.currencySymbol == 0 {
		numStrDto.currencySymbol = '$'
	}

	lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

	if lenAbsAllNumRunes == 0 {
		err = errors.New(ePrefix + "\n" +
			"- Error: 'numStrDto' Number string is a ZERO length string!\n")
		return isValid, err
	}

	if int(numStrDto.precision) >= lenAbsAllNumRunes {
		err = errors.New(ePrefix + "\n" +
			"Error: precision does match number string. Type is Corrupted!")
		return isValid, err
	}

	if numStrDto.signVal != 1 && numStrDto.signVal != -1 {
		err = fmt.Errorf(ePrefix + "\n + " +
			"Sign Value is INVALID. Should be +1 or -1.\n" +
			"This Sign Value (numStrDto.signVal) is %v\n",
			numStrDto.signVal)

		return isValid, err
	}

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if numStrDto.absAllNumRunes[i] < '0' || numStrDto.absAllNumRunes[i] > '9' {

			err = errors.New(ePrefix + "\n" +
				"Error: Non-Numeric character found in " +
				"'numStrDto' number string!\n")
			return isValid, err
		}
	}

	return isValid, err
}

