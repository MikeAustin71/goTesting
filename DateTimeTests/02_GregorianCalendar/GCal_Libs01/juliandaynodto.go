package GCal_Libs01

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
// 2 451 545.
//
// The Julian date (JD), or Julian Day Number/Time, of any instant
// is the Julian day number/ plus the fraction of a day since the
// preceding noon in Universal Time. Julian dates are expressed as
// a Julian day number with a decimal fraction added. For example,
// the Julian Date for 00:30:00.0 UT January 1, 2013, is
// 2 456 293.520 833.
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
	julianDayNo         *big.Int   // Julian Day Number expressed as integer value
	julianDayNoFraction *big.Float // The Fractional Time value of the Julian
	//                                  Day No Time
	julianDayNoTime *big.Float // Julian Day Number Plus Time Fraction accurate to
	//                                  within nanoseconds
	julianDayNoNumericalSign int   // Sign of the Julian Day Number/Time value
	totalJulianNanoSeconds   int64 // Julian Day Number Time Value expressed in nanoseconds.
	//                                  Always represents a positive value less than 36-hours
	netGregorianNanoSeconds int64 // Gregorian nanoseconds. Always represents a value in
	//                                  nanoseconds which is less than 24-hours.
	hours       int // Gregorian Hours
	minutes     int
	seconds     int
	nanoseconds int
	lock        *sync.Mutex
}

// CopyOut - Returns a deep copy of of the current
// JulianDayNoDto instance.
//
func (jDNDto *JulianDayNoDto) CopyOut() JulianDayNoDto {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	jDNDtoUtil := julianDayNoDtoUtility{}

	return jDNDtoUtil.copyOut(jDNDto)
}

func (jDNDto *JulianDayNoDto) Empty() {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	jDNDtoUtil := julianDayNoDtoUtility{}

	jDNDtoUtil.empty(jDNDto)

	return
}

