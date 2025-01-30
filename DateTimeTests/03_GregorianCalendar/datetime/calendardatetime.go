package datetime

import (
	"math/big"
	"sync"
)

type CalendarDateTime struct {
	dateTimeDto     ADateTimeDto       // Date Time data fields
	julianDayNumber JulianDayNoDto     // Encapsulates Julian Day Number/Time
	dateTimeFmt        string          // Date Time Format String. Empty string or default is
	//                                 //   "2006-01-02 15:04:05.000000000 -0700 MST"
	tag                string          // Tag Description string
	lock              *sync.Mutex      // Used for coordinating thread safe operations.
}

// CopyIn - Populates the current CalendarDateTime instance with a deep copy
// of member data elements extracted from the the incoming CalendarDateTime
// instance, 'incomingCalDTime'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  incomingCalDTime  *CalendarDateTime
//     - Data elements from parameter, 'incomingCalDTime' will be used
//       to populate the current CalendarDateTime instance. When successfully
//       completed, all member data variables from 'incomingCalDTime' and the
//       current CalendarDateTime instance will be identical.
//
//
//  ePrefix           string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calDTime *CalendarDateTime) CopyIn(
	incomingCalDTime *CalendarDateTime,
	ePrefix string) error {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.CopyIn() "

	calDTimeUtil := calendarDateTimeUtility{}

	return calDTimeUtil.copyIn(calDTime, incomingCalDTime, ePrefix)
}

// CopyOut - Returns a deep copy of the current CalendarDateTime
// instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix       string
//   This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  CalendarDateTime
//     - A deep copy of the current CalendarDateTime instance.
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calDTime *CalendarDateTime) CopyOut(
	ePrefix string) (CalendarDateTime, error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.CopyOut() "

	calDTimeUtil := calendarDateTimeUtility{}

	newCalDTime := CalendarDateTime{}

	return calDTimeUtil.copyOut(
		&newCalDTime,
		ePrefix)
}

// GetCalendarSpecification - Returns the Calendar Specification
// associated with this CalendarDateTime instance.
//
// Input Value
//
//  ePrefix       string
//   This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods.
//
func (calDTime *CalendarDateTime) GetCalendarSpecification(
	ePrefix string) (CalendarSpec, error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.GetCalendarSpecification() "

	calDtMech := calendarDateTimeMechanics{}

	var err error

	_, err =calDtMech.testCalendarDateTimeValidity(
		calDTime,
		ePrefix)

	if err != nil {
		return CalendarSpec(0).None(), err
	}

	return calDTime.dateTimeDto.date.calendarBaseData.GetCalendarSpecification(), nil
}

// GetDate - Returns the date as year, month and day.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//       Gregorian, Julian, Revised Julian, or Revised Goucher-Parker.
//
//       Possible Enumeration Values:
//         CalendarSpec(0).Gregorian()
//         CalendarSpec(0).Julian()
//         CalendarSpec(0).RevisedJulian()
//         CalendarSpec(0).RevisedGoucherParker()
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
func (calDTime *CalendarDateTime) GetDate(
	ePrefix string) (
	year int64,
	yearType CalendarYearNumType,
	isLeapYear bool,
	calendarSystem CalendarSpec,
	month int,
	day int,
	err error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.GetDate() "

	year = 0
	yearType = CalYearType.None()
	month = -1
	day = -1
	calendarSystem = CalendarSpec(0).None()

	calDtMech := calendarDateTimeMechanics{}

	_, err = calDtMech.testCalendarDateTimeValidity(
		calDTime,
		ePrefix + "- Current CalendarDateTime instance. ")

	if err != nil {
		return year, yearType, isLeapYear, calendarSystem, month, day, err
	}

	calDTimeUtil := calendarDateTimeUtility{}

	year,
	yearType,
	isLeapYear,
	calendarSystem,
	month,
	day,
	err = calDTimeUtil.getCalendarDate(
		calDTime,
		ePrefix)

	return year, yearType, isLeapYear, calendarSystem, month, day, err
}

