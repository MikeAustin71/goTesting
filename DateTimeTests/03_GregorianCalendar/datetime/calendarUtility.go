package datetime

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

// Consists of low-level methods used in calendar
// date/time calculations.
//
type CalendarUtility struct {

	lock *sync.Mutex
}


// GetCompleteYearInterval - Computes the interval of whole years between a base date/time
// and a target date/time. To guarantee complete accuracy, the input ADateTimeDto
// objects must contain the correct setting for the internal member variable 'isLeapYear'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// baseDateTimeDto          ADateTimeDto
//     - Specifies the base date time. This value will be compared
//       with parameter 'targetDateDto' to compute the number of
//       whole or complete years between the two date times.
//
//
// targetDateTimeDto        ADateTimeDto
//     - Specifies the target date time. This value will be compared
//       with parameter 'baseDateTimeDto' to compute the number of
//       whole or complete years between the two date times.
//
//
//  ePrefix                 string
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
//  completedYearsInterval  int64
//     - If this method completes successfully, this value will contain
//       the total number of consecutive whole years between the date
//       times specified by input parameters, 'baseDateTimeDto' and
//       'targetDateTimeDto'.
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calUtil *CalendarUtility) GetCompleteYearInterval(
	baseDateTimeDto ADateTimeDto,
	targetDateTimeDto ADateTimeDto,
	ePrefix string) (completedYearsInterval int64, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	completedYearsInterval = -999999999

	err = nil

	ePrefix += "CalendarUtility.NewGetCompleteYearInterval() "

	if !baseDateTimeDto.IsValidInstance() {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'baseDateTimeDto' is INVALID!\n")
		return completedYearsInterval, err
	}

	if !targetDateTimeDto.IsValidInstance() {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'targetDateTimeDto' is INVALID!\n")
		return completedYearsInterval, err
	}

	if baseDateTimeDto.CompareYears(&targetDateTimeDto) == 0 {
		// Years are equal. Nothing to do.
		completedYearsInterval = 0
		return completedYearsInterval, err
	}

	var baseTargetComparisonResult int

	var tempBaseDateTimeDto, tempTargetDateTimeDto ADateTimeDto

	tempBaseDateTimeDto, err = baseDateTimeDto.CopyOut(ePrefix + "- baseDateTimeDto ")

	if err != nil {
		return completedYearsInterval, err
	}

	tempTargetDateTimeDto, err = targetDateTimeDto.CopyOut(ePrefix + "- targetDateTimeDto ")

	if err != nil {
		return completedYearsInterval, err
	}

	baseTargetComparisonResult, err = tempBaseDateTimeDto.Compare(&tempTargetDateTimeDto, ePrefix)

	if err != nil {
		return completedYearsInterval, err
	}

	if baseTargetComparisonResult == 1 {

		err = tempBaseDateTimeDto.ExchangeValues(&tempTargetDateTimeDto, ePrefix)

		if err != nil {
			return completedYearsInterval, err
		}
	}

	fmt.Println()
	fmt.Printf("  Temp Base Year: %v\n", tempBaseDateTimeDto.GetYear())
	fmt.Printf("Temp Target Year: %v\n", tempTargetDateTimeDto.GetYear())

	// baseYear is now less than targetYear

	// Now determine whether base year is a full year
	// or partial year and make necessary adjustments.
	//

	completedYearsInterval =
		tempTargetDateTimeDto.GetYear() -
			tempBaseDateTimeDto.GetYear() - 1


	return completedYearsInterval, err
}