// GetDayNoTimeBigFloat - Returns a *big.Float type representing
// the Julian Day Number/Time which is accurate to the nanosecond.
// Therefore, this return value contains both the Julian Day Number
// and the time fraction.
//
// If the current instance of type JulianDayNoDto has been incorrectly
// initialized, this method will return an error.
//
func (jDNDto *JulianDayNoDto) GetDayNoTimeBigFloat() (*big.Float, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto) GetDayNoTimeBigFloat() "

	result := big.NewFloat(0.0)

	if jDNDto.julianDayNoTime == nil {
		return result, errors.New(ePrefix + "\n" +
			"Error: This instance 'JulianDayNoDto' was not initialized.\n" +
			"'JulianDayNoDto' is INVALID!\n")
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

// GetDayNoTimeFloat64 - Returns a float64 value representing
// the Julian Day Number/Time to the nearest second.
//
// Typically a Julian Day Number/Time value can be represented
// by a float64 with 15-decimals to the right of the decimal
// place.
//
func (jDNDto *JulianDayNoDto) GetDayNoTimeFloat64() (
	float64, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.GetDayNoTimeFloat64() "

	var float64Result float64

	float64Result = 0.0

	if jDNDto.julianDayNoTime == nil {
		return float64Result,
			errors.New(ePrefix + "\n" +
				"Error: This instance of JulianDayNoDto was " +
				"incorrectly initialized and is invalid.\n" +
				"'julianDayNoTime' is nil!")
	}

	if jDNDto.julianDayNo == nil {
		return float64Result,
			errors.New(ePrefix + "\n" +
				"Error: This instance of JulianDayNoDto was " +
				"incorrectly initialized and is invalid.\n" +
				"'julianDayNo' is nil!")
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

	var err error

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
func (jDNDto *JulianDayNoDto) GetDayNoTimeFloat32() (
	float32, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.GetDayNoTimeFloat64() "

	var float32Result float32

	float32Result = 0.0

	if jDNDto.julianDayNoTime == nil {
		return float32Result,
			errors.New(ePrefix + "\n" +
				"Error: This instance of JulianDayNoDto was " +
				"incorrectly initialized and is invalid.\n" +
				"'julianDayNoTime' is nil!")
	}

	if jDNDto.julianDayNo == nil {
		return float32Result,
			errors.New(ePrefix + "\n" +
				"Error: This instance of JulianDayNoDto was " +
				"incorrectly initialized and is invalid.\n" +
				"'julianDayNo' is nil!")
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

	var err error
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
func (jDNDto *JulianDayNoDto) GetJulianDayInt64() int64 {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	if jDNDto.julianDayNo == nil ||
		!jDNDto.julianDayNo.IsInt64() {
		return -1
	}

	result := jDNDto.julianDayNo.Int64()

	if jDNDto.julianDayNoNumericalSign == -1 {
		result *= -1
	}

	return result
}

// GetJulianDay - Returns the Julian Day Number as a
// type *big.Int.
//
func (jDNDto *JulianDayNoDto) GetJulianDayBigInt() (*big.Int, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.GetJulianDayBigInt() "

	if jDNDto.julianDayNo == nil ||
		!jDNDto.julianDayNo.IsInt64() {
		return big.NewInt(-1),
			errors.New(ePrefix + "\n" +
				"Error: This instance of 'JulianDayNoDto' is NOT initialized!\n")
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
	intWidth int) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.GetJulianDayNoTimeStr() "
	julianDayNumTimeStr = ""
	strWidth = -1
	intWidth = -1

	if jDNDto.julianDayNoTime == nil ||
		jDNDto.julianDayNo == nil {
		julianDayNumTimeStr = ePrefix + "JulianDayNoDto instance INVALID!\n"
		return julianDayNumTimeStr, strWidth, intWidth
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
		julianDayNumTimeStr = ePrefix + "\n" +
			"Error: julianDayNumTimeStr does NOT contain a decimal point!\n"
		strWidth = -1
		intWidth = -1
	}

	return julianDayNumTimeStr, strWidth, intWidth
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
func (jDNDto *JulianDayNoDto) GetJulianTimeFraction() (*big.Float, error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.GetJulianTimeFraction() "

	if jDNDto.julianDayNoFraction == nil {
		return big.NewFloat(0.0).SetInf(true),
			errors.New(ePrefix + "\n" +
				"Error: This 'JulianDayNoDto' instance is NOT initialized!\n")
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

// New - Returns a new, populated instance of type
// JulianDayNoDto.
//
func (jDNDto JulianDayNoDto) New(
	julianDayNo int64,
	julianDayNoTimeFraction *big.Float) (
	JulianDayNoDto,
	error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.New() "

	julianDayNoDto := JulianDayNoDto{}

	jDNDtoUtil := julianDayNoDtoUtility{}

	err := jDNDtoUtil.setDto(
		&julianDayNoDto,
		julianDayNo,
		julianDayNoTimeFraction,
		ePrefix)

	return julianDayNoDto, err
}

// NewFromFloat64 - Computes and returns a new instance
// of JulianDayNoDto based on the value of a float64
// input parameter. 'float64' time fractions are only
// accurate to a second. Fractional subMicrosecondNanoseconds are
// rounded to the nearest second.
//
func (jDNDto JulianDayNoDto) NewFromFloat64(
	julianDayNoTime float64) (
	JulianDayNoDto,
	error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.NewFromFloat64() "

	julianDayNoDto := JulianDayNoDto{}

	jDNDtoUtil := julianDayNoDtoUtility{}

	err := jDNDtoUtil.setDtoFromFloat64(
		&julianDayNoDto,
		julianDayNoTime,
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
//     - A string containing the names of the calling functions
//       which invoked this method. The last character in this
//       string should be a blank space.
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
//   The integer portion of this number (digits to left of
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
//       If errors are encountered this error Type will encapsulate
//       an error message.
//
func (jDNDto JulianDayNoDto) NewFromGregorianDate(
	gregorianDateTime time.Time) (
	gregorianDateTimeUtc time.Time,
	julianDayNoDto JulianDayNoDto,
	err error) {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	ePrefix := "JulianDayNoDto.NewFromGregorianDate() "
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

// NewZero - Returns a new instance of JulianDayNoDto with
// all internal data elements initialized to their zero
// values.
//
func (jDNDto JulianDayNoDto) NewZero() JulianDayNoDto {

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.lock.Lock()

	defer jDNDto.lock.Unlock()

	newJDNDto := JulianDayNoDto{}

	jDNDtoUtil := julianDayNoDtoUtility{}

	jDNDtoUtil.rationalizeJulianDayNoDto(&newJDNDto)

	return newJDNDto
}