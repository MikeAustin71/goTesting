package datetime

import (
	"math/big"
	"sync"
)

// MathBigFloatHelper - Provides helper methods for
// arithmetic operations using type *big.Float. For more
// information on type *big.Float reference:
//   https://golang.org/pkg/math/big/
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
type MathBigFloatHelper struct {
	lock *sync.Mutex
}

// Abs - Returns the absolute value of a *big.Float floating
// point value.
//
// The absolute value of a number is always a positive value.
//
// "In mathematics, the absolute value or modulus of a real number
// x, denoted |x|, is the non-negative value of x without regard
// to its sign." Wikipedia:
//   https://en.wikipedia.org/wiki/Absolute_value
//
// In addition to the absolute value, an integer is returned specifying
// the numeric sign of the original value before conversion to
// the corresponding absolute value. The returned value is either
// one of two values:
//
// 1. If the original value is is greater than or equal to zero,
//    the returned 'originalNumSign' is set to plus one ('+1').
//
// 2. If the original value is is less than zero, the returned
//    'originalNumSign' is set to minus one ('-1').
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  bigFloatNum           *big.Int
//     - This method will convert this floating number to its
//       corresponding absolute value.
//
//
//  precision             uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on the input parameters as
//       well as specifying the numeric precision incorporated in the
//       returned floating point value, 'absValue'. The 'precision'
//       value applies to the internal accuracy maintained by type
//       *big.Float floating point values. For those seeking to
//       maximize accuracy. Try a value of '1024'. For more information
//       on precision and type *big.Float floating point numbers,
//       reference:
//           https://golang.org/pkg/math/big/
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  absValue              *big.Float
//     - Contains the absolute value of input parameter 'bigFloatNum'.
//       Absolute values are always positive.
//
//
//  originalNumSign       int
//     - Returns the number sign of the original value, 'bigFloatNum'.
//       The returned 'originalNumSign' parameter will be set to one
//       of two values depending on value of input parameter
//       'bigFloatNum':
//
//         bigFloatNum             originalNumSign
//             > 0                        1
//             = 0                        1
//             < 0                       -1
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//    bigFloatNum        absValue       originalNumSign
//   7853.123456      7853.123456             1
//      0.123456         0.123456             1
//      0.000000         0.000000             1
//     -0.123456         0.123456            -1
//  -7853.123456      7853.123456            -1
//
func (mathBFloatHlpr *MathBigFloatHelper) Abs(
	bigFloatNum *big.Float,
	precision uint) (
	absValue *big.Float,
	originalNumSign int) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	originalNumSign = 1

	absValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)


	if bigFloatNum == nil {

		return absValue, originalNumSign
	}

	 originalNumSign = bigFloatNum.Sign()

	 if originalNumSign == 0 {
	 	originalNumSign = 1
		 return absValue, originalNumSign
	 }

	 if originalNumSign == -1 {
		absValue =
			big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
				SetPrec(precision).
				Neg(bigFloatNum)
	 } else {
		 absValue =
			 big.NewFloat(0.0).
				 SetMode(big.ToNearestAway).
				 SetPrec(precision).
				 Set(bigFloatNum)
	 }

	absValue = absValue.SetPrec(absValue.Prec())

	return absValue, originalNumSign
}

