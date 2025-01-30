package datetime

import (
	"fmt"
	"math/big"
	"sync"
)

// mathBigFloatNanobot - Provides helper methods for type
// MathBigFloatHelper.
//
type mathBigFloatNanobot struct {
	lock *sync.Mutex
}


// ceiling - Provides a standard 'ceiling' function for
// type *big.Float floating point number values.
//
// For a discussion of the 'ceiling' function, reference:
//  https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum       *big.Float
//     - This method will calculate and return the 'ceiling' value
//       associated with this parameter.
//
//
//  precision         uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on input parameter 'bigFloatNum'
//       as well as specifying the numeric precision incorporated in the
//       returned floating point value, 'ceiling'.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  ceiling           *big.Float
//     - This value represents the ceiling value associated with input
//       parameter 'bigFloatNum'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//       Input Parameter       Return Parameter
//         bigFloatNum             ceiling
//       ---------------       ----------------
//          2.3                      3.0
//         -2.3                     -2.0
//          7.5555555555             8.0
//         -7.5555555555            -7.0
//
func (bigFloatNanobot *mathBigFloatNanobot) ceiling(
	bigFloatNum *big.Float,
	precision uint) (ceiling *big.Float) {

	if bigFloatNanobot.lock == nil {
		bigFloatNanobot.lock = new(sync.Mutex)
	}

	bigFloatNanobot.lock.Lock()

	defer bigFloatNanobot.lock.Unlock()

	ceiling =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return ceiling
	}

	newBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)


	if newBigFloat.IsInt() {
		ceiling =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(newBigFloat)

		ceiling = ceiling.SetMode(ceiling.Mode())

		return ceiling
	}

	var newInterimFloat *big.Float

	if newBigFloat.Sign() == -1 {
		// Input big float value is negative
		newInterimFloat =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(newBigFloat)

	} else {
		// newBigFloat must be positive

		bigFloat1 :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				SetInt(big.NewInt(1))

		newInterimFloat =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Add(newBigFloat, bigFloat1)
	}

	var intCeiling *big.Int

	intCeiling, _ = newInterimFloat.Int(nil)

	ceiling =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(intCeiling)

	ceiling = ceiling.SetMode(ceiling.Mode())

	return ceiling
}

// floatFractionalValue - Returns the absolute value of the
// fraction contained in a type *big.Float, floating point
// number. Again, the returned fractional number is always
// positive. In addition, an integer is returned signaling
// the numeric sign of the original floating point number,
// 'bigFloatNum'.
//
// If the input parameter, 'bigFloatNum' is equal to zero,
// the return parameter, 'numSign' is set to zero. If 'bigFloatNum'
// is less than zero, 'numSign' is set to -1. And, if 'bigFloatNum'
// is positive, the returned 'numSign' is set to +1.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum           *big.Float
//     - This method will calculate and return the 'fractional' component
//       of this parameter.
//
//
//  precision             uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on input parameter 'bigFloatNum'
//       as well as specifying the numeric precision incorporated in the
//       returned floating point value, 'floatFractionalValue'. The
//       'precision' value applies to the internal accuracy maintained
//        by type *big.Float floating point values.  For more information
//        on precision and type *big.Float floating point numbers,
//        reference:
//           https://golang.org/pkg/math/big/
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  floatFractionalValue  *big.Float
//     - This value represents the fractional component of input parameter
//       'bigFloatNum'.
//
//
//  numSign               int
//     - This returned parameter will be set to one of three values depending
//       on the numeric sign presented by return parameter 'floatFractionalValue'
//         floatFractionalValue     numSign Value
//             > 0                        1
//             0.0                        0
//             < 0                       -1
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//      Input                  Return             Return
//    Parameter               Parameter          Parameter
//                                                numSign
//   bigFloatNum         floatFractionalValue      Value
//   -----------         --------------------     -------
//   7853.123456               0.123456             +1
//      0.000000               0.000000              0
//  -7853.123456               0.123456             -1
//
func (bigFloatNanobot *mathBigFloatNanobot) floatFractionalValue(
	bigFloatNum *big.Float,
	precision uint) (
	floatFractionalValue *big.Float,
	numSign int) {

	if bigFloatNanobot.lock == nil {
		bigFloatNanobot.lock = new(sync.Mutex)
	}

	bigFloatNanobot.lock.Lock()

	defer bigFloatNanobot.lock.Unlock()

	numSign = 0

	floatFractionalValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return floatFractionalValue, numSign
	}

	var newBigFloatNum *big.Float

	if bigFloatNum.Sign() == -1 {
		numSign = -1
		newBigFloatNum =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Neg(bigFloatNum)

	} else {
		// newBigFloatNum.Sign() == +1

		numSign = 1

		newBigFloatNum =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(bigFloatNum)

	}

	bigIntVal, _ := newBigFloatNum.Int(nil)

	bigIntFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(bigIntVal)

	floatFractionalValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Sub(newBigFloatNum, bigIntFloat)

	floatFractionalValue = floatFractionalValue.SetMode(floatFractionalValue.Mode())

	return floatFractionalValue, numSign
}

