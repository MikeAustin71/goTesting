package datetime

import (
	"errors"
	"math/big"
	"sync"
)

type mathBigFloatUtility struct {
	lock *sync.Mutex
}

// getFloatNumText - Receives format specifications for converting a
// *big.Float floating point number to a text string. The text string
// is returned to the caller along with calculated parameters necessary
// for console display or further text formatting.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  bigFloatNum              *big.Float
//     - The floating point number which will be formatted as a string of
//       text.
//
//
//  numberFieldLength        int
//     - The length of the number field in which the floating point text
//       string should be right justified. If this is less than the actual
//       length of the floating point text, an alternate field length will
//       be calculated.
//
//
//  precision                uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on input parameter 'bigFloatNum'
//       as well as specifying the numeric precision incorporated in the
//       the output text string. 'precision' should not be confused with
//       parameter 'roundToDecPlaces'. The term 'precision' applies to
//       the internal accuracy maintained by type *big.Float floating
//       point values.  For more information on precision and type
//       *big.Float floating point numbers, reference:
//           https://golang.org/pkg/math/big/
//
//
//  roundToDecPlaces  uint
//     - This parameter specifies the number of digits to the right of the
//       decimal place which will be contained in the returned text string.
//
//
//  ePrefix                  string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//  bigFloatTextDto          MathFloatTextDto
//     - If this method completes successfully, it will populate and
//       return a data structure containing the formatting information
//       and text string necessary to present the input parameter
//       'bigFloatNum' as a text string. The actual text string is
//       contained in the last data field, 'MathFloatTextDto.FloatTextStr'
//
//       type MathFloatTextDto struct {
//          MinimumFloatFieldLength  int
//          ActualFloatFieldLength   int
//          RoundToDecimalPlaces     uint
//          FloatPrecision           uint
//          IntegerNumLength         int
//          FractionalNumLength      int
//          NumberSign               int
//          OriginalFloatNum         *big.Float
//          RoundedFloatNum          *big.Float
//          FloatFmtStr              string
//          FloatTextStr             string
//          IsValidInstanceError          bool
//       }
//
//
//  err                      error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message. Note this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//
func (bigFloatUtil *mathBigFloatUtility) getFloatNumText(
	bigFloatNum *big.Float,
	numberFieldLength int,
	precision uint,
	roundToDecPlaces uint,
	ePrefix string) (
	bigFloatTextDto MathFloatTextDto,
	err error) {

	if bigFloatUtil.lock == nil {
		bigFloatUtil.lock = new(sync.Mutex)
	}

	bigFloatUtil.lock.Lock()

	defer bigFloatUtil.lock.Unlock()

	ePrefix += "mathBigFloatUtility.GetFloatNumText() "

	bigFloatTextDto = MathFloatTextDto{}
	err = nil

	if bigFloatNum == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'bigFloatNum' is a 'nil' pointer!\n")
		return bigFloatTextDto, err
	}

	return bigFloatTextDto, err

}