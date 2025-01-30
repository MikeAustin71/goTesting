package datetime

import (
	"errors"
	"math/big"
	"sync"
	"time"
)

type julianDayNoDtoMechanics struct {

	lock *sync.Mutex
}

// empty - Receives a pointer to a type JulianDayNoDto,
// 'jDNDto'. The method then proceeds to set all internal
// member variables to their 'invalid' or uninitialized
// values.
//
// BE CAREFUL!
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  jDNDto             *JulianDayNoDto
//     - All data fields in this JulianDayNoDto instance will be set
//       to invalid invalid values. Data field pointers will be set
//       to 'nil'.
//
//  ePrefix            string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
//
// ------------------------------------------------------------------------
//
// Return Value
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note: This error
//       message will incorporate the text passed by the Error Prefix
//       input parameter, 'ePrefix'.
//
func (jDNMech *julianDayNoDtoMechanics) empty(
	jDNDto *JulianDayNoDto,
	ePrefix string) error {

	if jDNMech.lock == nil {
		jDNMech.lock = new(sync.Mutex)
	}

	jDNMech.lock.Lock()

	defer jDNMech.lock.Unlock()

	ePrefix += "julianDayNoDtoMechanics) empty("

	if jDNDto == nil {
		return errors.New(ePrefix + "\n" +
			"Input parameter 'jDNDto' is a 'nil' pointer!\n")
	}

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.julianDayNo = nil
	jDNDto.julianDayNoFraction = nil
	jDNDto.julianDayNoTime = nil
	jDNDto.julianDayNoNumericalSign = 0
	jDNDto.totalJulianNanoSeconds = -1
	jDNDto.netGregorianNanoSeconds = -1
	jDNDto.hasLeapSecond = false
	jDNDto.hours = -1
	jDNDto.minutes = -1
	jDNDto.seconds = -1
	jDNDto.nanoseconds = -1
	jDNDto.isThisInstanceValid = false

	return nil
}

