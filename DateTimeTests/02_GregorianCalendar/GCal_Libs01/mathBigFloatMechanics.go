package GCal_Libs01

import (
	"fmt"
	"math/big"
	"sync"
)

// mathBigFloatMechanics - Provides helper methods for
// arithmetic operations.
//
//
// Rounding Mode Notes:
//   https://golang.org/pkg/math/big/#RoundingMode
//
//     type RoundingMode byte
//
//      These constants define supported rounding modes.
//     const (
//     	ToNearestEven RoundingMode = iota // == IEEE 754-2008 roundTiesToEven
//     	ToNearestAway                     // == IEEE 754-2008 roundTiesToAway
//     	ToZero                            // == IEEE 754-2008 roundTowardZero
//     	AwayFromZero                      // no IEEE 754-2008 equivalent
//     	ToNegativeInf                     // == IEEE 754-2008 roundTowardNegative
//     	ToPositiveInf                     // == IEEE 754-2008 roundTowardPositive
//     )
//
//     x         ToNearestEven  ToNearestAway  ToZero  AwayFromZero  ToNegativeInf  ToPositiveInf
//    2.6              3              3          2           3            2              3
//    2.5              2              3          2           3            2              3
//    2.1              2              2          2           3            2              3
//   -2.1             -2             -2         -2           3            3              2
//   -2.5             -2             -3         -2           3            3              2
//   -2.6             -3             -3         -2           3            3              2
//
type mathBigFloatMechanics struct {
	lock *sync.Mutex
}

// FloatFractionalValue - Returns the absolute value of the
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
// For a discussion of the 'ceiling' function, reference:
//  https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
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
// ------------------------------------------------------------------------
//
// Example Usage
//  bigFloatNum          floatFractionalValue  numSign
//   7853.123456               0.123456          1
//      0.000000               0.000000          0
//  -7853.123456               0.123456         -1
//
//
func (mathBFMech *mathBigFloatMechanics) FloatFractionalValue(
	bigFloatNum *big.Float,
	precision uint) (floatFractionalValue *big.Float, numSign int) {

	if mathBFMech.lock == nil {
		mathBFMech.lock = new(sync.Mutex)
	}

	mathBFMech.lock.Lock()

	defer mathBFMech.lock.Unlock()

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

// FloatIntFracValues - Returns both the integer and fractional components
// of a type *big.Float floating point number as two separate *big.Float
// values.
//
//
func (mathBFMech *mathBigFloatMechanics) FloatIntFracValues(
	bigFloatNum *big.Float,
	precision uint) (
	floatIntegerValue *big.Float,
	floatFractionalValue *big.Float) {

	if mathBFMech.lock == nil {
		mathBFMech.lock = new(sync.Mutex)
	}

	mathBFMech.lock.Lock()

	defer mathBFMech.lock.Unlock()

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


	// FloatIntegerValue - Returns the integer value of
// a type *big.Float as another *big.Float.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum       *big.Float
//     - This method will extract the integer component from this
//       *big.Float input value.
//
//
//  precision         uint
//     - This unsigned integer value will determine the precision of internal
//       calculations performed on input parameter 'bigFloatNum' as well as
//       specifying the precision incorporated in the return integer value
//       expressed as a type *big.Float, floating point number.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  floatIntegerValue   *big.Float
//     - This value represents the integer component of input parameter
//       'bigFloatNum' expressed as a floating point number.
//
func (mathBFMech *mathBigFloatMechanics) FloatIntegerValue(
	bigFloatNum *big.Float,
	precision uint) (floatIntegerValue *big.Float) {

	if mathBFMech.lock == nil {
		mathBFMech.lock = new(sync.Mutex)
	}

	mathBFMech.lock.Lock()

	defer mathBFMech.lock.Unlock()

	floatIntegerValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return floatIntegerValue
	}

	var newBigFloat *big.Float

	newBigFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)

	if bigFloatNum.IsInt() {
		floatIntegerValue =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(newBigFloat)

		return floatIntegerValue
	}

	var intValOfFloat *big.Int

	intValOfFloat, _ = bigFloatNum.Int(nil)

	floatIntegerValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(intValOfFloat)

	floatIntegerValue = floatIntegerValue.SetMode(floatIntegerValue.Mode())

	return floatIntegerValue
}