// GetDateTimeStr - Returns the equivalent date time string
// reflecting the date time value of the current CalendarDateTime
// instance. The Date Time Format was previously specified and
// is extracted from internal member variable, 'calDTime.dateTimeFmt'.
//
// To manage the Date Time Format for the current CalendarDateTime
// instance, reference method CalendarDateTime.SetDateTimeFormat().
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//  string
//     - If successful, this method will return a string containing
//       the date time formatted according to the Date Time Format
//       associated with this CalendarDateTime instance. Reference
//       method CalendarDateTime.SetDateTimeFormat().
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calDTime *CalendarDateTime) GetDateTimeStr(
	ePrefix string) (string, error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.GetDateTimeStr() "

	calDtMech := calendarDateTimeMechanics{}

	var err error

	_, err =calDtMech.testCalendarDateTimeValidity(
		calDTime,
		ePrefix)

	if err != nil {
		return "", err
	}

	calDTimeUtil := calendarDateTimeUtility{}

	var usDayOfWeekNo UsDayOfWeekNo

	usDayOfWeekNo,
	err = calDTime.dateTimeDto.date.GetUsDayOfWeekNo(
		calDTime.julianDayNumber, ePrefix)

	if err != nil {
		return "", err
	}


	return calDTimeUtil.generateDateTimeStr(
		calDTime.dateTimeDto.date.astronomicalYear,
		calDTime.dateTimeDto.date.month,
		calDTime.dateTimeDto.date.day,
		usDayOfWeekNo,
		calDTime.dateTimeDto.time.hour,
		calDTime.dateTimeDto.time.minute,
		calDTime.dateTimeDto.time.second,
		calDTime.dateTimeDto.time.nanosecond,
		calDTime.dateTimeFmt,
		ePrefix)
}

// GetDaysInYear - Returns the the number of days in the year identified
// by this CalendarDateTime instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//  int
//     - An integer value representing the number of days in the
//       the year specified by this CalendarDateTime instance.
//
//
//  err      error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calDTime *CalendarDateTime) GetDaysInYear(
	ePrefix string) (int, error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.GetDaysInYear() "

	calDtMech := calendarDateTimeMechanics{}

	var err error

	_, err =calDtMech.testCalendarDateTimeValidity(
		calDTime,
		ePrefix)

	if err != nil {
		return 0, err
	}

	return calDTime.dateTimeDto.date.calendarBaseData.GetDaysInLeapYear(), nil
}

// GetOrdinalDayNo - Returns the Ordinal Day Number for this Calendar
// Date Time instance. This Ordinal Day Number is the day number within
// the year specified by the current CalendarDateTime instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - An integer value representing the Ordinal Day Number within
//       the year specified by this CalendarDateTime instance.
//
//
//  err      error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calDTime *CalendarDateTime) GetOrdinalDayNo(
	ePrefix string) (int, error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.GetOrdinalDayNo() "

	calDtMech := calendarDateTimeMechanics{}

	var err error

	_, err =calDtMech.testCalendarDateTimeValidity(
		calDTime,
		ePrefix)

	if err != nil {
		return 0, err
	}

	return calDTime.dateTimeDto.date.calendarBaseData.GetOrdinalDayNumber(
		calDTime.dateTimeDto.date.GetIsLeapYear(),
		calDTime.dateTimeDto.date.month,
		calDTime.dateTimeDto.date.day,
		ePrefix)
}

// GetRemainingDaysInYear - Returns the number of days remaining
// in the year.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix       string
//    This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - The number of days remaining in the year. For example, January 1, 2021,
//       is the first day of a non-leap year or standard year containing 365 days.
//       Therefore there are 364-days remaining in this year.
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calDTime *CalendarDateTime) GetRemainingDaysInYear(
	ePrefix string) (int, error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.GetRemainingDaysInYear() "

	calDtMech := calendarDateTimeMechanics{}

	var err error

	_, err =calDtMech.testCalendarDateTimeValidity(
		calDTime,
		ePrefix)

	if err != nil {
		return 0, err
	}

	return calDTime.dateTimeDto.date.calendarBaseData.GetRemainingDaysInYear(
		calDTime.dateTimeDto.date.astronomicalYear,
		calDTime.dateTimeDto.date.yearNumType,
		calDTime.dateTimeDto.date.month,
		calDTime.dateTimeDto.date.day,
		ePrefix)
}

