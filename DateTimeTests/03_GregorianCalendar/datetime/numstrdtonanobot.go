package datetime

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
)

type numStrDtoNanobot struct {
	lock *sync.Mutex
}

// findSignificantDigitLimits - Analyzes an array of characters which
// constitute a number string and returns the significant digits in the
// form of a new NumStrDto instance. This operation will effectively
// eliminate leading zeros from the integer value and trailing zeros from
// the fractional value.
//
// See the section below on Example Usage.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numSeparators       NumericSeparatorDto
//     - The numeric separator values contained in this input
//       parameter will be copied to the returned input parameter
//       'newNumStrDto', a newly created instance of NumStrDto.
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
func (nStrDtoNanobot *numStrDtoNanobot) findSignificantDigitLimits(
	numSeparators NumericSeparatorDto,
	absAllRunes []rune,
	precision uint,
	signVal int,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.findSignificantDigitLimits() "

	err = nil

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)


	if len(absAllRunes) == 0 {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'absAllRunes' is INVALID!\n" +
			"'absAllRunes' is a zero length array!\n")

		return newNumStrDto, err
	}

	if signVal != 1 &&
			signVal != -1 {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'signVal' is INVALID!\n" +
			"'signVal' represents the numeric sign of a number.\n" +
			"Valid 'signVal' values are +1 or -1.\n" +
			"signVal='%v'\n",
			signVal)
		return newNumStrDto, err
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	numSeparators.SetToUSADefaultsIfEmpty()

	err = nStrDtoElectron.setNumericSeparatorsDto(
		&newNumStrDto,
		numSeparators,
		ePrefix + "Setting 'newNumStrDto' numeric separators ")

	if err != nil {
		return newNumStrDto, err
	}

	iPrecision := int(precision)
	firstIntIdx := -1
	lastIntIdx := -1
	lastFracIdx := -1

	isFractional := false

	if iPrecision > 0 {
		isFractional = true
	}

	lenAbsAllRunes := len(absAllRunes)
	lenAbsFracRunes := iPrecision
	lenAbsIntRunes := lenAbsAllRunes - lenAbsFracRunes

	for i := 0; i < lenAbsAllRunes; i++ {

		if i < lenAbsIntRunes {

			if firstIntIdx == -1 &&
				absAllRunes[i] > '0' &&
				absAllRunes[i] <= '9' {

				firstIntIdx = i

			}

			lastIntIdx = i
		}

		if isFractional &&
			i >= lenAbsIntRunes &&
			absAllRunes[i] > '0' &&
			absAllRunes[i] <= '9' {

			lastFracIdx = i

		}

	}

	if firstIntIdx == -1 {

		firstIntIdx = lastIntIdx

	}

	if isFractional &&
		lastFracIdx == -1 {

		lastFracIdx = lenAbsIntRunes

	}

	numStrOut := ""

	if signVal < 0 {
		numStrOut = "-"
	}

	numStrOut +=
		string(absAllRunes[firstIntIdx : lastIntIdx+1])

	if isFractional {

		numStrOut +=
			string(numSeparators.DecimalSeparator)

		numStrOut +=
			string(absAllRunes[lastIntIdx+1 : lastFracIdx+1])

	}

	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto,
	err = nStrDtoAtom.parseNumStr(
		numStrOut,
		numSeparators,
		ePrefix + "numStrOut ")


	return newNumStrDto, err
}

// NewBigFloat - Creates a new NumStrDto instance from a Big Float value
// (*big.Float) and a precision specification.
//
// For more information on the *big.Float floating point numeric value,
// reference:
//   https://golang.org/pkg/math/big/
//
//
// See the 'Example Usage' section below.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numSeps             NumericSeparatorDto
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
//
//  err                error
//     - If this method completes successfully, the returned error Type
//       is set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. Said text will be prefixed
//       to the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     numSepsDto := NumericSeparatorDto{}
//     numSepsDto.SetToUSADefaults()
//
//     f64Num := float64(123.456)
//     bigFloatNum := big.NewFloat(f64Num)
//     precision := uint(2)
//     ePrefix := "calling method name "
//
//          nDto, err  :=
//              numStrDtoUtility.newBigFloat(
//              numSepsDto,
//              bigFloatNum,
//              precision,
//              ePrefix)
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
func (nStrDtoNanobot *numStrDtoNanobot) newBigFloat(
	numSepsDto NumericSeparatorDto,
	bigFloatNum *big.Float,
	precision uint,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.newBigFloat() "

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	numSepsDto.SetToUSADefaultsIfEmpty()

	if bigFloatNum == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'bigFloatNum' has a 'nil' pointer!\n")
		return newNumStrDto, err
	}

	numStr := bigFloatNum.Text('f', int(precision))

	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto,
		err = nStrDtoAtom.parseNumStr(
		numStr,
		numSepsDto,
		ePrefix + "numStr ")

	return newNumStrDto, err
}

