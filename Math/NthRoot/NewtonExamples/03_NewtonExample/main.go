package main

import (
  "fmt"
  "math"
  "math/big"
  "time"
)

func main() {

  TestNewtonInitGuess01()

  return
}

func TestNumDecDigitsFromPrecisionBits() {

  ePrefix := "TestNumDecDigitsFromPrecisionBits()"

  precisionBits := uint(2048)
  //safetyDigits := uint(2)
  safetyDigits := uint(0)

  numOfDigits := ComputeBigFloatDecimalDigits(precisionBits, safetyDigits)

  fmt.Printf("\n%v\n"+
    "precisionBits = %v\n"+
    "safetyDigits = %v\n"+
    "Calculated Number Of Digits = %v\n",
    ePrefix, precisionBits, safetyDigits, numOfDigits)

}

func TestNewtonInitGuess01() {

  ePrefix := "TestNumDecDigitsFromPrecisionBits()"

  nRoot := int64(83)

  radicandString := "2632183"

  maxInternalNumOfDecDigits := uint(256)

  accuracyThresholdStr := "0.0000000000000000000000005"

  expectedResultStr := "1.194959553693090978422686714"

  bFloatRadicand, isOk := new(big.Float).
    SetMode(big.AwayFromZero).
    SetString(radicandString)

  if isOk == false {
    fmt.Printf("\n%v\n"+
      "SetString(radicandString) failed.\n"+
      "nRoot = %v\n"+
      "radicandString = %v\n\n",
      ePrefix, nRoot, radicandString)
    return
  }

  accuracyThreshold, isOk := new(big.Float).
    SetMode(big.AwayFromZero).
    SetString(accuracyThresholdStr)

  if isOk == false {
    fmt.Printf("\n%v\n"+
      "SetString(accuracyThresholdStr) failed.\n"+
      "accuracyThresholdStr = %v\n\n",
      ePrefix, accuracyThresholdStr)
    return
  }

  startTime := time.Now()

  guess,
    raisedToPower,
    absRadicandDelta,
    compareResult,
    cycleCount,
    err := NewtonInitialGuess01(
    nRoot,
    bFloatRadicand,
    accuracyThreshold,
    maxInternalNumOfDecDigits)

  if err != nil {
    fmt.Printf("\n%v\n"+
      "NewtonInitialGuess01() failed.\n"+
      "nRoot = %v\n"+
      "radicandString = %v\n"+
      "accuracyThresholdStr = %v\n"+
      "maxInternalNumOfDecDigits = %v\n"+
      "Error= '%v'\n\n",
      ePrefix,
      nRoot,
      radicandString,
      accuracyThresholdStr,
      maxInternalNumOfDecDigits,
      err.Error())

    return
  }

  endTime := time.Now()

  elapsedTime := DurationBreakdown(startTime, endTime)

  fmt.Printf("\n\n%v\n"+
    "Successful Completion\n"+
    "nRoot = %v\n"+
    "radicandString = %v\n"+
    "maxInternalNumOfDecDigits = %v\n"+
    "cycleCount = %v\n"+
    "compareResult = %v\n"+
    "accuracyThresholdStr = %v\n"+
    "guess                = %v\n"+
    "expected guess       = %v\n"+
    "raisedToPower = %v\n"+
    "absRadicandDelta = %v\n"+
    "Elapsed Time: %v\n\n",
    ePrefix,
    nRoot,
    radicandString,
    maxInternalNumOfDecDigits,
    cycleCount,
    compareResult,
    accuracyThresholdStr,
    guess.Text('f', 20),
    expectedResultStr,
    raisedToPower.Text('f', 20),
    absRadicandDelta.Text('f', 20),
    elapsedTime)

}