// GetElapsedWholeOrdinalDaysInYear - Returns the number of elapsed
// 24-hour days in a year based on specified month, day, hour, minute,
// second and nanosecond.
//
func (calUtil *CalendarUtility) GetElapsedWholeOrdinalDaysInYear(
	isLeapYear bool,
	month,
	day int,
	hour int,
	minute int,
	second int,
	nanosecond int,
	ePrefix string) (elapsedCompleteDays int, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	elapsedCompleteDays = math.MinInt32
	err = nil

	ePrefix += "CalendarUtility.GetElapsedWholeOrdinalDaysInYear() "

	calMech2 := CalendarUtility{}

	var ordinalDayNumber int


	ordinalDayNumber, err = calMech2.GetOrdinalDayNumber(
		isLeapYear,
		month,
		day,
		ePrefix)

	if err != nil {
		return elapsedCompleteDays, err
	}

	isFullDay := false

	if hour == 23 &&
		minute == 59 &&
		(second == 59 || second == 60) &&
		nanosecond== 999999999 {
		isFullDay = true
	}

	if ordinalDayNumber > 0 {
		if isFullDay {
			elapsedCompleteDays = ordinalDayNumber
		} else {
			// This is a partial day.
			elapsedCompleteDays = ordinalDayNumber - 1
		}

	} else {
		elapsedCompleteDays = 0
	}


	return elapsedCompleteDays, err
}

// GetJulianDayNoTimeFraction - Computes the Julian Day Number time
// fraction. The Julian Day starts a 12:00:00-hours or Noon.
//
// For more information on the Julian Day Number, reference:
//   https://en.wikipedia.org/wiki/Julian_day
//
//
// Note: If the 'second' value is set to '60' a leap second is assumed
// and the calculation is adjusted accordingly.
//
// For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second              int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of 'leap seconds'.
//       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond          int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  julianDayNoTimeFraction   *big.Float
//     - If successful this method will return the Julian Day Number
//       time fraction as a floating point value. Obviously, the value
//       only contains the time fraction and therefore does NOT contain
//       the integer portion of the Julian Day Number.
//
//  julianDayNoAdjustment     int64
//     - If successful this method returns an integer value signaling
//       whether an adjustment is the integer Julian Day Number is
//       required. If the time value represented by the input parameters
//       occurs before Noon (12:00:00.000000000), it signals that the
//       moment occurred on the previous Julian Day Number. Therefore,
//       if the time value occurs before Noon (12:00:00.000000000), this
//       return value is set to -1 (minus one). Otherwise, the return
//       value is set to zero.
//
//  err                       error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calUtil *CalendarUtility) GetJulianDayNoTimeFraction(
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	julianDayNoTimeFraction *big.Float,
	julianDayNoAdjustment int64,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "CalendarUtility.GetJulianDayNoTimeFraction() "

	julianDayNoTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	julianDayNoAdjustment = 0

	err = nil

	timeMech := TimeMechanics{}

	var totalDateTimeNanoSecs, noonNanoSecs int64

	noonNanoSecs = int64(time.Hour) * 12
	var isGreaterThan24Hours bool

	totalDateTimeNanoSecs,
		isGreaterThan24Hours,
		err =
		timeMech.ComputeTotalTimeNanoseconds(
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

		if err != nil {
			return julianDayNoTimeFraction, julianDayNoAdjustment, err
		}

		if isGreaterThan24Hours {
			err = fmt.Errorf(ePrefix + "\n" +
				"Error: The combined time value of 'hour', 'minute',\n" +
				"second and nanosecond exceeds 24-hours!\n" +
				"Total Time in Nanoseconds='%v'\n",
				totalDateTimeNanoSecs)

			return julianDayNoTimeFraction, julianDayNoAdjustment, err
		}

	if totalDateTimeNanoSecs == noonNanoSecs {
		// Fraction is zero.
		return julianDayNoTimeFraction, julianDayNoAdjustment, err
	}

	if totalDateTimeNanoSecs  < noonNanoSecs {
		julianDayNoAdjustment = -1

		totalDateTimeNanoSecs += noonNanoSecs

	} else {
		// totalDateTimeNanoSecs  > noonNanoSecs
		julianDayNoAdjustment = 0
		totalDateTimeNanoSecs -= noonNanoSecs
	}

	var twentyFourHourNanoseconds *big.Float

	if second == 60 {

		twentyFourHourNanoseconds =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64((int64(time.Hour) * 24) + 1)

	} else {

		twentyFourHourNanoseconds =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(int64(time.Hour) * 24)
	}

	actualNanoSec :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(totalDateTimeNanoSecs)

	julianDayNoTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(actualNanoSec, twentyFourHourNanoseconds)

	return julianDayNoTimeFraction, julianDayNoAdjustment, err
}

