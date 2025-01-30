package datetime

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

// JulianDayNoDto - This type is used to transfer information
// on a Julian Day Number/Time.
//
// Julian day is the continuous count of days since the beginning
// of the Julian Period and is used primarily by astronomers, and
// in software for easily calculating elapsed days between two events
// (e.g. food production date and sell by date).
//
// ------------------------------------------------------------------------
//
// Background
//
// The Julian Day Number (JDN) is the integer assigned to a
// whole solar day in the Julian day count starting from noon
// Universal time, with Julian day number 0 assigned to the
// day starting at noon on Monday, January 1, 4713 BC, proleptic
// Julian calendar (November 24, 4714 BC, in the proleptic
// Gregorian calendar), a date which preceded any dates in
// recorded history. For example, the Julian day number for
// the day starting at 12:00 UT (noon) on January 1, 2000, was
// '2451545.0' on the Gregorian Calendar.
//
// The Julian date (JD), or Julian Day Number/Time, of any instant
// is the Julian day number/ plus the fraction of a day since the
// preceding noon in Universal Time. Julian dates are expressed as
// a Julian day number with a decimal fraction added. For example,
// the Julian Date for 00:30:00.0 UT January 1, 2013, is
// 2456293.52083 (Gregorian Calendar).
//
// The Julian day number is based on the Julian Period proposed
// by Joseph Scaliger, a classical scholar, in 1583 (one year after
// the Gregorian calendar reform) as it is the product of three
// calendar cycles used with the Julian calendar.
//
// Examples
//
// Gregorian Date -4713-11-24 12:00:00.000000000 +0000 UTC yields a
// Julian Day Number of '0.0'.
//
// ------------------------------------------------------------------------
//
// Technical Considerations
//
// The Julian Day Number Time is a floating point number with an integer
// to the left of the decimal point representing the Julian Day Number
// and the fraction to the right of the decimal point representing time
// in hours minutes and seconds.
//
// Julian Day numbers start on day zero at noon. This means that Julian
// Day Number Times are valid for all dates on or after noon on Monday,
// January 1, 4713 BCE, in the proleptic Julian calendar or November 24,
// 4714 BCE, in the proleptic Gregorian calendar. Remember that the Golang
// 'time.Time' type uses Astronomical Year numbering with the Gregorian
// Calendar. In other words, the 'time.Time' type recognizes the year
// zero. Dates expressed in the 'Common Era' ('BCE' Before Common Era
// or 'CE' Common Era). Therefore a 'time.Time' year of '-4713' is equal
// to the year '4714 BCE'
//
// The 'JulianDayNoDto' type provides a Julian Day Number/Time as a float64.
// This version of the Julian Day Number/Time is accurate to within 1-second.
// In addition, the 'JulianDayNoDto' type also provides a Julian Day Number/Time
// stored as *big.Float type. This version of the Julian Day Number/Time is
// accurate to within one nanosecond.
//
//
//
// ------------------------------------------------------------------------
//
// Resources
//
// For more information on the Julian Day Number/Time see:
//   https://en.wikipedia.org/wiki/Julian_day
//
type JulianDayNoDto struct {
	julianDayNo             *big.Int    // Julian Day Number expressed as integer value
	julianDayNoFraction     *big.Float  // The Fractional Time value of the Julian
	//                                  //   Day No Time
	julianDayNoTime         *big.Float  // Julian Day Number Plus Time Fraction accurate to
	//                                  //   within nanoseconds
	julianDayNoNumericalSign int        // Sign of the Julian Day Number/Time value. This value
	//                                  //   is either '+1' or '-1'
	totalJulianNanoSeconds   int64      // Julian Day Number Time Value expressed in nanoseconds.
	//                                  //   Always represents a positive value less than 36-hours
	netGregorianNanoSeconds  int64      // Gregorian nanoseconds. Always represents a value in
	//                                  //    nanoseconds which is less than 24-hours.
	hasLeapSecond      bool // If set to 'true' it signals that the day identified
	//                      //   by this Julian Day Number has a duration go 24-hours
	//                      //   + 1-second.
	hours               int // Gregorian Hours
	minutes             int // Gregorian Minutes
	seconds             int // Gregorian Seconds
	nanoseconds         int // Gregorian Nanoseconds

	isThisInstanceValid bool        // Is this object valid flag
	lock                *sync.Mutex // Used for coordinating thread safe operations.
}
// CopyIn - Receives a pointer to an incoming JulianDayNoDto instance
// (JulianDayNoDto2) and performs a deep copy of all internal data
// field values to the current JulianDayNoDto instance (jDNDto).
//
//
// ------------------------------------------------------------------------
//
// Input Values
//
//  JulianDayNoDto2    *JulianDayNoDto
//    - Data field values from this JulianDayNoDto instance will
//      be copied into the current JulianDayNoDto instance. If the
//      method completes successfully, data fields in both
//      instances will have identical values.
//
//  ePrefix            string
//    - Error Prefix. This is an error prefix which is included in all returned
//      error messages. Usually, it contains the names of the calling
//      method or methods. If no error prefix is desired, simply provide
//      an empty string for this parameter.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (jDNDto *JulianDayNoDto) CopyIn(
	JulianDayNoDto2 *JulianDayNoDto,
	ePrefix string) error {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.CopyIn() "

	jDNDtoUtil := julianDayNoDtoUtility{}

	return jDNDtoUtil.copyIn(
		jDNDto,
		JulianDayNoDto2,
		ePrefix)
}