func NewtonInitialGuessX(
  nthRoot int64,    // degree - a.k.a. nthRoot
  alpha *big.Float, // radicand - a.k.a. value
  accuracyThreshold *big.Float,
  prec uint) (
  guess *big.Float,
  raisedToPower *big.Float,
  absRadicandDelta *big.Float,
  compareResult int,
  cycleCount uint64) {

  // SetMode sets z's rounding mode to mode and returns an exact z.
  // z remains unchanged otherwise.
  // z.SetMode(z.Mode()) is a cheap way to set z's accuracy to Exact.

  numOfLowGuesses := uint64(0)
  numOfHighGuesses := uint64(0)

  absRadicandDelta = new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  absPercentRadicandDelta := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  absGuessDelta := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  half := new(big.Float).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  half.SetString("0.5")

  guess = new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  guess.SetString("1.015")

  lastLowGuess := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastLowGuess.SetString("0.0")

  lastLowRaisedToPower := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastLowRaisedToPower.SetString("0.0")

  lastLowAbsAlphaDelta := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastLowAbsAlphaDelta.SetString("0.0")

  lastHighGuess := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastHighGuess.SetString("1.5")

  lastHighRaisedToPower := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastHighRaisedToPower.SetString("1.5")

  lastHighAbsAlphaDelta := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastHighAbsAlphaDelta.SetString("1.5")

  numberOne := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  numberOne.SetString("1.0")

  numberTwo := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  numberTwo.SetString("2.0")

  deltaIncrement := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  deltaIncrement.SetString("0.0")

  tempCompare := 0

  cycleCount = uint64(0)

  closeEnough := false

  for closeEnough == false {

    raisedToPower = BigFloatPower01(
      guess,
      nthRoot,
      prec)

    compareResult = raisedToPower.Cmp(alpha)

    if compareResult == 0 {

      absRadicandDelta.Sub(alpha, raisedToPower)

      // On target
      return guess,
        raisedToPower,
        absRadicandDelta,
        compareResult,
        cycleCount

    } else if compareResult == -1 {
      // Guess is low. Alpha is higher than
      // raisedToPower

      lastLowGuess.Set(guess)

      lastLowRaisedToPower.Set(raisedToPower)

      absRadicandDelta.Sub(alpha, lastLowRaisedToPower)

      absPercentRadicandDelta.Quo(absRadicandDelta, alpha)

      //fmt.Printf("lastLowGuess  = %v\n"+
      //	"lastLowRaisedToPower = %v\n"+
      //	"       absAlphaDelta = %v\n"+
      //	"absPercentAlphaDelta = %v\n",
      //	lastLowGuess.Text('f', 20),
      //	lastLowRaisedToPower.Text('f', 20),
      //	absAlphaDelta.Text('f', 20),
      //	absPercentAlphaDelta.Text('f', 20))

      deltaIncrement.Set(half)

      if numOfLowGuesses > 2 {

        if absPercentRadicandDelta.Cmp(half) == 1 {

          deltaIncrement.Set(numberTwo)

        }

      }

      lastLowAbsAlphaDelta.Set(absRadicandDelta)

      tempCompare = lastLowGuess.Cmp(lastHighGuess)

      if tempCompare == 1 {
        // Last Low Guess is higher than
        // Last High Guess
        absGuessDelta.Mul(guess, deltaIncrement)

      } else if tempCompare == 0 {
        // Last Low Guess is Equal To
        // Last High Guess
        absGuessDelta.Mul(guess, deltaIncrement)

      } else {
        // Last High Guess is higher than
        //	Last Low Guess
        absGuessDelta.Sub(lastHighGuess, lastLowGuess)

        absGuessDelta.Mul(absGuessDelta, deltaIncrement)

      }

      //fmt.Printf("Old Guess = %v\n",
      //	guess.Text('f', -1))
      //
      //fmt.Printf("absGuessDelta = %v\n",
      //	absGuessDelta.Text('f', -1))

      guess.Add(guess, absGuessDelta)

      //fmt.Printf("New Guess = %v\n",
      //	guess.Text('f', -1))

      numOfLowGuesses++

    } else {
      // MUST BE -
      //	Guess is high. raisedToPower is higher
      //	than Alpha

      lastHighGuess.Set(guess)

      lastHighRaisedToPower.Set(raisedToPower)

      absRadicandDelta.Sub(lastHighRaisedToPower, alpha)

      absPercentRadicandDelta.Quo(absRadicandDelta, alpha)

      deltaIncrement.Set(half)

      if numOfHighGuesses > 2 {

        if absPercentRadicandDelta.Cmp(half) == 1 {

          deltaIncrement.Set(numberTwo)

        }

      }

      //fmt.Printf("lastHighGuess  = %v\n"+
      //	"lastHighRaisedToPower = %v\n",
      //	lastHighGuess.Text('f', 10),
      //	lastHighRaisedToPower.Text('f', 10))

      lastHighAbsAlphaDelta.Set(absRadicandDelta)

      tempCompare = lastLowGuess.Cmp(lastHighGuess)

      if tempCompare == 1 {
        // Low Guess is higher than High Guess
        absGuessDelta.Mul(lastHighGuess, half)
        guess.Sub(guess, absGuessDelta)

      } else if tempCompare == 0 {
        absGuessDelta.Mul(lastHighGuess, half)
        guess.Sub(guess, absGuessDelta)

      } else {
        // High Guess is higher than Low Guess
        absGuessDelta.Sub(lastHighGuess, lastLowGuess)
        absGuessDelta.Mul(absGuessDelta, half)
      }

      guess.Sub(guess, absGuessDelta)

      //fmt.Printf("New Guess = %v\n",
      //	guess.Text('f', -1))

      numOfHighGuesses++
    }

    cycleCount++

    if compareResult == -1 &&
      absRadicandDelta.Cmp(accuracyThreshold) < 1 {

      guess.Set(lastLowGuess)
      raisedToPower.Set(lastLowRaisedToPower)

      closeEnough = true
    }

    if cycleCount > 10000 {

      if compareResult == 1 {
        guess.Set(lastLowGuess)
        raisedToPower.Set(lastLowRaisedToPower)
        absRadicandDelta.Sub(alpha, raisedToPower)

      }

      closeEnough = true

    }

  }

  compareResult = raisedToPower.Cmp(alpha)

  absRadicandDelta.Sub(raisedToPower, alpha)

  if absRadicandDelta.Sign() == -1 {
    absRadicandDelta.Neg(absRadicandDelta)
  }

  return guess,
    raisedToPower,
    absRadicandDelta,
    compareResult,
    cycleCount
}