// GetTag - Returns the tag description string associated with
// the current instance of CalendarDateTime. The tag description
// is stored in internal member variable, 'CalendarDateTime.tag'.
// This member variable is used to store any descriptive information
// required by the user.
//
// NOTE: To insert descriptive information in this tag member
// variable, call method TimeTransferDto.SetTagDescription().
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//          -- NONE --
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string - The returned string contains the tag description string
//           stored in internal member variable, 'CalendarDateTime.tag'.
//
func (calDTime CalendarDateTime) GetTag() string {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	return calDTime.tag
}

// GetTime - Returns the time of day encapsulated by this CalendarDateTime
// instance as hours, minutes, seconds and nanoseconds. The returned
// value 'timeZone' specifies the time zone related to this time value.
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
func (calDTime *CalendarDateTime) GetTime(
	ePrefix string) (
	hour int,
	minute int,
	second int,
	nanosecond int,
	timeZone TimeZoneDefinition,
	hasLeapSecond bool,
	err error ) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.GetTime() "

	hour = -1
	minute = -1
	second = -1
	nanosecond = -1
	timeZone = TimeZoneDefinition{}
	hasLeapSecond = false

	calDtMech := calendarDateTimeMechanics{}

	_, err = calDtMech.testCalendarDateTimeValidity(
		calDTime,
		ePrefix + "- Testing calDTime (CalendarDateTime) instance validity. ")

	if err != nil {
		return hour, minute, second, nanosecond, timeZone, hasLeapSecond, err
	}

	calDTimeUtil := calendarDateTimeUtility{}

	hour,
	minute,
	second,
	nanosecond,
	timeZone,
	hasLeapSecond,
	err = calDTimeUtil.getCalendarTimeOfDay(
		calDTime,
		ePrefix)

	return hour, minute, second, nanosecond, timeZone, hasLeapSecond, err
}

// GetYear - Returns the calendar year value.
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//  err             error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (calDTime *CalendarDateTime) GetYear(
	ePrefix string) (
	year int64,
	yearType CalendarYearNumType,
	isLeapYear bool,
	calendarSystem CalendarSpec,
	err error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.GetYearWithType() "

	year = 0

	yearType = CalYearType.None()

	calDtMech := calendarDateTimeMechanics{}

	_, err = calDtMech.testCalendarDateTimeValidity(
		calDTime,
		ePrefix + "- Current CalendarDateTime instance. ")

	if err != nil {
		return year, yearType, isLeapYear, calendarSystem, err
	}

	year,
	yearType,
	isLeapYear,
	calendarSystem,
	err =
		calDtMech.getCalendarYear(
			calDTime,
			ePrefix)

	return year, yearType, isLeapYear, calendarSystem, err
}

// IsValidInstance - Tests the current CalendarDateTime instance
// for validity.
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//  bool
//     - If the current CalendarDateTime instance is valid and
//       properly initialized, this method will return 'true'.
//       If the current CalendarDateTime instance is invalid,
//       a value of 'false' will be returned.
//
func (calDTime *CalendarDateTime) IsValidInstance() bool {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	calDtMech := calendarDateTimeMechanics{}

	var isValid bool

	isValid, _ = calDtMech.testCalendarDateTimeValidity(
		calDTime,
		"")

	return isValid
}

// IsValidInstanceError - Similar to method CalendarDateTime.IsValidInstance().
// However, this method returns a error message.
//
// This method will test the current CalendarDateTime instance for validity
// and return an error if the instance is invalid.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix       string
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
//     - This method will analyze and test the current instance of CalendarDateTime
//       for validity. If the instance is invalid, a type 'error' will be returned
//       encapsulating an appropriate error message. The error message will be prefixed
//       with the error prefix string (ePrefix) passed as an input parameter.
//
//       If the current CalendarDateTime is valid an properly initialized, this
//       returned error type will be set to 'nil'.
//
func (calDTime *CalendarDateTime) IsValidInstanceError(
	ePrefix string) error {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.IsValidInstanceError() "

	var err error

	calDtMech := calendarDateTimeMechanics{}

	_, err = calDtMech.testCalendarDateTimeValidity(
		calDTime,
		ePrefix)

	return err
}