// CopyOut - Returns a deep copy of of the current
// JulianDayNoDto instance.
//
// ------------------------------------------------------------------------
//
// Input Parameter
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
func (jDNDto *JulianDayNoDto) CopyOut(
	ePrefix string) (JulianDayNoDto, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.CopyOut() "

	jDNDtoUtil := julianDayNoDtoUtility{}

	return jDNDtoUtil.copyOut(jDNDto, ePrefix)
}

// Empty - Sets all data fields in the current JulianDayNoDto
// instance to invalid values. Data field pointers will be set
// to 'nil'.
//
// BE CAREFUL!
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
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
func (jDNDto *JulianDayNoDto) Empty(ePrefix string) error {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.Empty() "

	jDNMech := julianDayNoDtoMechanics{}

	return jDNMech.empty(jDNDto, ePrefix)
}

// ExchangeValues - Receives a pointer to an incoming JulianDayNoDto
// object and proceeds to exchange the data values of all internal
// member variables with those contained in the current JulianDayNoDto
// instance.
//
// If the method completes successfully, the current JulianDayNoDto
// instance will be populated with the original data values from
// input parameter 'jDNDtoTwo'. Likewise, 'jDNDtoTwo' will be
// populated with the original values copied from the current
// JulianDayNoDto instance.
//
// If either the current JulianDayNoDto instance or the input parameter
// 'jDNDtoTwo' are judged invalid, this method will return an
// error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  jDNDtoTwo                *JulianDayNoDto
//     - A pointer to an instance of JulianDayNoDto. This is one of the
//       JulianDayNoDto objects used in the data exchange. Data values from
//       this instance will be copied to the current JulianDayNoDto
//       instance while the original values from the current JulianDayNoDto
//       instance will be copied to this instance, 'jDNDtoTwo'.
//
//       If 'jDNDtoTwo' is an invalid instance, an error will be returned.
//
//
//  ePrefix                  string
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
func (jDNDto *JulianDayNoDto) ExchangeValues(
	jDNDtoTwo *JulianDayNoDto,
	ePrefix string) error {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.ExchangeValues() "

	jDNMech := julianDayNoDtoMechanics{}

	return jDNMech.exchangeValues(
		jDNDto,
		jDNDtoTwo,
		ePrefix)
}

// GetHasLeapSecond - Returns the value of the internal data field
// 'hasLeapSecond'.  The standard 'day' has a duration of 24-hours.
// If this member variable is set to 'true' is signals that the day
// identified by this Julian Day Number consists of 24-hours + 1-second.
//
// The default value for 'hasLeapSecond' is 'false', signaling that
// the day identified by this DateTime consists of exactly 24-hours.
//
// A leap second is a one-second adjustment that is occasionally applied
// to Coordinated Universal Time (UTC) in order to accommodate the
// difference between precise time (as measured by atomic clocks) and
// imprecise observed solar time (known as UT1 and which varies due
// to irregularities and long-term slowdown in the Earth's rotation).
//
// For more information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
// IMPORTANT
//
// This method does NOT validate the current JulianDayNoDto instance
// before returning the value. To run a validity check on the
// JulianDayNoDto instance first call one of the two following
// methods:
//
//  JulianDayNoDto.IsValidInstance() bool
//                OR
//  JulianDayNoDto.IsValidInstanceError(ePrefix string) error
func (jDNDto *JulianDayNoDto) GetHasLeapSecond() bool {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	return jDNDto.hasLeapSecond
}


// GetDayNoTimeBigFloat - Returns a *big.Float type representing
// the Julian Day Number/Time which is accurate to the nanosecond.
// Therefore, this return value contains both the Julian Day Number
// and the time fraction.
//
// If the current instance of type JulianDayNoDto has been incorrectly
// initialized, this method will return an error.
//
func (jDNDto *JulianDayNoDto) GetDayNoTimeBigFloat(
	ePrefix string) (*big.Float, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.GetDayNoTimeBigFloat() "

	result := big.NewFloat(0.0)

	jDNNanobot := julianDayNoNanobot{}

	_, err := jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing current JulianDayNoDto instance validity. ")

	if err != nil {
		return result, err
	}

	result =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(jDNDto.julianDayNoTime.Prec()).
			Copy(jDNDto.julianDayNoTime)

	if jDNDto.julianDayNoNumericalSign == -1 {
		result =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(jDNDto.julianDayNoTime.Prec()).
				Neg(result)
	}

	return result, nil
}

