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

type calendarDateTimeUtility struct {
	lock   *sync.Mutex
}


// copyIn - Makes a deep copy of incoming CalendarDateTime instance
// 'incomingCalDTime' and stores the data in the internal member variables
// of 'oldCalDTime', the original CalendarDateTime instance. All member
// variable data values in 'oldCalDTime' will be overwritten.
//
// If this method completes successfully, the internal member variable
// data values for both 'oldCalDTime' and 'incomingCalDTime' will be
// identical.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  oldCalDTime         *CalendarDateTime
//     - A pointer to an instance of CalendarDateTime. This method will
//       perform a deep copy of member variable data values from input
//       parameter 'incomingCalDTime' to 'oldCalDTime'. All original
//       member variable data values contained in 'oldCalDTime' will be
//       overwritten.
//
//
//  incomingCalDTime    *CalendarDateTime
//     - A pointer to an instance of CalendarDateTime. This method will
//       perform a deep copy and transfer member variable data values
//       from 'incomingCalDTime' to 'oldCalDTime'.
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
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//
func (calDTimeUtil *calendarDateTimeUtility) copyIn(
	oldCalDTime *CalendarDateTime,
	incomingCalDTime *CalendarDateTime,
	ePrefix string) (err error) {


	if calDTimeUtil.lock == nil {
		calDTimeUtil.lock = new(sync.Mutex)
	}

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	err = nil

	ePrefix += "calendarDateTimeUtility.copyIn() "

	if oldCalDTime == nil {
		err = errors.New(ePrefix +
			"\nError: Input parameter 'oldCalDTime' is INVALID!\n" +
			"oldCalDTime=nil!")
		return err
	}

	if oldCalDTime.lock == nil {
		oldCalDTime.lock = new(sync.Mutex)
	}

	if incomingCalDTime == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'incomingCalDTime' is INVALID!\n" +
			"incomingCalDTime=nil!")
		return err
	}

	if incomingCalDTime.lock == nil {
		incomingCalDTime.lock = new(sync.Mutex)
	}

	oldCalDTime.dateTimeDto, err =
		incomingCalDTime.dateTimeDto.CopyOut(ePrefix + "- oldCalDTime.dateTimeDto ")

	if err != nil {
		return err
	}

	oldCalDTime.julianDayNumber, err =
		incomingCalDTime.julianDayNumber.CopyOut(ePrefix)

	if err != nil {
		return err
	}

	oldCalDTime.dateTimeFmt = incomingCalDTime.dateTimeFmt

	oldCalDTime.tag = incomingCalDTime.tag

	return err
}

// copyOut - Returns a deep copy of input parameter 'calDTime' which is a
// pointer to an instance of 'CalendarDateTime'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  calDTime            *CalendarDateTime
//     - A pointer to an instance of CalendarDateTime. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives.
//
//
//  ePrefix            string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newCalDTime         CalendarDateTime
//     - If successful this method returns a deep copy of the input
//       parameter, 'calDTime'. When the copy operation is completed
//       member variables contained in return value, 'newCalDTime'
//       will be identical to those contained in input parameter,
//      'calDTime'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (calDTimeUtil *calendarDateTimeUtility) copyOut(
	calDTime *CalendarDateTime,
	ePrefix string) ( newCalDTime CalendarDateTime, err error) {

	if calDTimeUtil.lock == nil {
		calDTimeUtil.lock = new(sync.Mutex)
	}

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.copyOut() "

	err = nil
	newCalDTime = CalendarDateTime{}

	if calDTime == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'calDTime' is a nil pointer!")

		return newCalDTime, err
	}

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	newCalDTime.dateTimeDto, err =
		calDTime.dateTimeDto.CopyOut(ePrefix + "- newCalDTime.dateTimeDto ")

	if err != nil {
		newCalDTime = CalendarDateTime{}
		return newCalDTime, err
	}

	newCalDTime.julianDayNumber, err =
		calDTime.julianDayNumber.CopyOut(ePrefix)

	if err != nil {
		newCalDTime = CalendarDateTime{}
		return newCalDTime, err
	}

	newCalDTime.dateTimeFmt = calDTime.dateTimeFmt

	newCalDTime.tag = calDTime.tag

	newCalDTime.lock = new(sync.Mutex)

	return newCalDTime, err
}

