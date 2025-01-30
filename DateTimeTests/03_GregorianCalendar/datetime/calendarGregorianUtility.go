package datetime

import (
	"fmt"
	"strings"
	"sync"
)

// CalendarGregorianUtility - This type contains methods
// used to process date arithmetic associated with the
// Gregorian Calendar
//
// Reference:
//  https://en.wikipedia.org/wiki/Gregorian_calendar
//
type CalendarGregorianUtility struct {

	lock *sync.Mutex
}

// GetJulianDayNumber - Returns values defining the Julian Day Number
// and Time for a date/time in the Gregorian Calendar.
//
// All time input parameters are assumed to be expressed in Coordinated
// Universal Time (UTC). For more information on Coordinated Universal
// Time, reference:
//   https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//
// Julian Day Number is used to define a standard time duration, and
// perform date/time conversions, between differing calendar systems.
//
// The base date/time for Julian Day Number zero on the Gregorian
// Calendar is November 24, -4713 12:00:00.000000000 UTC (Noon) or
// the equivalent November 24, 4714 BCE 12:00:00.000000000 UTC (Noon).
//
// For more information on the Julian Day Number, reference:
//  https://en.wikipedia.org/wiki/Julian_day
//
// For more information on the Gregorian Calendar, reference:
//  https://en.wikipedia.org/wiki/Gregorian_calendar
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
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  julianDayNoDto     JulianDayNoDto
//     - If successful, this method will return a fully populated instance
//       of JulianDayNoDto containing the Julian Day Number as well as the
//       associated time value.  The integer julian day number will be
//       calculated for the Gregorian Calendar date specified by input
//       parameters 'targetYear', 'targetMonth' and 'targetDay'. This
//       value equals the number of days elapsed between the base date,
//       November 24, 4714 BCE 12:00:00 UTC (Noon), and the target date/time
//       specified by the input parameters. Both base and target date/times
//       represent moments on the Gregorian Calendar.
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
func (calGregUtil *CalendarGregorianUtility) GetJulianDayNumber(
	targetYear int64,
	targetMonth int,
	targetDay int,
	targetHour int,
	targetMinute int,
	targetSecond int,
	targetNanosecond int,
	ePrefix string) (
	julianDayNoDto JulianDayNoDto,
	err error) {

	if calGregUtil.lock == nil {
		calGregUtil.lock = &sync.Mutex{}
	}

	calGregUtil.lock.Lock()

	defer calGregUtil.lock.Unlock()

	ePrefix += "CalendarGregorianUtility.GetJulianDayNumber() "

	err = nil

	julianDayNoDto = JulianDayNoDto{}

	calGregMech := calendarGregorianMechanics{}

	isLeapYear := calGregMech.isLeapYear(targetYear)

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

		return julianDayNoDto, err
	}

	var targetDateTimeDto ADateTimeDto

	targetDateTimeDto, err =  ADateTimeDto{}.New(
		CalendarSpec(0).Gregorian(),
		targetYear,
		CalendarYearNumType(0).Astronomical(),
		targetMonth,
		targetDay,
		false,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		"UTC",
		"",
		"",
		ePrefix)

	if err != nil {
		return julianDayNoDto, err
	}

	calCycleCfg := calGregMech.getCalendarCyclesConfig()

	calEng := CalendarEngines{}

	calEng.SetCodeDebug(false)

	julianDayNoDto,
	err = calEng.DateTimeToJulianDayNumber(
		targetDateTimeDto,
		calCycleCfg,
		1024,
		ePrefix)

	return julianDayNoDto, err
}

// GetYearDays - Returns the number of days in the year
// identified by input parameter 'year' under the Gregorian
// Calendar.
//
// If the year is a standard year, this method will return 365-days.
// If the year is a leap year, this method will return 365-days.
//
// For more information on the Gregorian Calendar and leap years,
// reference:
//
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//   https://www.timeanddate.com/date/leapyear.html
//
//
func (calGregUtil *CalendarGregorianUtility) GetYearDays(
	year int64) int {

	if calGregUtil.lock == nil {
		calGregUtil.lock = &sync.Mutex{}
	}

	calGregUtil.lock.Lock()

	defer calGregUtil.lock.Unlock()

	calGregMech := calendarGregorianMechanics{}

	isLeapYear := calGregMech.isLeapYear(year)

	if isLeapYear {
		return 366
	}

	return 365
}