// floatIntFracValues - Returns both the integer and fractional components
// as *big.Float floating point numbers. These are returned as two separate
// *big.Float values.
//
// The numerical sign of input parameter 'bigFloatNum' will always be
// preserved in the integer component return parameter, 'floatIntegerValue'.
// If input parameter 'bigFloatNum' is a negative value, the returned
// 'floatIntegerValue' will be a negative value. Likewise, if 'bigFloatNum'
// is a positive value, the returned 'floatIntegerValue' will be a positive
// value.
//
// The returned fractional component, 'floatFractionalValue', will always
// be returned as an absolute or positive value.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum              *big.Float
//     - This method will calculate and return the integer and 'fractional'
//       components of this parameter.
//
//
//  precision                uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on input parameter 'bigFloatNum'
//       as well as specifying the numeric precision incorporated in the
//       returned floating point value, 'floatFractionalValue'. The
//       'precision' value applies to the internal accuracy maintained
//        by type *big.Float floating point values.  For more information
//        on precision and type *big.Float floating point numbers,
//        reference:
//           https://golang.org/pkg/math/big/
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  floatIntegerValue        *big.Float
//     - Contains the integer component of the input parameter 'bigFloatNum'
//       formatted as a type *big.Float. Note that the numerical sign of
//       input parameter 'bigFloatNum' will be preserved in this return
//       value. If 'bigFloatNum' is a negative value, the returned
//       'floatIntegerValue' will be a negative value. Likewise, if
//       'bigFloatNum' is a positive value, the returned 'floatIntegerValue'
//       will be a positive value.
//
//
//  floatFractionalValue     *big.Float
//     - Contains the fractional component of the input parameter
//       'bigFloatNum' formatted as a type *big.Float. This returned value
//       will always be a positive value. The numeric sign of input
//       parameter 'bigFloatNum' is preserved in return value
//       'floatIntegerValue'. This return parameter, 'floatFractionalValue'
//       will always be an absolute or positive value.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//        Input            Return               Return
//      Parameter         Parameter            Parameter
//     bigFloatNum    floatIntegerValue    floatFractionalValue
//     -----------    -----------------    --------------------
//     7853.123456         7853.0               0.123456
//        0.000000            0.0               0.000000
//    -7853.123456        -7853.0               0.123456
//
func (bigFloatNanobot *mathBigFloatNanobot) floatIntFracValues(
	bigFloatNum *big.Float,
	precision uint) (
	floatIntegerValue *big.Float,
	floatFractionalValue *big.Float) {

	if bigFloatNanobot.lock == nil {
		bigFloatNanobot.lock = new(sync.Mutex)
	}

	bigFloatNanobot.lock.Lock()

	defer bigFloatNanobot.lock.Unlock()

	floatIntegerValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	floatFractionalValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return floatIntegerValue, floatFractionalValue
	}

	var numSign int

	var newBigFloatNum *big.Float

	if bigFloatNum.Sign() == -1 {
		numSign = -1
		newBigFloatNum =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Neg(bigFloatNum)

	} else {
		// newBigFloatNum.Sign() == +1

		numSign = 1

		newBigFloatNum =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(bigFloatNum)

	}

	bigIntVal, _ := newBigFloatNum.Int(nil)

	floatIntegerValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(bigIntVal)

	floatFractionalValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Sub(newBigFloatNum, floatIntegerValue)

	if numSign == -1 {

		floatIntegerValue =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Neg(floatIntegerValue)

	}

	floatIntegerValue = floatIntegerValue.SetMode(floatIntegerValue.Mode())

	floatFractionalValue = floatFractionalValue.SetMode(floatFractionalValue.Mode())

	return floatIntegerValue, floatFractionalValue
}