// newBigInt - receives a signed Big Integer number (type *big.Int)
// and  precision parameter. This method then proceeds to generate
// and return a new NumStrDto type encapsulating the numeric value
// of the passed Big Integer number.
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
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//
//  err                error
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
//     numSeps := NumericSeparatorDto{}
//     numSeps.SetToUSADefaults()
//     precision := uint(3)
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//          nDto, err :=
//            nStrDtoNanobot.newBigInt(
//            numSeps,
//            bigIntNum,
//            precision,
//            "calling method name ")
//
//           nDto is now equal to 123.456
//
//  Example #2
//     int64Num := int64(123456)
//     bigIntNum := big.NewInt(int64Num)
//     numSeps := NumericSeparatorDto{}
//     numSeps.SetToUSADefaults()
//     precision := uint(0)
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//          nDto, err :=
//            nStrDtoNanobot.newBigInt(
//            numSeps,
//            bigIntNum,
//            precision,
//            "calling method name ")
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
func (nStrDtoNanobot *numStrDtoNanobot) newBigInt(
	numSepsDto NumericSeparatorDto,
	bigIntNum *big.Int,
	precision uint,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.newBigInt() "

	err = nil
	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	numSepsDto.SetToUSADefaultsIfEmpty()

	if bigIntNum == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'bigIntNum' has a 'nil' pointer!\n")

		return newNumStrDto, err
	}

	err = nStrDtoElectron.setNumericSeparatorsDto(
		&newNumStrDto,
		numSepsDto,
		ePrefix + "newNumStrDto ")

	if err != nil {
		return newNumStrDto, err
	}

	newNumStrDto.precision = precision
	scratchNum := big.NewInt(0).Set(bigIntNum)
	bigZero := big.NewInt(0)
	newNumStrDto.signVal = 1

	if scratchNum.Cmp(bigZero) == -1 {
		scratchNum.Neg(scratchNum)
		newNumStrDto.signVal = -1
	}

	bigTen := big.NewInt(int64(10))
	modulo := big.NewInt(0)
	newNumStrDto.absAllNumRunes =
		make([]rune, 0, 100)

	if scratchNum.Cmp(bigZero) == 0 {

		newNumStrDto.absAllNumRunes =
			append(newNumStrDto.absAllNumRunes,
				'0')

	} else {

		for scratchNum.Cmp(bigZero) == 1 {
			modulo = big.NewInt(0).Rem(scratchNum, bigTen)
			scratchNum = big.NewInt(0).Quo(scratchNum, bigTen)
			newNumStrDto.absAllNumRunes =
				append(newNumStrDto.absAllNumRunes,
					rune(modulo.Int64()+int64(48)))
		}
	}

	lenAllNumRunes :=
		len(newNumStrDto.absAllNumRunes)

	if int(newNumStrDto.precision) >=
			lenAllNumRunes {

		deltaNumRunes :=
			int(newNumStrDto.precision) - lenAllNumRunes + 1

		for k := 0; k < deltaNumRunes; k++ {
			newNumStrDto.absAllNumRunes =
				append(newNumStrDto.absAllNumRunes,
					'0')
			lenAllNumRunes++
		}

	}

	tRune := rune(0)

	if lenAllNumRunes > 1 {
		xLen := lenAllNumRunes - 1
		sortLimit := xLen / 2
		yCnt := 0
		for i := xLen; i > sortLimit; i-- {
			tRune =
				newNumStrDto.absAllNumRunes[yCnt]

			newNumStrDto.absAllNumRunes[yCnt] =
				newNumStrDto.absAllNumRunes[i]

			newNumStrDto.absAllNumRunes[i] = tRune

			yCnt++
		}
	}

	err = nStrDtoElectron.setNumericSeparatorsDto(
		&newNumStrDto,
		numSepsDto,
		ePrefix + "newNumStrDto #2 ")

	if err != nil {
		return newNumStrDto, err
	}

	_,
	err =
	nStrDtoElectron.testNumStrDtoValidity(
		&newNumStrDto,
		ePrefix + "Final Validity Check-newNumStrDto ")

	return newNumStrDto, err
}