// exchangeValues - Receives pointers to two JulianDayNoDto objects
// and proceeds to exchange the data values of all internal member
// variables.
//
// If the method completes successfully, input parameter 'jDNDtoOne'
// will be populated with the data values from input parameter
// 'jDNDtoTwo'. Likewise, 'jDNDtoTwo' will be populated with
// the values copied from 'jDNDtoOne'.
//
// If either of the two JulianDayNoDto instances are judged invalid, this
// method will return an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  jDNDtoOne                *JulianDayNoDto
//     - A pointer to an instance of JulianDayNoDto. This is one of the
//       two JulianDayNoDto objects used in the data exchange. Data values
//       from this instance will be copied to input parameter 'jDNDtoTwo'.
//       Likewise, the data values from 'jDNDtoTwo' to this input
//       parameter, 'jDNDtoOne'.
//
//       If 'jDNDtoOne' is an invalid instance, an error will
//       be returned.
//
//
//  jDNDtoTwo                *JulianDayNoDto
//     - A pointer to an instance of JulianDayNoDto. This is one of the
//       two JulianDayNoDto objects used in the data exchange. Data values
//       from this instance will be copied to input parameter 'jDNDtoOne'.
//       Likewise, data values from 'jDNDtoOne' will be copied to this
//      instance, 'jDNDtoTwo'.
//
//       If 'jDNDtoTwo' is an invalid instance, an error will
//       be returned.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (jDNMech *julianDayNoDtoMechanics) exchangeValues(
	jDNDtoOne *JulianDayNoDto,
	jDNDtoTwo *JulianDayNoDto,
	ePrefix string) (err error) {

	if jDNMech.lock == nil {
		jDNMech.lock = new(sync.Mutex)
	}

	jDNMech.lock.Lock()

	defer jDNMech.lock.Unlock()

	ePrefix += "julianDayNoDtoMechanics.rationalizeJulianDayNoDto() "

	if jDNDtoOne == nil {
		err = errors.New(ePrefix + "\n" +
			"Input parameter 'jDNDtoOne' is a 'nil' pointer!\n")

		return err
	}

	if jDNDtoOne.lock == nil {
		jDNDtoOne.lock = new(sync.Mutex)
	}

	if jDNDtoTwo == nil {
		err = errors.New(ePrefix + "\n" +
			"Input parameter 'jDNDtoTwo' is a 'nil' pointer!\n")

		return err
	}

	if jDNDtoTwo.lock == nil {
		jDNDtoTwo.lock = new(sync.Mutex)
	}

	jDNNanobot := julianDayNoNanobot{}

	_,
	err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDtoOne,
		ePrefix +
			"Testing Input 'jDNDtoOne' validity. ")

	if err != nil {
		return err
	}

	_,
	err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDtoTwo,
		ePrefix +
			"Testing Input 'jDNDtoTwo' validity. ")

	if err != nil {
		return err
	}


	tempJDNDto := JulianDayNoDto{}

	tempJDNDto.julianDayNo =
		big.NewInt(0).
			Set(jDNDtoOne.julianDayNo)

	tempJDNDto.julianDayNoFraction =
		big.NewFloat(0.0).
			Set(jDNDtoOne.julianDayNoFraction)

	tempJDNDto.julianDayNoTime =
		big.NewFloat(0.0).
			Set(jDNDtoOne.julianDayNoTime)

	tempJDNDto.julianDayNoNumericalSign =
		jDNDtoOne.julianDayNoNumericalSign

	tempJDNDto.totalJulianNanoSeconds =
		jDNDtoOne.totalJulianNanoSeconds

	tempJDNDto.netGregorianNanoSeconds =
		jDNDtoOne.netGregorianNanoSeconds

	tempJDNDto.hasLeapSecond =
		jDNDtoOne.hasLeapSecond

	tempJDNDto.hours =
		jDNDtoOne.hours

	tempJDNDto.minutes =
		jDNDtoOne.minutes

	tempJDNDto.seconds =
		jDNDtoOne.seconds

	tempJDNDto.nanoseconds =
		jDNDtoOne.nanoseconds

	tempJDNDto.isThisInstanceValid =
		jDNDtoOne.isThisInstanceValid

	jDNDtoOne.julianDayNo =
		big.NewInt(0).
			Set(jDNDtoTwo.julianDayNo)

	jDNDtoOne.julianDayNoFraction =
		big.NewFloat(0.0).
			Set(jDNDtoTwo.julianDayNoFraction)

	jDNDtoOne.julianDayNoTime =
		big.NewFloat(0.0).
			Set(jDNDtoTwo.julianDayNoTime)

	jDNDtoOne.julianDayNoNumericalSign =
		jDNDtoTwo.julianDayNoNumericalSign

	jDNDtoOne.totalJulianNanoSeconds =
		jDNDtoTwo.totalJulianNanoSeconds

	jDNDtoOne.netGregorianNanoSeconds =
		jDNDtoTwo.netGregorianNanoSeconds

	jDNDtoOne.hasLeapSecond =
		jDNDtoTwo.hasLeapSecond

	jDNDtoOne.hours =
		jDNDtoTwo.hours

	jDNDtoOne.minutes =
		jDNDtoTwo.minutes

	jDNDtoOne.seconds =
		jDNDtoTwo.seconds

	jDNDtoOne.nanoseconds =
		jDNDtoTwo.nanoseconds

	jDNDtoOne.isThisInstanceValid =
		jDNDtoTwo.isThisInstanceValid

	jDNDtoTwo.julianDayNo =
		big.NewInt(0).
			Set(tempJDNDto.julianDayNo)

	jDNDtoTwo.julianDayNoFraction =
		big.NewFloat(0.0).
			Set(tempJDNDto.julianDayNoFraction)

	jDNDtoTwo.julianDayNoTime =
		big.NewFloat(0.0).
			Set(tempJDNDto.julianDayNoTime)

	jDNDtoTwo.julianDayNoNumericalSign =
		tempJDNDto.julianDayNoNumericalSign

	jDNDtoTwo.totalJulianNanoSeconds =
		tempJDNDto.totalJulianNanoSeconds

	jDNDtoTwo.netGregorianNanoSeconds =
		tempJDNDto.netGregorianNanoSeconds

	jDNDtoTwo.hasLeapSecond =
		tempJDNDto.hasLeapSecond

	jDNDtoTwo.hours =
		tempJDNDto.hours

	jDNDtoTwo.minutes =
		tempJDNDto.minutes

	jDNDtoTwo.seconds =
		tempJDNDto.seconds

	jDNDtoTwo.nanoseconds =
		tempJDNDto.nanoseconds

	jDNDtoTwo.isThisInstanceValid =
		tempJDNDto.isThisInstanceValid


	_,
		err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDtoOne,
		ePrefix +
			"Testing Output 'jDNDtoOne' validity. ")

	if err != nil {
		return err
	}

	_,
		err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDtoTwo,
		ePrefix +
			"Testing Output 'jDNDtoTwo' validity. ")


	return err
}

