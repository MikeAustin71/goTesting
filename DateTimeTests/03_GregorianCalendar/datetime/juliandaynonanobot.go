package datetime

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type julianDayNoNanobot struct {
	lock     *sync.Mutex

}


// testJulianDayNoDtoValidity - Tests a JulianDayNoDto instance
// to verify validity.
//
func (jDNNanobot *julianDayNoNanobot) testJulianDayNoDtoValidity(
	jDNDto *JulianDayNoDto,
	ePrefix string) (bool, error){

	if jDNNanobot.lock == nil {
		jDNNanobot.lock = new(sync.Mutex)
	}

	jDNNanobot.lock.Lock()

	defer jDNNanobot.lock.Unlock()

	ePrefix += "julianDayNoNanobot.testJulianDayNoDtoValidity() "

	if jDNDto == nil {
		return false,
			errors.New(ePrefix + "\n" +
				"Input parameter 'jDNDto' is a 'nil' pointer!\n")
	}

	if jDNDto.lock == nil {
		jDNDto.lock = new(sync.Mutex)
	}

	jDNDto.isThisInstanceValid = false

	if jDNDto.julianDayNo == nil {
		return false,
			errors.New(ePrefix + "\n" +
				"Data Field 'julianDayNo' is INVALID!\n" +
				"jDNDto.julianDayNo == nil\n")
	}

	if jDNDto.julianDayNoFraction == nil {
		return false,
			errors.New(ePrefix + "\n" +
				"Data Field 'julianDayNoFraction' is INVALID!\n" +
				"jDNDto.julianDayNoFraction == nil\n")
	}

	if jDNDto.julianDayNoTime == nil {
		return false,
			errors.New(ePrefix + "\n" +
				"Data Field 'julianDayNoTime' is INVALID!\n" +
				"jDNDto.julianDayNoTime == nil\n")
	}

	if jDNDto.julianDayNoNumericalSign != 1 &&
		jDNDto.julianDayNoNumericalSign != -1 {
		return false,
			fmt.Errorf(ePrefix + "\n" +
				"Data Field 'julianDayNoNumericalSign' is INVALID!\n" +
				"jDNDto.julianDayNoNumericalSign == %v\n",
				jDNDto.julianDayNoNumericalSign)
	}

	if jDNDto.julianDayNoTime == nil {
		return false,
			errors.New(ePrefix + "\n" +
				"Data Field 'julianDayNoTime' is INVALID!\n" +
				"jDNDto.julianDayNoTime == nil\n")
	}

	if jDNDto.totalJulianNanoSeconds >= int64(time.Hour * 36) ||
		jDNDto.totalJulianNanoSeconds < 0 {
		return false,
			fmt.Errorf(ePrefix + "\n" +
				"Data Field 'totalJulianNanoSeconds' is INVALID!\n" +
				"jDNDto.totalJulianNanoSeconds == %v\n",
				jDNDto.totalJulianNanoSeconds)
	}

	if jDNDto.netGregorianNanoSeconds >= int64(time.Hour * 24) ||
		jDNDto.netGregorianNanoSeconds < 0 {
		return false,
			fmt.Errorf(ePrefix + "\n" +
				"Data Field 'netGregorianNanoSeconds' is INVALID!\n" +
				"jDNDto.netGregorianNanoSeconds == %v\n",
				jDNDto.netGregorianNanoSeconds)
	}

	if jDNDto.hours < 0 || jDNDto.hours >= 23 {
		return false,
			fmt.Errorf(ePrefix + "\n" +
				"Data Field 'hours' is INVALID!\n" +
				"jDNDto.hours == %v\n",
				jDNDto.hours)
	}

	if jDNDto.minutes < 0 || jDNDto.minutes >= 59 {
		return false,
			fmt.Errorf(ePrefix + "\n" +
				"Data Field 'minutes' is INVALID!\n" +
				"jDNDto.minutes == %v\n",
				jDNDto.minutes)
	}

	maxSeconds := 59

	if jDNDto.hasLeapSecond == true {
		maxSeconds = 60
	}

	if jDNDto.seconds < 0 || jDNDto.seconds > maxSeconds {
		return false,
			fmt.Errorf(ePrefix + "\n" +
				"Data Field 'seconds' is INVALID!\n" +
				"Maximum Number Of Seconds='%v'\n" +
				"Minumum Number Of Seconds='1\n" +
				"Has Leap Second='%v'\n" +
				"jDNDto.seconds == %v\n",
				maxSeconds,
				jDNDto.hasLeapSecond,
				jDNDto.seconds)
	}

	if jDNDto.nanoseconds < 0 || jDNDto.nanoseconds > 999999999 {
		return false,
			fmt.Errorf(ePrefix + "\n" +
				"Data Field 'nanoseconds' is INVALID!\n" +
				"jDNDto.nanoseconds == %v\n",
				jDNDto.nanoseconds)
	}

	jDNDto.isThisInstanceValid = true

	return true, nil
}