// IsLeapYear - Returns a boolean value signaling whether the year
// value passed as an input parameter is a leap year (366-days)
// under the Gregorian Calendar.
//
// If the method returns 'true' the input parameter 'year' qualifies
// as a leap year consisting of 366-days. If the method returns 'false',
// the input parameter 'year' is a standard year consisting of 365-days.
//
// Methodology:
//
// In the Gregorian calendar, three criteria must be taken
// into account to identify leap years:
//
// 1. The year must be evenly divisible by 4;
//
// 2. If the year can also be evenly divided by 100, it is not
//    a leap year, unless...
//
//  3. The year is evenly divisible by 100 and the year is also
//     evenly divisible by 400. Then it is a leap year.
//
// According to these rules, the years 2000 and 2400 are leap years,
// while 1800, 1900, 2100, 2200, 2300, and 2500 are not leap years.
//
// For more information on the Gregorian Calendar and leap years,
// reference:
//
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//   https://www.timeanddate.com/date/leapyear.html
//
//
func (calGregUtil *CalendarGregorianUtility) IsLeapYear(
	year int64) bool {

	if calGregUtil.lock == nil {
		calGregUtil.lock = &sync.Mutex{}
	}

	calGregUtil.lock.Lock()

	defer calGregUtil.lock.Unlock()
	calGregMech := calendarGregorianMechanics{}

	return calGregMech.isLeapYear(year)
}

// DateTimeFromJulianDateTime - Converts a Julian Day Number/Time to its
// corresponding Gregorian Calendar date/time.
//
func (calGregUtil *CalendarGregorianUtility) DateTimeFromJulianDateTime(
	julianDayNumberTime JulianDayNoDto,
	ePrefix string) (ADateTimeDto, error) {

	if calGregUtil.lock == nil {
		calGregUtil.lock = &sync.Mutex{}
	}

	calGregUtil.lock.Lock()

	defer calGregUtil.lock.Unlock()

	ePrefix += "CalendarGregorianUtility.DateTimeFromJulianDateTime() "

	calGregMech := calendarGregorianMechanics{}

	calCyclesConfig := calGregMech.getCalendarCyclesConfig()

	calEng := CalendarEngines{}

	calEng.SetCodeDebug(true)

	return calEng.JulianDayNoToDateTime(
		julianDayNumberTime,
		calCyclesConfig,
		ePrefix)
}

// NumCalendarDaysForWholeYearsInterval - Computes the total
// number of 24-hour days in a period of years specified by
// input parameter 'wholeYearsInterval'. The number of total
// days is calculated in accordance with the Gregorian Calendar.
//
// Methodology:
//
// In the Gregorian calendar, three criteria must be taken
// into account to correctly identify leap years:
//
// 1. The year must be evenly divisible by 4;
//
// 2. If the year can also be evenly divided by 100, it is not
//    a leap year, unless...
//
//  3. The year is evenly divisible by 100 and the year is also
//     evenly divisible by 400. Then it is a leap year.
//
// According to these rules, the years 2000 and 2400 are leap years,
// while 1800, 1900, 2100, 2200, 2300, and 2500 are not leap years.
//
// For more information on the Gregorian Calendar and leap years,
// reference:
//
//   https://en.wikipedia.org/wiki/Gregorian_calendar
//   https://www.timeanddate.com/date/leapyear.html
//
// The input parameter 'wholeYearsInterval' is defined as a series of
// contiguous whole, or complete, years consisting of either 365-days
// or 366-days (in the case of leap years).
//
// No partial years should be included in this interval.
//
//
func (calGregUtil *CalendarGregorianUtility) NumCalendarDaysForWholeYearsInterval(
	wholeYearsInterval int64) (totalDays int64) {

	if calGregUtil.lock == nil {
		calGregUtil.lock = &sync.Mutex{}
	}

	calGregUtil.lock.Lock()

	defer calGregUtil.lock.Unlock()

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

		numOfCycles := wholeYearsInterval / 400

		totalDays = numOfCycles * 146097

		fmt.Printf("  Number of 400-Year Cycles: %v\n", numOfCycles)
		fmt.Printf("Number of Days in %v-Cycles: %v\n", numOfCycles, totalDays)

		wholeYearsInterval = wholeYearsInterval - (numOfCycles * 400)

		fmt.Printf("  Number of Remainder Years: %v\n", wholeYearsInterval)
		fmt.Println(separator)
		fmt.Println()

	}

	totalDays += wholeYearsInterval * 365

	leapDays := wholeYearsInterval / 4

	skipLeapDays := wholeYearsInterval / 100

	totalDays += leapDays - skipLeapDays

	fmt.Println(separator)
	fmt.Printf("Total Days In wholeYearsInterval: %v\n",
		totalDays)
	fmt.Println(separator)
	fmt.Println()

	return totalDays
}