// GetRoundedDayNoTimeBigFloat - Returns a *big.Float instance
// representing the Julian Day Number/Time. The time fraction
// will be rounded to the number of decimal places specified
// in the input parameter, 'roundToDecPlaces'. Note that
// 'roundToDecPlaces' is an unsigned integer.
//
// The rounding algorithm applied by this method is known as,
// 'Round half away from zero'. For a discussion of rounding
// algorithms reference:
//   https://en.wikipedia.org/wiki/Rounding
//
// The return value is a type *big.Float floating point number
// which contains both the Julian Day Number and the time
// fraction.
//
// If the current instance of type JulianDayNoDto has been
// incorrectly initialized, this method will return an error.
//
func (jDNDto *JulianDayNoDto) GetRoundedDayNoTimeBigFloat(
	roundToDecPlaces uint,
	ePrefix string) (roundedJDN *big.Float, err error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.GetRoundedDayNoTimeBigFloat() "

	roundedJDN = big.NewFloat(0.0)

	jDNFloat := big.NewFloat(0.0)

	jDNNanobot := julianDayNoNanobot{}

	_, err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing current JulianDayNoDto instance validity. ")

	if err != nil {
		return roundedJDN, err
	}

	jDNFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(jDNDto.julianDayNoTime.Prec()).
			Copy(jDNDto.julianDayNoTime)

	if jDNDto.julianDayNoNumericalSign == -1 {
		jDNFloat =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(jDNDto.julianDayNoTime.Prec()).
				Neg(jDNFloat)
	}

	precision := uint(1024)

	if jDNFloat.Prec() > precision {
		precision = jDNFloat.Prec()
	}

	mathBFMech := MathBigFloatHelper{}

	roundedJDN = mathBFMech.RoundHalfAwayFromZero(
		jDNFloat,
		precision,
		roundToDecPlaces)

	return roundedJDN, err
}

// GetDayNoTimeFloat64 - Returns a float64 value representing
// the Julian Day Number/Time to the nearest second.
//
// Typically a Julian Day Number/Time value can be represented
// by a float64 with 15-decimals to the right of the decimal
// place.
//
func (jDNDto *JulianDayNoDto) GetDayNoTimeFloat64(
	ePrefix string) (
	float64, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.GetDayNoTimeFloat64() "

	var float64Result float64

	float64Result = 0.0

	jDNNanobot := julianDayNoNanobot{}

	_, err := jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing current JulianDayNoDto instance validity. ")

	if err != nil {
		return float64Result, err
	}

	if !jDNDto.julianDayNo.IsInt64() {
		return float64Result,
			errors.New("Error: 'jDNDto.julianDayNo' could not be converted to type 'int64'!\n")
	}

	actualNanoSecsBigInt := int64(jDNDto.nanoseconds)

	netSecsBigInt := jDNDto.totalJulianNanoSeconds - actualNanoSecsBigInt

	netSecsBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		SetInt64(netSecsBigInt)

	dayNoNanoSecsBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		SetInt64(int64(time.Hour) * 24)

	secsTimeFracBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		Quo(netSecsBigFloat, dayNoNanoSecsBigFloat)

	julianDayNoBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		SetInt(jDNDto.julianDayNo)

	adjustedJulianDayNo := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		Add(julianDayNoBigFloat, secsTimeFracBigFloat)

	// Example Target "%.20f"
	fStr1 := "%"
	fStr3 := "f"
	fStr2 := fmt.Sprintf(".%d", 15)

	fStr := fStr1 + fStr2 + fStr3

	julianDayNumTimeStr :=
		fmt.Sprintf(fStr, adjustedJulianDayNo)

	float64Result,
		err =
		strconv.ParseFloat(julianDayNumTimeStr, 64)

	if err != nil {
		float64Result = 0.0
		return float64Result, fmt.Errorf(ePrefix+"\n"+
			"Error returned by strconv.ParseFloat(julianDayNumTimeStr, 64).\n"+
			"julianDayNumTimeStr='%v'\n"+
			"Error='%v'\n",
			julianDayNumTimeStr, err.Error())
	}

	if jDNDto.julianDayNoNumericalSign == -1 {
		float64Result *= -1.0
	}

	return float64Result, nil
}

