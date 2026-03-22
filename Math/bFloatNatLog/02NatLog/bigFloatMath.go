package naturalLogCalcs

// ***********************
// ** bigFloatMath.go   **
// ** 013Fix additions  **
// ***********************

import (
	"math"
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
)

type BigFloatMath struct{}

// IntExpoBigFloat
//
//	Calculates a big.Float base raised to an integer exponent with
//	specified precision in expressed in bits for use with big.Float.
//
//	bFloatBase: the base as a *big.Float, must not be nil.
//
//	intExponent: the integer exponent value. Handles zero and
//	negative exponents per mathematical conventions.
//
//	calculationPrecisionBits: big.Float precision for the calculation
//	expressed in bits. This controls the number of decimal digits
//	achievable for the result.
//
//	Returns an error if the base is nil or division by zero
//	occurs for a negative exponent.
//
//	On success, returns a *big.Float with precision
//	'calculationPrecisionBits'.
func (bFloatMath *BigFloatMath) IntExpoBigFloat(
	bFloatBase *big.Float,
	intExponent int64,
	calculationPrecisionBits uint) (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigFloatMath.IntExpoBigFloat()",
		"")

	if err != nil {
		return new(big.Float), err
	}

	return new(bigFloatMathMacrobot).intExpoBigFloat(
		bFloatBase,
		intExponent,
		calculationPrecisionBits,
		ePrefix)
}

// IntExpoBigFloatDigits
//
//	Calculates a big.Float base raised to an integer exponent with
//	big.Float precision bits calculated from input parameter decimal
//	digit precision (calculationDecimalDigits). Decimal digits are
//	defined as those digits to the right of the decimal point.
//
//	bFloatBase: The base as a *big.Float, must not be nil.
//
//	intExponent: The integer exponent supports zero and negative
//	exponents per mathematical conventions.
//
//	calculationDecimalDigits: The number of decimal digits to the
//	right of the decimal point which will be used in generating the
//	calculated results.
//
//	Returns an error if the 'bFloatBase' is nil or division by zero
//	occurs for a negative exponent.
//
//	On success, returns a *big.Float with precision derived from
//	the 'calculationDecimalDigits' input parameter.
func (bFloatMath *BigFloatMath) IntExpoBigFloatDigits(
	bFloatBase *big.Float,
	intExponent int64,
	calculationDecimalDigits uint) (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigFloatMath.IntExpoBigFloatDigits()",
		"")

	if err != nil {
		return new(big.Float), err
	}

	safetyMarginDigits := uint(math.Ceil(float64(calculationDecimalDigits) * 0.5))

	if safetyMarginDigits < 4 {
		safetyMarginDigits = 4
	}

	calculationPrecisionBits := new(BigFloatHelper).
		ComputeBigFloatPrecisionBits(calculationDecimalDigits+safetyMarginDigits, 1)

	return new(bigFloatMathMacrobot).intExpoBigFloat(
		bFloatBase,
		intExponent,
		calculationPrecisionBits,
		ePrefix)
}

// SqrtBigFloat
//
//	 Computes sqrt(x) using the classic Newton iteration:
//
//			  y_{k+1} = 1/2 * (y_k + x / y_k)
//
//	 Preconditions:
//			  - x must be non‑nil
//			  - x must be >= 0
//
//	 'integerDigits' is the number of integer digits to the
//	 left of the decimal point in the input parameter,
//	 'bFloatXValue'.
//
//		If integerDigits == 0, the number of integer digits is
//		calculated for the input parameter 'bFloatXValue'.
//
//	 If the input parameter, 'bFloatXValue' is zero, the result
//	 is zero.
//
//	 If the input parameter, 'bFloatXValue' is negative, an error
//	 is returned.
//
//	 On success, returns a *big.Float with a precision bits value
//	 of 'calculationPrecisionBits'.
func (bFloatMath *BigFloatMath) SqrtBigFloat(
	bFloatXValue *big.Float,
	integerDigits uint,
	calculationPrecisionBits uint) (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigFloatMath.SqrtBigFloat()",
		"")

	if err != nil {
		return new(big.Float), err
	}

	return new(bigFloatMathMacrobot).sqrtBigFloat(bFloatXValue, integerDigits, calculationPrecisionBits, ePrefix)
}

// SqrtBigFloatDigits
//
// Calculates the square root of a big.Float value with precision
// derived from decimal digits of precision (calculationDecimalDigits).
//
// bFloatXValue: The input value as a *big.Float, must not be nil or
// negative. Zero returns zero without error.
//
// calculationDecimalDigits: The number of decimal digits to the right
// of the decimal place used for precision during calculation. This
// value is converted to precision bits used in big.Float calculations.
//
// Returns an error if the input value is nil or negative.
//
// On success, returns the square root of 'bFloatXValue' as a
// type *big.Float with a precision bits value derived from
// 'calculationDecimalDigits'.
func (bFloatMath *BigFloatMath) SqrtBigFloatDigits(
	bFloatXValue *big.Float,
	calculationDecimalDigits uint) (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigFloatMath.SqrtBigFloatDigits()",
		"")

	if err != nil {
		return new(big.Float), err
	}

	safetyMarginDigits := uint(math.Ceil(float64(calculationDecimalDigits) * 0.5))

	if safetyMarginDigits < 4 {
		safetyMarginDigits = 4
	}

	calculationPrecisionBits := new(BigFloatHelper).
		ComputeBigFloatPrecisionBits(calculationDecimalDigits+safetyMarginDigits, 1)

	intXValueStr := bFloatXValue.Text('f', 0)

	integerDigits := uint(len(intXValueStr))

	return new(bigFloatMathMacrobot).sqrtBigFloat(bFloatXValue, integerDigits, calculationPrecisionBits, ePrefix)
}
