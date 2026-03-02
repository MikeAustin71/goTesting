package bigIntEx

import (
  "fmt"
  "math/big"
  "strings"
  "sync"
)

type BigIntMathPowerMolecule struct {
  lock *sync.Mutex
}

func (bIMathPwrMolecule *BigIntMathPowerMolecule) BigRat(
  numerator *big.Int, denominator *big.Int) error {

  ePrefix := "BigIntMathPowerMolecule.BigRat()"

  if numerator == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix,
      ParameterName: "'numerator'",
    }
  }

  if denominator == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix,
      ParameterName: "'denominator'",
    }
  }

  bigIntOne := big.NewInt(1)

  bRatNumerator := big.NewRat(1, 1)

  bRatNumerator.SetFrac(numerator, bigIntOne)

  bRatDenominator := big.NewRat(1, 1)

  bRatDenominator.SetFrac(denominator, bigIntOne)

  quotient := big.NewRat(0, 1)

  _ = quotient.Quo(bRatNumerator, bRatDenominator)

  quotientStr := quotient.FloatString(51)

  quotientRatStr := quotient.RatString()

  numOfDecimalDigits, isExact := quotient.FloatPrec()

  floatStr := quotient.FloatString(51)

  breakStr := strings.Repeat("=", 50)

  subBreakStr := strings.Repeat("-", 50)

  fmt.Printf("'%v'\n"+
    "%v\n"+
    "%v\n"+
    "numerator= '%v'\n"+
    "denominator= '%v'\n"+
    "numOfDecimalDigits= '%v'\n"+
    "isExact='%v'\n"+
    "   floatStr= '%v'\n"+
    "quotientStr= '%v'\n"+
    "quotientRatStr= '%v'\n"+
    "%v\n"+
    "%v\n\n",
    breakStr,
    ePrefix,
    subBreakStr,
    numerator.Text(10),
    denominator.Text(10),
    numOfDecimalDigits,
    isExact,
    floatStr,
    quotientStr,
    quotientRatStr,
    subBreakStr,
    breakStr)

  return nil
}

func (bIMathPwrMolecule *BigIntMathPowerMolecule) BigIntDivide(
  numerator *big.Int,
  denominator *big.Int) (err error) {

  ePrefix := "BigIntMathPowerMolecule.BigIntDivide()"

  quotient := big.NewInt(0)
  remainder := big.NewInt(0)

  if numerator == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix,
      ParameterName: "'numerator'",
    }
  }

  if denominator == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix,
      ParameterName: "'denominator'",
    }
  }

  quotient, remainder = big.NewInt(0).QuoRem(numerator, denominator, remainder)

  mod := big.NewInt(0).Mod(numerator, denominator)

  breakStr := strings.Repeat("=", 50)

  subBreakStr := strings.Repeat("-", 50)

  fmt.Printf("'%v'\n"+
    "%v\n"+
    "%v\n"+
    "numerator= '%v'\n"+
    "denominator= '%v'\n"+
    "quotient= '%v'\n"+
    "remainder='%v'\n"+
    "mod= %v\n"+
    "%v\n"+
    "%v\n\n",
    breakStr,
    ePrefix,
    subBreakStr,
    numerator.Text(10),
    denominator.Text(10),
    quotient.Text(10),
    remainder.Text(10),
    mod.Text(10),
    subBreakStr,
    breakStr)

  return nil
}

func (bIMathPwrMolecule *BigIntMathPowerMolecule) BigIntNumRaiseToNegativeFractionalPower(
  base *big.Int,
  exponent *big.Int) (err error) {

  ePrefix := "BigIntMathPowerMolecule.BigIntNumRaiseToNegativeFractionalPower()"

  if base == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix,
      ParameterName: "'base'",
    }
  }

  if exponent == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix,
      ParameterName: "'exponent'",
    }
  }

  exponentSignValue := exponent.Sign()

  if exponentSignValue != -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix,
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("'exponent' MUST BE a negative number!\n"+
        "exponent='%v'\n", exponent.Text(10)),
    }

  }

  absoluteExponent := big.NewInt(0).Abs(exponent)

  /*

  	temBaseBigInt := big.NewInt(0).Set(base)

  	tempBaseOriginalIntSignValue := temBaseBigInt.Sign()
  	temBaseBigIntSignValue := big.NewInt(int64(tempBaseOriginalIntSignValue))

  	if tempBaseOriginalIntSignValue == -1 {
  		temBaseBigInt.Neg(temBaseBigInt)
  	}
  */
  denominator := big.NewInt(0).Exp(base, absoluteExponent, nil)

  numeratorOne := big.NewInt(1)

  bigIntOne := big.NewInt(1)

  bRatNumerator := big.NewRat(1, 1)

  bRatNumerator.SetFrac(numeratorOne, bigIntOne)

  bRatDenominator := big.NewRat(1, 1)

  bRatDenominator.SetFrac(denominator, bigIntOne)

  quotient := big.NewRat(0, 1)

  _ = quotient.Quo(bRatNumerator, bRatDenominator)

  //quotientStr := quotient.FloatString(51)

  //posOne := big.NewInt(1)

  negOne := big.NewInt(-1)

  result := big.NewInt(0).Exp(denominator, negOne, nil)

  breakStr := strings.Repeat("=", 50)

  subBreakStr := strings.Repeat("-", 50)

  fmt.Printf("'%v'\n"+
    "%v\n"+
    "%v\n"+
    "base= '%v'\n"+
    "exponent= '%v'\n"+
    "denominator= '%v'\n"+
    "result='%v'\n"+
    "%v\n"+
    "%v\n\n",
    breakStr,
    ePrefix,
    subBreakStr,
    base.Text(10),
    exponent.Text(10),
    denominator.Text(10),
    result.Text(10),
    subBreakStr,
    breakStr)

  return nil
}