// floor - Provides a standard 'floor' function for
// type *big.Float.
//
// For a discussion of the 'floor' function, reference:
//  https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum       *big.Float
//     - This method will calculate and return the 'floor' value
//       associated with this parameter.
//
//
//  precision         uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on input parameter 'bigFloatNum'
//       as well as specifying the numeric precision incorporated in the
//       returned floating point value, 'floor'.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  floor   *big.Float
//     - This value represents the floor value associated with input
//       parameter 'bigFloatNum'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//       Input Parameter       Return Parameter
//         bigFloatNum              floor
//       ---------------       ----------------
//          2.3                      2.0
//         -2.3                     -3.0
//          7.5555555555             7.0
//         -7.5555555555            -8.0
//
func (bigFloatNanobot *mathBigFloatNanobot) floor(
	bigFloatNum *big.Float,
	precision uint) (floor *big.Float) {

	if bigFloatNanobot.lock == nil {
		bigFloatNanobot.lock = new(sync.Mutex)
	}

	bigFloatNanobot.lock.Lock()

	defer bigFloatNanobot.lock.Unlock()

	floor =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return floor
	}

	newBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)

	if newBigFloat.IsInt() {

		floor =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(newBigFloat)

		floor = floor.SetMode(floor.Mode())

		return floor
	}

	bigIntFloor, _ := newBigFloat.Int(nil)

	if newBigFloat.Sign() == 1 {
		// Numeric sign of  bigFloatNum is positive (+).

		floor =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				SetInt(bigIntFloor)

	} else {
		// newBigFloat.Sign() == -1
		// bigFloatNum is LESS THAN zero

		bigIntFloor =
			big.NewInt(0).
				Add(bigIntFloor,
					big.NewInt(-1))

		floor =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				SetInt(bigIntFloor)
	}

	floor = floor.SetMode(floor.Mode())

	return floor
}

// getIntegerLength - Receives a floating point number of type
// *big.Float and determines the length or number of digits in the
// integer component of the *big.Float floating point number.
//
// The return values will specify the number of digits in the integer
// portion of the input parameter floating point number as well as
// the numerical sign of said floating point value (+, -, or zero
// value).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum         *big.Float
//     - This method will calculate and return the length, or
//       number of digits, contained in the integer component
//       of this parameter.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  intLength           int
//     - This value represents the the integer length or the number of
//       digits in the input parameter 'bigFloatNum'. This return value
//       will always be greater than zero. At minimum, the integer
//       component of 'bigFloatNum' will have at least one numeric digit.
//       Note that this integer length return value will never include
//       any leading or non-numeric signs such as plus ('+'), minus ('-'),
//       or decimal ('.'). 'intLength' will only specify the number of
//       integer numeric digits in the 'bigFloatNum' floating point value.
//
//
//  numSign             int
//     - This returned integer will describe the numerical sign associated
//       with input parameter 'bigFloatNum'. It will be set to one of three
//       values:
//
//       +1 - Signals that 'bigFloatNum' is a positive value greater than
//            zero ('bigFloatNum' > 0 ).
//
//        0 - Signals that 'bigFloatNum' has a zero value.
//            ('bigFloatNum' == 0 ).
//
//       -1 - Signals that 'bigFloatNum' is a negative value less than zero
//            ('bigFloatNum' < 0 ).
//
func (bigFloatNanobot *mathBigFloatNanobot) getIntegerLength(
	bigFloatNum *big.Float) (
	intLength int,
	numSign int) {

	if bigFloatNanobot.lock == nil {
		bigFloatNanobot.lock = new(sync.Mutex)
	}

	bigFloatNanobot.lock.Lock()

	defer bigFloatNanobot.lock.Unlock()

	intLength = 0

	//	-1 if x <   0
	//	 0 if x is ±0
	//	+1 if x >   0

	numSign = bigFloatNum.Sign()

	var intValOfFloat *big.Int

	intValOfFloat, _ = bigFloatNum.Int(nil)

	intValStr := fmt.Sprintf("%v",
		intValOfFloat.Text(10))

	if intValStr[0] == '+' ||
		intValStr[0] == '-' {
		intValStr = intValStr[1:]
	}

	intLength = len(intValStr)

	return intLength, numSign
}