// ChangeSign - Changes the numeric sign of a floating
// point number. If the floating point number passed to
// this method is negative, the corresponding positive
// value is returned. Likewise positive values are converted
// and returned as negative values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  bigFloatNum           *big.Int
//     - This method will change the numeric sign of this value and
//       return the altered value. Positive values are converted to
//       negative values and negative values are converted to positive
//       values
//
//
//  precision             uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on the input parameter as
//       well as specifying the numeric precision incorporated in the
//       returned floating point value, 'resultFloatNum'. The 'precision'
//       value applies to the internal accuracy maintained by type
//       *big.Float floating point values.  For more information on
//       precision and type *big.Float floating point numbers,
//       reference:
//           https://golang.org/pkg/math/big/
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  resultFloatNum        *big.Float
//     - Contains the result of converting the numeric sign of input
//       parameter, 'bigFloatNum'.
//
//
//  originalNumSign       int
//     - Returns the number sign of the original value, 'bigFloatNum'.
//       The returned 'originalNumSign' parameter will be set to one
//       of three values depending on value of input parameter
//       'bigFloatNum':
//
//         bigFloatNum            originalNumSign
//             > 0                        1
//             0.0                        0
//             < 0                       -1
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//    bigFloatNum    resultFloatNum     originalNumSign
//   7853.123456      -7853.123456            1
//      0.123456        -0.123456             1
//      0.000000         0.000000             0
//     -0.123456         0.123456            -1
//  -7853.123456      7853.123456            -1
//
func (mathBFloatHlpr *MathBigFloatHelper) ChangeSign(
	bigFloatNum *big.Float,
	precision uint) (
	resultFloatNum *big.Float,
	originalNumSign int) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	originalNumSign = 0

	resultFloatNum =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil {
		return resultFloatNum, originalNumSign
	}

	originalNumSign = bigFloatNum.Sign()

	if originalNumSign == 0 {
		return resultFloatNum, originalNumSign
	}

	resultFloatNum =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Neg(bigFloatNum)


	resultFloatNum =
		resultFloatNum.SetPrec(resultFloatNum.Prec())

	return resultFloatNum, originalNumSign
}

// CombineIntFracValues - Receives a *big.Int integer value and
// a *big.Float floating point value. The method then proceeds
// to combine the two values together creating and returning a
// *big.Float floating point result.
//
// The numeric sign (plus '+' or minus '-') of the returned floating
// point value is determined by the integer input parameter,
// 'numSign'.  If 'numSign' is equal to '-1', the returned floating
// point value is formatted as a negative floating point value.
// Otherwise, the returned value is positive.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  bigIntegerValue       *big.Int
//     - An integer value which will be added to parameter
//       'bigFloatFracValue' to create and return a floating
//       point result value.
//
//
//  bigFloatFracValue     *big.Float
//     - This floating point value contains the fractional
//       portion of a final result which will be produced by
//       adding this value to parameter 'bigIntegerValue'.
//
//
//  numSign               int
//     - Determines the numeric sign of the returned floating
//       point value 'combinedFloatValue'. If this value is set
//       to '-1', the returned floating point value will be negative.
//       otherwise, the returned value will be positive.
//
//       Valid values for 'numSign' are '+1' and '-1'. Invalid
//       values will be converted to '+1'.
//
//
//  precision             uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on the input parameters as
//       well as specifying the numeric precision incorporated in the
//       returned floating point value, 'combinedFloatValue'. The
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
//  combinedFloatValue    *big.Float
//     - Contains the combined integer and fractional values supplied by
//       input parameters 'bigIntegerValue' and 'bigFloatFracValue'. This
//       value is computed by adding 'bigIntegerValue' and 'bigFloatFracValue'.
//       The numeric sign of this returned value is dependent on input
//       parameter, 'numSign'.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//  bigIntegerValue  bigFloatFracValue    numSign      combinedFloatValue
//         1             0.123456             1             1.123456
//         3             0.000000             1             3.000000
//     -7853           0.123456               1          7853.123456
//     -7853          -0.123456               1          7853.123456
//     -7853           0.123456              -1         -7853.876544
//      7853          -0.123456              -1         -7853.876544
//         0           0.123456              -1            -0.123456
//
func (mathBFloatHlpr *MathBigFloatHelper) CombineIntFracValues(
	bigIntegerValue *big.Int,
	bigFloatFracValue *big.Float,
	numSign int,
	precision uint) (
	combinedFloatValue *big.Float) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	combinedFloatValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)


	if bigIntegerValue == nil {
		return combinedFloatValue
	}

	if bigFloatFracValue == nil  {
		return combinedFloatValue
	}

	if numSign != 1 && numSign != -1 {
		numSign = 1
	}

	paramSignVal := bigIntegerValue.Sign()

	if paramSignVal == 0 {
		paramSignVal = 1
	}

	if paramSignVal != numSign {
		bigIntegerValue =
			big.NewInt(0).
				Neg(bigIntegerValue)
	}

	bigIntFloatVal :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(bigIntegerValue)


	paramSignVal = bigFloatFracValue.Sign()

	if paramSignVal == 0 {
		paramSignVal = 1
	}

	var newFloatFracVal *big.Float

	if paramSignVal != numSign {

		newFloatFracVal =
			big.NewFloat(0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Neg(bigFloatFracValue)

	} else {

		newFloatFracVal =
			big.NewFloat(0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(bigFloatFracValue)

	}


	combinedFloatValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Add(bigIntFloatVal,
			newFloatFracVal)

	return combinedFloatValue
}


// FloatFractionalValue - Returns the absolute value of the
// fractional component contained in a type *big.Float,
// floating point number. Again, the returned fractional
// number is always positive. In addition, an integer is
// returned signaling the numeric sign of the original
// floating point number, 'bigFloatNum'.
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
func (mathBFloatHlpr *MathBigFloatHelper) FloatFractionalValue(
	bigFloatNum *big.Float,
	precision uint) (
	floatFractionalValue *big.Float,
	numSign int) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	bigFloatNanobot := mathBigFloatNanobot{}

	floatFractionalValue,
	numSign = bigFloatNanobot.floatFractionalValue(
		bigFloatNum,
		precision)


	return floatFractionalValue, numSign
}