// getCalendarTimeOfDay - Returns the time of day encapsulated by
// the CalendarDateTime input parameter, 'calDTime'.
//
// This method returns the time of day contained in 'calDTime' as
// hours, minutes, seconds and nanoseconds. The returned value
// 'timeZone' specifies the time zone related to this time value.
//
//  **IMPORTANT** - No validity checking is performed on input
//                  parameter, 'calDTime'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  calDTime          *CalendarDateTime
//     - A pointer to a CalendarDateTime instance. The calendar
//       year data will be extracted from this instance.
//       **IMPORTANT** - No validity checking is performed on
//       input parameter, 'calDTime'.
//
//
//  ePrefix       string
//    This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  hour           int
//     - The hours component of the time value.
//
//
//  minute         int
//     - The minutes component of the time value.
//
//
//  second         int
//     - The seconds component of the time value.
//
//
//  nanosecond     int
//     - The nanoseconds component of the time value.
//
//
//  timeZone       TimeZoneDefinition
//     - The time zone associated with the returned
//       time value.
//
//
//  hasLeapSecond  bool
//     - If this return value is set to 'true', it signals that the time value
//       includes a "Leap Second". This parameter is rarely used and is almost
//       always set to 'false'.
//
//       A leap second is a one-second adjustment that is occasionally applied
//       to Coordinated Universal Time (UTC) in order to accommodate the
//       difference between precise time (as measured by atomic clocks) and
//       imprecise observed solar time (known as UT1 and which varies due to
//       irregularities and long-term slowdown in the Earth's rotation). If
//       this return parameter is set to 'true',      the time calculation
//       will assume the duration of the relevant 'day' is 24-hours plus one
//       second. Otherwise, the duration of a day is assumed to consist of
//       exactly 24-hours. For more information on the 'leap second',
//       reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  err      error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (calDTimeUtil *calendarDateTimeUtility) getCalendarTimeOfDay(
	calDTime *CalendarDateTime,
	ePrefix string) (
	hour int,
	minute int,
	second int,
	nanosecond int,
	timeZone TimeZoneDefinition,
	hasLeapSecond bool,
	err error ) {

	if calDTimeUtil.lock == nil {
		calDTimeUtil.lock = new(sync.Mutex)
	}

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.getCalendarYear() "
	hour = -1
	minute = -1
	second = -1
	nanosecond = -1
	timeZone = TimeZoneDefinition{}
	hasLeapSecond = false

	if calDTime == nil {
		err = errors.New(ePrefix + "\n" +
			"Input parameter 'calDTime' is INVALID!\n" +
			"calDTime == nil\n")
		return hour, minute, second, nanosecond, timeZone, hasLeapSecond, err
	}

	hour = calDTime.dateTimeDto.time.hour
	minute = calDTime.dateTimeDto.time.minute
	second = calDTime.dateTimeDto.time.second
	nanosecond = calDTime.dateTimeDto.time.nanosecond
	timeZone = calDTime.dateTimeDto.time.timeZone.CopyOut()
	hasLeapSecond = calDTime.dateTimeDto.time.hasLeapSecond

	return hour, minute, second, nanosecond, timeZone, hasLeapSecond, err
}

