package datetime

import (
	"errors"
	"sync"
)

type aDateTimeDtoNanobot struct {
	lock *sync.Mutex
}


// testDateTransferDtoValidity - Tests the validity of a
// ADateTimeDto instance.
//
func (aDateTimeDtoNanobot *aDateTimeDtoNanobot) testDateTransferDtoValidity(
	aDateTimeDto *ADateTimeDto,
	ePrefix string) (bool, error) {

	if aDateTimeDtoNanobot.lock == nil {
		aDateTimeDtoNanobot.lock = new(sync.Mutex)
	}

	aDateTimeDtoNanobot.lock.Lock()

	defer aDateTimeDtoNanobot.lock.Unlock()

	ePrefix += "aDateTimeDtoNanobot.testDateTransferDtoValidity() "

	if aDateTimeDto == nil {

		return false, errors.New(ePrefix + "\n" +
			"Input parameter 'aDateTimeDto' " +
			"has a 'nil' pointer!\n")
	}

	var err error

	err = aDateTimeDto.date.IsValidInstanceError(
		ePrefix + "Testing 'date' for 'aDateTimeDto'. ")

	if err != nil {
		return false, err
	}

	err = aDateTimeDto.time.IsValidInstanceError(
		ePrefix + "Testing 'date' for 'aDateTimeDto'. ")

	if err != nil {
		return false, err
	}

	if aDateTimeDto.julianDayNumber.isThisInstanceValid == true {

		err =
			aDateTimeDto.julianDayNumber.IsValidInstanceError(
				ePrefix + "Julian Day Number was set. Testing validity. ")

		if err != nil {
			return false, err
		}

	}

	if aDateTimeDto.dateTimeFmt == "" {
		dtMech := DTimeNanobot{}
		aDateTimeDto.dateTimeFmt =
			dtMech.PreProcessDateFormatStr("")
	}

	return true, nil
}