// FloatIntFracValues - Returns both the integer and fractional components
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
func (mathBFloatHlpr *MathBigFloatHelper) FloatIntFracValues(
	bigFloatNum *big.Float,
	precision uint) (
	floatIntegerValue *big.Float,
	floatFractionalValue *big.Float) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	bigFloatNanobot := mathBigFloatNanobot{}

	floatIntegerValue,
	floatFractionalValue = bigFloatNanobot.floatIntFracValues(
		bigFloatNum,
		precision)

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
func (mathBFloatHlpr *MathBigFloatHelper) FloatIntegerValue(
	bigFloatNum *big.Float,
	precision uint) (floatIntegerValue *big.Float) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

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
//       Input Parameter       Return Parameter
//         bigFloatNum             ceiling
//       ---------------       ----------------
//          2.3                      3.0
//         -2.3                     -2.0
//          7.5555555555             8.0
//         -7.5555555555             7.0
//
func (mathBFloatHlpr *MathBigFloatHelper) Ceiling(
	bigFloatNum *big.Float,
	precision uint) (ceiling *big.Float) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	bigFloatNanobot := mathBigFloatNanobot{}

	ceiling = bigFloatNanobot.ceiling(
		bigFloatNum,
		precision)

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
func (mathBFloatHlpr *MathBigFloatHelper) Floor(
	bigFloatNum *big.Float,
	precision uint) (floor *big.Float) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	bigFloatNanobot := mathBigFloatNanobot{}

	floor = bigFloatNanobot.floor(
		bigFloatNum,
		precision)

	return floor
}

// GetIntegerLength - Receives a floating point number of type
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
//`
//        0 - Signals that 'bigFloatNum' has a zero value.
//            ('bigFloatNum' == 0 ).
//
//       -1 - Signals that 'bigFloatNum' is a negative value less than zero
//            ('bigFloatNum' < 0 ).
//
func (mathBFloatHlpr *MathBigFloatHelper) GetIntegerLength(
	bigFloatNum *big.Float) (
	intLength int,
	numSign int) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	bigFloatNanobot := mathBigFloatNanobot{}

	intLength,
	numSign =
		bigFloatNanobot.getIntegerLength(
		bigFloatNum)

	return intLength, numSign
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
func (mathBFloatHlpr *MathBigFloatHelper) GetBigIntValue(
	bigFloatNum *big.Float) (intValOfFloat *big.Int) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	intValOfFloat =	big.NewInt(0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return intValOfFloat
	}

	intValOfFloat, _ = bigFloatNum.Int(nil)

	return intValOfFloat
}