// GetDayNoTimeFloat32 - Returns a float32 value representing
// the Julian Day Number/Time to the nearest second.
//
// Typically a Julian Day Number/Time value can be represented
// by a float32 with 6-decimals to the right of the decimal
// place.
//
func (jDNDto *JulianDayNoDto) GetDayNoTimeFloat32(
	ePrefix string) (
	float32, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.GetDayNoTimeFloat64() "

	var float32Result float32

	float32Result = 0.0

	jDNNanobot := julianDayNoNanobot{}

	_, err := jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing current JulianDayNoDto instance validity. ")

	if err != nil {
		return float32Result, err
	}

	if !jDNDto.julianDayNo.IsInt64() {
		return float32Result,
			errors.New("Error: 'jDNDto.julianDayNo' could not be converted to type 'int64'!\n")
	}

	actualNanoSecsBigInt := int64(jDNDto.nanoseconds)

	netSecsBigInt := jDNDto.totalJulianNanoSeconds - actualNanoSecsBigInt

	netSecsBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		SetInt64(netSecsBigInt)

	dayNoNanoSecsBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		SetInt64(int64(time.Hour) * 24)

	secsTimeFracBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		Quo(netSecsBigFloat, dayNoNanoSecsBigFloat)

	julianDayNoBigFloat := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		SetInt(jDNDto.julianDayNo)

	adjustedJulianDayNo := big.NewFloat(0.0).
		SetMode(big.ToNearestAway).
		SetPrec(512).
		Add(julianDayNoBigFloat, secsTimeFracBigFloat)

	// Example Target "%.20f"
	fStr1 := "%"
	fStr3 := "f"
	fStr2 := fmt.Sprintf(".%d", 6)

	fStr := fStr1 + fStr2 + fStr3

	julianDayNumTimeStr :=
		fmt.Sprintf(fStr, adjustedJulianDayNo)

	var float64Result float64

	float64Result,
		err =
		strconv.ParseFloat(julianDayNumTimeStr, 32)

	if err != nil {
		float32Result = 0.0
		return float32Result, fmt.Errorf(ePrefix+"\n"+
			"Error returned by strconv.ParseFloat(julianDayNumTimeStr, 64).\n"+
			"julianDayNumTimeStr='%v'\n"+
			"Error='%v'\n",
			julianDayNumTimeStr, err.Error())
	}

	float32Result = float32(float64Result)

	if jDNDto.julianDayNoNumericalSign == -1 {
		float32Result *= -1.0
	}

	return float32Result, nil
}

// GetJulianDay - Returns the Julian Day Number as a
// type int64.
//
func (jDNDto *JulianDayNoDto) GetJulianDayInt64(
	ePrefix string) (julianDayNo int64, err error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	julianDayNo = 0

	err = nil

	ePrefix += "JulianDayNoDto.GetJulianDayInt64() "

	jDNNanobot := julianDayNoNanobot{}

	_, err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing current JulianDayNoDto instance validity. ")

	if err != nil {
		return julianDayNo, err
	}

	if !jDNDto.julianDayNo.IsInt64() {
		bIntVal :=
			big.NewInt(0).
				Set(jDNDto.julianDayNo)

		if jDNDto.julianDayNoNumericalSign == -1 {
			bIntVal =
				big.NewInt(0).
					Neg(bIntVal)
		}

		err = fmt.Errorf(ePrefix + "\n" +
			"Error: The Julian Day Number integer is too large to be" +
			"represented by an Int64 value.\n" +
			"Julian Day Number='%v'\n",
			bIntVal.Text(10))

		return julianDayNo, err
	}

	julianDayNo = jDNDto.julianDayNo.Int64()

	if jDNDto.julianDayNoNumericalSign == -1 {
		julianDayNo *= -1
	}

	return julianDayNo, err
}

// GetJulianDay - Returns the Julian Day Number as a
// type *big.Int.
//
func (jDNDto *JulianDayNoDto) GetJulianDayBigInt(
	ePrefix string) (*big.Int, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.GetJulianDayBigInt() "

	jDNNanobot := julianDayNoNanobot{}

	_, err := jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing current JulianDayNoDto instance validity. ")

	if err != nil {
		return big.NewInt(-1), err
	}

	if !jDNDto.julianDayNo.IsInt64() {
		return big.NewInt(-1),
			errors.New(ePrefix + "\n" +
				"Error: jDNDto.julianDayNo is NOT Int64!\n")
	}

	result := big.NewInt(0).Set(jDNDto.julianDayNo)

	if jDNDto.julianDayNoNumericalSign == -1 {
		result = big.NewInt(0).Neg(result)
	}

	return result, nil
}

