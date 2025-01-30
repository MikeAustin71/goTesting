package datetime

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
)

type CalendarNewCycleConfiguration struct {
	mainCycleBaseDateForJDNNo CalendarDateTime         // Main Cycle Base Date is always less than
	//                                                 // Official Calendar Julian Day Number Base Date.
	mainCycleAdjustmentYearsForJDNNo *big.Int          // Adjustment Years for Main Cycle Base Date and
	//                                                 // Official Calendar JDN Base Date.
	mainCycleAdjustmentDaysForJDNNo  *big.Int          // Adjustment Days for Main Cycle Base Date and
	//                                                 // Official Calendar JDN Base Date.
	calJdnBaseDate CalendarDateTime                    // The Official Calendar Julian Day Number Base Date
	//

	ordinalNumberFixedBaseDateTime CalendarDateTime    // Base Start date time for Ordinal Fixed Date

	mainCycleConfig CalendarCycleDto                   // Calendar cycle for main calendar cycle.
	//
	calendarCyclesConfig []CalendarCycleDto            // This array contains all of the calendar cycles
	//
	calendarBaseData ICalendarBaseData // Contains methods for calculating calendar base data

	lock *sync.Mutex                                   // Internal lock
}

// GetCalendarBaseData - Returns Calendar Base Data Object. Calendar Base Data
// includes methods for calculating calendar specific data and events.
//
func(calNewCycCfg *CalendarNewCycleConfiguration) GetCalendarBaseData(ePrefix string) (ICalendarBaseData, error) {

	if calNewCycCfg.lock == nil {
		calNewCycCfg.lock = new(sync.Mutex)
	}

	calNewCycCfg.lock.Lock()

	defer calNewCycCfg.lock.Unlock()

	ePrefix += "CalendarNewCycleConfiguration.GetCalendarBaseData() "

	if calNewCycCfg.calendarBaseData == nil {
		return nil, errors.New(ePrefix + "\n" +
			"Error: CalendarCycleConfiguration improperly configured!\n" +
			" calCycCfg.calendarBaseData is 'nil'!\n")
	}

	newCalBaseData := calNewCycCfg.calendarBaseData.New()

	return newCalBaseData, nil
}

// GetCalendarCycleConfigurations - Returns a array of CalendarCycleDto objects
// covering all of the calendar cycles associated with the Gregorian
// calendar.
//
func(calNewCycCfg *CalendarNewCycleConfiguration) GetCalendarCycleConfigurations() []CalendarCycleDto {

	if calNewCycCfg.lock == nil {
		calNewCycCfg.lock = new(sync.Mutex)
	}

	calNewCycCfg.lock.Lock()

	defer calNewCycCfg.lock.Unlock()

	lenArray := len(calNewCycCfg.calendarCyclesConfig)

	calConfigs := make([]CalendarCycleDto, lenArray)

	if lenArray == 0 {
		return calConfigs
	}

	for i:=0; i < lenArray; i++ {
		calConfigs[i] = calNewCycCfg.calendarCyclesConfig[i].CopyOut()
	}

	return calConfigs
}

// GetCalendarJDNBaseDate - Returns the official Calendar Julian Day
// Number base date/time.
//
func(calNewCycCfg *CalendarNewCycleConfiguration) GetCalendarJDNBaseDate(
	ePrefix string) (CalendarDateTime, error) {

	if calNewCycCfg.lock == nil {
		calNewCycCfg.lock = new(sync.Mutex)
	}

	calNewCycCfg.lock.Lock()

	defer calNewCycCfg.lock.Unlock()

	ePrefix += "CalendarNewCycleConfiguration.GetCalendarJDNBaseDate() "

	err := calNewCycCfg.calJdnBaseDate.IsValidInstanceError(ePrefix)

	if err != nil {
		return CalendarDateTime{},
			fmt.Errorf(ePrefix + "\n" +
				"calNewCycCfg.calJdnBaseDate is INVALID!\n" +
				"Error= %v\n", err.Error())
	}

	return calNewCycCfg.calJdnBaseDate.CopyOut(ePrefix)
}