// newFloat64 - Creates a new NumStrDto instance from a float64
// and precision specification.
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
//  numSepsDto := NumericSeparatorDto{}
//  numSepsDto.SetToUSADefaults()
//
//           f64 := float64(123.456)
//     precision := uint(2)
//
//          nDto, err  :=
//              numStrDtoUtility.newFloat64(
//              numSepsDto,
//              f64,
//              precision,
//              "calling method name ")
//
//           nDto is now equal to 123.46
//
//  Examples:
//  ---------
//                                newNumStrDto
//    f64        precision           Result
//  ------------------------------------------
//
//   12.3456         4                  12.3456
//   123456.5        0              123457
//   1234.56         1                1234.6
//
func (nStrDtoNanobot *numStrDtoNanobot) newFloat64(
	numSepsDto NumericSeparatorDto,
	f64 float64,
	precision uint,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.newFloat64() "

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	numSepsDto.SetToUSADefaultsIfEmpty()

	// Number string rounded to 'precision' decimal
	// places.
	numStr := strconv.FormatFloat(f64,
		'f',
		int(precision),
		64)

	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto,
		err = nStrDtoAtom.parseNumStr(
		numStr,
		numSepsDto,
		ePrefix)

	return newNumStrDto, err
}

// newInt64 - Creates a new NumStrDto instance from an int64 and a
// precision specification.
//
//
// --------------------------------------------------------------------------------------------------
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
//  newNumStrDto        NumStrDto
//     - A new instance of NumStrDto encapsulating the numeric value
//       calculated from the input parameters.
//
//
//  err                error
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
//     numSeps := NumericSeparatorDto{}
//     numSeps.SetToUSADefaults()
//     precision := uint(3)
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//          nDto, err :=
//            nStrDtoNanobot.newInt64(
//            numSeps,
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
func (nStrDtoNanobot *numStrDtoNanobot) newInt64(
	numSepsDto NumericSeparatorDto,
	int64Num int64,
	precision uint,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.newInt64() "

	err = nil
	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	numSepsDto.SetToUSADefaultsIfEmpty()

	numStr := strconv.FormatInt(int64Num, 10)

	nStrDtoMolecule := numStrDtoMolecule{}

	newNumStrDto,
		err = nStrDtoMolecule.setPrecision(
		numSepsDto,
		numStr,
		precision,
		true,
		ePrefix + "numSepsDto -> newNumStrDto ")

	return newNumStrDto, err
}


// NewInt64Exponent - Returns a new NumStrDto instance. The numeric
// value is set using an int64 value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//    numeric value = int64 X 10^exponent
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numSeps             NumericSeparatorDto
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
//     numSeps := NumericSeparatorDto{}
//     numSeps.SetToUSADefaults()
//     int64Num := int64(123456)
//     exponent := -3
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//  nDto, err := nStrDtoNanobot.newInt64Exponent(
//               int64Num,
//               exponent)
//
//  -- nDto is now equal to "123.456", precision = 3
//
//     numSeps := NumericSeparatorDto{}
//     numSeps.SetToUSADefaults()
//     int64Num := int64(123456)
//     exponent := 3
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//  nDto, err := nStrDtoNanobot.newInt64Exponent(
//                 int64Num,
//                 exponent)
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
func (nStrDtoNanobot *numStrDtoNanobot) newInt64Exponent(
	numSepsDto NumericSeparatorDto,
	int64Num int64,
	exponent int,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.newInt64Exponent() "

	err = nil
	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	numSepsDto.SetToUSADefaultsIfEmpty()

	numStr := strconv.FormatInt(int64Num, 10)

	if exponent > 0 {
		for i := 0; i < exponent; i++ {
			numStr += "0"
		}
	}

	if exponent < 0 {
		exponent = exponent * -1
	}

	if exponent == 0 {
		nStrDtoAtom := numStrDtoAtom{}

		newNumStrDto,
		err = nStrDtoAtom.parseNumStr(
			numStr,
			numSepsDto,
			ePrefix)

	} else {

		nStrDtoMolecule := numStrDtoMolecule{}

		newNumStrDto,
			err = nStrDtoMolecule.shiftPrecisionLeft(
			numSepsDto,
			numStr,
			uint(exponent),
			ePrefix)

	}

	return newNumStrDto, err
}