// rationalizeJulianDayNoDto - Receives a pointer to a
// JulianDayNoDto instance, 'JulianDayNoDto'. The method
// will then test internal Big.Int and Big.Float pointers
// to ensure that all such pointers are valid. Invalid
// pointers are initialized to new objects with a zero value.
//
// If 'jDNDto' is nil, this method will return an error.
//
func (jDNMech *julianDayNoDtoMechanics) rationalizeJulianDayNoDto(
	jDNDto *JulianDayNoDto,
	ePrefix string) error {

	if jDNMech.lock == nil {
		jDNMech.lock = new(sync.Mutex)
	}

	jDNMech.lock.Lock()

	defer jDNMech.lock.Unlock()

	ePrefix += "julianDayNoDtoMechanics.rationalizeJulianDayNoDto() "

	if jDNDto == nil {
		return errors.New(ePrefix + "\n" +
			"Input parameter 'jDNDto' is a 'nil' pointer!\n")
	}

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	if jDNDto.julianDayNo == nil {
		jDNDto.julianDayNo = big.NewInt(0)
	}

	if jDNDto.julianDayNoFraction == nil {
		jDNDto.julianDayNoFraction = big.NewFloat(0.0)
	}

	if jDNDto.julianDayNoTime == nil {
		jDNDto.julianDayNoTime = big.NewFloat(0.0)
	}

	if jDNDto.julianDayNoNumericalSign != 1 &&
		jDNDto.julianDayNoNumericalSign != -1 {
		jDNDto.julianDayNoNumericalSign = 1
	}

	if jDNDto.totalJulianNanoSeconds < 0 {
		jDNDto.totalJulianNanoSeconds = 0
	}

	if jDNDto.totalJulianNanoSeconds < 0 {
		jDNDto.totalJulianNanoSeconds = 0
	}

	if jDNDto.netGregorianNanoSeconds < 0 {
		jDNDto.netGregorianNanoSeconds = 0
	}

	if jDNDto.hours < 0 {
		jDNDto.hours = 0
	}

	if jDNDto.minutes < 0 {
		jDNDto.minutes = 0
	}

	if jDNDto.seconds < 0 {
		jDNDto.seconds = 0
	}

	if jDNDto.nanoseconds < 0 {
		jDNDto.nanoseconds = 0
	}

	jDNDto.isThisInstanceValid = true

	return nil
}