func NewtonInitialGuess01(
  nthRoot int64,       // degree - a.k.a. nthRoot
  radicand *big.Float, // radicand - a.k.a. value
  accuracyThreshold *big.Float,
  maxInternalNumOfDecDigits uint) (
  guess *big.Float,
  raisedToPower *big.Float,
  absRadicandDelta *big.Float,
  compareResult int,
  cycleCount uint64,
  err error) {

  // SetMode sets z's rounding mode to mode and returns an exact z.
  // z remains unchanged otherwise.
  // z.SetMode(z.Mode()) is a cheap way to set z's accuracy to Exact.

  ePrefix := "NewtonInitialGuess01()"

  numOfLowGuesses := uint64(0)
  numOfHighGuesses := uint64(0)

  prec := ComputeBigFloatPrecisionBits(maxInternalNumOfDecDigits, 0)

  absRadicandDelta = new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  absPercentRadicandDelta := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  absGuessDelta := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  half := new(big.Float).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  half.SetString("0.5")

  guess = new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  guess.SetString("1.015")

  lastLowGuess := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastLowGuess.SetString("0.0")

  lastLowRaisedToPower := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastLowRaisedToPower.SetString("0.0")

  lastLowAbsRadicandDelta := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastLowAbsRadicandDelta.SetString("0.0")

  lastHighGuess := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastHighGuess.SetString("1.5")

  lastHighRaisedToPower := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastHighRaisedToPower.SetString("1.5")

  lastHighAbsAlphaDelta := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  lastHighAbsAlphaDelta.SetString("1.5")

  numberOne := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  numberOne.SetString("1.0")

  numberTwo := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  numberTwo.SetString("2.0")

  deltaIncrement := new(big.Float).
    SetInt64(0).
    SetPrec(prec).
    SetMode(big.AwayFromZero)

  deltaIncrement.SetString("0.0")

  tempCompare := 0

  cycleCount = uint64(0)

  closeEnough := false

  var targetPrec uint

  for closeEnough == false {

    guess.SetPrec(guess.MinPrec())

    targetPrec = ComputeBigFloatPrecisionBits(maxInternalNumOfDecDigits, uint(nthRoot))

    raisedToPower = BigFloatPower01(
      guess,
      nthRoot,
      targetPrec)

    compareResult = raisedToPower.Cmp(radicand)

    if compareResult == 0 {

      absRadicandDelta.Sub(radicand, raisedToPower)

      // On target
      return guess,
        raisedToPower,
        absRadicandDelta,
        compareResult,
        cycleCount,
        err

    } else if compareResult == -1 {
      // Guess is low. Radicand is higher than
      // raisedToPower

      lastLowGuess.Set(guess)

      lastLowRaisedToPower.Set(raisedToPower)

      absRadicandDelta.Sub(radicand, lastLowRaisedToPower)

      absPercentRadicandDelta.Quo(absRadicandDelta, radicand)

      //fmt.Printf("lastLowGuess  = %v\n"+
      //	"lastLowRaisedToPower = %v\n"+
      //	"       absAlphaDelta = %v\n"+
      //	"absPercentAlphaDelta = %v\n",
      //	lastLowGuess.Text('f', 20),
      //	lastLowRaisedToPower.Text('f', 20),
      //	absAlphaDelta.Text('f', 20),
      //	absPercentAlphaDelta.Text('f', 20))

      deltaIncrement.Set(half)

      if numOfLowGuesses > 2 {

        if absPercentRadicandDelta.Cmp(half) == 1 {

          deltaIncrement.Set(numberTwo)

        }

      }

      lastLowAbsRadicandDelta.Set(absRadicandDelta)

      tempCompare = lastLowGuess.Cmp(lastHighGuess)

      if tempCompare == 1 {
        // Last Low Guess is higher than
        // Last High Guess
        absGuessDelta.Mul(guess, deltaIncrement)

      } else if tempCompare == 0 {
        // Last Low Guess is Equal To
        // Last High Guess
        absGuessDelta.Mul(guess, deltaIncrement)

      } else {
        // Last High Guess is higher than
        //	Last Low Guess
        absGuessDelta.Sub(lastHighGuess, lastLowGuess)

        absGuessDelta.Mul(absGuessDelta, deltaIncrement)

      }

      //fmt.Printf("Old Guess = %v\n",
      //	guess.Text('f', -1))
      //
      //fmt.Printf("absGuessDelta = %v\n",
      //	absGuessDelta.Text('f', -1))

      guess.Add(guess, absGuessDelta)

      //fmt.Printf("New Guess = %v\n",
      //	guess.Text('f', -1))

      numOfLowGuesses++

    } else {
      // MUST BE -
      //	Guess is high. raisedToPower is higher
      //	than Alpha

      lastHighGuess.Set(guess)

      lastHighRaisedToPower.Set(raisedToPower)

      absRadicandDelta.Sub(lastHighRaisedToPower, radicand)

      absPercentRadicandDelta.Quo(absRadicandDelta, radicand)

      deltaIncrement.Set(half)

      if numOfHighGuesses > 2 {

        if absPercentRadicandDelta.Cmp(half) == 1 {

          deltaIncrement.Set(numberTwo)

        }

      }

      //fmt.Printf("lastHighGuess  = %v\n"+
      //	"lastHighRaisedToPower = %v\n",
      //	lastHighGuess.Text('f', 10),
      //	lastHighRaisedToPower.Text('f', 10))

      lastHighAbsAlphaDelta.Set(absRadicandDelta)

      tempCompare = lastLowGuess.Cmp(lastHighGuess)

      if tempCompare == 1 {
        // Low Guess is higher than High Guess
        absGuessDelta.Mul(lastHighGuess, half)
        guess.Sub(guess, absGuessDelta)

      } else if tempCompare == 0 {
        absGuessDelta.Mul(lastHighGuess, half)
        guess.Sub(guess, absGuessDelta)

      } else {
        // High Guess is higher than Low Guess
        absGuessDelta.Sub(lastHighGuess, lastLowGuess)
        absGuessDelta.Mul(absGuessDelta, half)
      }

      guess.Sub(guess, absGuessDelta)

      //fmt.Printf("New Guess = %v\n",
      //	guess.Text('f', -1))

      numOfHighGuesses++
    }

    cycleCount++

    if compareResult == -1 &&
      absRadicandDelta.Cmp(accuracyThreshold) < 1 {

      guess.Set(lastLowGuess)
      raisedToPower.Set(lastLowRaisedToPower)

      closeEnough = true
    }

    if cycleCount > 10000 {

      if compareResult == 1 {
        guess.Set(lastLowGuess)
        raisedToPower.Set(lastLowRaisedToPower)
        absRadicandDelta.Sub(radicand, raisedToPower)

      }

      closeEnough = true

    }

  }

  if raisedToPower == nil {

    err = fmt.Errorf("%v\n"+
      "raisedToPower is nil\n", ePrefix)

    return guess,
      raisedToPower,
      absRadicandDelta,
      compareResult,
      cycleCount,
      err
  }

  compareResult = raisedToPower.Cmp(radicand)

  absRadicandDelta.Sub(raisedToPower, radicand)

  if absRadicandDelta.Sign() == -1 {
    absRadicandDelta.Neg(absRadicandDelta)
  }

  return guess,
    raisedToPower,
    absRadicandDelta,
    compareResult,
    cycleCount,
    err
}