// NewNumStr - Used to create a populated NumStrDto instance using a
// valid number string as an input parameter.
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
//
// ----------------------------------------------------------------
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
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (nStrDtoNanobot *numStrDtoNanobot) newNumStr(
	numSepsDto NumericSeparatorDto,
	numStr string,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.newNumStr() "

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	if len(numStr) == 0 {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'numStr' is a zero length string!\n")

		return newNumStrDto, err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()


	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto,
		err = nStrDtoAtom.parseNumStr(
		numStr,
		numSepsDto,
		ePrefix)

	return newNumStrDto, err
}

// newRational - Creates a new NumStrDto instance from a rational
// number and a precision specification.
//
// For information on Big Rational Numbers (*big.Rat), reference:
//    https://golang.org/pkg/math/big/
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numSeps             NumericSeparatorDto
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
//  precision       uint
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
func (nStrDtoNanobot *numStrDtoNanobot) newRational(
	numSeps NumericSeparatorDto,
	bigRatNum *big.Rat,
	precision uint,
	ePrefix string)(
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.newRational() "

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	if bigRatNum == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'bigRatNum' has a nil pointer!\n")

		return newNumStrDto, err
	}

	numSeps.SetToUSADefaultsIfEmpty()

	numStr := bigRatNum.FloatString(int(precision))

	nStrDtoAtom := numStrDtoAtom{}

	newNumStrDto,
	err = nStrDtoAtom.parseNumStr(
		numStr,
		numSeps,
		ePrefix)

	return newNumStrDto, err
}

// NewUint64 - Creates a new NumStrDto instance from an uint64 and a
// precision specification.
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
//  uint64Num           uint64
//     - This numeric value will be converted to a new instance of
//       type NumStrDto.
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
//     uint64Num := uint64(123456)
//     numSepsDto := NumericSeparatorDto{}
//     numSeps.SetToUSADefaults()
//     precision := uint(3)
//
//          nDto, err :=
//            numStrDtoUtility.newUint64(
//            numSepsDto,
//            uint64Num,
//            precision,
//            "calling method name ")
//
//           nDto is now equal to 123.456
//
//  Examples:
//  ---------
//
//  <-- Input Parameters -->     <-- Output -->
//                                newNumStrDto
//  uint64Num     precision          Result
//  ------------------------------------------
//
//   123456          4                  12.3456
//   123456          0              123456
//   123456          1               12345.6
//
func (nStrDtoNanobot *numStrDtoNanobot) newUint64(
	numSepsDto NumericSeparatorDto,
	uint64Num uint64,
	precision uint,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.newUint64() "

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	numSepsDto.SetToUSADefaultsIfEmpty()

	numStr := strconv.FormatUint(uint64Num, 10)

	nStrDtoMolecule := numStrDtoMolecule{}

	newNumStrDto,
		err = nStrDtoMolecule.setPrecision(
		numSepsDto,
		numStr,
		precision,
		true,
		ePrefix + "numSepsDto -> newNumStrDto ")

	return newNumStrDto, err
}