// GetMainCycleAdjustmentDays - Returns the number of days between
// the Main Cycle Base Date and the official Calendar Julian Day Number
// Base Date.
//
func(calNewCycCfg *CalendarNewCycleConfiguration) GetMainCycleAdjustmentDays(
	ePrefix string) (mainCycleAdjDays *big.Int, err error) {

	if calNewCycCfg.lock == nil {
		calNewCycCfg.lock = new(sync.Mutex)
	}

	calNewCycCfg.lock.Lock()

	defer calNewCycCfg.lock.Unlock()

	ePrefix += "CalendarNewCycleConfiguration.GetMainCycleAdjustmentDays() "

	mainCycleAdjDays = big.NewInt(0)
	err = nil

	if calNewCycCfg.mainCycleAdjustmentDaysForJDNNo == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Main Cycle Adjustment Days is 'nil'.\n" +
			"calNewCycCfg.mainCycleAdjustmentDaysForJDNNo == nil\n")
		return mainCycleAdjDays, err
	}

	mainCycleAdjDays =
		big.NewInt(0).
			Set(calNewCycCfg.mainCycleAdjustmentDaysForJDNNo)

	return mainCycleAdjDays, err
}

// GetMainCycleAdjustmentYears - Returns the number of years between
// the Main Cycle Base Date and the official Calendar Julian Day Number
// Base Date.
//
func(calNewCycCfg *CalendarNewCycleConfiguration) GetMainCycleAdjustmentYears(
	ePrefix string) (mainCycleAdjYears *big.Int, err error) {

	if calNewCycCfg.lock == nil {
		calNewCycCfg.lock = new(sync.Mutex)
	}

	calNewCycCfg.lock.Lock()

	defer calNewCycCfg.lock.Unlock()

	ePrefix += "CalendarNewCycleConfiguration.GetMainCycleAdjustmentYears() "

	mainCycleAdjYears = big.NewInt(0)
	err = nil

	if calNewCycCfg.mainCycleAdjustmentYearsForJDNNo == nil {
		err = errors.New(ePrefix + "\n" +
			"Error: Main Cycle Adjustment Years is 'nil'.\n" +
			"calNewCycCfg.mainCycleAdjustmentYearsForJDNNo == nil\n")
		return mainCycleAdjYears, err
	}

	mainCycleAdjYears =
		big.NewInt(0).
			Set(calNewCycCfg.mainCycleAdjustmentYearsForJDNNo)

	return mainCycleAdjYears, err
}

// GetMainCycleBaseDate - Returns a CalendarDateTime instance describing
// the Main Cycle Base Date and Time.
//
func(calNewCycCfg *CalendarNewCycleConfiguration) GetMainCycleBaseDate(
	ePrefix string) (CalendarDateTime, error) {

	if calNewCycCfg.lock == nil {
		calNewCycCfg.lock = new(sync.Mutex)
	}

	calNewCycCfg.lock.Lock()

	defer calNewCycCfg.lock.Unlock()

	ePrefix += "CalendarNewCycleConfiguration.GetMainCycleBaseDate() "

	err := calNewCycCfg.mainCycleBaseDateForJDNNo.IsValidInstanceError(ePrefix)

	if err != nil {
		return CalendarDateTime{},
		fmt.Errorf(ePrefix + "\n" +
			"calNewCycCfg.mainCycleBaseDateForJDNNo is INVALID!\n" +
			"Error= %v\n", err.Error())
	}

	return calNewCycCfg.mainCycleBaseDateForJDNNo.CopyOut(ePrefix)
}

// GetOrdinalFixedBaseDate - Returns the base date for multi-year
// ordinal day numbering.
func(calNewCycCfg *CalendarNewCycleConfiguration) GetOrdinalFixedBaseDate(
	ePrefix string) (CalendarDateTime, error) {

	if calNewCycCfg.lock == nil {
		calNewCycCfg.lock = new(sync.Mutex)
	}

	calNewCycCfg.lock.Lock()

	defer calNewCycCfg.lock.Unlock()

	ePrefix += "CalendarNewCycleConfiguration.GetOrdinalFixedBaseDate() "

	err := calNewCycCfg.ordinalNumberFixedBaseDateTime.IsValidInstanceError(ePrefix)

	if err != nil {
		return CalendarDateTime{},
			fmt.Errorf(ePrefix + "\n" +
				"calNewCycCfg.ordinalNumberFixedBaseDateTime is INVALID!\n" +
				"Error= %v\n", err.Error())
	}

	return calNewCycCfg.ordinalNumberFixedBaseDateTime.CopyOut(ePrefix)
}