// GetStandardDayTimeFraction - Computes the time for a standard
// day as a fractional value.
//
// The time value represented by the input parameters is converted
// to total nanoseconds and divided by the number of nanoseconds
// in a standard 24-hour day.
//
// The month and day input parameters are used to validate 'leap seconds'.
//
// Important: If the 'second' value is set to '60' a leap second is assumed
// and the total number of nanoseconds is divided by 24-hours plus one
// second.
//
// For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  month              int
//     - The month time component for this time value.
//       The valid range is 1 - 12 inclusive. This parameter
//       is used to validate 'leap seconds'.
//
//
//  day                int
//     - The day time component for this time value.
//       The valid range is 1 - 31 inclusive. This parameter
//       is used to validate 'leap seconds'.
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of leap seconds.
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  julianDayNoTimeFraction   *big.Float
//     - If successful this method will return the Julian Day Number
//       time fraction as a floating point value. Obviously, the value
//       only contains the time fraction and therefore does NOT contain
//       the integer portion of the Julian Day Number.
//
//  err                       error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calUtil *CalendarUtility) GetStandardDayTimeFraction(
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	standardDayTimeFraction *big.Float,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	standardDayTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	ePrefix += "CalendarUtility.GetStandardDayTimeFraction() "

	if month < 0 || month > 12 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'month' is INVALID!\n" +
			"month='%v'\n", month)

		return standardDayTimeFraction, err
	}

	if day < 1 || day > 31 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'day' is INVALID!\n" +
			"day='%v'\n", day)

		return standardDayTimeFraction, err
	}

	if hour < 0 || hour > 23 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'hour' is INVALID!\n" +
			"hour='%v'\n", hour)

		return standardDayTimeFraction, err
	}

	if minute < 0 || minute > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'minute' is INVALID!\n" +
			"minute='%v'\n", minute)

		return standardDayTimeFraction, err
	}

	if nanosecond < 0 || nanosecond > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'nanosecond' is INVALID!\n" +
			"nanosecond='%v'\n", nanosecond)

		return standardDayTimeFraction, err
	}

	calUtil2 := CalendarUtility{}

	isValidSecond := calUtil2.IsValidSecond(
		month,
		day,
		hour,
		minute,
		second)

	if !isValidSecond {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'nanosecond' is INVALID!\n" +
			"nanosecond='%v'\n", nanosecond)

		return standardDayTimeFraction, err
	}

	var totalDateTimeNanoSecs int64

	totalDateTimeNanoSecs, err =
		calUtil2.GetTimeTotalNanoseconds(
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

	var twentyFourHourNanoseconds *big.Float

	if second == 60 {

		twentyFourHourNanoseconds =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64((int64(time.Hour) * 24) + 1)

	} else {

		twentyFourHourNanoseconds =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(int64(time.Hour) * 24)
	}

	actualNanoSec :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(totalDateTimeNanoSecs)

	standardDayTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(actualNanoSec, twentyFourHourNanoseconds)

	return standardDayTimeFraction, err
}

