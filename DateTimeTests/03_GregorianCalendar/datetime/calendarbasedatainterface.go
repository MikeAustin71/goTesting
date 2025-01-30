package datetime

type ICalendarBaseData interface {

	GetCalendarSpecification() CalendarSpec

	GetISODayOfWeekNo(
		julianDayNoDto JulianDayNoDto,
		ePrefix string) (
		isoDayOfWeekNo ISO8601DayOfWeekNo,
		err error)

	GetUsDayOfWeekNo(
		julianDayNoDto JulianDayNoDto,
		ePrefix string) (
		usDayOfWeekNo UsDayOfWeekNo,
		err error)

	GetDaysInLeapYear() int

	GetDaysInStandardYear() int

	GetDaysInYear(
		year int64,
		yearNumType CalendarYearNumType,
		ePrefix string) (
		daysInYear int,
		err error)

	GetDaysOfWeekNames(
		dayOfWeekNoSysType DayOfWeekNumberingSystemType,
		ePrefix string) (
		daysOfWeekNames map[int] string,
		err error)

	GetDaysOfWeekNameAbbreviations(
		dayOfWeekNoSysType DayOfWeekNumberingSystemType,
		numberOfCharsInAbbreviation int,
		ePrefix string) (
		weekDayNameAbbrvs map[int] string,
		err error)

	GetLeapYearOrdinalDays() map[int] int

	GetLeapYearMonthDays() map[int] int

	GetMonthDayFromOrdinalDayNo(
		ordinalDate int,
		isLeapYear bool,
		ePrefix string)(
		yearAdjustment int,
		month int,
		day int,
		err error)

	GetOrdinalDayNumber(
		isLeapYear bool,
		month int,
		day int,
		ePrefix string) (
		ordinalDayNo int,
		err error)

	GetOrdinalDayNoFromDate(
		year int64,
		yearNumType CalendarYearNumType,
		month int,
		day int,
		ePrefix string) (
		ordinalDayNo int,
		err error)

	GetRemainingDaysInYear(
		year int64,
		yearNumType CalendarYearNumType,
		month int,
		day int,
		ePrefix string) (
		remainingDaysOfYear int,
		err error)

	GetStandardYearOrdinalDays() map[int] int

	GetStandardYearMonthDays() map[int] int

	GetYearMonthDayFromOrdinalDayNo(
		ordinalDate int,
		year int64,
		yearNumType CalendarYearNumType,
		ePrefix string)(
		astronomicalYear int64,
		month int,
		day int,
		err error)

	IsLeapYear(
		year int64,
		yearNumType CalendarYearNumType,
		ePrefix string) (
		isLeapYear bool,
		err error)

	IsValidDate(
		year int64,
		yearNumType CalendarYearNumType,
		month int,
		day int,
		ePrefix string) (
		isValid bool,
		err error)

	New() ICalendarBaseData

}

