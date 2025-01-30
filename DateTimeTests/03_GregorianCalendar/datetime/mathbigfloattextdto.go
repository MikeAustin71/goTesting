package datetime

import "math/big"

type MathFloatTextDto struct {
	MinimumFloatFieldLength  int         // Minimum number of spaces necessary to accommodate
	//                                   //   display of floating point number 'RoundedFloatNum'.
	ActualFloatFieldLength   int         // Actual length of number of spaces used to construct
	//                                   //   the numeric text field for display of 'RoundedFloatNum'.
	RoundToDecimalPlaces     uint        // The number of digits to the right of the decimal which
	//                                   //   will remain in 'RoundedFloatNum' after rounding.
	FloatPrecision           uint        // 'precision' value used for internal floating point
	//                                   //   calculations. See https://golang.org/pkg/math/big/
	IntegerNumLength         int         // The number of integer digits, excluding any leading number
	//                                   //   number signs, in the 'RoundedFloatNum' floating point value.
	FractionalNumLength      int         // The actual number of digits to the right of the decimal place.
	NumberSign               int         // One of three Values describing 'RoundedFloatNum':
	//                                   //   +1 - Signals that 'RoundedFloatNum' is a positive value greater than
	//                                   //        zero ('RoundedFloatNum' > 0 ).
	//                                   //    0 - Signals that 'RoundedFloatNum' has a zero value.
	//                                   //        ('RoundedFloatNum' == 0 ).
	//                                   //   -1 - Signals that 'RoundedFloatNum' is a negative value less than zero
	//                                   //        ('RoundedFloatNum' < 0 ).
	OriginalFloatNum         *big.Float  // The original floating point value before rounding
	RoundedFloatNum          *big.Float  // The final floating point value after rounding used for conversion
	//                                   //   to a text string.
	FloatFmtStr              string      // The printf format which may be used to display 'RoundedFloatNum'
	//                                   //   in accordance with 'ActualFloatFieldLength'.
	FloatTextStr             string      // The text string containing the final formatted value of 'RoundedFloatNum'
	IsValidInstance          bool        // If set to 'true', this signals that all data fields in this data
	//                                   //   structure have been properly initialized and are valid. A value of false
	//                                   //   says the data contained in these data fields is invalid.
}