// GetMonthDayFromOrdinalDayNo - Receives an Ordinal Day Number and returns
// the associated month and day number. The input parameter 'isLeapYear'
// specifies whether the Ordinal Day Number is included in a standard year
// (365-Days) or a Leap Year (366-Days).
//
// The return value 'yearAdjustment' this value will be populated with one
// of two values: Zero (0) or Minus One (-1). A value of Zero signals that
// no prior year adjustment is required. A value of Minus One indicates that
// the ordinal date, passed as an input parameter, represents December 31st
// of the prior year. In this event, an adjustment to the original year value
// is necessary.
//
// For more information on Ordinal Date, reference:
//    https://en.wikipedia.org/wiki/Ordinal_date
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  ordinalDate        int64
//     - A value with a valid range of 1-366 inclusive which specifies
//       the day of the year expressed as an ordinal day number or ordinal
//       date. For more information on 'ordinal dates', reference:
//         https://en.wikipedia.org/wiki/Ordinal_date
//
//
//  isLeapYear         bool
//     - If 'true' it signals that the input parameter 'ordinalDate' represents
//       an ordinal date within a leap year.
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// yearAdjustment     int
//     - If the method completes successfully, this value will
//       be populated with one of three values:
//
//           One (+1)       -
//             A value of plus one (+1) signals that the ordinal day
//             day number represents the first day of the next year
//             (January 1st of the next year). Effectively the year
//             value is equal to year + 1, month is equal to '1' and
//             the day value is equal to '1'.
//
//           Zero (0)       -
//             A value of Zero signals that no year adjustment is
//             required. The ordinal day number was converted to
//             month and day number in the current year.
//
//           Minus One (-1) -
//             A value of Minus One indicates that the ordinal date
//             represents December 31st of the prior year. 'year'
//             is therefore equal to year - 1, month = 12 and day = 31.
//
//  month             int
//     - If the method completes successfully, this value
//       will contain the month number associated with the
//       input parameter 'ordinalDate'.
//
//
//  day                int
//     - If successful this value will contain the day number
//       associated with the input parameter, 'ordinalDate'.
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calUtil *CalendarUtility) GetMonthDayFromOrdinalDayNo(
	ordinalDate int64,
	isLeapYear bool,
	ePrefix string)(
	yearAdjustment int,
	month int,
	day int,
	err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "CalendarUtility.GetMonthDayFromOrdinalDayNo() "

	yearAdjustment = math.MinInt32
	month = -1
	day = -1
	err = nil
	var ordDays []int64
	var mthDays []int

	if isLeapYear {

		if ordinalDate < 0 ||
			ordinalDate > 366 {
			err = fmt.Errorf("\n" +ePrefix + "Error:\n" +
				"Input Parameter 'ordinalDate' is INVALID!\n" +
				"ordinalDate='%v'\n",
				ordinalDate)
			return yearAdjustment, month, day, err
		}

		ordDays = []int64 {
			0,
			31,
			60,
			91,
			121,
			152,
			182,
			213,
			244,
			274,
			305,
			335,
		}
		mthDays = []int {
			-1,
			31,
			29,
			31,
			30,
			31,
			30,
			31,
			31,
			30,
			31,
			30,
			31,
		}

	} else {
		// This is NOT a leap year

		if ordinalDate < 0 ||
			ordinalDate > 365 {
			err = fmt.Errorf("\n" +ePrefix + "Error:\n" +
				"Input Parameter 'ordinalDate' is INVALID!\n" +
				"ordinalDate='%v'\n",
				ordinalDate)
			return yearAdjustment, month, day, err
		}

		ordDays = []int64 {
			0,
			31,
			59,
			90,
			120,
			151,
			181,
			212,
			243,
			273,
			304,
			334,
		}

		mthDays = []int {
			-1,
			31,
			28,
			31,
			30,
			31,
			30,
			31,
			31,
			30,
			31,
			30,
			31,
		}

	}

	mthDays = []int {
		-1,
		31,
		28,
		31,
		30,
		31,
		30,
		31,
		31,
		30,
		31,
		30,
		31,
	}

	yearAdjustment = 0

	for i:=11; i > -1; i-- {

		if ordDays[i] <= ordinalDate {

			if ordDays[i] == ordinalDate {
				day = 1
			} else {
				day = int(ordinalDate - ordDays[i])
			}

			month = i + 1

			if day < 1 ||
				day > mthDays[i+1] {
				err = fmt.Errorf(ePrefix + "\n" +
					"Invalid Ordinal Day Number Result!\n" +
					"month='%v'  Ordinal Day No='%v'\n" +
					"Original Ordinal Date='%v'\n",
					month, day, ordinalDate)
				return yearAdjustment, month, day, err
			}

			break
			}
		}

	return yearAdjustment, month, day, err
}

