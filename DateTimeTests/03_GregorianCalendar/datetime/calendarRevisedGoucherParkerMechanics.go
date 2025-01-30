package datetime

import "sync"

type calendarRevisedGoucherParkerMechanics struct {
	lock *sync.Mutex
}

// isLeapYear - Returns 'true' if the year value is a leap year under
// the Revised Goucher-Parker calendar.
//
// Reference:
//  Documentation for Type CalendarSpec: datetime\calendarspecenum.go
//
//   https://www.inverse.com/article/12152-how-to-make-a-better-leap-year-with-math
//   https://www.theguardian.com/science/2011/feb/28/leap-year-alex-bellos
//   https://www.youtube.com/watch?v=qkt_wmRKYNQ
//
// Summary
//
// 1. If year is divisible by 4, it IS a leap year; add on day to February
// 2. If year is divisible by 128 it is NOT a leap year and no day is added to February.
// 3. If year is divisible by 454,545, it IS a leap year; add a day to February.
//
func (calRevGoucherParkerMech *calendarRevisedGoucherParkerMechanics) isLeapYear(
	year int64) bool {

	if calRevGoucherParkerMech.lock == nil {
		calRevGoucherParkerMech.lock = new(sync.Mutex)
	}

	calRevGoucherParkerMech.lock.Lock()

	defer calRevGoucherParkerMech.lock.Unlock()

	var by4Remainder, by128Remainder, byCycleRemainder int64

	byCycleRemainder = year % 454545

	if byCycleRemainder == 0 {
		return true
	}

	by128Remainder = year % 128

	if by128Remainder == 0 {
		return false
	}

	by4Remainder = year % 4

	if by4Remainder == 0 {
		return true
	}

	return false

}