// roundHalfAwayFromZero - Performs a rounding operation on floating
// point numbers of type *big.Float. The calling function specifies the number
// of digits to the right of the decimal point which will be contained the
// returned, 'rounded value'.
//
// The rounding algorithm applied by this method is known as, 'Round half away
// from zero'.
//
// For a discussion of rounding algorithms reference:
//   https://en.wikipedia.org/wiki/Rounding
//
// For more information on type *big.Float and its associated internal
// 'precision', reference:
//   https://golang.org/pkg/math/big/
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum       *big.Float
//     - This method will calculate and return the rounded value
//       of this parameter.
//
//
//  precision         uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on input parameter 'bigFloatNum'
//       as well as specifying the numeric precision incorporated in the
//       returned floating point value, 'roundedFloat'. 'precision' should
//       not be confused with parameter 'roundToDecPlaces'. The term 'precision'
//       applies to the internal accuracy maintained by type *big.Float
//       floating point values.  For more information on precision and
//       type *big.Float floating point numbers, reference:
//           https://golang.org/pkg/math/big/
//
//
//  roundToDecPlaces  uint
//     - This parameter specifies the number of digits to the right of the
//       decimal place which will be contained in the returned value,
//       'roundedFloat'.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  roundedFloat      *big.Float
//     - This value represents the rounded value of input parameter
//       'bigFloatNum'. It will contain the number of digits to the
//       right of the decimal point specified by input parameter,
//       'roundToDecPlaces'
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  bigFloatNum      roundToDecPlaces       roundedFloat
//   7853.1234567          5                7853.12346
//   7853.1234567          3                7853.123
//  -7853.1234567          5               -7853.12346
//  -7853.1234567          3               -7853.123
//   7853.5234567          0                7854.0
//   7853.4234567          0                7853.0
//  -7853.5234567          0               -7854.0
//  -7853.4234567          0               -7853.0
//
//
func (bigFloatNanobot *mathBigFloatNanobot) roundHalfAwayFromZero(
	bigFloatNum *big.Float,
	precision uint,
	roundToDecPlaces uint) (roundedFloat *big.Float) {

	if bigFloatNanobot.lock == nil {
		bigFloatNanobot.lock = new(sync.Mutex)
	}

	bigFloatNanobot.lock.Lock()

	defer bigFloatNanobot.lock.Unlock()

	if roundToDecPlaces > precision {
		precision = roundToDecPlaces + 100
	}

	roundedFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	roundedFloat = roundedFloat.SetMode(roundedFloat.Mode())

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return roundedFloat
	}

	newBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)

	var bigInt5 *big.Int

	roundedValue :=
		big.NewInt(0).
			Exp(big.NewInt(10),
				big.NewInt(int64(roundToDecPlaces+1)), nil)

	if newBigFloat.Sign() == -1 {

		bigInt5 = big.NewInt(-5)

	} else {
		bigInt5 = big.NewInt(5)
	}

	ratRound := big.NewRat(1,1).
		SetFrac(bigInt5, roundedValue)

	fracRound :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetRat(ratRound)

	newNumFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Add(newBigFloat, fracRound)

	roundedValue =
		big.NewInt(0).
			Exp(big.NewInt(10),
				big.NewInt(int64(roundToDecPlaces)), nil)

	roundFrac :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(roundedValue)

	newIntNumFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Mul(newNumFloat, roundFrac)

	newIntNum, _ :=
		newIntNumFloat.Int(nil)

	ratResult :=
		big.NewRat(1,1).
			SetFrac(newIntNum, roundedValue)

	roundedFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetRat(ratResult)

	roundedFloat = roundedFloat.SetMode(roundedFloat.Mode())

	return roundedFloat
}