// Ceiling - Provides a standard 'ceiling' function for
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
//  bigFloatNum         ceiling
//   2.3                  3
//  -2.3                 -2
//   7.5555555555         8
//  -7.5555555555         7
//
func (mathBFMech *mathBigFloatMechanics) Ceiling(
	bigFloatNum *big.Float,
	precision uint) (ceiling *big.Float) {

	if mathBFMech.lock == nil {
		mathBFMech.lock = new(sync.Mutex)
	}

	mathBFMech.lock.Lock()

	defer mathBFMech.lock.Unlock()

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

// Floor - Provides a standard 'floor' function for
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
func (mathBFMech *mathBigFloatMechanics) Floor(
	bigFloatNum *big.Float,
	precision uint) (floor *big.Float) {

	if mathBFMech.lock == nil {
		mathBFMech.lock = new(sync.Mutex)
	}

	mathBFMech.lock.Lock()

	defer mathBFMech.lock.Unlock()

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

// GetIntegerLength - Receives a floating point number of
// type *big.Float and returns both the length of the integer digits
// contained therein and the leading number sign prefixed to those
// integer digits if such number sign exists.
//
// If a number sign is present, the value of the returned 'leading
// number sign' string will be either plus ('+') or minus ('-'). If
// no leading number sign is found (such as in the case of positive
// values where the plus sign is inferred), the returned leading
// number sign is an empty string.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum       *big.Float
//     - This method will calculate and return the length, or
//       number of digits, contained in the integer component
//       of this parameter.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  intLength    int
//     - This value represents the floor value associated with input
//       parameter 'bigFloatNum'.
//
//
//  leadingSign  string
//     - This returned string contains the numeric sign (plus '+' or
//       minus '-') associated with the input parameter, 'bigFloatNum'.
//       With positive values, the plus sign is often implied and
//       therefore it is not present in the numeric value. In this
//       case, 'leadingSign' is an empty string. For negative values,
//       the leading minus sign is always present therefore it is
//       returned in 'leadingSign' as the the string value "-".
//
func (mathBFMech *mathBigFloatMechanics) GetIntegerLength(
	bigFloatNum *big.Float) (
	intLength int,
	leadingSign string)  {

	if mathBFMech.lock == nil {
		mathBFMech.lock = new(sync.Mutex)
	}

	mathBFMech.lock.Lock()

	defer mathBFMech.lock.Unlock()

	intLength = 0
	leadingSign = ""

	if bigFloatNum == nil {
		return intLength, leadingSign
	}

	var intValOfFloat *big.Int

	intValOfFloat, _ = bigFloatNum.Int(nil)

	intValStr := fmt.Sprintf("%v",
		intValOfFloat.Text(10))

	if intValStr[0] == '+' ||
		intValStr[0] == '-' {
		leadingSign = intValStr[0:1]
	}

	return len(intValStr), leadingSign
}

// GetBigIntValue - Returns the integer portion of a type
// *big.Float floating point number. The returned integer value
// is of type *big.Int.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigFloatNum       *big.Float
//     - This method will calculate and return the integer value
//       component extracted from this floating point number.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  intValOfFloat     *big.Int
//     - This value represents the integer value of input
//       parameter 'bigFloatNum'.
//
func (mathBFMech *mathBigFloatMechanics) GetBigIntValue(
	bigFloatNum *big.Float) (intValOfFloat *big.Int) {

	if mathBFMech.lock == nil {
		mathBFMech.lock = new(sync.Mutex)
	}

	mathBFMech.lock.Lock()

	defer mathBFMech.lock.Unlock()

	intValOfFloat =	big.NewInt(0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return intValOfFloat
	}

	intValOfFloat, _ = bigFloatNum.Int(nil)

	return intValOfFloat
}

// RoundHalfAwayFromZero - Performs a rounding operation on floating
// point numbers of type *big.Float. The calling function specifies the number
// of digits to the right of the decimal point which will be contained the
// returned, 'rounded value'.
//
// The rounding algorithm applied by this method is named, 'Round half away
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
func (mathBFMech *mathBigFloatMechanics) RoundHalfAwayFromZero(
	bigFloatNum *big.Float,
	precision uint,
	roundToDecPlaces uint) (roundedFloat *big.Float) {

	if mathBFMech.lock == nil {
		mathBFMech.lock = new(sync.Mutex)
	}

	mathBFMech.lock.Lock()

	defer mathBFMech.lock.Unlock()

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

	roundValue :=
		big.NewInt(0).
			Exp(big.NewInt(10),
				big.NewInt(int64(roundToDecPlaces+1)), nil)

	if newBigFloat.Sign() == -1 {

		bigInt5 = big.NewInt(-5)

	} else {
		bigInt5 = big.NewInt(5)
	}

	ratRound := big.NewRat(1,1).
		SetFrac(bigInt5, roundValue)

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

	roundValue =
		big.NewInt(0).
			Exp(big.NewInt(10),
				big.NewInt(int64(roundToDecPlaces)), nil)

	roundFrac :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(roundValue)

	newIntNumFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Mul(newNumFloat, roundFrac)

	newIntNum, _ :=
		newIntNumFloat.Int(nil)

	ratResult :=
		big.NewRat(1,1).
			SetFrac(newIntNum, roundValue)

	roundedFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetRat(ratResult)

	roundedFloat = roundedFloat.SetMode(roundedFloat.Mode())

	return roundedFloat
}


// Truncate - Receives a type *big.Float floating point  number and
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
func (mathBFMech *mathBigFloatMechanics) Truncate(
	bigFloatNum *big.Float,
	precision uint,
	truncateToDecPlaces uint) (truncatedFloat *big.Float) {

	if mathBFMech.lock == nil {
		mathBFMech.lock = new(sync.Mutex)
	}

	mathBFMech.lock.Lock()

	defer mathBFMech.lock.Unlock()


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