// GetOrdinalDayNumber - Computes the ordinal day number
// for any given month and day. Input parameter 'isLeapYear'
// indicates whether the year encompassing the specified
// month and day is a 'leap year' containing 366-days
// instead of the standard 365-days.
//
// Reference
//    https://en.wikipedia.org/wiki/Ordinal_date
//
func (calUtil *CalendarUtility) GetOrdinalDayNumber(
	isLeapYear bool,
	month int,
	day int,
	ePrefix string) (ordinalDayNo int, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "CalendarUtility.GetOrdinalDayNumber() "

	ordinalDayNo = -1
	err = nil

	if month < 1 || month > 12 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input Parameter 'month' is INVALID!\n" +
			"month='%v'\n", month)
		return ordinalDayNo, err
	}


	ordDays := []int {
		0,
		31,
		59,
		90,
		120,
		151,
		181,
		212,
		243,
		273,
		304,
		334,
	}

	mthDays := []int {
		-1,
		31,
		28,
		31,
		30,
		31,
		30,
		31,
		31,
		30,
		31,
		30,
		31,
	}

	if month == 1 &&
		day == 0 {
		ordinalDayNo = 0
		return ordinalDayNo, err
	}

	monthDays := mthDays[month]

	if month==2 && isLeapYear {
		monthDays++
	}

	if day > monthDays || day < 1 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input parameter 'day' is INVALID!\n" +
			"month='%v'\n day='%v'\n",
			month, day)
		return ordinalDayNo, err
	}


	if month == 1 {
		ordinalDayNo = day
		return ordinalDayNo, err
	} else {

		ordinalDayNo = ordDays[month-1] + day

		if isLeapYear && month > 2 {
			ordinalDayNo++
		}
	}

	return ordinalDayNo, err
}


// GetTimeTotalNanoseconds - Computes the total time in nanoseconds for
// a given time of day expressed in Hours, Minutes, Seconds, and Nanoseconds.
//
// If the calculated total time nanoseconds exceeds 24-hours, an error
// is returned.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 should only used in the case of leap seconds.
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  totalTimeNanoseconds     int64
//     - If successful, this method will return the total time value in
//       nanoseconds for the combined input parameters 'hour', 'minute',
//       'second' and 'nanosecond'. If no errors occurred during method
//       execution, this returned value is always greater than or equal
//       to zero.
//
//  err                      error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calUtil *CalendarUtility) GetTimeTotalNanoseconds(
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (totalTimeNanoseconds int64, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "CalendarUtility.GetTimeTotalNanoseconds() "

	var isGreaterThan24Hours bool

	timeMech :=TimeMechanics{}

	totalTimeNanoseconds,
	isGreaterThan24Hours,
	err = timeMech.ComputeTotalTimeNanoseconds(
		hour,
		minute,
		second,
		nanosecond,
		ePrefix)

	if isGreaterThan24Hours {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Total Time in nanoseconds exceeds 24-hours\n" +
			"Total Nanosecond='%v'\n", totalTimeNanoseconds)
	}

	return totalTimeNanoseconds, err
}

// IsLastNanosecondBeforeMidnight - Determines if a date time is precisely
// equivalent to the nanosecond before Midnight:
//       (23:59:59.999999999).
//
// This moment is literally the last nanosecond of the current day.
//
// In addition this method will validate if a leap second
// if it occurs on the correct month and day:
//       (23:59:60.999999999)
//
// If the combination of date time components is invalid,
// this method will return an error.
//
// If this method returns true for month=12 and day=31 it means that
// this date/time marks the end of a complete year.
//
// For a comparison, see the companion method 'IsMidnight()'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
// isLeapYear          bool
//     - Used to validates days and leap seconds.
//
//
//  month              int
//     - The month time component for this time value.
//       The valid range is 1 - 12 inclusive.
//
//
//  day                int
//     - The day time component for this time value.
//       The valid range is 1 - 31 inclusive.
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of leap seconds. The
//       calling method is responsible for supplying the
//       correct 'second' or 'leap second' value.
//
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isLastNanosecondBeforeMidnight  bool
//     - If successful this boolean flag signals whether the date/time
//       passed by the input parameters precisely represents the last
//       nanosecond before midnight of the next day. As such this means
//       that the passed date/time is literally the last nanosecond of
//       the complete current day.
//
//  err                             error
//     - If successful, and the date/time values specified by the input
//       parameters are valid, the returned error Type is set equal to
//       'nil'.
//
//       If errors are encountered during processing, or if the date/time
//       values specified by the input parameters are invalid, the returned
//       error Type will encapsulate an appropriate error message. Note this
//       error message will be prefixed with the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calUtil *CalendarUtility) IsLastNanosecondBeforeMidnight(
	isLeapYear bool,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	isLastNanosecondBeforeMidnight bool, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	isLastNanosecondBeforeMidnight = false

	err = nil

	ePrefix += "CalendarUtility.IsLastNanosecondBeforeMidnight() "

	calUtil2 := CalendarUtility{}

	err = calUtil2.IsValidDateTimeComponents(
		isLeapYear,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		ePrefix)

	if err != nil {
		return isLastNanosecondBeforeMidnight, err
	}

	if hour == 23 &&
		minute == 59 &&
		(second == 59 || second==60) &&
		nanosecond == 999999999 {
		isLastNanosecondBeforeMidnight = true
	}

	return isLastNanosecondBeforeMidnight, err
}