// newUint64Exponent - Creates a new NumStrDto instance from a type uint64 and a
// precision specification.
//
// See the Example Usage section below.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numSeps             NumericSeparatorDto
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
//     numSeps := NumericSeparatorDto{}
//     numSeps.SetToUSADefaults()
//     uint64Num := uint64(123456)
//     exponent := -3
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//  nDto,err := nStrDtoNanobot.newUint64Exponent(
//          numSeps,
//          uint64Num,
//          exponent,
//          "calling method name ")
//   -- nDto is now equal to "123.456", precision = 3
//
//
//     numSeps := NumericSeparatorDto{}
//     numSeps.SetToUSADefaults()
//     uint64Num := uint64(123456)
//     exponent := 3
//     nStrDtoNanobot := numStrDtoNanobot{}
//
//  nDto,err := nStrDtoNanobot.newUint64Exponent(
//          numSeps,
//          uint64Num,
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
//   uint64Num          exponent            Result
//   123456               -3               123.456
//   123456                3            123456.000
//   123456                0            123456
//
func (nStrDtoNanobot *numStrDtoNanobot) newUint64Exponent(
	numSeps NumericSeparatorDto,
	uint64Num uint64,
	exponent int,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.newUint64Exponent() "

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	err = nil

	numSeps.SetToUSADefaultsIfEmpty()

	numStr := strconv.FormatUint(uint64Num, 10)

	if exponent > 0 {
		for i := 0; i < exponent; i++ {
			numStr += "0"
		}
	}

	if exponent < 0 {
		exponent = exponent * -1
	}

	if exponent == 0 {

		nStrDtoAtom := numStrDtoAtom{}

		newNumStrDto,
			err = nStrDtoAtom.parseNumStr(
			numStr,
			numSeps,
			ePrefix + "exponent == 0 ")

	} else {

		nStrDtoMolecule := numStrDtoMolecule{}

		newNumStrDto,
			err = nStrDtoMolecule.shiftPrecisionLeft(
			numSeps,
			numStr,
			uint(exponent),
			ePrefix + "numStr ")

	}

	return newNumStrDto, err
}

// scaleNumStr - Receives a signed number string and proceeds
// to shifts the position of the decimal point left or right
// depending on the value of input parameter 'scaleMode'.
//
// The scaled number string will then be converted to a new
// instance of NumStrDto and returned to the caller.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  numSeparators       NumericSeparatorDto
//     - The numeric separator values contained in this input
//       parameter will be copied to the returned input parameter
//       'newNumStrDto', a newly created instance of NumStrDto.
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
//       numStrDtoMolecule.shiftPrecisionLeft() for additional
//       information.
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
//     - If this method completes successfully, the result of the numeric
//       scaling operation performed by this method will be returned in
//       the form of a new 'NumStrDto' instance.
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (nStrDtoNanobot *numStrDtoNanobot) scaleNumStr(
	numSeparators NumericSeparatorDto,
	signedNumStr string,
	shiftPrecision uint,
	scaleMode PrecisionScaleMode,
	ePrefix string) (
	newNumStrDto NumStrDto,
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.scaleNumStr() "

	err = nil

	nStrDtoElectron := numStrDtoElectron{}

	newNumStrDto =
		nStrDtoElectron.newBaseZeroNumStrDto(0)

	if len(signedNumStr) == 0 {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'signedNumStr' is INVALID!\n" +
			"'signedNumStr' is a zero length number string!\n")

		return newNumStrDto, err
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	numSeparators.SetToUSADefaultsIfEmpty()

	err = nStrDtoElectron.setNumericSeparatorsDto(
		&newNumStrDto,
		numSeparators,
		ePrefix + "Setting 'newNumStrDto' numeric separators ")

	if err != nil {
		return newNumStrDto, err
	}

	nStrDtoMolecule := numStrDtoMolecule{}

	var err2 error

	if scaleMode == SCALEPRECISIONLEFT {

		newNumStrDto,
		err2 = nStrDtoMolecule.shiftPrecisionLeft(
			numSeparators,
			signedNumStr,
			shiftPrecision,
			ePrefix)

		if err2 != nil {
			err = fmt.Errorf(ePrefix + "\n" +
				"Error returned from nStrDtoMolecule.ShiftPrecisionLeft" +
				"(signedNumStr, shiftPrecision)\n"+
				"signedNumStr='%v'\n" +
				"shiftPrecision='%v'\n" +
				"scaleMode='%v'\n" +
				"Error='%v'\n",
				signedNumStr,
				shiftPrecision,
				scaleMode.String(),
				err2.Error())

			return newNumStrDto, err
		}

	} else if scaleMode == SCALEPRECISIONRIGHT {

		newNumStrDto, err2 =
			nStrDtoMolecule.shiftPrecisionRight(
			numSeparators,
			signedNumStr,
			shiftPrecision,
			ePrefix)


		if err2 != nil {
			err =	fmt.Errorf(ePrefix + "\n" +
				"Error returned from nStrDtoMolecule.ShiftPrecisionRight" +
				"(signedNumStr, shiftPrecision)\n"+
				"signedNumStr='%v'\n" +
				"shiftPrecision='%v'\n" +
				"scaleMode='%v'\n" +
				"Error='%v'\n",
				signedNumStr,
				shiftPrecision,
				scaleMode.String(),
				err2.Error())

			return newNumStrDto, err
		}

	} else {

		err = fmt.Errorf(ePrefix + "\n" +
			"Error! Scale Mode is INVALID!\n"+
			"Scale Mode is NOT Equal to SCALEPRECISIONLEFT or SCALEPRECISIONRIGHT.\n" +
			"scaleMode='%v'\n",
			scaleMode.String())

		return newNumStrDto, err
	}

	return newNumStrDto, err
}

