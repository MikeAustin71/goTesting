package naturalLogCalcs

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

type BigFloatHelper struct{}

// TODO: Unfinished. FIX THIS!
func (bFloatHlpr *BigFloatHelper) EulersNoToPower(exponent uint) *big.Float {

	nlSharedMech := new(naturalLogSharedMechanics)

	eulersNumber := nlSharedMech.getEulersNum(0)

	basePrec := eulersNumber.Prec()

	bFloatResult := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(basePrec + 1024).
		Set(eulersNumber)

	bigFloatExponent := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(basePrec).
		SetUint64(uint64(exponent))

	for i := uint(1); i < exponent; i++ {

		bFloatResult.Mul(bFloatResult, bigFloatExponent)
	}

	return bFloatResult
}

// ComputeBigFloatPrecisionBits ...
func (bFloatHlpr *BigFloatHelper) ComputeBigFloatPrecisionBits(resultDecDigits uint, multiplyCount uint) uint {

	baseBits := float64(resultDecDigits) * 3.321928094887362
	safety := 64.0 + (float64(multiplyCount) * 4.0)

	return uint(math.Ceil(baseBits + safety))
}

// ComputeBigFloatDecimalDigits ...
func (bFloatHlpr *BigFloatHelper) ComputeBigFloatDecimalDigits(
	precisionBits uint,
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
		SetUint64(uint64(precisionBits))

	bFloatDecDigits := new(big.Float).Quo(bFloatPrecisionBits, bFloatLog210)
	bFloatDecDigits.Add(bFloatDecDigits, bFloatDigitSafetyMargin)

	uint64NumOfDigits, _ := bFloatDecDigits.Uint64()

	return uint(uint64NumOfDigits)
}

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
