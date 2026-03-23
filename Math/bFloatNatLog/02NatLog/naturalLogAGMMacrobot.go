// naturalLogAGMMacrobot.go

package naturalLogCalcs

import (
  "math/big"

  ePref "github.com/MikeAustin71/errpref"
)

// naturalLogAGMMacrobot
//
// High-level AGM engine: drives the full AGM iteration
// from initial (a0, b0) to the final mean.
type naturalLogAGMMacrobot struct{}

// agmCore
//
// Computes the arithmetic-geometric mean of (a0, b0) using the
// classical iteration:
//
//	a_{n+1} = (a_n + b_n) / 2
//	b_{n+1} = sqrt(a_n * b_n)
//
// Preconditions:
//   - a0, b0 non-nil
//   - a0, b0 > 0
//   - precBits > 0
//
// On success, returns AGM(a0, b0) as *big.Float with precision = precBits.
func (agmMacro *naturalLogAGMMacrobot) agmCore(
  a0 *big.Float,
  b0 *big.Float,
  precBits uint,
  errPrefDto *ePref.ErrPrefixDto) (*big.Float, error) {

  ePrefix, err := ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "naturalLogAGMMacrobot.agmCore()",
    "")
  if err != nil {
    return new(big.Float), err
  }

  if a0 == nil {
    return nil, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ErrContext:    "",
      ParameterName: "'a0'",
    }
  }

  if b0 == nil {
    return nil, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ErrContext:    "",
      ParameterName: "'b0'",
    }
  }

  if a0.Sign() <= 0 || b0.Sign() <= 0 {
    return nil, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "agmCore(a0, b0, precBits)",
      ErrContext: "AGM requires strictly positive inputs.",
      ErrMessage: "non-positive AGM arguments",
    }
  }

  if precBits == 0 {
    precBits = 64
  }

  // Slightly elevated working precision for internal iteration.
  workingPrec := precBits + 16

  newF := func() *big.Float {
    return new(big.Float).
      SetMode(big.AwayFromZero).
      SetPrec(workingPrec)
  }

  a := newF().Set(a0)
  b := newF().Set(b0)

  mol := new(naturalLogAGMMolecule)

  // Conservative upper bound on iterations.
  maxIter := uint64(workingPrec*4 + 64)

  for i := uint64(0); i < maxIter; i++ {

    aNext, bNext, err := mol.agmIterateOnce(
      a,
      b,
      workingPrec,
      ePrefix)
    if err != nil {
      return new(big.Float), err
    }

    if mol.agmConverged(aNext, bNext) {
      a = aNext
      b = bNext
      break
    }

    a = aNext
    b = bNext
  }

  // Round result to requested precision.
  out := new(big.Float).
    SetMode(big.AwayFromZero).
    SetPrec(precBits)

  out.Set(a)

  return out, nil
}