func BigFloatPower01(
  base *big.Float,
  power int64,
  prec uint) (
  raisedToPower *big.Float) {

  raisedToPower =
    new(big.Float).
      SetPrec(prec).
      SetMode(big.AwayFromZero)

  var ok bool
  _,
    ok = raisedToPower.SetString("1.0")

  if !ok {
    panic("raisedToPower.SetString(\"1.0\") Failed!")
  }

  for i := int64(0); i < power; i++ {

    raisedToPower.Mul(raisedToPower, base)
  }

  return raisedToPower
}

// ComputeBigFloatPrecisionBits
//
//	This function computes the number of bits required to represent a
//	floating-point number with the specified number of decimal digits.
//
//	The function uses the following formula to compute the number of bits:
//
//	baseBits = decDigits * log2(10) + safety
//
//	where:
//
//	decDigits - The number of decimal digits to be represented.
//
//	log2(10) - The base-2 logarithm of 10.
//
//	safety - A safety margin for exponentiation.
//
//	Important:
//
//	Remember, You must compute the total decimal digits (decDigits) in
//	the floating point number to be supported with the computed and
//	returned precision value.
//
//	Input Parameters:
//
//	resultDecDigits - The number of decimal digits to be represented.
//
//	multiplyCount - The number of times the base number multiplies by
//	itself.
func ComputeBigFloatPrecisionBits(resultDecDigits uint, multiplyCount uint) uint {

  // Convert decimal digits → bits
  // baseBits := float64(decDigits) * math.Log2(10)
  //
  // math.Log2(10) =
  //  3.321 928 094 887 362 347 870 319 429 489 390 175 864 831 393 024 6
  // float64 maximum 15-digits of precision
  baseBits := float64(resultDecDigits) * 3.321928094887362

  // Add safety margin for exponentiation
  safety := 64.0 + (float64(multiplyCount) * 4.0)

  // Round up to nearest whole bit
  return uint(math.Ceil(baseBits + safety))
}