type CalendarCycleConfiguration struct {

	mainCycleStartDateForPositiveJDNNo ADateTimeDto // Main Cycle Start Date is less than JDN Base Date
	//                                                            //  Used for positive JDN Number calculations
	mainCycleAdjustmentYearsForPositiveJDNNo *big.Int             // Adjustment Years for JDN Base Date and positive
	//                                                            //  JDN Number calculations. Always a negative value.
	mainCycleAdjustmentDaysForPositiveJDNNo  *big.Int             // Adjustment Years for JDN Base Date and positive
	//                                                            //  JDN Number Calculations. Always a negative value.
	mainCycleStartDateForNegativeJDNNo ADateTimeDto // Main Cycle Start Date is greater than JDN Base
	//                                                            //  Date. Used for Negative JDN Number calculations
	mainCycleAdjustmentYearsForNegativeJDNNo  *big.Int            // Adjustment Years for JDN Base Date and negative
	//                                                            //  JDN Number calculations. Always a negative value.
	mainCycleAdjustmentDaysForNegativeJDNNo   *big.Int            // Adjustment Years for JDN Base Date and negative
	//                                                            //  JDN Number calculations.
	jdnBaseStartYearDateTime          ADateTimeDto // Base Start date time for JDN calculations
	ordinalFixedDateStartYearDateTime ADateTimeDto // Base Start date time for Ordinal Fixed Date
	//                                                            //  calculations.
	mainCycleConfig      CalendarCycleDto   // Calendar cycle for main calendar cycle.
	calendarCyclesConfig []CalendarCycleDto // Ths array contains all of the calendar cycles
	calendarBaseData     ICalendarBaseData  // Contains methods for calculating calendar base data
	//                                                           //   such as Leap Years and number of days in year.
	lock                                      *sync.Mutex
}

func( calCycCfg *CalendarCycleConfiguration) GetCalendarBaseData(ePrefix string) (ICalendarBaseData, error) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	ePrefix += "CalendarCycleConfiguration.GetCalendarBaseData() "

	if calCycCfg.calendarBaseData == nil {
		return nil, errors.New(ePrefix + "\n" +
			"Error: CalendarCycleConfiguration improperly configured!\n" +
			" calCycCfg.calendarBaseData is 'nil'!\n")
	}

	newCalBaseData := calCycCfg.calendarBaseData.New()

	return newCalBaseData, nil
}

// GetCalendarCycleConfigurations - Returns a array of CalendarCycleDto objects
// covering all of the calendar cycles associated with the Gregorian
// calendar.
//
func( calCycCfg *CalendarCycleConfiguration) GetCalendarCycleConfigurations() []CalendarCycleDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	lenArray := len(calCycCfg.calendarCyclesConfig)

	calConfigs := make([]CalendarCycleDto, lenArray)

	if lenArray == 0 {
		return calConfigs
	}

	for i:=0; i < lenArray; i++ {
		calConfigs[i] = calCycCfg.calendarCyclesConfig[i].CopyOut()
	}

	return calConfigs
}

// GetJDNBaseStartYearDateTime - Returns the base starting year and
// date/time for Julian Day Number calculations.
//
func( calCycCfg *CalendarCycleConfiguration) GetJDNBaseStartYearDateTime() ADateTimeDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	jdnBaseStartYearDateTime, _ := calCycCfg.jdnBaseStartYearDateTime.CopyOut("")

	return jdnBaseStartYearDateTime
}

// GetMainCycleAdjustmentYearsForPositiveJDNNo - Returns Main Cycle Adjustment
// Years used in calculating positive Julian Day Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleAdjustmentYearsForPositiveJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if calCycCfg.mainCycleAdjustmentYearsForPositiveJDNNo == nil {
		calCycCfg.mainCycleAdjustmentYearsForPositiveJDNNo = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycCfg.mainCycleAdjustmentYearsForPositiveJDNNo)
}

// GetMainCycleAdjustmentDaysForPositiveJDNNo - Returns Main Cycle Adjustment
// Days used in calculating positive Julian Day Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleAdjustmentDaysForPositiveJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if calCycCfg.mainCycleAdjustmentDaysForPositiveJDNNo == nil {
		calCycCfg.mainCycleAdjustmentDaysForPositiveJDNNo = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycCfg.mainCycleAdjustmentDaysForPositiveJDNNo)
}