// NewGregorianDate - Creates a new instance of 'CalendarDateTime' formatted
// for a Gregorian Date Time.
//
// Taken collectively, the 'input' parameters years, months, days, hours,
// minutes, seconds and subMicrosecondNanoseconds represents a Gregorian date/time using
// the time zone specified by input parameter 'timeZoneLocation'.
//
// Note That the default Year Numbering Mode is 'Astronomical'. For other Year
// Numbering Modes, use method CalendarDateTime.SetYearNumberMode().
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  year               int64
//    - The year number expressed as an int64 value. This 'years' value
//      should be formatted using Astronomical Year Numbering; that is,
//      a year numbering system which includes year zero.
//
//      Under the Astronomical Year Numbering System, the date
//      January 1, year 1 is immediately preceded by the date
//      December 31, year 0. Years prior to year zero (0) are negative
//      values prefixed with a minus sign. For example, years before
//      year zero (0) are numbered as -1, -2, -3 etc.
//
//      As its name implies, Astronomical Year Numbering is frequently
//      used in astronomical calculations. For additional information on
//      the Astronomical Year Numbering System, reference:
//        https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
//
//  month              int
//    - The month number. This is a month on the Gregorian Calendar.
//
//
//  day                int
//    - The day number. This is a day number on the Gregorian Calendar.
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
//      hours, minutes, seconds and nanoseconds. Example: "UTC"
//
//
//  dateTimeFmt        string
//    - This string contains the date/time format which will be used to
//      to format date/time output values. Example:
//          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  ePrefix            string
//    - This is an error prefix which is included in all returned
//      error messages. Usually, it contains the names of the calling
//      method or methods.
//
func (calDTime CalendarDateTime) NewGregorianDate(
	year int64,
	yearNumberType CalendarYearNumType,
	month,
	day,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	applyLeapSecond bool,
	timeZoneLocation string,
	dateTimeFmt string,
	ePrefix string) (calDateTime CalendarDateTime, err error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.NewGregorianDate() "

	calDTimeUtil := calendarDateTimeUtility{}

	calDateTime = CalendarDateTime{}

	err = calDTimeUtil.setCalDateTime(
		&calDateTime,
		year,
		month,
		day,
		hours,
		minutes,
		seconds,
		nanoseconds,
		applyLeapSecond,
		timeZoneLocation,
		CalSpec.Gregorian(),
		yearNumberType,
		dateTimeFmt,
		ePrefix)

	return calDateTime, err
}

// NewJulianDate - Creates a new instance of 'CalendarDateTime' formatted
// for a Julian Date Time.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  years              int64
//    - The year number expressed as an int value. This 'years' value
//      should be formatted using Astronomical Year Numbering; that is,
//      a year numbering system which includes year zero. Year values which
//      are less than -4713, using the using Astronomical Year Numbering System,
//      are invalid and will generate an error. This is a year on the Julian
//      calendar.
//
//  months             int
//    - The month number. This is a month on the Julian Calendar.
//
//  days               int
//    - The day number
//
//  hours              int
//    - The hour number expressed on a 24-hour time scale.
//      Example: 3:00PM is passed as the hour 15
//
//  minutes            int
//    - The minutes number
//
//  seconds            int
//    - The number of seconds
//
//  nanoseconds        int
//    - The number of nanoseconds
//
//  timeZoneLocation   string
//    - This string identifies the Time Zone associated with parameters
//      hours, minutes, seconds and nanoseconds.
//
//  dateTimeFmt        string
//    - This string contains the date/time format which will be used to
//      to format date/time output values. Example:
//          "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  ePrefix            string
//    This is an error prefix which is included in all returned
//    error messages. Usually, it contains the names of the calling
//    method or methods.
//
func (calDTime CalendarDateTime) NewJulianDate(
	year int64,
	yearNumberType CalendarYearNumType,
	month,
	day,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	timeZoneLocation string,
	dateTimeFmt string,
	ePrefix string) (calDateTime CalendarDateTime, err error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.NewJulianDate() "

	calDTimeUtil := calendarDateTimeUtility{}

	calDateTime = CalendarDateTime{}

	err = calDTimeUtil.setCalDateTime(
		&calDateTime,
		year,
		month,
		day,
		hours,
		minutes,
		seconds,
		nanoseconds,
		false,
		timeZoneLocation, CalendarSpec(0).Julian(),
		yearNumberType,
		dateTimeFmt,
		ePrefix)

	return calDateTime, err
}