// GetJulianDayNoTimeStr - Returns a string containing a floating point
// number representing the Julian Day Number/Time. The calling routine
// specifies the number of digits to the right of the decimal in the
// returned floating point value. In addition to floating point numeric
// value, this method also returns the total string width of the floating
// point number along with the with of the integer value contained in
// that string.
//
// If an error is encountered returned parameters 'strWidth' and 'intWidth'
// are set equal to '-1'.
//
func (jDNDto *JulianDayNoDto) GetJulianDayNoTimeStr(
	digitsToRightOfDecimal int) (
	julianDayNumTimeStr string,
	strWidth,
	intWidth int,
	err error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.GetJulianDayNoTimeStr() "
	julianDayNumTimeStr = ""
	strWidth = -1
	intWidth = -1

	jDNNanobot := julianDayNoNanobot{}

	_, err = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing current JulianDayNoDto instance validity. ")

	if err != nil {
		return julianDayNumTimeStr, strWidth, intWidth, err
	}

	// Example Target "%.20f"
	fStr1 := "%"
	fStr3 := "f"
	fStr2 := fmt.Sprintf(".%d",
		digitsToRightOfDecimal)

	fStr := fStr1 + fStr2 + fStr3

	julianDayNumTimeStr =
		fmt.Sprintf(fStr, jDNDto.julianDayNoTime)

	if jDNDto.julianDayNoNumericalSign == -1 {
		julianDayNumTimeStr = "-" + julianDayNumTimeStr
	}

	strWidth = len(julianDayNumTimeStr)

	intWidth = strings.Index(julianDayNumTimeStr, ".")

	if intWidth < 0 {
		err = errors.New(ePrefix + "\n" +
			"Error: julianDayNumTimeStr does NOT contain a decimal point!\n")
		julianDayNumTimeStr = ""
		strWidth = -1
		intWidth = -1
	}

	return julianDayNumTimeStr, strWidth, intWidth, err
}

// GetGregorianHours - Returns the hours associated with this Julian
// Day Number Time instance. These hours are Gregorian Calendar
// Hours and therefore they may differ from Julian Day Number
// Time hours.
//
// Remember that the Julian Day starts a noon, 12:00:00.000000.
// The Gregorian Calendar day starts at midnight 24:00:00.000000 or
// 00:00:00.000000.
//
// Again this method returns Gregorian Calendar Hours. The value
// hours returned by this method is always less than or equal to
// 24-hours.
//
// The value returned by this method is always positive. This means
// it is always greater than or equal to zero.
//
func (jDNDto *JulianDayNoDto) GetGregorianHours() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	return jDNDto.hours
}

// GetJulianHours - Returns the julian hours associated
// with this Julian Day Number.
//
// Remember that the Julian Day starts a noon, 12:00:00.000000.
// The Gregorian Calendar day starts at midnight 24:00:00.000000 or
// 00:00:00.000000. Therefore, the value of hours returned by this
// method is always less than or equal to 36-hours.
//
// The value returned by this method is always positive. This means
// it is always greater than or equal to zero.
//
func (jDNDto *JulianDayNoDto) GetJulianHours() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	return int(jDNDto.totalJulianNanoSeconds / int64(time.Hour))
}

// GetJulianTimeFraction - Returns the fractional part of Julian Day
// Number/Time as a type *big.Float. The integer portion of this
// this fractional number is always zero. Only the time value is
// returned.
//
// This time fraction will convert to an accuracy of subMicrosecondNanoseconds.
// However, remember that this value represents Julian Time associated
// with a Julian Day. Julian Days start at noon whereas Gregorian Days
// start at midnight.
//
// If the current instance of type JulianDayNoDto has NOT been
// correctly initialized, this method will return an error.
//
// The value returned by this method is always positive. This means
// that it is always greater than or equal to zero.
//
func (jDNDto *JulianDayNoDto) GetJulianTimeFraction(
	ePrefix string) (*big.Float, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.GetJulianTimeFraction() "

	jDNNanobot := julianDayNoNanobot{}

	_, err := jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		ePrefix + "- Testing current JulianDayNoDto instance validity. ")

	if err != nil {
		return big.NewFloat(0.0), err
	}

	return big.NewFloat(0.0).Copy(jDNDto.julianDayNoFraction), nil
}

// GetJulianTotalNanoSeconds - Returns the total subMicrosecondNanoseconds
// associated with this Julian Day Time. As such the Julian Day
// Number is ignored and only the time of day is returned in subMicrosecondNanoseconds.
//
// The returned int64 value represents the total subMicrosecondNanoseconds equaling
// the sum of the hours, minutes, seconds and subMicrosecondNanoseconds encapsulated
// in this Julian Day Number/Time instance. The hours are Julian hours,
// not Gregorian Calendar Hours.
//
// Julian time represented by this total nanosecond value differs
// from Gregorian Calendar time because the Julian Day starts at
// noon (12:00:00.000000 12-hundred hours). By comparison, the
// Gregorian calendar day starts at midnight (00:00:00.000000 Zero hours).
//
// This method returns the Julian time of day in total subMicrosecondNanoseconds.
//
func (jDNDto *JulianDayNoDto) GetJulianTotalNanoSeconds() int64 {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	totalNanoSeconds := jDNDto.totalJulianNanoSeconds

	if jDNDto.julianDayNoNumericalSign == -1 {
		totalNanoSeconds *= -1
	}

	return totalNanoSeconds
}

