package datetime

type IDayOfWeekNumber interface {

	String() string

	XDayOfWeekName() string

	XDayOfWeekNameAbbreviation(
		numberOfCharsInAbbreviation int,
		ePrefix string) (
		dayOfWeekNameAbbrv string,
		err error)

	XDayOfWeekNumber() int

	XDayOfWeekNumberingSystemType() DayOfWeekNumberingSystemType

	XIsValid() bool

	XValueInt() int

	XZero() IDayOfWeekNumber
}