// GetMainCycleStartYearForNegativeJDNNo - Returns the starting Main Cycle
// Base Year for Negative Julian Day Numbers as a type *big.Int.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleStartYearForNegativeJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	return big.NewInt(
		calCycCfg.mainCycleStartDateForPositiveJDNNo.date.astronomicalYear)
}

// GetMainCycleStartDateForNegativeJDNNo - Returns Main Cycle starting
// date/time information to be used in calculating negative Julian Day
// Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleStartDateForNegativeJDNNo() ADateTimeDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	mainCycleStartDateForNegativeJdnNo, _ := calCycCfg.mainCycleStartDateForNegativeJDNNo.CopyOut("")

	return mainCycleStartDateForNegativeJdnNo
}

// GetMainCycleStartYearForPositiveJDNNo - Returns the starting Main Cycle
// Base Year for Positive Julian Day Numbers as a type *big.Int.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleStartYearForPositiveJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	return big.NewInt(
		calCycCfg.mainCycleStartDateForPositiveJDNNo.date.astronomicalYear)
}

// GetMainCycleStartDateForPositiveJDNNo - Returns Main Cycle starting
// date/time information to be used in calculating positive Julian Day
// Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleStartDateForPositiveJDNNo() ADateTimeDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	mainCycleStartDateForPositiveJDNNo, _ :=
	 	calCycCfg.mainCycleStartDateForPositiveJDNNo.CopyOut("")

	return mainCycleStartDateForPositiveJDNNo
}

// GetMainCycleAdjustmentYearsForNegativeJDNNo - Returns Main Cycle Adjustment
// Years used in calculating negative Julian Day Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleAdjustmentYearsForNegativeJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if calCycCfg.mainCycleAdjustmentYearsForNegativeJDNNo == nil {
		calCycCfg.mainCycleAdjustmentYearsForNegativeJDNNo = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycCfg.mainCycleAdjustmentYearsForNegativeJDNNo)
}

// GetMainCycleAdjustmentDaysForNegativeJDNNo - Returns Main Cycle Adjustment
// Days used in calculating negative Julian Day Numbers.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleAdjustmentDaysForNegativeJDNNo() *big.Int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if calCycCfg.mainCycleAdjustmentDaysForNegativeJDNNo == nil {
		calCycCfg.mainCycleAdjustmentDaysForNegativeJDNNo = big.NewInt(0)
	}

	return big.NewInt(0).
		Set(calCycCfg.mainCycleAdjustmentDaysForNegativeJDNNo)
}

// GetMainCycleConfiguration - Returns calendar main cycle configuration
// information.
//
func( calCycCfg *CalendarCycleConfiguration) GetMainCycleConfiguration() CalendarCycleDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	return calCycCfg.mainCycleConfig.CopyOut()
}

// Returns the number of Calendar Cycle Dto objects included in this
// calendar configuration.
//
func( calCycCfg *CalendarCycleConfiguration) GetNumberOfCalendarCycleConfigDtos() int {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	return len(calCycCfg.calendarCyclesConfig)
}

// GetOrdinalFixedDateBaseStartYearDateTime - Returns the base starting year and
// date/time for Ordinal or Fixed Day Number calculations.
//
func( calCycCfg *CalendarCycleConfiguration) GetOrdinalFixedDateBaseStartYearDateTime() ADateTimeDto {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	ordinalFixedDateStartYearDateTime, _ :=
		calCycCfg.ordinalFixedDateStartYearDateTime.CopyOut("")

	return ordinalFixedDateStartYearDateTime
}

// SetCalendarCycleConfigurations - Receives an array of CalendarCycleDto objects
// and proceeds to copy configuration to the internal member variable
// CalendarCycleConfiguration.calendarCyclesConfig.
//
func( calCycCfg *CalendarCycleConfiguration) SetCalendarCycleConfigurations(calConfigs []CalendarCycleDto )  {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	lenArray := len(calConfigs)

	if lenArray == 0 {
		return
	}

	calCycCfg.calendarCyclesConfig = make([]CalendarCycleDto, lenArray)

	for i:=0; i < lenArray; i++ {
		calCycCfg.calendarCyclesConfig[i] = calConfigs[i].CopyOut()
	}

	return
}

// SetJDNBaseStartYearDateTime - Sets internal member variable
// CalendarCycleConfiguration.jdnBaseStartYearDateTime.
//
func( calCycCfg *CalendarCycleConfiguration) SetJDNBaseStartYearDateTime(
	jdnBaseStartYearDateTime ADateTimeDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.jdnBaseStartYearDateTime, _ =
		jdnBaseStartYearDateTime.CopyOut("")

	return
}