// GetMinutes - Returns the internal data field
// 'minutes' from the current instance of 'JulianDayNoDto'.
//
// The value returned by this method is always greater than
// or equal to zero.
//
func (jDNDto *JulianDayNoDto) GetMinutes() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	minutesInt := jDNDto.minutes

	if jDNDto.julianDayNoNumericalSign == -1 {
		minutesInt *= -1
	}

	return minutesInt
}

// GetGregorianTotalNanosecs - Returns the total subMicrosecondNanoseconds
// associated with this Julian Day Time. The returned int64 value
// represents the total subMicrosecondNanoseconds equaling the sum of the hours,
// minutes, seconds and subMicrosecondNanoseconds encapsulated in this Julian Day
// Number/Time instance as converted to a Gregorian Calendar day.
//
// Gregorian time represented by this total nanosecond value differs
// from Julian Day time because the Gregorian Day starts at midnight
// (00:00:00.000000 Zero hours). Whereas the Day starts at noon
// (12:00:00.000000 12-hundred hours).
//
// This method returns the Gregorian time in total subMicrosecondNanoseconds which
// in turn represents a value which is always less than or equal to
// 24-hours.
//
// The value returned is always positive which means it is always
// greater than or equal to zero.
//
func (jDNDto *JulianDayNoDto) GetGregorianTotalNanosecs() int64 {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	return jDNDto.netGregorianNanoSeconds
}

// GetSeconds - Returns the internal data field
// 'seconds' from the current instance of 'JulianDayNoDto'.
//
// The value returned by this method is positive and
// therefore, it is always greater than or equal to zero.
//
func (jDNDto *JulianDayNoDto) GetSeconds() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	secondsInt := jDNDto.seconds

	return secondsInt
}

// GetNanoseconds - Returns the internal data field
// 'Nanoseconds' from the current instance of 'JulianDayNoDto'.
//
// This value is always positive, i.e. greater than or
// equal to zero.
//
func (jDNDto *JulianDayNoDto) GetNanoseconds() int {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	nanosecondsInt := jDNDto.nanoseconds

	return nanosecondsInt
}

// IsValidInstance - Returns a boolean value signaling whether the
// current JulianDayNoDto instance is valid.
//
// If the current instance is valid, this method returns 'true'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//   --- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - This boolean flag signals whether the current JulianDayNoDto
//       instance is invalid. If the return value is set to 'true' it
//       signals that the current JulianDayNoDto instance is valid. A
//       return value of 'false' signals that the current JulianDayNoDto
//       instance is invalid.
//
func (jDNDto *JulianDayNoDto) IsValidInstance() bool {


	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	jDNNanobot := julianDayNoNanobot{}

	var isValid bool

	isValid, _ = jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		"")

	return isValid
}

// IsValidInstanceError - Similar to method JulianDayNoDto.IsValidInstance().
// However, this method will return an error message if the current instance
// of JulianDayNoDto is invalid.
//
// If the current instance of JulianDayNoDto is valid, this method returns
// 'nil'. Otherwise the the current JulianDayNoDto instance is judged to be
// invalid and an appropriate error message is returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameter
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
//  error
//     - If the current JulianDayNoDto instance is valid, this method
//       will return 'nil'. Otherwise, the current JulianDayNoDto
//       instance is judged to be invalid, and an appropriate error
//       message is returned.
//
//       Returned error messages will include the Error Prefix passed
//       as an input parameter.
//
func (jDNDto *JulianDayNoDto) IsValidInstanceError(
	ePrefix string) error {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.IsValidInstanceError() "

	jDNNanobot := julianDayNoNanobot{}

	_, err := jDNNanobot.testJulianDayNoDtoValidity(
		jDNDto,
		"JulianDayNoDto.IsValidInstanceError() ")

	return err
}

// New - Returns a new, populated instance of type 'JulianDayNoDto' containing
// Julian Day Number information calculated from the input parameters listed
// below.
//
// For more information on Julian Day Numbers, reference:
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  julianDayNo              int64
//     - The integer value of the Julian Day Number. This effectively
//       identifies the year, month and day values, but not the time
//       of day.
//
//
//  julianDayNoTimeFraction  *big.Float
//       The Julian Day Number time fraction. This is the fractional
//       portion of a Julian Day Number. This fraction identifies the
//       time of day for a Julian Day. Remember that unlike a Universal
//       Coordinated Time, Julian Day Time begins at 12:00:00.000000000-hours
//       or noon.
//
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
//  JulianDayNoDto
//     - If successful, this method will return an instance of 'JulianDayNoDto'
//       populated with Julian Day Number information calculated from the input
//       parameters listed above.
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note: This error
//       message will incorporate the text passed by the Error Prefix
//       input parameter, 'ePrefix'.
//
func (jDNDto JulianDayNoDto) New(
	julianDayNo int64,
	julianDayNoTimeFraction *big.Float,
	ePrefix string) (
	JulianDayNoDto,
	error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.New() "

	julianDayNoDto := JulianDayNoDto{}

	jDNDtoUtil := julianDayNoDtoUtility{}

	err := jDNDtoUtil.setDto(
		&julianDayNoDto,
		julianDayNo,
		julianDayNoTimeFraction,
		1024,
		false,
		ePrefix)

	return julianDayNoDto, err
}