// NewCalDateTime - Creates and returns a populated 'CalendarDateTime' instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  calendarSystem     CalendarSpec
//    - This is a Calendar Specification designating the calendar system
//      which will be used to create a new instance of CalendarDateTime.
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
//  yearNumberType     CalendarYearNumType
//     - 'yearNumberType' classifies input parameter 'year' as one
//        of three year types:
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
func (calDTime CalendarDateTime) NewCalDateTime(
	calendarSystem CalendarSpec,
	year int64,
	yearNumberType CalendarYearNumType,
	month,
	day,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	applyLeapSecond bool,
	timeZoneLocation string,
	dateTimeFmt string,
	ePrefix string) (
	calDateTime CalendarDateTime,
	err error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.NewCalDateTime() "

	calDTimeUtil := calendarDateTimeUtility{}

	calDateTime = CalendarDateTime{}

	err = calDTimeUtil.setCalDateTime(
		&calDateTime,
		year,
		month,
		day,
		hours,
		minutes,
		seconds,
		nanoseconds,
		applyLeapSecond,
		timeZoneLocation,
		calendarSystem,
		yearNumberType,
		dateTimeFmt,
		ePrefix)

	return calDateTime, err
}

// NewCalDateTimeFromComponents - Creates a new instance of
// CalendarDateTime from individual data components.
//
//
// ------------------------------------------------------------------------
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
//  CalendarDateTime
//     - If successful, this method will return a properly initialized
//       CalendarDateTime instance constructed from the input parameters.
//
//
//  error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (calDTime CalendarDateTime) NewCalDateTimeFromComponents(
	calendarSystem CalendarSpec,
	year int64,
	yearType CalendarYearNumType,
	month int,
	day int,
	dateHasLeapSecond bool,
	hour int,
	minute int,
	second int,
	nanosecond int,
	timeZoneLocation string,
	julianDayNo *big.Float,
	dateTimeFmt string,
	tag string,
	ePrefix string) (CalendarDateTime, error) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.NewCalDateTimeFromComponents() "

	calDTimeUtil := calendarDateTimeUtility{}

	return calDTimeUtil.newCalDateTimeFromComponents(
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
		julianDayNo,
		dateTimeFmt,
		tag,
		ePrefix)
	}

// SetDateTimeFormat - Sets the Date Time Format for the current
// CalendarDateTime instance. The format string is stored in
// internal member variable, 'calDTime.dateTimeFmt'.
//
// This Date Time Format string controls the Date Time format for
// Date and Time information returned for the current CalendarDateTime
// instance.
//
// If input parameter 'dateTimeFmt' is an empty string, the Default
// Date Time Format will be applied.
//
//      DEFAULTDATETIMEFORMAT =
//             "2006-01-02 15:04:05.000000000 -0700 MST"
//
// Outside of processing empty strings, this method performs virtually
// no error checking on the input parameter, 'dateTimeFmt'. If you are
// unsure about the correct syntax see the date time format constants
// provided in source file:
//      /datetime/constantsdatetimeformat.go
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  dateTimeFmt        string
//    - This string contains the date/time format which will be used to
//      to format date/time output values. Example:
//          "2006-01-02 15:04:05.000000000 -0700 MST"
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
func (calDTime CalendarDateTime) SetDateTimeFormat(
	dateTimeFmt string,
	ePrefix string) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	ePrefix += "CalendarDateTime.SetDateTimeFormat() "

	dtMech := DTimeNanobot{}

	calDTime.dateTimeFmt =
	 	dtMech.PreProcessDateFormatStr(dateTimeFmt)

	return
}

// SetTagDescription - Sets the tag description associated with
// the current instance of CalendarDateTime.
//
// This method will set the internal member variable CalendarDateTime.tag
// with a description string as designated by the calling function.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  tagDescription  string
//     - A tag description used to set internal member variable
//       CalendarDateTime.tag.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//          -- NONE --
//
//
func (calDTime CalendarDateTime) SetTagDescription(
	tagDescription string) {

	if calDTime.lock == nil {
		calDTime.lock = new(sync.Mutex)
	}

	calDTime.lock.Lock()

	defer calDTime.lock.Unlock()

	calDTime.tag = tagDescription
}