// getCalendarDate - Returns the date as year, month and day. Date
// information is extracted from the CalendarDateTime instance
// passed as input parameter, 'calDTime'
//
//  **IMPORTANT** - No validity checking is performed on input
//                  parameter, 'calDTime'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  calDTime          *CalendarDateTime
//     - A pointer to a CalendarDateTime instance. The member variables
//       of this instance will be populated according to the following
//       input parameters
//
//
//  ePrefix       string
//    This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  year            int64
//     - The properly formatted year value. Years are formatted
//       according to Astronomical Year Numbering or Common Era
//       era numbering depending on the value of returned parameter
//       'yearType' described below.
//
//
//  yearType        CalendarYearNumType
//     - 'yearType' designates the returned year value described above
//       as one of three year types:
//         1. Astronomical Year
//         2. BCE - Before Common Era
//         3. CE  - Common Era
//       For more information on Astronomical and Common Era Year
//       Numbering, reference:
//           Source File: datetime\calendaryearnumbertypeenum.go
//           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//           https://en.wikipedia.org/wiki/Common_Era
//
//
//  isLeapYear      bool
//     - If 'true' it signals that the return parameter 'year' is a
//       leap year.
//
//
//  calendarSystem  CalendarSpec
//     - An enumeration type designating the Calendar System
//       associated with the generated date. Reference:
//       Source File: datetime\calendarspecenum.go
//
//       Possible Calendar System values include:
//       Gregorian, Julian, Revised Julian, or
//       Revised Goucher-Parker.
//
//
//  month           int
//     - The month number
//
//  day             int
//     - The day number
//
//
//  err      error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calDTimeUtil *calendarDateTimeUtility) getCalendarDate(
	calDTime *CalendarDateTime,
	ePrefix string) (
	year int64,
	yearType CalendarYearNumType,
	isLeapYear bool,
	calendarSystem CalendarSpec,
	month int,
	day int,
	err error) {

	if calDTimeUtil.lock == nil {
		calDTimeUtil.lock = new(sync.Mutex)
	}

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.getCalendarDate() "

	year = 0
	yearType = CalYearType.None()
	month = -1
	day = -1
	isLeapYear = false
	calendarSystem = CalendarSpec(0).None()

	if calDTime == nil {
		err = errors.New(ePrefix + "\n" +
			"Input parameter 'calDTime' is INVALID!\n" +
			"calDTime == nil\n")
		return year, yearType, isLeapYear, calendarSystem, month, day, err
	}

	calDtMech := calendarDateTimeMechanics{}

	year,
	yearType,
	isLeapYear,
	calendarSystem,
	err =
		calDtMech.getCalendarYear(
			calDTime,
			ePrefix)

	if err != nil {
		return year, yearType, isLeapYear, calendarSystem, month, day, err
	}

	month = calDTime.dateTimeDto.date.month

	day = calDTime.dateTimeDto.date.day

	return year, yearType, isLeapYear, calendarSystem, month, day, err
}