// GetNumericSign - Returns the numeric sign of the floating point
// value passed to this method.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  bigFloatNum           *big.Int
//     - This floating point value will be analysed to determine its
//       numeric sign. The result will be returned as an integer value,
//       'numSign'
//       corresponding absolute value.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  numSign               int
//     - Returns the number sign of the original value, 'bigFloatNum'.
//       The returned 'originalNumSign' parameter will be set to one
//       of three values depending on value of input parameter
//       'bigFloatNum':
//
//         bigFloatNum                 numSign
//             > 0                        1
//             0.0                        0
//             < 0                       -1
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//    bigFloatNum           numSign
//   7853.123456               1
//      0.123456               1
//      0.000000               0
//     -0.123456              -1
//  -7853.123456              -1
//
func (mathBFloatHlpr *MathBigFloatHelper) GetNumericSign(
	bigFloatNum *big.Float) (numSign int) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	numSign = 0

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return numSign
	}

	numSign = bigFloatNum.Sign()

	return numSign
}

// GetNumericSignAsPlusMinus - Instances of type *big.Float typically return
// an integer numeric sign which is one of three values: '+1', '0' or '-1'.
// Often, a binary designation is required instead of this trio of values.
//
// This method will examine the *big.Float input parameter, 'bigFloatNum' and
// return an integer set to one of two values: '+1' or '-1'.
//
// If 'bigFloatNum' is less than zero, a 'numSign' value of '-1' is returned.
//
// If 'bigFloatNum' is equal to or greater than zero, a 'numSign' value of
// '+1' is returned.
//
func (mathBFloatHlpr *MathBigFloatHelper) GetNumericSignAsPlusMinus(
	bigFloatNum *big.Float) (numSign int) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	numSign = 1

	if bigFloatNum == nil {
		return numSign
	}

	numSign = bigFloatNum.Sign()

	if numSign == 0 {
		numSign = 1
	}

	return numSign
}

// GetFloatNumText
func (mathBFloatHlpr *MathBigFloatHelper) GetFloatNumText(
	bigFloatNum *big.Float,
	numberFieldLength int,
	precision uint,
	roundToDecPlaces uint,
	ePrefix string) (
	bigFloatTextDto MathFloatTextDto,
	err error) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	ePrefix += "MathBigFloatHelper.GetFloatNumText() "

	bigFloatUtil := mathBigFloatUtility{}

	bigFloatTextDto,
	err = bigFloatUtil.getFloatNumText(
		bigFloatNum,
		numberFieldLength,
		precision,
		roundToDecPlaces,
		ePrefix)

	return bigFloatTextDto, err
}

// RoundHalfAwayFromZero - Performs a rounding operation on floating
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
func (mathBFloatHlpr *MathBigFloatHelper) RoundHalfAwayFromZero(
	bigFloatNum *big.Float,
	precision uint,
	roundToDecPlaces uint) (roundedFloat *big.Float) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	bigFloatNanobot := mathBigFloatNanobot{}

	roundedFloat =
		bigFloatNanobot.roundHalfAwayFromZero(
		bigFloatNum,
		precision,
		roundToDecPlaces)

	return roundedFloat
}


// Truncate - Receives a type *big.Float floating point number and
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
func (mathBFloatHlpr *MathBigFloatHelper) Truncate(
	bigFloatNum *big.Float,
	precision uint,
	truncateToDecPlaces uint) (truncatedFloat *big.Float) {

	if mathBFloatHlpr.lock == nil {
		mathBFloatHlpr.lock = new(sync.Mutex)
	}

	mathBFloatHlpr.lock.Lock()

	defer mathBFloatHlpr.lock.Unlock()

	bigFloatNanobot := mathBigFloatNanobot{}

	truncatedFloat =
		bigFloatNanobot.truncate(
			bigFloatNum,
			precision,
			truncateToDecPlaces)

	return truncatedFloat
}