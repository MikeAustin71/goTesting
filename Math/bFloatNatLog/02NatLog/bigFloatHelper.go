package naturalLogCalcs

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

type BigFloatHelper struct{}

// ComputeBigFloatPrecisionBits
//
//	Calculates the big.Float bit precision needed to achieve the
//	desired number of decimal digits of accuracy.
func (bFloatHlpr *BigFloatHelper) ComputeBigFloatPrecisionBits(resultDecimalDigits uint, multiplyCount uint) uint {

	baseBits := float64(resultDecimalDigits) * 3.321928094887362

	safety := 64.0 + (float64(multiplyCount) * 4.0)

	return uint(math.Ceil(baseBits + safety))
}

// ComputeBigFloatDecimalDigits
//
//	Calculates the number of decimal digits achievable for a
//	given big.Float precision value in bits.
func (bFloatHlpr *BigFloatHelper) ComputeBigFloatDecimalDigits(
	bigFloatPrecisionBits uint,
	digitsSafetyMargin uint) uint {

	bFloatRoundValue := new(big.Float).
		SetMode(big.AwayFromZero).
		SetFloat64(0.5)

	bFloatDigitSafetyMargin := new(big.Float).
		SetMode(big.AwayFromZero).
		SetUint64(uint64(digitsSafetyMargin))

	bFloatDigitSafetyMargin.Add(bFloatDigitSafetyMargin, bFloatRoundValue)

	bFloatLog210 := new(big.Float).
		SetMode(big.AwayFromZero).
		SetFloat64(3.321928094887362)

	bFloatPrecisionBits := new(big.Float).
		SetMode(big.AwayFromZero).
		SetUint64(uint64(bigFloatPrecisionBits))

	bFloatDecimalDigits :=
		new(big.Float).Quo(bFloatPrecisionBits, bFloatLog210)

	bFloatDecimalDigits.
		Add(bFloatDecimalDigits, bFloatDigitSafetyMargin)

	uint64NumOfDecimalDigits, _ := bFloatDecimalDigits.Uint64()

	return uint(uint64NumOfDecimalDigits)
}

func (bFloatHlpr *BigFloatHelper) CleanNumberString(rawNumberStr string) string {

	var cleanRuneArray = make([]rune, 0, len(rawNumberStr))

	for _, v := range rawNumberStr {

		if v >= '0' && v <= '9' {

			cleanRuneArray = append(cleanRuneArray, v)
		}

		if v == '.' || v == '-' || v == '+' {
			cleanRuneArray = append(cleanRuneArray, v)
		}

	}

	return string(cleanRuneArray)
}

// CountDigits
//
//	Takes a string representation of a number and a decimal
//	separator and counts the integer and decimal digits contained
//	in input parameter 'numStr'.
func (bFloatHlpr *BigFloatHelper) CountDigits(numStr string, decimalSeparator rune) (intDigits int, decDigits int) {

	intDigits = 0
	decDigits = 0

	var isInt bool

	isInt = true

	for _, v := range numStr {

		if v == decimalSeparator {
			isInt = false
			continue
		}

		if v >= '0' && v <= '9' {

			if isInt {
				// Integer digit to left of decimal point
				intDigits++
			} else {
				// Decimal digit to right of decimal point
				decDigits++
			}
		}

	}

	return intDigits, decDigits
}

// DurationBreakdown ...
func (bFloatHlpr *BigFloatHelper) DurationBreakdown(start, end time.Time) string {
	d := end.Sub(start)

	minutes := d / time.Minute
	d -= minutes * time.Minute

	seconds := d / time.Second
	d -= seconds * time.Second

	milliseconds := d / time.Millisecond
	d -= milliseconds * time.Millisecond

	microseconds := d / time.Microsecond
	d -= microseconds * time.Microsecond

	nanoseconds := d

	return fmt.Sprintf(
		"%d min, %d sec, %d milliseconds, %d microseconds, %d nanoseconds",
		minutes, seconds, milliseconds, microseconds, nanoseconds,
	)
}
