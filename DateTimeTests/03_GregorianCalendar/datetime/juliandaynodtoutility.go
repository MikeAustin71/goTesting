package datetime

import (
	"errors"
	"math/big"
	"sync"
)

type julianDayNoDtoUtility struct {

	lock     *sync.Mutex

}

// copyIn - Copies data from an incoming jDNDtoIncoming instance
// to the current jDNDtoIncoming instance, 'jDNDto'
//
func (jDNDtoUtil *julianDayNoDtoUtility) copyIn(
	jDNDto *JulianDayNoDto,
	jDNDtoIncoming *JulianDayNoDto,
	ePrefix string) error {

	if jDNDtoUtil.lock == nil {
		jDNDtoUtil.lock = new(sync.Mutex)
	}

	jDNDtoUtil.lock.Lock()

	defer jDNDtoUtil.lock.Unlock()

	ePrefix += "julianDayNoDtoUtility.copyIn() "

	if jDNDto == nil {
		return errors.New(ePrefix + "\n" +
		 	"Input parameter 'jDNDto' is a 'nil' pointer!\n")
	}

	if jDNDtoIncoming == nil {
		return errors.New(ePrefix + "\n" +
			"Input parameter 'jDNDtoIncoming' is a 'nil' pointer!\n")
	}

	if jDNDtoIncoming.lock == nil {
		jDNDtoIncoming.lock = new(sync.Mutex)
	}

	jDNNanobot := julianDayNoNanobot{}

	_, err := jDNNanobot.testJulianDayNoDtoValidity(
		jDNDtoIncoming,
		ePrefix + "- Checking Input Parameter 'jDNDtoIncoming' Validity ")

	if err != nil {
		return err
	}

	jDNMech := julianDayNoDtoMechanics{}

	err = jDNMech.empty(jDNDto, ePrefix)

	if err != nil {
		return err
	}

	jDNDto.julianDayNo =
		big.NewInt(0).
			Set(jDNDtoIncoming.julianDayNo)

	jDNDto.julianDayNoFraction =
		big.NewFloat(0.0).
			Set(jDNDtoIncoming.julianDayNoFraction)

	jDNDto.julianDayNoTime =
		big.NewFloat(0.0).
			Set(jDNDtoIncoming.julianDayNoTime)

	jDNDto.julianDayNoNumericalSign =
		jDNDtoIncoming.julianDayNoNumericalSign

	jDNDto.totalJulianNanoSeconds =
		jDNDtoIncoming.totalJulianNanoSeconds

	jDNDto.netGregorianNanoSeconds =
		jDNDtoIncoming.netGregorianNanoSeconds

	jDNDto.hasLeapSecond =
		jDNDtoIncoming.hasLeapSecond

	jDNDto.hours =
		jDNDtoIncoming.hours

	jDNDto.minutes =
		jDNDtoIncoming.minutes

	jDNDto.seconds =
		jDNDtoIncoming.seconds

	jDNDto.nanoseconds =
		jDNDtoIncoming.nanoseconds

	_,
	err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Checking Output 'jDNDto' Validity ")

	return nil
}

// copyOut - Returns a deep copy of input parameter
// 'jDNDto' which is a pointer to a type 'JulianDayNoDto'.
//
// If 'jDNDto' is nil, this method will panic.
//
func (jDNDtoUtil *julianDayNoDtoUtility) copyOut(
	jDNDto *JulianDayNoDto,
	ePrefix string) (JulianDayNoDto, error) {

	if jDNDtoUtil.lock == nil {
		jDNDtoUtil.lock = new(sync.Mutex)
	}

	jDNDtoUtil.lock.Lock()

	defer jDNDtoUtil.lock.Unlock()

	ePrefix += "julianDayNoDtoUtility.copyOut() "

	if jDNDto == nil {
		return JulianDayNoDto{},
		errors.New(ePrefix + "\n" +
			"Input parameter 'jDNDto' is a 'nil' pointer!\n")
	}

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNNanobot := julianDayNoNanobot{}

	_, err := jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing Input Parameter 'jDNDto' ")

	if err != nil {
		return JulianDayNoDto{}, err
	}

	newJDNDto := JulianDayNoDto{}

	newJDNDto.julianDayNo =
		big.NewInt(0).
			Set(jDNDto.julianDayNo)

	newJDNDto.julianDayNoFraction =
		big.NewFloat(0.0).
			Set(jDNDto.julianDayNoFraction)

	newJDNDto.julianDayNoTime =
		big.NewFloat(0.0).
			Set(jDNDto.julianDayNoTime)

	newJDNDto.julianDayNoNumericalSign =
		jDNDto.julianDayNoNumericalSign

	newJDNDto.totalJulianNanoSeconds =
		jDNDto.totalJulianNanoSeconds

	newJDNDto.netGregorianNanoSeconds =
		jDNDto.netGregorianNanoSeconds

	newJDNDto.hasLeapSecond =
		jDNDto.hasLeapSecond

	newJDNDto.hours =
		jDNDto.hours

	newJDNDto.minutes =
		jDNDto.minutes

	newJDNDto.seconds =
		jDNDto.seconds

	newJDNDto.nanoseconds =
		jDNDto.nanoseconds

	newJDNDto.lock =  new(sync.Mutex)

	newJDNDto.isThisInstanceValid = jDNDto.isThisInstanceValid

	_,
	err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing Output Parameter 'newJDNDto' ")

	if err != nil {
		return JulianDayNoDto{}, err
	}

	return newJDNDto, nil
}

