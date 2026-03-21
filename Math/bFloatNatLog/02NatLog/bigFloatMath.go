package naturalLogCalcs

// ***********************
// ** bigFloatMath.go   **
// ** 013Fix additions  **
// ***********************

import (
  "math/big"
)

type BigFloatMath struct{}

// IntExpoBigFloat
//
// 013Fix
//
// Computes x^n for a *big.Float base x and integer exponent n,
// using exponentiation by squaring.
//
//   - x	 must be non-nil
//   - precBits is the desired output precision in bits
//   - n may be negative, zero, or positive
//
// On success, returns a *big.Float with precision 'precBits'.
func (bFloatMath *BigFloatMath) IntExpoBigFloat(
  x *big.Float,
  n int64,
  precBits uint) (*big.Float, error) {

  ePrefix := "BigFloatMath.IntExpoBigFloat()"

  if x == nil {
    return nil, &InputPtrNilError{
      ErrPrefix:     ePrefix,
      ErrContext:    "",
      ParameterName: "'x'",
    }
  }

  // Handle exponent == 0: x^0 = 1 (even if x == 0, we follow the usual convention)
  if n == 0 {
    return new(big.Float).
      SetMode(big.AwayFromZero).
      SetPrec(precBits).
      SetFloat64(1.0), nil
  }

  // Working precision slightly above requested precision
  workingPrec := precBits + 8

  newF := func() *big.Float {
    return new(big.Float).
      SetMode(big.AwayFromZero).
      SetPrec(workingPrec)
  }

  // Copy base to working precision
  base := newF().Set(x)

  // Track sign of exponent
  negExponent := n < 0
  if negExponent {
    n = -n
  }

  // result = 1
  result := newF().SetFloat64(1.0)

  // Exponentiation by squaring
  for n > 0 {
    if (n & 1) == 1 {
      result.Mul(result, base)
    }
    base.Mul(base, base)
    n >>= 1
  }

  // If exponent was negative, take reciprocal
  if negExponent {
    one := newF().SetFloat64(1.0)
    // Guard against division by zero
    if result.Sign() == 0 {
      return nil, &FuncReturnError{
        ErrPrefix:  ePrefix,
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
    SetPrec(precBits)

  out.Set(result)

  return out, nil
}

// SqrtBigFloat
//
//	Computes sqrt(x) using the classic Newton iteration:
//
//	  y_{k+1} = 1/2 * (y_k + x / y_k)
//
//	Preconditions:
//	  - x must be non‑nil
//	  - x must be >= 0
//	On success, returns a *big.Float with precision 'precBits'.
func (bFloatMath *BigFloatMath) SqrtBigFloat(
  x *big.Float,
  precBits uint) (*big.Float, error) {

  ePrefix := "bFloatMath.SqrtBigFloat()"

  if x == nil {
    return nil, &InputPtrNilError{
      ErrPrefix:     ePrefix,
      ErrContext:    "",
      ParameterName: "'x'",
    }
  }

  // Handle x == 0 quickly
  if x.Sign() == 0 {
    return new(big.Float).
      SetMode(big.AwayFromZero).
      SetPrec(precBits).
      SetFloat64(0.0), nil
  }

  // Negative input is invalid for real sqrt
  if x.Sign() < 0 {
    return nil, &FuncReturnError{
      ErrPrefix:  ePrefix,
      ReturnFunc: "SqrtBigFloat(x, precBits)",
      ErrContext: "Input parameter 'x' is negative.",
      ErrMessage: "square root of negative number is undefined in this context",
    }
  }

  // Working precision: a bit higher than requested to improve convergence
  workingPrec := precBits + 8

  // Create helpers with consistent mode/precision
  newF := func() *big.Float {
    return new(big.Float).
      SetMode(big.AwayFromZero).
      SetPrec(workingPrec)
  }

  // Initial guess: x / 2
  half := newF().SetFloat64(0.5)
  y := newF().Mul(x, half)

  // If x is very small, avoid zero initial guess
  if y.Sign() == 0 {
    y.SetFloat64(1.0)
  }

  preBits := newF()
  tmp := newF()
  xOverY := newF()

  // const maxIter = 128

  bFloatLog210 := new(big.Float).
    SetMode(big.AwayFromZero).
    SetFloat64(3.321928094887362)

  bFloatDecimalDigits :=
    new(big.Float).Quo(preBits, bFloatLog210)

  uint64NumOfDecimalDigits, _ := bFloatDecimalDigits.Uint64()

  var i, maxIter int

  if uint64NumOfDecimalDigits > 8192 {
    maxIter = 1024
  } else if uint64NumOfDecimalDigits > 4096 {
    maxIter = 512
  } else if uint64NumOfDecimalDigits > 2048 {
    maxIter = 256
  } else {
    maxIter = 128
  }

  for i = 0; i < maxIter; i++ {

    // xOverY = x / y
    xOverY.Quo(x, y)

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
    SetPrec(precBits)

  result.Set(y)

  return result, nil
}
