package RJC_Libs02

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"
)

// CalendarRevisedJulianUtility - This type contains methods
// used to process date arithmetic associated with the Revised
// Julian Calendar.
//
// Reference:
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
type CalendarRevisedJulianUtility struct {

	lock *sync.Mutex
}

// GetJulianDayNumber - Returns values defining the Julian Day Number
// and Time for a date/time in the Revised Julian Calendar.
//
// All time input parameters are assumed to be expressed in Coordinated
// Universal Time (UTC). For more information on Coordinated Universal
// Time, reference:
//   https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// Julian Day Number is used to define a standard time duration, and
// perform date/time conversions, between differing calendar systems.
//
// The base date/time for Julian Day Number zero on the Revised Julian
// Calendar is February 8, -4713 12:00:00 UTC (Noon) or February 8,
// 4714 BCE 12:00:00 UTC (Noon).
//
// For more information on the Julian Day Number, reference:
//  https://en.wikipedia.org/wiki/Julian_day
//
// For more information on the Revised Julian Calendar, reference:
//  //  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetYear         int64
//     - The year number associated with this date/time specification.
//       The year value may be positive or negative. The year value must
//       conform to the astronomical year numbering system. This means
//       that year zero is valid and recognized. Example: 1/1/0000. The
//       astronomical year value -4712 is therefore equivalent to
//       -4713 BCE. All year values submitted to this method must use
//       the astronomical year numbering system. For more information
//       on the astronomical year numbering system, reference:
//              https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//  targetMonth        int
//     - The month number for this date/time specification.
//       The valid range is 1 - 12 inclusive.
//
//  targetDay          int
//     - The day number for this date/time specification. The day
//       number must fall within the limits of the month number
//       submitted above in input parameter, 'targetMonth'.
//
//  targetHour         int
//     - The hour time component for this date/time specification.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00. All time
//       parameters are assumed to be expressed in Universal
//       Coordinated Time (UTC).
//
//  targetMinute       int
//     - The minute time component for this date/time specification.
//       The valid range is 0 - 59 inclusive.  All time
//       parameters are assumed to be expressed in Universal
//       Coordinated Time (UTC).
//
//  targetSecond       int
//     - The second time component for this date/time specification.
//       The valid range is 0 - 60 inclusive. The value 60 is only
//       used in the case of leap seconds.  All time parameters are
//       assumed to be expressed in Universal Coordinated Time (UTC).
//
//  targetNanosecond   int
//     - The nanosecond time component for this date/time specification.
//       The valid range is 0 - 999,999,999 inclusive.  All time
//       parameters are assumed to be expressed in Universal
//       Coordinated Time (UTC).
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  julianDayNumber             int64
//     - The integer julian day number of the date specified by input
//       parameters 'targetYear', 'targetMonth' and 'targetDay'. This
//       value equals the number of days elapsed between the base date,
//       February 8, 4714 BCE 12:00:00 UTC (Noon), and the target date/time
//       specified by the input parameters. Both base and target date/times
//       represent moments on the Revised Julian Calendar.
//
//
//  julianDayNumberTime         *big.Float
//     - The combined Julian Day number and fractional time value represented
//       by the date/time specified by input parameters, targetYear,
//       targetMonth, targetDay, targetHour, targetMinute, targetSecond
//       and targetNanosecond. This value equals the duration in day numbers
//       calculated from the start date of February 8, 4714 BCE 12:00:00 UTC
//       (Noon) on the Revised Julian Calendar.
//
//
//  julianDayNumberTimeFraction *big.Float
//     - The fractional time value associated with the the ordinal
//       day number. This value does NOT contain the integer ordinal
//       day number. Instead, it only contains the time value
//       represented by the date/time input parameters, targetHour,
//       targetMinute, targetSecond and targetNanosecond. This value
//       equals the number of nanoseconds since midnight of the ordinal
//       day number divided by the number of nanoseconds in a
//       24-hour day.
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
func (calRJM *CalendarRevisedJulianUtility) GetJulianDayNumber(
	targetYear int64,
	targetMonth int,
	targetDay int,
	targetHour int,
	targetMinute int,
	targetSecond int,
	targetNanosecond int,
	ePrefix string) (
	julianDayNumber int64,
	julianDayNumberTime *big.Float,
	julianDayNumberTimeFraction *big.Float,
	err error) {

	if calRJM.lock == nil {
		calRJM.lock = &sync.Mutex{}
	}

	calRJM.lock.Lock()

	defer calRJM.lock.Unlock()

	julianDayNumber = 0

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	julianDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	calRJM2 := CalendarRevisedJulianUtility{}

	isLeapYear := calRJM2.IsLeapYear(targetYear)

	calUtil := CalendarUtility{}

	err = calUtil.IsValidDateTimeComponents(
		isLeapYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {

		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	var baseDateTimeDto, targetDateTimeDto DateTimeTransferDto

	targetDateTimeDto, err = DateTimeTransferDto{}.New(
		isLeapYear,
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	// Old Disproven Base Date Time:  February 8, -4713 12:00:00 UTC
	//
	// Now Using Base Date Time:
	//  November 23, -4713 12:00:00.000000000 UTC (Noon)
	//  -4713-11-23 12:00:00:00.000000000 UTC (Noon)
	baseDateTimeDto, err = DateTimeTransferDto{}.New(
		false,
		int64(-4713),
		11,
		23,
		12,
		0,
		0,
		0,
		ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	var baseTargetComparisonResult, baseTargetYearsComparison int

	baseTargetComparisonResult, err = baseDateTimeDto.Compare(&targetDateTimeDto, ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	baseTargetYearsComparison = baseDateTimeDto.CompareYears(&targetDateTimeDto)

	// baseDateTimeDto is now less than targetDateTimeDto

	baseDateTimeDto.SetTag("baseDateTimeDto")
	targetDateTimeDto.SetTag("targetDateTimeDto")
	// Compute wholeYearInterval
	var wholeYearsInterval int64

	if baseTargetComparisonResult == 0 {
	// base and target have equivalent date/times
		// Julian Day Number is Zero and time fraction is Zero.

		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err

	} else if baseTargetYearsComparison == 0 {
		// base and target years are equal
		wholeYearsInterval = 0

	} else if baseTargetComparisonResult == 1 {
		// base is greater than target date/time
		// target date/time must be negative
		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() + 1

	} else {
		// target is greater than base date/time
		// target date/time could be positive or
		// negative
		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() - 1

	}

	if wholeYearsInterval < 0 {
		wholeYearsInterval *= -1
	}

	var wholeYearDays int64

	if wholeYearsInterval == 0 {

		wholeYearDays = 0

	} else {

		wholeYearDays =
			calRJM2.NumCalendarDaysForWholeYearsInterval(wholeYearsInterval)

	}

	var targetYearOrdinalNumber, baseYearOrdinalNumber int

	targetYearOrdinalNumber, err = targetDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	baseYearOrdinalNumber, err = baseDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	var lastWholeTargetYear int64
	var lastWholeTargetYearDeltaDays, targetPartialYearDeltaDays int64


	if baseTargetYearsComparison == 1 {
		// base year is greater than target year
		// target year is negative

		lastWholeTargetYear = targetDateTimeDto.GetYear() + 1

	} else {
		// baseTargetYearsComparison == -1
		// base year is less than target year.
		// target could be positive or negative.
		//
		if targetDateTimeDto.GetYearNumberSign() < 0 {
			// target year is negative
			lastWholeTargetYear = targetDateTimeDto.GetYear() + 1

		} else {
			// target year is positive
			lastWholeTargetYear = targetDateTimeDto.GetYear() - 1
		}
	}

	if calRJM2.IsLeapYear(lastWholeTargetYear) {
		lastWholeTargetYearDeltaDays = int64(366) - int64(baseYearOrdinalNumber)
	} else {
		lastWholeTargetYearDeltaDays = int64(365) - int64(baseYearOrdinalNumber)
	}

	targetPartialYearDeltaDays = int64(targetYearOrdinalNumber)

	julianDayNumber = wholeYearDays +
		lastWholeTargetYearDeltaDays +
		targetPartialYearDeltaDays

	noonNanoseconds := int64(time.Hour * 12)

	var targetDayTotalNanoseconds int64

	targetDayTotalNanoseconds, err =
			targetDateTimeDto.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	julianDayAdjustment := 0

	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)


	if targetDayTotalNanoseconds == noonNanoseconds {
		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(0.0)

	} else if targetDayTotalNanoseconds < noonNanoseconds {
		julianDayNumber--
		targetDayTotalNanoseconds += noonNanoseconds
		julianDayAdjustment--

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)


	} else {
		// targetDayTotalNanoseconds > noonNanoseconds

		targetDayTotalNanoseconds -= noonNanoseconds

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)

	}

	julianDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(julianDayNumber)

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Add(julianDayNoFloat, julianDayNumberTimeFraction)

	separator := strings.Repeat("-", 65)
	fmt.Println()
	fmt.Println(separator)

	fmt.Printf("julianDayAdjustment: %v\n",
		julianDayAdjustment)

	fmt.Println(separator)
	fmt.Println(baseDateTimeDto.String())
	fmt.Println(targetDateTimeDto.String())
	fmt.Println(separator)

	return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
}

// GetOrdinalDayNumberTime - Computes and returns the 'Fixed Day'
// or ordinal day number time under the Revised Julian Calendar.
//
// The ordinal day number time identifies the number of days and
// and fraction times since a fixed start date. The start date
// for the Revised Julian Calendar is Monday, 1 January 1 AD 00:00:00
// (midnight). This is defined as the beginning of the first ordinal
// day. All ordinal day numbers use this start date as the base base
// reference point.
//
// This method receives a series of input parameters specifying
// a target date and time. The times are always assumed to be
// Universal Coordinated Times (UTC). The method then returns
// the ordinal day number as integer (int64) and a type *big.Float
// which defines both the ordinal day number and the time expressed
// as a fraction of a 24-hour day.
//
// Note that the input parameter 'targetYear' is a type int64 which
// must be configured under the astronomical year numbering system.
// This system recognizes the year zero as a legitimate year value.
// Under the astronomical year numbering system the year 4713 BCE is
// formatted as -4712.
//
// Reference:
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//  https://en.wikipedia.org/wiki/Rata_Die
//  https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetYear         int64
//     - The year number associated with this date/time specification.
//       The year value may be positive or negative. The year value must
//       conform to the astronomical year numbering system. This means
//       that year zero is valid and recognized. Example: 1/1/0000. The
//       astronomical year value -4712 is therefore equivalent to
//       -4713 BCE. All year values submitted to this method must use
//       the astronomical year numbering system. For more information
//       on the astronomical year numbering system, reference:
//              https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//  targetMonth        int
//     - The month number for this date/time specification.
//       The valid range is 1 - 12 inclusive.
//
//  targetDay          int
//     - The day number for this date/time specification. The day
//       number must fall within the limits of the month number
//       submitted above in input parameter, 'targetMonth'.
//
//  targetHour         int
//     - The hour time component for this date/time specification.
//       The valid range is 0 - 23 inclusive. The 24th hour should
//       should be expressed as zero hour, 00:00:00. 'targetHour'
//       is assumed to represent Coordinated Universal Time (UTC).
//
//  targetMinute       int
//     - The minute time component for this date/time specification.
//       The valid range is 0 - 59 inclusive. 'targetMinute' is assumed
//       to represent  Coordinated Universal Time (UTC).
//
//  targetSecond       int
//     - The second time component for this date/time specification.
//       The valid range is 0 - 60 inclusive. The value 60 is only
//       used in the case of leap seconds.
//
//  targetNanosecond   int
//     - The nanosecond time component for this date/time specification.
//       The valid range is 0 - 999,999,999 inclusive
//
//  ePrefix            string
//     - A string consisting of the method chain used to call
//       this method. In case of error, this text string is included
//       in the error message.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ordinalDayNumber              int64
//     - The ordinal day number of the date represented by the date
//       specified by input parameters targetYear, targetMonth and
//       targetDay. This value equals the number of days elapsed,
//       plus 1-day, since January 1, 1 CE on the Revised Julian
//       Calendar.
//
//
//  ordinalDayNumberTime          *big.Float
//     - The combined ordinal day number and fractional time
//       value for the ordinal day number represented by the
//       date/time specified by input parameters, targetYear,
//       targetMonth, targetDay, targetHour, targetMinute,
//       targetSecond and targetNanosecond. This value equals
//       the ordinal day number calculated from the start date
//       of January 1, 1 CE on the Revised Julian Calendar.
//
//
//  ordinalDayNumberTimeFraction  *big.Float
//     - The fractional time value associated with the the ordinal
//       day number. This value does NOT contain the integer ordinal
//       day number. Instead, it only contains the time value
//       represented by the date/time input parameters, targetHour,
//       targetMinute, targetSecond and targetNanosecond. This value
//       equals the number of nanoseconds since midnight of the ordinal
//       day number divided by the number of nanoseconds in a
//       24-hour day.
//
//
//  err                           error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       in input parameter, 'ePrefix'.
//
//
func (calRJM *CalendarRevisedJulianUtility) GetOrdinalDayNumberTime(
	targetYear int64,
	targetMonth,
	targetDay,
	targetHour,
	targetMinute,
	targetSecond,
	targetNanosecond int,
	ePrefix string) (
	ordinalDayNumber int64,
	ordinalDayNumberTime *big.Float,
	ordinalDayNumberTimeFraction *big.Float,
	err error) {

	if calRJM.lock == nil {
		calRJM.lock = &sync.Mutex{}
	}

	calRJM.lock.Lock()

	defer calRJM.lock.Unlock()

	ordinalDayNumber = 0

	ordinalDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	ordinalDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil

	calRJM2 := CalendarRevisedJulianUtility{}

	isTargetLeapYear := calRJM2.IsLeapYear(targetYear)

	calUtil := CalendarUtility{}

	err = calUtil.IsValidDateTimeComponents(
		isTargetLeapYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {

		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	var baseDateTimeDto, targetDateTimeDto DateTimeTransferDto

	targetDateTimeDto, err = DateTimeTransferDto{}.New(
		isTargetLeapYear,
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	// Base Date Time:  January 1, 0001 00:00:00.000000000 UTC
	baseDateTimeDto, err = DateTimeTransferDto{}.New(
		false,
		int64(1),
		1,
		1,
		0,
		0,
		0,
		0,
		ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	var baseTargetComparisonResult, baseTargetYearsComparison int

	baseTargetComparisonResult, err = baseDateTimeDto.Compare(&targetDateTimeDto, ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	baseTargetYearsComparison = baseDateTimeDto.CompareYears(&targetDateTimeDto)

	// baseDateTimeDto is now less than targetDateTimeDto

	baseDateTimeDto.SetTag("baseDateTimeDto")
	targetDateTimeDto.SetTag("targetDateTimeDto")
	// Compute wholeYearInterval
	var wholeYearsInterval int64

	var targetYearOrdinalNumber, baseYearOrdinalNumber int

	targetYearOrdinalNumber, err = targetDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	wholeYearsInterval = 0

	if baseTargetComparisonResult == 0 {
		// base and target have equivalent date/times
		// Julian Day Number is Zero and time fraction is Zero.

		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err

	} else if baseTargetYearsComparison == 0 {
		// base and target years are equal
		ordinalDayNumber = int64(targetYearOrdinalNumber)




	} else if baseTargetComparisonResult == 1 {
		// base is greater than target date/time
		// target date/time must be negative
		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() + 1

	} else {
		// target is greater than base date/time
		// target date/time could be positive or
		// negative
		wholeYearsInterval =
			targetDateTimeDto.GetYear() -
				baseDateTimeDto.GetYear() - 1

	}

	if wholeYearsInterval < 0 {
		wholeYearsInterval *= -1
	}

	var wholeYearDays int64

	if wholeYearsInterval == 0 {

		wholeYearDays = 0

	} else {

		wholeYearDays =
			calRJM2.NumCalendarDaysForWholeYearsInterval(wholeYearsInterval)

	}

	baseYearOrdinalNumber, err = baseDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	var lastWholeTargetYear int64
	var lastWholeTargetYearDeltaDays, targetPartialYearDeltaDays int64


	if baseTargetYearsComparison == 1 {
		// base year is greater than target year
		// target year is negative

		lastWholeTargetYear = targetDateTimeDto.GetYear() + 1

	} else {
		// baseTargetYearsComparison == -1
		// base year is less than target year.
		// target could be positive or negative.
		//
		if targetDateTimeDto.GetYearNumberSign() < 0 {
			// target year is negative
			lastWholeTargetYear = targetDateTimeDto.GetYear() + 1

		} else {
			// target year is positive
			lastWholeTargetYear = targetDateTimeDto.GetYear() - 1
		}
	}

	if calRJM2.IsLeapYear(lastWholeTargetYear) {
		lastWholeTargetYearDeltaDays = int64(366) - int64(baseYearOrdinalNumber)
	} else {
		lastWholeTargetYearDeltaDays = int64(365) - int64(baseYearOrdinalNumber)
	}

	targetPartialYearDeltaDays = int64(targetYearOrdinalNumber)

	ordinalDayNumber = wholeYearDays +
		lastWholeTargetYearDeltaDays +
		targetPartialYearDeltaDays

	var targetDayTotalNanoseconds int64

	targetDayTotalNanoseconds, err =
		targetDateTimeDto.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
	}

	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)

	targetNanosecondFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(targetDayTotalNanoseconds)

	ordinalDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)

	ordinalDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(ordinalDayNumber)

	ordinalDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Add(ordinalDayNoFloat, ordinalDayNumberTimeFraction)

	separator := strings.Repeat("-", 65)
	fmt.Println()
	fmt.Println(separator)

	fmt.Println(separator)
	fmt.Println(baseDateTimeDto.String())
	fmt.Println(targetDateTimeDto.String())
	fmt.Println(separator)

	return ordinalDayNumber, ordinalDayNumberTime, ordinalDayNumberTimeFraction, err
}

// GetYearDays - Returns the number of days in the year
// identified by input parameter 'year' under the Revised
// Julian Calendar.
//
// If the year is a standard year, this method will return
// 365-days.
//
// If the year is a leap year, this method will return
// 365-days.
//
// For more information on the Revised Julian Calendar and
// leap years, reference:
//
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
//
func (calRJM *CalendarRevisedJulianUtility) GetYearDays(
	year int64) int {
	if calRJM.lock == nil {
		calRJM.lock = &sync.Mutex{}
	}

	calRJM.lock.Lock()

	defer calRJM.lock.Unlock()

	var isLeapYear bool

	if year < 0 {
		year *= -1
	}

	if year % int64(4) == 0 {

		isLeapYear = true

		if year % 100 == 0 {

			isLeapYear = false

			mod900 := year % int64(900)

			if  mod900 == 200 ||
				mod900 == 600 {
				isLeapYear = true
			}

		}

	} else {
		isLeapYear = false
	}

	if isLeapYear {
		return 366
	}

	return 365

}

// IsLeapYear - Returns a boolean value signaling whether the year
// value passed as an input parameter is a leap year (366-days) under
// the Revised Julian Calendar.
//
// Methodology:
//
// 1. Years evenly divisible by 4 are leap years unless they are
//    century years (evenly divisible by 100).
//
// 2. Years evenly divisible by 100 are not leap years unless...
//
// 3. Years evenly divisible by 100 when divided by 900 have
//    remainders of 200 or 600.  In this case they are leap years.
//
// For more information on the Revised Julian Calendar, reference:
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
func (calRJM *CalendarRevisedJulianUtility) IsLeapYear(year int64) bool {

	if calRJM.lock == nil {
		calRJM.lock = &sync.Mutex{}
	}

	calRJM.lock.Lock()

	defer calRJM.lock.Unlock()

	var isLeapYear bool

	if year < 0 {
		year *= -1
	}

	if year % int64(4) == 0 {

		isLeapYear = true

		if year % 100 == 0 {

			isLeapYear = false

			mod900 := year % int64(900)

			if  mod900 == 200 ||
				mod900 == 600 {
				isLeapYear = true
			}

		}

	} else {
		isLeapYear = false
	}

	return isLeapYear
}

// NumCalendarDaysForWholeYearsInterval - Computes the total
// number of 24-hour days in a period of years specified
// by input parameter 'wholeYearsInterval'. The number of total
// days is calculated in accordance with the Revised Julian
// Calendar.
//
// Methodology:
//
// 1. Years evenly divisible by 4 are leap years unless they are
//    century years.
//
// 2. Years evenly divisible by 100 are not leap years unless when
//    divided by 900 those years have remainders of 200 or 600 in
//    which case they are leap years.
//
// For more information on the Revised Julian Calendar, reference:
//  https://en.wikipedia.org/wiki/Revised_Julian_calendar
//
// 'wholeYearsInterval' is defined as a series of contiguous
// whole, or complete, years consisting of either 365-days
// or 366-days (in the case of leap years).
//
// No partial years should be included in this interval.
//
func (calRJM *CalendarRevisedJulianUtility) NumCalendarDaysForWholeYearsInterval(
	wholeYearsInterval int64) (totalDays int64) {

	if calRJM.lock == nil {
		calRJM.lock = &sync.Mutex{}
	}

	calRJM.lock.Lock()

	defer calRJM.lock.Unlock()

	totalDays = 0

	if wholeYearsInterval < 0 {
		wholeYearsInterval *= -1
	}

	if wholeYearsInterval == 0 {
		return 0
	}

	separator := strings.Repeat("*", 65)

	fmt.Println()
	fmt.Println("NumCalendarDaysForWholeYearsInterval() ")
	fmt.Println(separator)
	fmt.Printf("       Whole Years Interval: %v\n", wholeYearsInterval)

	if wholeYearsInterval >= 900 {

		numOfCycles := wholeYearsInterval / 900

		totalDays = numOfCycles * 328718

		fmt.Printf("  Number of 900-Year Cycles: %v\n", numOfCycles)
		fmt.Printf("Number of Days in %v-Cycles: %v\n", numOfCycles, totalDays)

		wholeYearsInterval = wholeYearsInterval - (numOfCycles * 900)

		fmt.Printf("  Number of Remainder Years: %v\n", wholeYearsInterval)
		fmt.Println(separator)
		fmt.Println()

	}

	totalDays += wholeYearsInterval * 365

	leapDays := wholeYearsInterval / 4

	skipLeapDays := wholeYearsInterval / 100

	addLeapDays := int64(0)

	if wholeYearsInterval >= 200 {
		addLeapDays++
	}

	if wholeYearsInterval >= 600 {
		addLeapDays++
	}

	totalDays += leapDays + addLeapDays - skipLeapDays

	fmt.Println(separator)
	fmt.Printf("Total Days In wholeYearsInterval: %v\n",
		totalDays)
	fmt.Println(separator)
	fmt.Println()

	return totalDays
}

// VerifyDayCountForWholeYearInterval - Uses a loop technique to verify
// algorithms for computing the total number of days in a series
// of contiguous whole years. No partial years should be included
// in the input parameter 'wholeYearsInterval'.
//
//
func (calRJM *CalendarRevisedJulianUtility) VerifyDayCountForWholeYearInterval(
	wholeYearsInterval int64,
	ePrefix string) (totalDays int64, err error) {

	if calRJM.lock == nil {
		calRJM.lock = &sync.Mutex{}
	}

	calRJM.lock.Lock()

	defer calRJM.lock.Unlock()

	totalDays = 0

	err = nil

	ePrefix += "CalendarRevisedJulianUtility. VerifyDayCountForWholeYearInterval() "

	if wholeYearsInterval < 1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'wholeYearsInterval' is a negative value and therefore, INVALID!\n" +
			"wholeYearsInterval='%v'\n")

		return totalDays, err
	}

	if wholeYearsInterval < 0 {
		wholeYearsInterval *= -1
	}

	wholeYearsInterval++

	calRJM2 := CalendarRevisedJulianUtility{}

	var yearDays int64

	for i:= int64(1); i < wholeYearsInterval; i++ {

		yearDays = 365

		if calRJM2.IsLeapYear(i) {
			yearDays = 366
		}

		totalDays += yearDays
	}

	return totalDays, err
}