// IsMidnight - Determines if a date time is precisely
// equivalent to Midnight, the beginning of the current
// day: (00:00:00.000000000).
//
// This moment is literally the first nanosecond of the
// current day.
//
// If the combination of date time components is invalid,
// this method will return an error.
//
// If this method returns true for month=1 and day=1 it means that
// this date/time marks the beginning of the current year.
//
// For a comparison, see the companion method:
//     'IsLastNanosecondBeforeMidnight()'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
// isLeapYear          bool
//     - Used to validates days and leap seconds.
//
//
//  month              int
//     - The month time component for this time value.
//       The valid range is 1 - 12 inclusive.
//
//
//  day                int
//     - The day time component for this time value.
//       The valid range is 1 - 31 inclusive.
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of leap seconds. The
//       calling method is responsible for supplying the
//       correct 'second' or 'leap second' value.
//
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isLastNanosecondBeforeMidnight  bool
//     - If successful this boolean flag signals whether the date/time
//       passed by the input parameters precisely represents the last
//       nanosecond before midnight of the next day. As such this means
//       that the passed date/time is literally the last nanosecond of
//       the complete current day.
//
//  err                             error
//     - If successful, and the date/time values specified by the input
//       parameters are valid, the returned error Type is set equal to
//       'nil'.
//
//       If errors are encountered during processing, or if the date/time
//       values specified by the input parameters are invalid, the returned
//       error Type will encapsulate an appropriate error message. Note this
//       error message will be prefixed with the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calUtil *CalendarUtility) IsMidnight(

	isLeapYear bool,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	ePrefix string) (
	isMidnight bool, err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	isMidnight = false

	err = nil

	ePrefix += "CalendarUtility.IsMidnight() "

	calUtil2 := CalendarUtility{}

	err = calUtil2.IsValidDateTimeComponents(
		isLeapYear,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		ePrefix)

	if err != nil {
		return isMidnight, err
	}

	if hour == 0 &&
		minute == 0 &&
		second == 0 &&
		nanosecond == 0 {
		isMidnight = true
	}

	return isMidnight, err
}

// IsValidSecond - Returns true if the 'second' qualifies
// as a valid 'leap second'.
//
// The typical range for a 'second' value is:
//    Greater Than or Equal to zero
//         AND
//    Less Than or Equal to '59'
//
// In addition to applying the standard criteria for a
// 'second', this method also tests to determine whether
// 'second' qualifies as a valid 'leap second'. Valid
// 'leap seconds' may have a value of '60'.
//
// A valid leap second must meet the following criteria:
//
//  (1) The value must be equal to '60'.
//
//  (2) The leap second insertion must occur in June,
//      December, March or September on the last day of the
//      month.
//
//  (3) The leap second must occur at the last second of the
//      the day. Example: 23:59:60.
//
//  (4) The time value of the input parameters is assumed to
//      represent Coordinated Universal Time (UTC). UTC leap
//      seconds occur simultaneously worldwide.
//
//      For more information on Coordinated Universal Time,
//      reference:
//      https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// For more information on 'leap second' reference:
//  https://en.wikipedia.org/wiki/Leap_second
//
func (calUtil *CalendarUtility) IsValidSecond(
	month,
	day,
	hour,
	minute,
	second int) bool {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	if second >= 0 && second <= 59 {
		return true
	}

	// Check to see if this is a valid
	// leap second.
	if month == 3 &&
		day == 31 &&
		hour == 23 &&
		minute == 59 &&
		second == 60 {
		return true
	}

	if month == 6 &&
		day == 30 &&
		hour == 23 &&
		minute == 59 &&
		second == 60 {
		return true
	}

	if month == 9 &&
		day == 30 &&
		hour == 23 &&
		minute == 59 &&
		second == 60 {
		return true
	}

	if month == 12 &&
		day == 31 &&
		hour == 23 &&
		minute == 59 &&
		second == 60 {
		return true
	}

	return false
}