// setDtoFromFloat64 - Receives a pointer to a type JulianDayNoDto
// and proceeds to compute an populate its Julian Day Number
// and Time data elements using a float64 input parameter.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  jDNDto            *JulianDayNoDto
//     - This is a pointer to a an instance Type /JulianDayNoDto'. This object
//       will be populated with Julian Day Number/Time information based on
//       the value of input parameter, 'julianDayNoTime'.
//
//
//  julianDayNoTime    float64
//     - This floating point value contains the integer Julian Day Number, the
//       Julian Day Number Time Fraction and the numerical sign ('+' or '-')
//       indicating whether the Julian Day Number is upstream or downstream
//       the base date/time. Be advised that the float64 value only provides
//       sufficient precision to accurately define seconds, and not nanoseconds,
//       of a designated time value.  To accurately identify time to the nanosecond
//       level of granularity, use methods which incorporate *big.Float floating
//       point values.
//
//
//  requestedPrecision        uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on the input parameters. The
//       'requestedPrecision' value applies to the internal accuracy
//        maintained by type *big.Float floating point values. Julian
//        Day Number/Time floating point values are stored in type
//        'JulianDayNoDto' as *big.Float objects.  For those seeking
//        to maximize accuracy. Try values of '512' or '1024'. For
//        more information on precision and type *big.Float floating
//        point numbers, reference:
//              https://golang.org/pkg/math/big/
//
//
//  hasLeapSecond           bool
//     - Set this boolean parameter to 'true' if, and ONLY if, the specified
//       day for which this Julian Day Number Time value will be calculated
//       includes a 'leap second'. A leap second is a one-second adjustment
//       that is occasionally applied to Coordinated Universal Time (UTC)
//       in order to accommodate the difference between precise time (as
//       measured by atomic clocks) and imprecise observed solar time
//       (known as UT1 and which varies due to irregularities and long-term
//       slowdown in the Earth's rotation). If this parameter is set to 'true',
//       the time calculation will assume the duration of the relevant 'day'
//       is 24-hours plus one second. Otherwise, the duration of a day is
//       assumed to consist of exactly 24-hours. For more information on
//       the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  ePrefix               string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (jDNDtoUtil *julianDayNoDtoUtility) setDtoFromFloat64(
	jDNDto *JulianDayNoDto,
	julianDayNoTime float64,
	requestedPrecision uint,
	applyLeapSecond bool,
	ePrefix string) (err error) {

	if jDNDtoUtil.lock == nil {
		jDNDtoUtil.lock = new(sync.Mutex)
	}

	jDNDtoUtil.lock.Lock()

	defer jDNDtoUtil.lock.Unlock()

	err = nil

	ePrefix += "julianDayNoDtoUtility.setDtoFromFloat64() "

	if jDNDto == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "jDNDto",
			inputParameterValue: "",
			errMsg: "Input parameter 'jDNDto' " +
				"is a 'nil' pointer!",
			err: nil,
		}
			return err
	}

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	julianBigFloatDayNoTime :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(requestedPrecision).
			SetFloat64(julianDayNoTime)

	jDNMech := julianDayNoDtoMechanics{}

	err = jDNMech.empty(jDNDto, ePrefix)

	if err != nil {
		return err
	}

	return jDNMech.setBigValDto(
		jDNDto,
		julianBigFloatDayNoTime,
		requestedPrecision,
		applyLeapSecond,
		ePrefix)
}