// NewFromFloat64 - Computes and returns a new instance
// of JulianDayNoDto based on the value of a float64
// input parameter. 'float64' time fractions are only
// accurate to a second. Fractional subMicrosecondNanoseconds are
// rounded to the nearest second.
//
// For greater accuracy as to time of day, use method
// JulianDayNoDto.NewFromJulianDayNoTime().
//
// For more information on the Julian Day Number, reference:
//      https://en.wikipedia.org/wiki/Julian_day
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  julianDayNoTime          float64
//     - The Julian Day Number. This effectively identifies the year,
//       month and day plus the time-of-day values.  Time-of-day is
//       formatted in the context of a Julian day which begins at
//       12:00:00.000000000-hours or noon.
//
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
//  JulianDayNoDto
//     - If successful, this method will return an instance of 'JulianDayNoDto'
//       populated with Julian Day Number information calculated from the input
//       parameters listed above.
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note: This error
//       message will incorporate the text passed by the Error Prefix
//       input parameter, 'ePrefix'.
//
func (jDNDto JulianDayNoDto) NewFromFloat64(
	julianDayNoTime float64,
	ePrefix string) (
	JulianDayNoDto,
	error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.NewFromFloat64() "

	julianDayNoDto := JulianDayNoDto{}

	jDNDtoUtil := julianDayNoDtoUtility{}

	err := jDNDtoUtil.setDtoFromFloat64(
		&julianDayNoDto,
		julianDayNoTime,
		1024,
		false,
		ePrefix)

	return julianDayNoDto, err
}

// NewFromGregorianDateTime - Receives a Gregorian Date Time
// and returns that Gregorian Date Time Converted to Universal
// Coordinated Time (UTC) plus a new JulianDayNoDto containing
// the Julian Day Number and Time.
//
// Julian Day Numbers require that Gregorian Date Time first be
// expressed in Universal Coordinated Time (UTC) before being
// converted to Julian Day Number/Time values.
//
// Reference Wikipedia:
//   https://en.wikipedia.org/wiki/Julian_day
//
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  gregorianDateTime  time.Time
//     - This date time value will be converted to Universal
//       Coordinated Time (UTC) before conversion to a Julian
//       Day Number/Time.
//
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
//  gregorianDateUtc   time.Time
//     - The input parameter 'gregorianDateTime' converted to Universal
//       Coordinated Time (UTC). This is the date time used to compute
//       the Julian Day Number (JDN)
//
//
//  julianDayNoDto     JulianDayNoDto
//     - This returned type contains the data elements of a Julian Day
//       Number/Time value. Note that key Julian Day Number and Time values
//       are stored as *big.Int and *big.Float
//
//        type JulianDayNoDto struct {
//           julianDayNo             *big.Int   // Julian Day Number expressed as integer value
//           julianDayNoFraction     *big.Float // The Fractional Time value of the Julian
//                                              //   Day No Time
//           julianDayNoTime         *big.Float // JulianDayNo Plus Time Fraction accurate to
//                                              //   within subMicrosecondNanoseconds
//           julianDayNoNumericalSign         int        // Sign of the Julian Day Number/Time value
//           totalJulianNanoSeconds        *big.Int   // Julian Day Number Time Value expressed in nano seconds.
//                                              //   Always represents a value less than 24-hours
//                                              // Julian Hours
//           hours                   int
//           minutes                 int
//           seconds                 int
//           subMicrosecondNanoseconds             int
//           lock                    *sync.Mutex
//        }
//
//       The integer portion of this number (digits to left of
//       the decimal) represents the Julian day number and is
//       stored in 'JulianDayNoDto.julianDayNo'. The fractional
//       digits to the right of the decimal represent elapsed time
//       since noon on the Julian day number and is stored in
//       'JulianDayNoDto.julianDayNoFraction'. The combined Julian
//       Day Number Time value is stored in 'JulianDayNoDto.julianDayNoTime'.
//       All time values are expressed as Universal Coordinated Time (UTC).
//
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (jDNDto JulianDayNoDto) NewFromGregorianDate(
	gregorianDateTime time.Time,
	ePrefix string) (
	gregorianDateTimeUtc time.Time,
	julianDayNoDto JulianDayNoDto,
	err error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.NewFromGregorianDate() "
	calendarMech := calendarMechanics{}

	gregorianDateTimeUtc = gregorianDateTime.UTC()

	julianDayNoDto,
		err =
		calendarMech.gregorianDateToJulianDayNoTime(
			int64(gregorianDateTimeUtc.Year()),
			int(gregorianDateTimeUtc.Month()),
			gregorianDateTimeUtc.Day(),
			gregorianDateTimeUtc.Hour(),
			gregorianDateTimeUtc.Minute(),
			gregorianDateTimeUtc.Second(),
			gregorianDateTimeUtc.Nanosecond(),
			ePrefix)

	return gregorianDateTimeUtc, julianDayNoDto, err
}