// ComputeBigFloatDecimalDigits
//
//	Calculates the number of decimal digits represented
//	by a BigFloat with the given bit precision.
func ComputeBigFloatDecimalDigits(
  precisionBits uint,
  digitsSafetyMargin uint) uint {

  bFloatRoundValue := new(big.Float).
    SetMode(big.AwayFromZero).
    SetFloat64(0.5)

  bFloatDigitSafetyMargin := new(big.Float).
    SetMode(big.AwayFromZero).
    SetUint64(uint64(digitsSafetyMargin))

  bFloatDigitSafetyMargin.Add(bFloatDigitSafetyMargin, bFloatRoundValue)

  // math.Log2(10) =
  //  3.321 928 094 887 362 347 870 319 429 489 390 175 864 831 393 024 6
  // float64 maximum 15-digits of precision
  bFloatLog210 := new(big.Float).
    SetMode(big.AwayFromZero).
    SetFloat64(3.321928094887362)

  bFloatPrecisionBits := new(big.Float).
    SetMode(big.AwayFromZero).
    SetUint64(uint64(precisionBits))

  bFloatDecDigits := new(big.Float).Quo(bFloatPrecisionBits, bFloatLog210)

  bFloatDecDigits.Add(bFloatDecDigits, bFloatDigitSafetyMargin)

  uint64NumOfDigits, _ := bFloatDecDigits.Uint64()

  return uint(uint64NumOfDigits)
}

// DurationBreakdown returns a formatted string describing the
// elapsed time between start and end in minutes, seconds,
// milliseconds, microseconds, and nanoseconds.
func DurationBreakdown(start, end time.Time) string {
  d := end.Sub(start)

  // Extract components
  minutes := d / time.Minute
  d -= minutes * time.Minute

  seconds := d / time.Second
  d -= seconds * time.Second

  milliseconds := d / time.Millisecond
  d -= milliseconds * time.Millisecond

  microseconds := d / time.Microsecond
  d -= microseconds * time.Microsecond

  nanoseconds := d // remaining

  return fmt.Sprintf(
    "%d min, %d sec, %d milliseconds, %d microseconds, %d nanoseconds",
    minutes, seconds, milliseconds, microseconds, nanoseconds,
  )
}
