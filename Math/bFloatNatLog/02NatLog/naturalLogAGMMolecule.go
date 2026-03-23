// naturalLogAGMMolecule.go
package naturalLogCalcs

import (
  "math/big"

  ePref "github.com/MikeAustin71/errpref"
)

// naturalLogAGMMolecule
//
// Implements the AGM iteration using Atom-level helpers.
type naturalLogAGMMolecule struct{}

// agmIterateOnce
//
// Performs a single AGM iteration step:
//
//	aNext = (a + b) / 2
//	bNext = sqrt(a * b)
func (agmMol *naturalLogAGMMolecule) agmIterateOnce(
  a *big.Float,
  b *big.Float,
  precBits uint,
  errPrefDto *ePref.ErrPrefixDto) (
  aNext *big.Float,
  bNext *big.Float,
  err error) {

  ePrefix, err := ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "naturalLogAGMMolecule.agmIterateOnce()",
    "")
  if err != nil {
    return new(big.Float), new(big.Float), err
  }

  if a == nil {
    return nil, nil, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ErrContext:    "",
      ParameterName: "'a'",
    }
  }

  if b == nil {
    return nil, nil, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ErrContext:    "",
      ParameterName: "'b'",
    }
  }

  atom := new(naturalLogAGMAtom)

  aNext = atom.arithmeticMean(a, b, precBits)

  bNext, err = atom.geometricMean(a, b, precBits, ePrefix)
  if err != nil {
    return new(big.Float), new(big.Float), err
  }

  return aNext, bNext, nil
}

// agmConverged
//
// Returns true if a and b are equal at the given precision.
// This is a strong convergence test at working precision.
func (agmMol *naturalLogAGMMolecule) agmConverged(
  a *big.Float,
  b *big.Float) bool {

  if a == nil || b == nil {
    return false
  }

  return a.Cmp(b) == 0
}