// NewFromJulianDayNoTime - Returns a new instance of JulianDayNoDto
// based on the single parameter, 'julianDayNoTime'. This parameter is
// of type 'big.Float' identifying the Julian Day Number, the fractional
// time value and the numerical sign of the Julian Day Number. As such,
// the 'julianDayNoTime' may be set to either a positive or negative
// value.
//
// Through the use of the 'big.Float' type, extremely accurate
// time-of-day values may calculated. Time values using the 'big.Float'
// type can achieve nanosecond accuracy.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  julianDayNoTimeFraction   *big.Float
//     - This floating point value represents the Julian Day Number and the
//       Julian Day Number time fraction. Julian Day Number/Time with a
//       fractional value of zero represent 12:00:00.000000000-Hours, or
//       Noon. This is due to the fact that Julian Day Numbers begin at
//       12:00:00-hours or noon. For more information on the Julian Day
//       Number/Time see:
//          https://en.wikipedia.org/wiki/Julian_day
//
//
//  hasLeapSecond           bool
//     - Set this boolean parameter to 'true' if, and ONLY if, the specified
//       day for which this Julian Day Number will be calculated includes a leap
//       second. A leap second is a one-second adjustment that is occasionally
//       applied to Coordinated Universal Time (UTC) in order to accommodate
//       the difference between precise time (as measured by atomic clocks)
//       and imprecise observed solar time (known as UT1 and which varies due
//       to irregularities and long-term slowdown in the Earth's rotation).
//       If this parameter is set to 'true', the time calculation will assume
//       the duration of a 'day' is 24-hours and one second. Otherwise, the
//       duration of a day is assumed to consist of exactly 24-hours. For more
//       information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  ePrefix                   string
//       An error prefix with will be added to any returned error messages.
//       Generally, this string contains the name of the calling method or
//       methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
// newJDNo              JulianDayNoDto
//     - If successful, this method will return a new instance
//       of type 'JulianDayNoDto' correctly populated with Julian
//       Day Number information computed from the input parameters,
//       shown above.
//
//  err                 error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (jDNDto JulianDayNoDto) NewFromJulianDayNoTime(
	julianDayNoTime *big.Float,
	applyLeapSecond bool,
	ePrefix string) (newJDNo JulianDayNoDto, err error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.NewFromJulianDayNoTime() "

	newJDNo = JulianDayNoDto{}

	jDNMech := julianDayNoDtoMechanics{}

	err =jDNMech.rationalizeJulianDayNoDto(&newJDNo, ePrefix)

	if err != nil {
		return newJDNo, err
	}

	err =
	jDNMech.setBigValDto(
		&newJDNo,
		julianDayNoTime,
		1024,
		applyLeapSecond,
		ePrefix)

	return newJDNo, err
}


// NewZero - Returns a new instance of JulianDayNoDto with
// all internal data elements initialized to their zero
// values. The returned JulianDayNoDto is in all respects,
// valid.
//
// Input Value
//
// ePrefix string - An error prefix which will be added to any
// returned error messages. Usually, this string contains the
// name of the calling method or methods. If no error prefix
// is desired, simply set this parameter to an empty string.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  JulianDayNoDto
//     - If successful this method will populate the return a valid
//       JulianDayNoDto instance populated with zero values.
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (jDNDto JulianDayNoDto) NewZero(ePrefix string) (JulianDayNoDto, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix += "JulianDayNoDto.NewZero()"

	newJDNDto := JulianDayNoDto{}

	jDNMech := julianDayNoDtoMechanics{}

	err := jDNMech.setJulianDayNoDtoToZero(
		&newJDNDto,
		ePrefix + "- newJDNDto ")

	return newJDNDto, err
}

// SetApplyLeapSecond - The 'JulianDayNoDto' type includes an internal
// data field named, 'hasLeapSecond'. The default value for this field
// is 'false'.
//
// The standard 'day' has a duration of exactly 24-hours. If this member
// internal data field is set to 'true' is signals that the day identified
// by this Julian Day Number consists of 24-hours + 1-second.
//
// A leap second is a one-second adjustment that is occasionally applied
// to Coordinated Universal Time (UTC) in order to accommodate the
// difference between precise time (as measured by atomic clocks) and
// imprecise observed solar time (known as UT1 and which varies due
// to irregularities and long-term slowdown in the Earth's rotation).
//
// For more information on the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
// This method can be used to set the value of internal member variable,
// 'hasLeapSecond'.
//
func (jDNDto *JulianDayNoDto) SetApplyLeapSecond(applyLeapSecond bool){

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	jDNDto.hasLeapSecond = applyLeapSecond
}