// setNumStrDtoPrecision - Sets or resets the precision for a passed instance
// of NumStrDto (Input parameter 'numStrDto').
//
// Input parameter 'precision' identifies the number of decimal places to the
// right of the decimal point which will be configured in 'numStrDto'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrDto           *NumStrDto
//     - A pointer to an instance of NumStrDto. This method WILL
//       CHANGE the values of internal member variables to achieve
//       the method's objectives. Member variables will be tested for
//       validity.
//
//       This method will set or reset the 'precision' of the numeric
//       value encapsulated by this instance of NumStrDto. 'precision',
//       as defined here, specifies the number of digits to the right
//       of the decimal point which will be formatted in the numeric
//       value encapsulated in parameter, 'numStrDto'.
//
//
//  precision           uint
//     - The number of numeric digits to the right of the decimal place
//       which will be configured in the numeric value encapsulated within
//       input parameter 'numStrDto'.
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
//       passed by input parameter, 'ePrefix'. The 'ePrefix' text will be
//       prefixed to the beginning of the returned error message.
//
func (nStrDtoNanobot *numStrDtoNanobot) setNumStrDtoPrecision(
	numStrDto *NumStrDto,
	precision uint,
	roundResult bool,
	ePrefix string) (
	err error) {

	if nStrDtoNanobot.lock == nil {
		nStrDtoNanobot.lock = new(sync.Mutex)
	}

	nStrDtoNanobot.lock.Lock()

	defer nStrDtoNanobot.lock.Unlock()

	ePrefix += "numStrDtoNanobot.setNumStrDtoPrecision() "

	err = nil

	var numStr string

	nStrDtoMolecule := numStrDtoMolecule{}

	numStr,
	err = nStrDtoMolecule.getNumStr(
		numStrDto,
		ePrefix + "Input parameter 'numStrDto' ")

	if err != nil {
		return err
	}

	nStrDtoAtom := numStrDtoAtom{}

	var numSepsDto NumericSeparatorDto

	numSepsDto,
	err =
	nStrDtoAtom.getNumericSeparatorsDto(
		numStrDto,
		ePrefix + "numStrDto -> numSepsDto ")

	if err != nil {
		return err
	}

	numSepsDto.SetToUSADefaultsIfEmpty()

	var n2 NumStrDto

	n2,
	err = nStrDtoMolecule.setPrecision(
		numSepsDto,
		numStr,
		precision,
		roundResult,
		ePrefix + "numStr ")

	if err != nil {
		return err
	}

	nStrDtoElectron := numStrDtoElectron{}

	err = nStrDtoElectron.setNumericSeparatorsDto(
		&n2,
		numSepsDto,
		ePrefix + "n2 ")

	if err != nil {
		return err
	}

	err = nStrDtoElectron.copyIn(
		numStrDto,
		&n2,
		ePrefix + "n2 -> numStrDto ")

	if err != nil {
		return err
	}

_,
err =
	nStrDtoElectron.testNumStrDtoValidity(
		numStrDto,
		ePrefix)

return err
}
