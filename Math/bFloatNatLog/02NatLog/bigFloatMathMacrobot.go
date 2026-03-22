package naturalLogCalcs

import (
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
)

type bigFloatMathMacrobot struct{}

// intExpoBigFloat
//
//	Calculates the result of raising a big.Float base to an
//	integer exponent with specified precision.
//
//	Returns the resulting big.Float and any errors encountered
//	during computation.
//
//	Handles zero and negative exponents per mathematical
//	conventions.
//
//	Returns an error if the base is nil or division by zero
//	occurs for a negative exponent.
//
//	On success, returns a *big.Float with precision
//	'calculationPrecisionBits'.
func (bFMathMacrobot *bigFloatMathMacrobot) intExpoBigFloat(
	bFloatBase *big.Float,
	intExponent int64,
	calculationPrecisionBits uint,
	errPrefDto *ePref.ErrPrefixDto) (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigFloatMathMacrobot.intExpoBigFloat()",
		"")

	if err != nil {
		return new(big.Float), err
	}

	if bFloatBase == nil {
		return nil, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'x'",
		}
	}

	// Handle exponent == 0: x^0 = 1 (even if x == 0, we follow the usual convention)
	if intExponent == 0 {
		return new(big.Float).
			SetMode(big.AwayFromZero).
			SetPrec(calculationPrecisionBits).
			SetFloat64(1.0), nil
	}

	// Working precision slightly above requested precision
	workingPrec := calculationPrecisionBits + 22

	newF := func() *big.Float {
		return new(big.Float).
			SetMode(big.AwayFromZero).
			SetPrec(workingPrec)
	}

	// Copy base to working precision
	base := newF().Set(bFloatBase)

	// Track sign of exponent
	negExponent := intExponent < 0

	if negExponent {
		intExponent = -intExponent
	}

	// result = 1
	result := newF().SetFloat64(1.0)

	// Exponentiation by squaring
	for intExponent > 0 {

		if (intExponent & 1) == 1 {
			result.Mul(result, base)
		}

		base.Mul(base, base)
		intExponent >>= 1

	}

	// If exponent was negative, take reciprocal
	if negExponent {

		one := newF().SetFloat64(1.0)

		// Guard against division by zero
		if result.Sign() == 0 {
			return nil, &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "IntExpoBigFloat(x, n, precBits)",
				ErrContext: "Attempted reciprocal of zero for negative exponent.",
				ErrMessage: "division by zero",
			}
		}

		result.Quo(one, result)
	}

	// Round to requested precision
	out := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(calculationPrecisionBits)

	out.Set(result)

	return out, nil
}

// sqrtBigFloat
//
//	Computes the square root of a *big.Float using Newton's
//	method with specified precision.
//
//	Returns an error if the input parameter, 'bFloatXValue'
//	is nil or negative.
//
//	'integerDigits' is the number of integer digits to the
//	left of the decimal point in the input parameter,
//	'bFloatXValue'.
//
//	`integerDigits` adjusts iterations based on input size.
//	If `integerDigits` 0, it is calculated internally.
//
//	If the input parameter, 'bFloatXValue' is zero, the result
//	is zero.
//
//	If the input parameter, 'bFloatXValue' is negative, an error
//	is returned.
//
//	On success, returns a *big.Float with a precision bits value
//	of 'calculationPrecisionBits'.
func (bFMathMacrobot *bigFloatMathMacrobot) sqrtBigFloat(
	bFloatXValue *big.Float,
	integerDigits uint,
	calculationPrecisionBits uint,
	errPrefDto *ePref.ErrPrefixDto) (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigFloatMathMacrobot.sqrtBigFloat()",
		"")

	if err != nil {
		return new(big.Float), err
	}

	if bFloatXValue == nil {
		return nil, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'x'",
		}
	}

	// Handle bFloatXValue == 0 quickly
	if bFloatXValue.Sign() == 0 {
		return new(big.Float).
			SetMode(big.AwayFromZero).
			SetPrec(calculationPrecisionBits).
			SetFloat64(0.0), nil
	}

	// Negative input is invalid for real sqrt
	if bFloatXValue.Sign() < 0 {
		return nil, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sqrtBigFloat(bFloatXValue, integerDigits, calculationPrecisionBits)",
			ErrContext: "Input parameter 'bFloatXValue' is negative.",
			ErrMessage: "Square root of negative number is undefined in this context!",
		}
	}

	bFloatLog210 := new(big.Float).
		SetMode(big.AwayFromZero).
		SetFloat64(3.321928094887362)

	// Working precision: 12% more precision bits
	// than requested to improve convergence and accuracy.
	workingPrec := calculationPrecisionBits + uint(float64(calculationPrecisionBits)*0.12)

	// Create helpers with consistent mode/precision
	newF := func() *big.Float {
		return new(big.Float).
			SetMode(big.AwayFromZero).
			SetPrec(workingPrec)
	}

	// Initial guess: XValue / 2
	half := newF().SetFloat64(0.5)

	y := newF().Mul(bFloatXValue, half)

	// If XValue is very small, avoid zero initial guess
	if y.Sign() == 0 {
		y.SetFloat64(1.0)
	}

	// base Working precision used to compute
	// number of equivalent decimal digits.
	preBits := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(0).
		SetUint64(uint64(workingPrec))

	tmp := newF()
	xOverY := newF()

	if integerDigits == 0 {
		intNumStr := bFloatXValue.Text('f', 0)

		integerDigits = uint(len(intNumStr))

	}

	bFloatDecimalDigits :=
		new(big.Float).Quo(preBits, bFloatLog210)

	uint64NumOfDecimalDigits, _ := bFloatDecimalDigits.Uint64()

	var i, maxIter uint64

	maxIter = (uint64(integerDigits) + uint64NumOfDecimalDigits) * uint64(100)

	for i = 0; i < maxIter; i++ {

		// xOverY = x / y
		xOverY.Quo(bFloatXValue, y)

		// tmp = y + x/y
		tmp.Add(y, xOverY)

		// yNext = 0.5 * tmp
		yNext := newF().Mul(tmp, half)

		// Check convergence: if yNext == y at this precision, stop
		if yNext.Cmp(y) == 0 {
			y = yNext
			break
		}

		y = yNext
	}

	// Round result to requested precision
	result := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(workingPrec).
		Set(y)

	return result, nil
}