// truncate - Receives a type *big.Float floating point number and
// proceeds to truncate that number to a specified number of digits
// after the decimal point. No rounding is performed.
//
// The calling function specifies the number of digits to the right
// of the decimal point which will be contained the returned truncated
// value.
//
// The calling also specifies the internal 'precision' to be applied
// to all internal calculations. For more information on type *big.Float
// and its associated internal precision values, reference:
//   https://golang.org/pkg/math/big/
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum          *big.Float
//     - This method will return a truncated version of this initial
//       value.
//
//
//  precision            uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on input parameter 'bigFloatNum'
//       as well as specifying the numeric precision incorporated in the
//       returned floating point value, 'truncatedFloat'. 'precision' should
//       not be confused with parameter 'truncateToDecPlaces'. The term
//       'precision' applies to the internal accuracy maintained by type
//       *big.Float floating point values.  For more information on precision
//       and type *big.Float floating point numbers, reference:
//           https://golang.org/pkg/math/big/
//
//
//  truncateToDecPlaces  uint
//     - This parameter specifies the number of digits to the right of the
//       decimal point which will be contained in the returned value,
//       'truncatedFloat'.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  truncatedFloat       *big.Float
//     - This value represents the truncate value of input parameter
//       'bigFloatNum'. It will contain the number of digits to the
//       right of the decimal point specified by input parameter,
//       'truncateToDecPlaces'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  bigFloatNum      truncateToDecPlaces   truncatedFloat
//   7853.1234567          5                7853.12345
//   7853.1234567          3                7853.123
//  -7853.1234567          5               -7853.12345
//  -7853.1234567          3               -7853.123
//   7853.5234567          0                7853.0
//   7853.4234567          0                7853.0
//  -7853.5234567          0               -7853.0
//  -7853.4234567          0               -7853.0
//
//
func (bigFloatNanobot *mathBigFloatNanobot) truncate(
	bigFloatNum *big.Float,
	precision uint,
	truncateToDecPlaces uint) (truncatedFloat *big.Float) {

	if bigFloatNanobot.lock == nil {
		bigFloatNanobot.lock = new(sync.Mutex)
	}

	bigFloatNanobot.lock.Lock()

	defer bigFloatNanobot.lock.Unlock()

	if truncateToDecPlaces > precision {
		precision = truncateToDecPlaces + 100
	}

	truncatedFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	truncatedFloat = truncatedFloat.SetMode(truncatedFloat.Mode())

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return truncatedFloat
	}

	newBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)

	truncateScale :=
		big.NewInt(0).
			Exp(big.NewInt(10),
				big.NewInt(int64(truncateToDecPlaces)), nil)

	truncateScaleFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(truncateScale)

	newIntBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Mul(newBigFloat, truncateScaleFloat)

	newInt, _ := newIntBigFloat.Int(nil)

	ratResult :=
		big.NewRat(1,1).
			SetFrac(newInt, truncateScale)

	truncatedFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetRat(ratResult)

	truncatedFloat = truncatedFloat.SetMode(truncatedFloat.Mode())

	return truncatedFloat
}