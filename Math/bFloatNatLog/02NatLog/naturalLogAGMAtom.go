// naturalLogAGMAtom.go

package naturalLogCalcs

import (
  "math/big"

  ePref "github.com/MikeAustin71/errpref"
)

// naturalLogAGMAtom
//
// Lowest-level AGM helpers: arithmetic and geometric means.
type naturalLogAGMAtom struct{}

// newFloat
//
// Convenience constructor for *big.Float with AwayFromZero mode
// and the specified precision in bits.
func (agmAtom *naturalLogAGMAtom) newFloat(
  precBits uint) *big.Float {

  return new(big.Float).
    SetMode(big.AwayFromZero).
    SetPrec(precBits)
}

// arithmeticMean
//
// Computes the arithmetic mean:
//
//	(a + b) / 2
func (agmAtom *naturalLogAGMAtom) arithmeticMean(
  a *big.Float,
  b *big.Float,
  precBits uint) *big.Float {

  f := agmAtom.newFloat(precBits)

  sum := agmAtom.newFloat(precBits)
  sum.Add(a, b)

  two := agmAtom.newFloat(precBits)
  two.SetFloat64(2.0)

  f.Quo(sum, two)

  return f
}

// geometricMean
//
// Computes the geometric mean:
//
//	sqrt(a * b)
//
// Uses the existing BigFloatMath.SqrtBigFloat for the square root.
func (agmAtom *naturalLogAGMAtom) geometricMean(
  a *big.Float,
  b *big.Float,
  precBits uint,
  errPrefDto *ePref.ErrPrefixDto) (*big.Float, error) {

  ePrefix, err := ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "naturalLogAGMAtom.geometricMean()",
    "")
  if err != nil {
    return agmAtom.newFloat(precBits), err
  }

  if a == nil {
    return nil, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ErrContext:    "",
      ParameterName: "'a'",
    }
  }

  if b == nil {
    return nil, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ErrContext:    "",
      ParameterName: "'b'",
    }
  }

  product := agmAtom.newFloat(precBits)
  product.Mul(a, b)

  if product.Sign() <= 0 {
    return nil, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "geometricMean(a, b, precBits)",
      ErrContext: "a*b must be > 0 for geometric mean.",
      ErrMessage: "non-positive product in geometric mean",
    }
  }

  sqrtCalc := new(BigFloatMath)

  gMean, err := sqrtCalc.SqrtBigFloat(
    product,
    0,        // let SqrtBigFloat compute integerDigits
    precBits, // working precision
  )
  if err != nil {
    return nil, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "SqrtBigFloat(product, 0, precBits)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return gMean, nil
}
