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
func (bFloatHlpr *BigFloatHelper) ComputeBigFloatPrecisionBits(resultDecDigits uint, multiplyCount uint) uint {

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
func (bFloatHlpr *BigFloatHelper) DurationBreakdown(start, end time.Time) string {
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