// setDto - Receives an instance of JulianDayNoDto and
// proceeds to compute and populate its data fields using
// a Julian Day Number integer and Time Fraction of type
// *big.Float.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  jDNDto                  *JulianDayNoDto
//     - This is a pointer to a an instance Type /JulianDayNoDto'. This object
//       will be populated with Julian Day Number/Time information based on
//       the value of input parameter, 'julianDayNoTime'.
//
//  julianDayNo              int64
//     - This integer value represents the Julian Day Number. The integer value
//       may be positive or negative and will dictate the numeric sign of the
//       final floating point Julian Day Number/Time value. The numerical sign
//       indicates whether the final Julian Day Number/Time value is upstream,
//       or downstream, from the base date/time.
//
//
//  julianDayNoTimeFraction  *big.Float
//     - This floating point value contains the fractional time value of the
//       Julian Day Number. The integer portion and the number sign of this
//       input parameter are ignored. Only the absolute fractional value is extracted
//       and combined with the input parameter 'julianDayNo' to form the composite
//       Julian Day Number/Time value.
//
//
//  requestedPrecision        uint
//     - This unsigned integer value will determine the precision of
//       internal calculations performed on the input parameters. The
//       'requestedPrecision' value applies to the internal accuracy
//        maintained by type *big.Float floating point values. Julian
//        Day Number/Time floating point values are stored in type
//        'JulianDayNoDto' as *big.Float objects.  For those seeking
//        to maximize accuracy. Try values of '512' or '1024'. For
//        more information on precision and type *big.Float floating
//        point numbers, reference:
//              https://golang.org/pkg/math/big/
//
//
//  hasLeapSecond           bool
//     - Set this boolean parameter to 'true' if, and ONLY if, the specified
//       day for which this Julian Day Number Time value will be calculated
//       includes a 'leap second'. A leap second is a one-second adjustment
//       that is occasionally applied to Coordinated Universal Time (UTC)
//       in order to accommodate the difference between precise time (as
//       measured by atomic clocks) and imprecise observed solar time
//       (known as UT1 and which varies due to irregularities and long-term
//       slowdown in the Earth's rotation). If this parameter is set to 'true',
//       the time calculation will assume the duration of the relevant 'day'
//       is 24-hours plus one second. Otherwise, the duration of a day is
//       assumed to consist of exactly 24-hours. For more information on
//       the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  ePrefix               string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (jDNDtoUtil *julianDayNoDtoUtility) setDto(
	jDNDto *JulianDayNoDto,
	julianDayNo int64,
	julianDayNoTimeFraction *big.Float,
	requestedPrecision uint,
	applyLeapSecond bool,
	ePrefix string) (err error) {

	if jDNDtoUtil.lock == nil {
		jDNDtoUtil.lock = new(sync.Mutex)
	}

	jDNDtoUtil.lock.Lock()

	defer jDNDtoUtil.lock.Unlock()

	err = nil

	ePrefix += "julianDayNoDtoUtility.setDto() "

	if jDNDto == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "jDNDto",
			inputParameterValue: "",
			errMsg:              "Input parameter 'jDNDto' " +
				"is a 'nil' pointer!",
			err:                 nil,
		}

			return err
	}

	mathBFMech := MathBigFloatHelper{}

	intNumSign := 1

	if julianDayNo < 0 {
		intNumSign = -1
	}

	julianDayNoTimeFraction, _ =
		mathBFMech.Abs(
			julianDayNoTimeFraction,
			requestedPrecision)

	julianDayNoTimeFraction, _ =
		mathBFMech.FloatFractionalValue(
			julianDayNoTimeFraction,
			requestedPrecision)

	julianDayNoFloat :=
		mathBFMech.CombineIntFracValues(
			big.NewInt(julianDayNo),
			julianDayNoTimeFraction,
			intNumSign,
			requestedPrecision)

	jDNMech := julianDayNoDtoMechanics{}

	return jDNMech.setBigValDto(
		jDNDto,
		julianDayNoFloat,
		requestedPrecision,
		applyLeapSecond,
		ePrefix)
}