// setBigValDto - Receives an instance of JulianDayNoDto and
// proceeds to compute and populate its data fields using
// a Julian Day Number integer of type *big.Int and a Time
// Fraction of type *big.Float.
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
func (jDNMech *julianDayNoDtoMechanics) setBigValDto(
	jDNDto *JulianDayNoDto,
	julianDayNoTime *big.Float,
	requestedPrecision uint,
	hasLeapSecond bool,
	ePrefix string) (err error) {

	if jDNMech.lock == nil {
		jDNMech.lock = new(sync.Mutex)
	}

	jDNMech.lock.Lock()

	defer jDNMech.lock.Unlock()

	err = nil

	ePrefix += "julianDayNoDtoMechanics.setBigValDto() "

	if jDNDto == nil {
		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "jDNDto",
			inputParameterValue: "",
			errMsg:              "Input parameter 'jDNDto' " +
				"is a nil pointer!",
			err:                 nil,
		}

		return err
	}

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	if julianDayNoTime == nil {

		err = &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "julianDayNoTime",
			inputParameterValue: "",
			errMsg:              "Error: 'julianDayNoTime' is a 'nil' pointer!",
			err:                 nil,
		}

		return err
	}

	if julianDayNoTime.Prec() > requestedPrecision {
		requestedPrecision = julianDayNoTime.Prec()
	}

	mathBFMech := MathBigFloatHelper{}

	jDNDto.julianDayNoTime,
	jDNDto.julianDayNoNumericalSign =
	 mathBFMech.Abs(julianDayNoTime, requestedPrecision)

	jDNDto.julianDayNo =
		mathBFMech.GetBigIntValue(jDNDto.julianDayNoTime)

	jDNDto.julianDayNoFraction,
		_ =
	mathBFMech.FloatFractionalValue(
		jDNDto.julianDayNoTime,
		requestedPrecision)

	twentyFourHourNanoseconds := int64(time.Hour) * 24

	if hasLeapSecond {
		twentyFourHourNanoseconds += int64(time.Second)
	}

	bigDayNanoSeconds := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		SetInt64(twentyFourHourNanoseconds)

	bigDayNanoSeconds = mathBFMech.RoundHalfAwayFromZero(
		bigDayNanoSeconds,
		requestedPrecision,
		0)

	grossNanoSecs := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		Mul(bigDayNanoSeconds, jDNDto.julianDayNoFraction)

	grossNanoSecs.
		SetMode(big.ToNearestAway).
		SetPrec(requestedPrecision).
		Add(big.NewFloat(0.5), grossNanoSecs)

	// Always less than or equal to 36-hours
	jDNDto.totalJulianNanoSeconds, _ = grossNanoSecs.Int64()

	// Always less than or equal to 24-hours
	jDNDto.netGregorianNanoSeconds = jDNDto.totalJulianNanoSeconds

	noonNanoSeconds := int64(time.Hour) * 12

	if jDNDto.netGregorianNanoSeconds >= noonNanoSeconds {
		jDNDto.netGregorianNanoSeconds -= noonNanoSeconds
	} else {
		jDNDto.netGregorianNanoSeconds += noonNanoSeconds
	}

	timeMech := TimeMechanics{}

	jDNDto.hours,
		jDNDto.minutes,
		jDNDto.seconds,
		jDNDto.nanoseconds,
		_ = timeMech.ComputeTimeElementsInt64(
		jDNDto.netGregorianNanoSeconds)

	jDNDto.hasLeapSecond = hasLeapSecond

	jDNNanobot := julianDayNoNanobot{}

	_,
	err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix)

	return  err
}

// setJulianDayNoDtoToZero - Sets the data fields of input parameter
// 'jDNDto' to their zero value. This differs from method
// julianDayNoDtoMechanics.empty() which sets all data fields to
// 'invalid' values. In contrast, this method,
// julianDayNoDtoMechanics.setJulianDayNoDtoToZero(), sets data fields
// to valid zero values. Essentially,'jDNDto' is converted to a valid
// JulianDayNoDto with all zero value.
//
// ------------------------------------------------------------------------
//
// Input Value
//
//  ePrefix            string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  JulianDayNoDto
//     - If successful, this method will return a deep copy of the current
//       JulianDayNoDto instance.
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (jDNMech *julianDayNoDtoMechanics) setJulianDayNoDtoToZero(
	jDNDto *JulianDayNoDto,
	ePrefix string) error {

	if jDNMech.lock == nil {
		jDNMech.lock = new(sync.Mutex)
	}

	jDNMech.lock.Lock()

	defer jDNMech.lock.Unlock()

	ePrefix += "julianDayNoDtoMechanics.rationalizeJulianDayNoDto() "

	if jDNDto == nil {
		return errors.New(ePrefix + "\n" +
			"Input parameter 'jDNDto' is a 'nil' pointer!\n")
	}

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.julianDayNo = big.NewInt(0)
	jDNDto.julianDayNoFraction = big.NewFloat(0.0)
	jDNDto.julianDayNoTime = big.NewFloat(0.0)
	jDNDto.julianDayNoNumericalSign = 1
	jDNDto.totalJulianNanoSeconds = 0
	jDNDto.netGregorianNanoSeconds = 0
	jDNDto.hasLeapSecond = false
	jDNDto.hours = 0
	jDNDto.minutes = 0
	jDNDto.seconds = 0
	jDNDto.nanoseconds = 0
	jDNDto.isThisInstanceValid = true

	return nil
}