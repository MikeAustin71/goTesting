package naturalLogCalcs

import (
	"math"
	"math/big"
)

var (
	lenEulersNumberStr              int
	lenPiNumberStr                  int
	lenNatLogNumberStr              int
	eulersNumBaseFull               *big.Float
	piFull                          *big.Float
	ln2Full                         *big.Float
	DEFAULT_BASE_NUM_DECIMAL_DIGITS uint
)

func init() {

	lenEulersNumberStr = len(eulersConstNumStr)

	lenPiNumberStr = len(pi20kDigitStr)

	lenNatLogNumberStr = len(natLog2Str20kDigits)

	DEFAULT_BASE_NUM_DECIMAL_DIGITS = 0

	var constPrec uint

	var baseBits float64

	// Set eulers number with default Number of
	// decimal digits. 4000 decimal digits of accuracy.
	// Length of lenEulersNumberStr MUST BE >= 4000
	baseBits = (float64(4000) * 3.321928094887362) + 128.0

	constPrec = uint(math.Ceil(baseBits))

	eulersNumBaseFull = big.NewFloat(0)

	eulersNumBaseFull.
		SetMode(big.AwayFromZero).
		SetPrec(constPrec).
		SetString(eulersConstNumStr[:4002])

	// Set pi and natural log constants to 2000 decimal digits
	// of accuracy.
	baseBits = (float64(2000) * 3.321928094887362) + 128.0

	constPrec = uint(math.Ceil(baseBits))

	piFull = new(big.Float).
		SetPrec(constPrec).
		SetMode(big.AwayFromZero)
	piFull.SetString(pi20kDigitStr[:2002])

	ln2Full = new(big.Float).
		SetPrec(constPrec).
		SetMode(big.AwayFromZero)
	ln2Full.SetString(natLog2Str20kDigits[:2002])

}
