package datetime

import "sync"

// CalendarJulianUtility - This type contains methods
// used to process date arithmetic associated with the
// Julian Calendar
//
// References:
//  https://en.wikipedia.org/wiki/Julian_calendar
//  https://en.wikipedia.org/wiki/Conversion_between_Julian_and_Gregorian_calendars
//  https://www.fourmilab.ch/documents/calendar/
//
type CalendarJulianUtility struct {

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
func (calJulianUtil *CalendarJulianUtility) GetJulianDayNumber(
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

	if calJulianUtil.lock == nil {
		calJulianUtil.lock = &sync.Mutex{}
	}

	calJulianUtil.lock.Lock()

	defer calJulianUtil.lock.Unlock()

	ePrefix += "CalendarJulianUtility.GetJulianDayNumber() "

	err = nil

	julianDayNoDto = JulianDayNoDto{}

	calJulianMech := calendarJulianMechanics{}

	isLeapYear := calJulianMech.isLeapYear(targetYear)

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

	targetDateTimeDto,
		err =
		ADateTimeDto{}.New(
			CalendarSpec(0).Julian(),
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

	calCycleCfg := calJulianMech.getCalendarCyclesConfig()

	calEng := CalendarEngines{}

	calEng.SetCodeDebug(true)

	julianDayNoDto,
		err = calEng.DateTimeToJulianDayNumber(
		targetDateTimeDto,
		calCycleCfg,
		1024,
		ePrefix)

	return julianDayNoDto, err
}

// IsLeapYear - Determines whether the input parameter
// 'year' is a leap year under the Julian Calendar.
//
// If 'year' is a Julian leap year, this method returns 'true'.
//
// Note: This method should NOT be used to determine leap years
// for the Revised Julian Calendar or the Goucher-Parker
// calendar.
//
// Reference:
//   https://en.wikipedia.org/wiki/Julian_calendar
//
func (calJulianUtil *CalendarJulianUtility) IsLeapYear(
	year int64) bool {

	if calJulianUtil.lock == nil {
		calJulianUtil.lock = &sync.Mutex{}
	}

	calJulianUtil.lock.Lock()

	defer calJulianUtil.lock.Unlock()

	calJulianMech := calendarJulianMechanics{}

	return calJulianMech.isLeapYear(year)
}
