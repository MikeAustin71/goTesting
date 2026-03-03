package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {

	return
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

func NewtonInitialGuess01(
	nthRoot int64, // 18
	alpha *big.Float, // 592
	accuracyThreshold *big.Float,
	prec uint) (
	guess *big.Float,
	raisedToPower *big.Float,
	absAlphaDelta *big.Float,
	compareResult int,
	cycleCount uint64) {

	// SetMode sets z's rounding mode to mode and returns an exact z.
	// z remains unchanged otherwise.
	// z.SetMode(z.Mode()) is a cheap way to set z's accuracy to Exact.

	numOfLowGuesses := uint64(0)
	numOfHighGuesses := uint64(0)

	absAlphaDelta = new(big.Float).
		SetInt64(0).
		SetPrec(prec).
		SetMode(big.AwayFromZero)

	absPercentAlphaDelta := new(big.Float).
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

			absAlphaDelta.Sub(alpha, raisedToPower)

			// On target
			return guess,
				raisedToPower,
				absAlphaDelta,
				compareResult,
				cycleCount

		} else if compareResult == -1 {
			// Guess is low. Alpha is higher than
			// raisedToPower

			lastLowGuess.Set(guess)

			lastLowRaisedToPower.Set(raisedToPower)

			absAlphaDelta.Sub(alpha, lastLowRaisedToPower)

			absPercentAlphaDelta.Quo(absAlphaDelta, alpha)

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

				if absPercentAlphaDelta.Cmp(half) == 1 {

					deltaIncrement.Set(numberTwo)

				}

			}

			lastLowAbsAlphaDelta.Set(absAlphaDelta)

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

			absAlphaDelta.Sub(lastHighRaisedToPower, alpha)

			absPercentAlphaDelta.Quo(absAlphaDelta, alpha)

			deltaIncrement.Set(half)

			if numOfHighGuesses > 2 {

				if absPercentAlphaDelta.Cmp(half) == 1 {

					deltaIncrement.Set(numberTwo)

				}

			}

			//fmt.Printf("lastHighGuess  = %v\n"+
			//	"lastHighRaisedToPower = %v\n",
			//	lastHighGuess.Text('f', 10),
			//	lastHighRaisedToPower.Text('f', 10))

			lastHighAbsAlphaDelta.Set(absAlphaDelta)

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
			absAlphaDelta.Cmp(accuracyThreshold) < 1 {

			guess.Set(lastLowGuess)
			raisedToPower.Set(lastLowRaisedToPower)

			closeEnough = true
		}

		if cycleCount > 10000 {

			if compareResult == 1 {
				guess.Set(lastLowGuess)
				raisedToPower.Set(lastLowRaisedToPower)
				absAlphaDelta.Sub(alpha, raisedToPower)

			}

			closeEnough = true

		}

	}

	compareResult = raisedToPower.Cmp(alpha)

	absAlphaDelta.Sub(raisedToPower, alpha)

	if absAlphaDelta.Sign() == -1 {
		absAlphaDelta.Neg(absAlphaDelta)
	}

	return guess,
		raisedToPower,
		absAlphaDelta,
		compareResult,
		cycleCount
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