// SetMainCycleAdjustmentDaysForNegativeJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleAdjustmentDaysForNegativeJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleAdjustmentDaysForNegativeJDNNo(
	mainCycleAdjustmentDaysForNegativeJDNNo *big.Int) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if mainCycleAdjustmentDaysForNegativeJDNNo == nil {
		mainCycleAdjustmentDaysForNegativeJDNNo = big.NewInt(0)
	}

	calCycCfg.mainCycleAdjustmentDaysForNegativeJDNNo =
		big.NewInt(0).
			Set(mainCycleAdjustmentDaysForNegativeJDNNo)

	return
}

// SetMainCycleAdjustmentDaysForPositiveJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleAdjustmentDaysForPositiveJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleAdjustmentDaysForPositiveJDNNo(
	mainCycleAdjustmentDaysForPositiveJDNNo *big.Int) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if mainCycleAdjustmentDaysForPositiveJDNNo == nil {
		mainCycleAdjustmentDaysForPositiveJDNNo = big.NewInt(0)
	}

	calCycCfg.mainCycleAdjustmentDaysForPositiveJDNNo =
		big.NewInt(0).
			Set(mainCycleAdjustmentDaysForPositiveJDNNo)
}

// SetMainCycleAdjustmentYearsForNegativeJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleAdjustmentYearsForNegativeJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleAdjustmentYearsForNegativeJDNNo(
	mainCycleAdjustmentYearsForNegativeJDNNo *big.Int) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if mainCycleAdjustmentYearsForNegativeJDNNo == nil {
		mainCycleAdjustmentYearsForNegativeJDNNo = big.NewInt(0)
	}

	calCycCfg.mainCycleAdjustmentYearsForNegativeJDNNo =
		big.NewInt(0).
			Set(mainCycleAdjustmentYearsForNegativeJDNNo)
}

// SetMainCycleAdjustmentYearsForPositiveJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleAdjustmentYearsForPositiveJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleAdjustmentYearsForPositiveJDNNo(
	mainCycleAdjustmentYearsForPositiveJDNNo *big.Int) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	if mainCycleAdjustmentYearsForPositiveJDNNo == nil {
		mainCycleAdjustmentYearsForPositiveJDNNo = big.NewInt(0)
	}
	
	calCycCfg.mainCycleAdjustmentYearsForPositiveJDNNo =
		big.NewInt(0).
			Set(mainCycleAdjustmentYearsForPositiveJDNNo)
}

// SetMainCycleConfig - Sets internal member variable
// CalendarCycleConfiguration.mainCycleConfig.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleConfig(
	mainCycleConfig CalendarCycleDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.mainCycleConfig =
		mainCycleConfig.CopyOut()

	return
}

// SetMainCycleStartDateForNegativeJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleStartDateForNegativeJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleStartDateForNegativeJDNNo(
	mainCycleStartDateForNegativeJDNNo ADateTimeDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.mainCycleStartDateForNegativeJDNNo, _ =
		mainCycleStartDateForNegativeJDNNo.CopyOut("")
}

// SetMainCycleStartDateForPositiveJDNNo - Sets internal member variable
// CalendarCycleConfiguration.mainCycleStartDateForPositiveJDNNo.
//
func( calCycCfg *CalendarCycleConfiguration) SetMainCycleStartDateForPositiveJDNNo(
	mainCycleStartDateForPositiveJDNNo ADateTimeDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.mainCycleStartDateForPositiveJDNNo, _ =
		mainCycleStartDateForPositiveJDNNo.CopyOut("")
}

// SetOrdinalFixedDateStartYearDateTime - Sets internal member variable
// CalendarCycleConfiguration.ordinalFixedDateStartYearDateTime.
//
func( calCycCfg *CalendarCycleConfiguration) SetOrdinalFixedDateStartYearDateTime(
	ordinalFixedDateStartYearDateTime ADateTimeDto) {

	if calCycCfg.lock == nil {
		calCycCfg.lock = new(sync.Mutex)
	}

	calCycCfg.lock.Lock()

	defer calCycCfg.lock.Unlock()

	calCycCfg.ordinalFixedDateStartYearDateTime, _ =
		ordinalFixedDateStartYearDateTime.CopyOut("")

	return
}