// IsValidDateTimeComponents - Returns a boolean flag signaling
// whether the date time values passed as input parameters are
// valid. To guarantee the accuracy of the result, the user must
// provide the true and correct value for parameter 'isLeapYear'.
//
// If the date time parameters are judged to be 'invalid', this
// method will return a type 'error' with an appropriate error
// message.
//
// Note that the 'year' value is not included in this validation.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
// isLeapYear          bool
//     - Used to validates days and leap seconds.
//
//
//  month              int
//     - The month time component for this time value.
//       The valid range is 1 - 12 inclusive.
//
//
//  day                int
//     - The day time component for this time value.
//       The valid range is 1 - 31 inclusive.
//
//
//  hour               int
//     - The hour time component for this time value.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00.
//
//  minute             int
//     - The minute time component for this time value.
//       The valid range is 0 - 59 inclusive
//
//  second             int
//     - The second time component for this time value.
//       The valid range is 0 - 60 inclusive. The value
//       60 is only used in the case of leap seconds.
///       For a discussion of 'leap seconds' reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//  nanosecond         int
//     - The nanosecond time component for time value.
//       The valid range is 0 - 999,999,999 inclusive
//
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message. Note: Be sure to leave a space at the
//       end of 'ePrefix'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                error
//     - If successful, and the date/time values specified by the input
//       parameters are valid, the returned error Type is set equal to
//       'nil'.
//
//       If errors are encountered during processing, or if the date/time
//       values specified by the input parameters are invalid, the returned
//       error Type will encapsulate an appropriate error message. Note this
//       error message will be prefixed with the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calUtil *CalendarUtility) IsValidDateTimeComponents(
	isLeapYear bool,
	month int,
	day int,
	hour int,
	minute int,
	second int,
	nanosecond int,
	ePrefix string) ( err error) {

	if calUtil.lock == nil {
		calUtil.lock = &sync.Mutex{}
	}

	calUtil.lock.Lock()

	defer calUtil.lock.Unlock()

	ePrefix += "CalendarUtility.IsValidDateTimeComponents() "

	if month < 1 || month > 12 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'month' is invalid!\n" +
			"month='%v'\n", month)
		return err
	}

	if day > 31 || day < 1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'day' is invalid!\n" +
			"day='%v'\n", day)
		return err
	}

	if month == 2 {
		// This is February - check it out!
		if isLeapYear &&
			day > 29 {

			err = fmt.Errorf(ePrefix + "\n" +
				"Error: Input parameter 'month' is invalid!\n" +
				"Febrary only has 29-days.\n" +
				"isLeapYear='%v' month='%v' day='%v'\n",
				isLeapYear, month, day)

			return err

		} else if !isLeapYear &&
			day > 28 {

			err = fmt.Errorf(ePrefix + "\n" +
				"Error: Input parameter 'month' is invalid!\n" +
				"Febrary only has 28-days.\n" +
				"isLeapYear='%v' month='%v' day='%v'\n",
				isLeapYear, month, day)

			return err
		}

	}

	if hour > 24 || hour < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hour' is invalid!\n" +
			"hour='%v'\n", hour)
		return err
	}

	if minute < 0 || minute > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minute' is invalid!\n" +
			"minute='%v'\n", minute)
		return err
	}

	calUtil2 := CalendarUtility{}

	isValidSecond := calUtil2.IsValidSecond(
		month,
		day,
		hour,
		minute,
		second)

	// Watch out for leap seconds
	if !isValidSecond {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'second' is invalid!\n" +
			"second='%v'\n", second)
		return err
	}

	if nanosecond < 0 || nanosecond > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecond' is invalid!\n" +
			"nanosecond='%v'\n", nanosecond)
		return err
	}

	err = nil

	return err
}