// generateDateTimeStr - Converts input years, months, days, hours,
// minutes, seconds and subMicrosecondNanoseconds to a formatted date time string
// the golang format string passed in input parameter 'dateFormatStr'.
//
func (calDTimeUtil *calendarDateTimeUtility) generateDateTimeStr(
	year int64,
	month,
	days int,
	usDayOfWeekNumber UsDayOfWeekNo,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	dateFormatStr,
	ePrefix string) (string, error) {

	if calDTimeUtil.lock == nil {
		calDTimeUtil.lock = new(sync.Mutex)
	}

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.generateDateTimeStr() "

	var err error

	/*	replacementTokens := map[string]string{
			"!YearFourDigit!":"",
			"!YearTwoDigit!":"",
			"!YearOneDigit!":"",
			"!DayOfWeek!":"",
			"!DayOfWeekAbbrv!":"",
			"!MonthName!":"",
			"!MonthNameAbbrv!":"",
			"!MonthNumberTwoDigit!":"",
			"!MonthNumberOneDigit!":"",
			"!DateDayTwoDigit!":"",
			"!DateDayLeadUnderScore!":"",
			"!DateDayOneDigit!":"",
			"!HourTwentyFourTwoDigit!":"",
			"!HourTwelveTwoDigit!":"",
			"!HourTwelveOneDigit!":"",
			"!AMPMUpperCase!",
			"!AMPMLowerCase!",
			"!MinutesTwoDigit!":"",
			"!MinutesOneDigit!":"",
			"!SecondsTwoDigit!":"",
			"!SecondsOneDigit!":"",
			"!NanosecondsTrailingZeros!":"",
			"!NanosecondsNoTrailingZeros!":"",
			"!MillisecondsTrailingZeros!":"",
			"!MillisecondsNoTrailingZeros!":"",
			"!MicrosecondsTrailingZeros!":"",
			"!MicrosecondsNoTrailingZeros!":"",
			"!OffsetUTC!":"",
			"!TimeZone!":"",
		}
	*/

	replacementTokens := map[string]string{}

	dtMech := DTimeNanobot{}

	resultStr := dtMech.PreProcessDateFormatStr(dateFormatStr)

	calDtMech := calendarDateTimeMechanics{}

	resultStr, err = calDtMech.processDayOfWeek(
		resultStr,
		usDayOfWeekNumber,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr = calDtMech.processYears(
		resultStr,
		year,
		replacementTokens)

	resultStr, err = calDtMech.processMonths(
		resultStr,
		month,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processDateDay(
		resultStr,
		days,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processHours(
		resultStr,
		hours,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processMinutes(
		resultStr,
		minutes,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processSeconds(
		resultStr,
		seconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processNanoseconds(
		resultStr,
		nanoseconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processMicroseconds(
		resultStr,
		nanoseconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processMilliseconds(
		resultStr,
		nanoseconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = calDtMech.processAmPm(
		resultStr,
		hours,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr = calDtMech.processOffset(
		resultStr,
		-999,
		-999,
		replacementTokens)

	resultStr = calDtMech.processTimeZone(
		resultStr,
		"UTC",
		replacementTokens)

	for key, value := range replacementTokens {

		resultStr = strings.Replace(resultStr,key,value,1)

	}

	return resultStr, err
}

// newCalDateTimeFromComponents - Returns a new instance of
// 'CalendarDateTime'. Data fields are populated from components
// passed by the calling function.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  calendarSystem           CalendarSpec
//     - An enumeration type designating the Calendar System
//       associated with the generated date. Reference:
//       Source File: datetime\calendarspecenum.go
//
//       Possible Calendar System values include:
//       Gregorian, Julian, Revised Julian, or Revised Goucher-Parker.
//
//       Possible Enumeration Values:
//         CalendarSpec(0).Gregorian()
//         CalendarSpec(0).Julian()
//         CalendarSpec(0).RevisedJulian()
//         CalendarSpec(0).RevisedGoucherParker()
//
//
//  year                     int64
//     - An int64 value containing the year number. Note that this year
//       value should be interpreted in the context the year type return
//       parameter discussed below. This year value could be formatted
//       according to the Astronomical Year Numbering System or the
//       Common Era Year Numbering System.
//
//
//  yearType                 CalendarYearNumType
//     - The year number type associated with the return value 'yearValue'
//       described above. 'yearType' classifies return parameter 'yearValue'
//       as one of three year types:
//
//         1. Astronomical Year
//         2. BCE - Before Common Era
//         3. CE  - Common Era
//
//       For more information on Astronomical and Common Era Year
//       Numbering, reference:
//           Source File: datetime\calendaryearnumbertypeenum.go
//           https://en.wikipedia.org/wiki/Astronomical_year_numbering
//           https://en.wikipedia.org/wiki/Common_Era
//
//
//  month                    int
//     - The month number
//
//
//  day                      int
//    - The day number. This is the day number within the
//      the 'month' identified in the 'month' input parameter,
//      above.
//
//
//  dateHasLeapSecond        bool
//     - If this parameter is set to 'true', it signals that the day identified
//       by year, month and day input parameters contains a leap second.
//       This parameter is rarely used and is almost always set to 'false'.
//
//       A leap second is a one-second adjustment that is occasionally applied
//       to Coordinated Universal Time (UTC) in order to accommodate the
//       difference between precise time (as measured by atomic clocks) and
//       imprecise observed solar time (known as UT1 and which varies due to
//       irregularities and long-term slowdown in the Earth's rotation). If
//       this return parameter is set to 'true', time and duration
//       calculations will assume the duration of the relevant 'day' is
//       24-hours plus one second. Otherwise, the duration of the day is
//       assumed to consist of exactly 24-hours. For more information on
//       the 'leap second', reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  hour                     int
//     - The hour component of a time value.
//
//
//  minute                   int
//     - The minute component of a time value.
//
//  second                   int
//     - The second component of a time value.
//
//
//  nanosecond               int
//       The nanosecond component of a time value.
//
//
//  timeZoneLocation         string
//     - A string containing the name of a valid time zone. Usually, this is
//       an IANA Time Zone as shown in the examples below:
//
//        Time Zone Strings           Time Zone
//        -------------------------------------
//
//       "America/New_York"         USA Eastern Time Zone
//       "America/Chicago"          USA Central Time Zone
//       "America/Denver"           USA Mountain Time Zone
//       "America/Los_Angeles"      USA Pacific Time Zone
//       "Local"                    The time zone on the host computer
//                                    where this code is executed.
//       "UTC"                      Coordinated Universal Time
//                                    https://en.wikipedia.org/wiki/Coordinated_Universal_Time
//       Reference:
//           https://golang.org/pkg/time/
//           https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//           https://en.wikipedia.org/wiki/Tz_database
//           https://en.wikipedia.org/wiki/List_of_military_time_zones
//           https://www.iana.org/time-zones
//
//
//  julianDayNo       *big.Float
//    - A properly formed Julian Day Number formatted as a floating-point value.
//      This Julian Day Number must equate the year, month, day, hours, minutes,
//      seconds, and nanoseconds provided as input parameters above.
//      For more information on Julian Day Numbers, reference:
//          https://en.wikipedia.org/wiki/Julian_day
//
//
//  dateTimeFmt        string
//    - This string contains the date/time format which will be used to
//      to format date/time output values. Example:
//          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  tag                      string
//     - A text description to be associated with the newly created CalendarDateTime
//       instance created by this method.
//
//
//  ePrefix                   string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
func (calDTimeUtil *calendarDateTimeUtility) newCalDateTimeFromComponents(
	calendarSystem CalendarSpec,
	year int64,
	yearType CalendarYearNumType,
	month int,
	day int,
	dateHasLeapSecond bool,
	hour,
	minute,
	second,
	nanosecond int,
	timeZoneLocation string,
	julianDayNo *big.Float,
	dateTimeFmt string,
	tag string,
	ePrefix string) (
	calDateTime CalendarDateTime,
	err error) {

	if calDTimeUtil.lock == nil {
		calDTimeUtil.lock = new(sync.Mutex)
	}

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.newCalDateTimeFromComponents() "

	calDateTime = CalendarDateTime{}

	err = nil

	if julianDayNo == nil {
		err = errors.New(ePrefix + "\n" +
			"Input parameter 'julianDayNo' is INVALID!\n" +
			"julianDayNo == nil\n")

		return calDateTime, err
	}

	if timeZoneLocation == "" {
		err = errors.New(ePrefix + "\n" +
			"Input parameter 'timeZoneLocation' is INVALID!\n" +
			"timeZoneLocation == \"\"\n")
		return calDateTime, err
	}


	calDateTime.dateTimeDto, err = ADateTimeDto{}.New(
		calendarSystem,
		year,
		yearType,
		month,
		day,
		dateHasLeapSecond,
		hour,
		minute,
		second,
		nanosecond,
		timeZoneLocation,
		dateTimeFmt,
		"",
		ePrefix)

	if err != nil {
		calDateTime = CalendarDateTime{}
		return calDateTime, err
	}

	calDateTime.julianDayNumber, err =
		JulianDayNoDto{}.NewFromJulianDayNoTime(
			julianDayNo,
			dateHasLeapSecond,
			ePrefix)

	if err != nil {
		calDateTime = CalendarDateTime{}
		return calDateTime, err
	}

	dtMech := DTimeNanobot{}

	calDateTime.dateTimeFmt = dtMech.PreProcessDateFormatStr(dateTimeFmt)

	calDateTime.tag = tag

	calDateTime.lock = new(sync.Mutex)

	return calDateTime, err
}

// setCalDateTime - populates a CalendarDateTime instance.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  calDTime          *CalendarDateTime
//     - A pointer to a CalendarDateTime instance. The member variables
//       of this instance will be populated according to the following
//       input parameters
//
//
//  year              int64
//    - The year number expressed as an int64 value. This 'years' value
//      should be formatted using Astronomical Year Numbering; that is,
//      a year numbering system which includes year zero. Year values which
//      are less than -4713, using the using Astronomical Year Numbering System,
//      are invalid and will generate an error.
//
//
//  month              int
//    - The month number. This is a month on the Julian Calendar.
//
//
//  day                int
//    - The day number
//
//
//  hours              int
//    - The hour number expressed on a 24-hour time scale.
//      Example: 3:00PM is passed as the hour 15
//
//
//  minutes            int
//    - The minutes number
//
//
//  seconds            int
//    - The number of seconds
//
//
//  nanoseconds        int
//    - The number of nanoseconds
//
//
//  applyLeapSecond    bool
//     - If this parameter is set to 'true', it signals that the time value
//       includes a "Leap Second". This parameter is rarely used and is almost
//       always set to 'false'.
//
//       A leap second is a one-second adjustment that is occasionally applied
//       to Coordinated Universal Time (UTC) in order to accommodate the
//       difference between precise time (as measured by atomic clocks) and
//       imprecise observed solar time (known as UT1 and which varies due to
//       irregularities and long-term slowdown in the Earth's rotation). If
//       this return parameter is set to 'true',      the time calculation
//       will assume the duration of the relevant 'day' is 24-hours plus one
//       second. Otherwise, the duration of a day is assumed to consist of
//       exactly 24-hours. For more information on the 'leap second',
//       reference:
//          https://en.wikipedia.org/wiki/Leap_second
//
//
//  timeZoneLocation   string
//    - This string identifies the Time Zone associated with parameters
//      hours, minutes, seconds and nanoseconds.
//
//
//  calendarSystem     CalendarSpec
//    - This is a Calendar Specification designating the calendar system
//      which will be used to create a new instance of CalendarDateTime.
//
//
//  yearNumberMode     CalendarYearNumMode
//    - This is Year Numbering Mode designating the type of year number
//      will be used when displaying date time value. The choices are:
//             CalendarYearNumMode(0).Astronomical()
//             CalendarYearNumMode(0).CommonEra()
//
//
//  dateTimeFmt        string
//    - This string contains the date/time format which will be used to
//      to format date/time output values. Example:
//          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  ePrefix            string
//    This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods.
//
func (calDTimeUtil *calendarDateTimeUtility) setCalDateTime(
	calDTime *CalendarDateTime,
	year int64,
	month,
	day,
	hour,
	minute,
	second,
	nanosecond int,
	applyLeapSecond bool,
	timeZoneLocation string,
	calendar CalendarSpec,
	yearNumberType CalendarYearNumType,
	dateTimeFmt,
	ePrefix string) (err error) {

	if calDTimeUtil.lock == nil {
		calDTimeUtil.lock = new(sync.Mutex)
	}

	calDTimeUtil.lock.Lock()

	defer calDTimeUtil.lock.Unlock()

	ePrefix += "calendarDateTimeUtility.setCalDateTime() "

	if calDTime == nil {
		err = errors.New(ePrefix + "\n" +
			"Input parameter 'calDTime' is INVALID!\n" +
			"calDTime == nil\n")
		return err
	}

	calDtMech := calendarDateTimeMechanics{}

	err = calDtMech.empty(calDTime, ePrefix)

	if err != nil {
		return err
	}

	if !calendar.XIsValid()  {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "calendar",
			inputParameterValue: calendar.String() ,
			errMsg:              "'calendar' is INVALID!",
			err:                 nil,
		}
	}

	timeZone := TimeZoneDefinition{}

	tzDefUtil := timeZoneDefUtility{}

	baseDateTime := time.Now().UTC()

	err = tzDefUtil.setFromTimeZoneName(
		&timeZone,
		baseDateTime,
		TimeZoneConversionType(0).Relative(),
		timeZoneLocation,
		ePrefix)

	if err != nil {
		return err
	}

	var jDayNoDto JulianDayNoDto

	var calBaseData ICalendarBaseData

	if !yearNumberType.XIsValid() {
		return &InputParameterError{
			ePrefix:             ePrefix,
			inputParameterName:  "yearNumberType",
			inputParameterValue: strconv.FormatInt(int64(yearNumberType.XValueInt()), 10) ,
			errMsg:              "'yearNumberType' is INVALID!",
			err:                 nil,
		}
	}

	calMech := calendarMechanics{}

	if calendar == CalendarSpec(0).Julian() {

		jDayNoDto, err = calMech.julianCalendarDateJulianDayNo(
			year,
			month,
			day,
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

		if err != nil {
			return err
		}

		// TODO - Build Julian Calendar Base Data
		calBaseData = &CalendarGregorianBaseData{}

	} else if calendar == CalendarSpec(0).Gregorian() {

		calGregUtil := CalendarGregorianUtility{}

		jDayNoDto,
		err = calGregUtil.GetJulianDayNumber(
			year,
			month,
			day,
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

		calBaseData = &CalendarGregorianBaseData{}

	} else if calendar == CalendarSpec(0).RevisedGoucherParker() {

		jDayNoDto, err = calMech.revisedGoucherParkerToJulianDayNo(
			year,
			month,
			day,
			hour,
			minute,
			second,
			nanosecond,
			ePrefix)

		if err != nil {
			return err
		}

		// TODO - Build Revised Goucher-Parker Calendar Base Data
		calBaseData = &CalendarGregorianBaseData{}

	} else if calendar == CalendarSpec(0).RevisedJulian() {
	// TODO - Fix Revised Julian Calendar initialization

		// TODO - Build Revised Julian Calendar Base Data
		calBaseData = &CalendarGregorianBaseData{}

	} else {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Invalid Calendar Specification.\n" +
			"Calendar Specification='%v'\n",
			calendar.String())

		return err
	}

	var isLeapYear bool

	isLeapYear, err =
		calBaseData.IsLeapYear(
			year,
			yearNumberType,
			ePrefix)

	if err != nil {
		return err
	}

	calUtil := CalendarUtility{}

	err = calUtil.IsValidDateTimeComponents(
		isLeapYear,
		month,
		day,
		hour,
		minute,
		second,
		nanosecond,
		ePrefix + " - dateTimeDto elements - ")

	if err != nil {
		return err
	}

	calDTime.dateTimeDto.date.calendarBaseData = calBaseData.New()
	calDTime.dateTimeDto.date.astronomicalYear = year
	calDTime.dateTimeDto.date.month = month
	calDTime.dateTimeDto.date.day = day
	calDTime.dateTimeDto.time.hour = hour
	calDTime.dateTimeDto.time.minute = minute
	calDTime.dateTimeDto.time.second = second
	calDTime.dateTimeDto.time.nanosecond = nanosecond
	calDTime.dateTimeDto.date.hasLeapSecond = applyLeapSecond
	calDTime.julianDayNumber, err = jDayNoDto.CopyOut(ePrefix)

	if err != nil {
		calDtMech := calendarDateTimeMechanics{}
		_ = calDtMech.empty(calDTime, "")
		return err
	}

	dtMech := DTimeNanobot{}

	calDTime.dateTimeFmt =
		dtMech.PreProcessDateFormatStr(dateTimeFmt)


	return nil
}

