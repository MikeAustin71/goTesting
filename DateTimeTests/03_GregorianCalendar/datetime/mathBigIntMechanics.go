package datetime

import (
	"math/big"
	"sync"
)

// MathBigIntHelper - Provides helper methods for
// arithmetic operations using the integer type *big.Int.
// For more information on type *big.Int reference:
//   https://golang.org/pkg/math/big/
//
//
type MathBigIntHelper struct {
	lock *sync.Mutex
}

// Abs - Returns the absolute value of a *big.Int integer value.
// In addition, another integer is returned specifying the numeric
// sign of the original value before it was converted to the
// corresponding absolute value.
//
// The absolute value of a number is always yields a positive value.
//
// "In mathematics, the absolute value or modulus of a real number
// x, denoted |x|, is the non-negative value of x without regard
// to its sign." Wikipedia:
//   https://en.wikipedia.org/wiki/Absolute_value
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  bigIntNum             *big.Int
//     - This  integer number will be converted to its corresponding
//       absolute value. The absolute value of a number is always
//       positive.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  absValue              *big.Int
//     - Contains the absolute value of input parameter 'bigIntNum'.
//       Absolute values are always positive.
//
//
//  originalNumSign       int
//     - Returns the number sign of the original value, 'bigIntNum'.
//       The returned 'originalNumSign' parameter will be set to one
//       of three values depending on value of input parameter
//       'bigIntNum':
//
//          bigIntNum             originalNumSign
//             > 0                        1
//             0.0                        0
//             < 0                       -1
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//    bigIntNum        absValue       originalNumSign
//       7853            7853                 1
//          0               0                 0
//      -7853            7853                -1
//
func(mathBIntMech *MathBigIntHelper) Abs(
	bigIntNum *big.Int) (
	absValue *big.Int,
	originalNumSign int) {

	if mathBIntMech.lock == nil {
		mathBIntMech.lock = new(sync.Mutex)
	}

	mathBIntMech.lock.Lock()

	defer mathBIntMech.lock.Unlock()

	originalNumSign = 0

	absValue =
		big.NewInt(0)

	if bigIntNum == nil {
		return absValue, originalNumSign
	}

	originalNumSign = bigIntNum.Sign()

	if originalNumSign == 0 {
		return absValue, originalNumSign
	}

	if originalNumSign == -1 {
		absValue =
			big.NewInt(0).
				Neg(bigIntNum)
	} else {
		absValue =
			big.NewInt(0).
				Set(bigIntNum)
	}

	return absValue, originalNumSign
}


// ChangeSign - Changes the numeric sign of a *big.Int integer
// number. If the *big.Int integer number passed to this method
// is negative, the corresponding positive value is returned.
// Likewise positive values are converted and returned as negative
// integer values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  bigIntNum             *big.Int
//     - This method will change the numeric sign of this value and
//       return the altered value. Positive values are converted to
//       negative values and negative values are converted to positive
//       values
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  resultIntNum          *big.Int
//     - Contains the result of converting the numeric sign of input
//       parameter, 'bigIntNum'.
//
//
//  originalNumSign       int
//     - Returns the number sign of the original value, 'bigIntNum'.
//       The returned 'originalNumSign' parameter will be set to one
//       of three values depending on value of input parameter
//       'bigIntNum':
//
//          bigIntNum             originalNumSign
//             > 0                        1
//             0.0                        0
//             < 0                       -1
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//    bigIntNum       resultIntNum     originalNumSign
//       7853             -7853               1
//          0                 0               0
//      -7853              7853              -1
//
func(mathBIntMech *MathBigIntHelper) ChangeSign(
	bigIntNum *big.Int) (
	resultIntNum *big.Int,
	originalNumSign int) {

	if mathBIntMech.lock == nil {
		mathBIntMech.lock = new(sync.Mutex)
	}

	mathBIntMech.lock.Lock()

	defer mathBIntMech.lock.Unlock()

	originalNumSign = 0

	resultIntNum =
		big.NewInt(0)

	if bigIntNum == nil {
		return resultIntNum, originalNumSign
	}

	originalNumSign = bigIntNum.Sign()

	if originalNumSign == 0 {
		return resultIntNum, originalNumSign
	}

	resultIntNum =
		big.NewInt(0).
			Neg(bigIntNum)

	return resultIntNum, originalNumSign
}


// GetBigFloatVal - Converts a *big.Int integer value passed
// to this method to a corresponding floating point value of
// type *big.Float.
//
// For more information on *big.Float floating point numbers,
// reference:
//  https://golang.org/pkg/math/big/
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  bigIntNum             *big.Int
//     - This method will convert this integer value to a floating
//       point number of type *big.Float.
//
//
//  precision             uint
//     - This unsigned integer value will specify the numeric precision
//       incorporated in the returned floating point value, 'resultFloat'.
//       The 'precision' value applies to the internal accuracy maintained
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
//  resultIntNum          *big.Int
//     - Contains the result of converting the numeric sign of input
//       parameter, 'bigIntNum'.
//
//
func(mathBIntMech *MathBigIntHelper) GetBigFloatVal(
	bigIntNum *big.Int,
	precision uint) (resultFloat *big.Float) {

	if mathBIntMech.lock == nil {
		mathBIntMech.lock = new(sync.Mutex)
	}

	mathBIntMech.lock.Lock()

	defer mathBIntMech.lock.Unlock()

	resultFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigIntNum == nil ||
			bigIntNum.Sign() == 0 {
		return resultFloat
	}

	resultFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(bigIntNum)

	return resultFloat
}