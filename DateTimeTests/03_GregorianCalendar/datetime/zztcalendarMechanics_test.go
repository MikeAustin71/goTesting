package datetime

import "testing"

func TestConvertAnyYearToAstronomicalYear01 (t *testing.T) {

	ePrefix := "TestConvertAnyYearToAstronomicalYear01() "

	var err error

	year := int64(4715)
	yearNumType := CalYearType.BCE()

	expectedYear := int64(-4716)

	calMech := calendarMechanics{}
	var astronomicalYearValue int64

	astronomicalYearValue,
	err= calMech.convertAnyYearToAstronomicalYear(
		year, yearNumType, ePrefix)

	if err != nil {
		t.Errorf("Error return from calMech.convertAnyYearToAstronomicalYear()\n" +
			"Error='%v'\n", err.Error() )
		return
	}

	if astronomicalYearValue != expectedYear {
		t.Errorf("Result INVALID!\n" +
			"astronomicalYearValue != expectedYear\n" +
			"astronomicalYearValue='%v'\n" +
			"expectedYear='%v'\n",
			astronomicalYearValue,
			expectedYear)
	}

	return
}

func TestConvertAnyYearToAstronomicalYear02 (t *testing.T) {

	ePrefix := "TestConvertAnyYearToAstronomicalYear02() "

	var err error

	year := int64(2020)
	yearNumType := CalYearType.CE()

	expectedYear := int64(2020)

	calMech := calendarMechanics{}
	var astronomicalYearValue int64

	astronomicalYearValue,
	err= calMech.convertAnyYearToAstronomicalYear(
		year, yearNumType, ePrefix)

	if err != nil {
		t.Errorf("Error return from calMech.convertAnyYearToAstronomicalYear()\n" +
			"Error='%v'\n", err.Error() )
		return
	}

	if astronomicalYearValue != expectedYear {
		t.Errorf("Result INVALID!\n" +
			"astronomicalYearValue != expectedYear\n" +
			"astronomicalYearValue='%v'\n" +
			"expectedYear='%v'\n",
			astronomicalYearValue,
			expectedYear)
	}

	return
}

func TestConvertAnyYearToAstronomicalYear03 (t *testing.T) {

	ePrefix := "TestConvertAnyYearToAstronomicalYear03() "

	var err error

	year := int64(-4713)
	yearNumType := CalYearType.Astronomical()

	expectedYear := int64(-4713)

	calMech := calendarMechanics{}
	var astronomicalYearValue int64

	astronomicalYearValue,
	err= calMech.convertAnyYearToAstronomicalYear(
		year, yearNumType, ePrefix)

	if err != nil {
		t.Errorf("Error return from calMech.convertAnyYearToAstronomicalYear()\n" +
			"Error='%v'\n", err.Error() )
		return
	}

	if astronomicalYearValue != expectedYear {
		t.Errorf("Result INVALID!\n" +
			"astronomicalYearValue != expectedYear\n" +
			"astronomicalYearValue='%v'\n" +
			"expectedYear='%v'\n",
			astronomicalYearValue,
			expectedYear)
	}

	return
}

func TestConvertAnyYearToAstronomicalYear04 (t *testing.T) {

	ePrefix := "TestConvertAnyYearToAstronomicalYear04() "

	var err error

	year := int64(-2020)
	yearNumType := CalYearType.CE()

	calMech := calendarMechanics{}

	_,
	err= calMech.convertAnyYearToAstronomicalYear(
		year, yearNumType, ePrefix)

	if err == nil {
		t.Error(ePrefix + "Error!\n" +
			"Expected an error return from calMech.convertAnyYearToAstronomicalYear()\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
		return
	}
}

func TestConvertAnyYearToAstronomicalYear05 (t *testing.T) {

	ePrefix := "TestConvertAnyYearToAstronomicalYear05() "

	var err error

	year := int64(0)
	yearNumType := CalYearType.CE()

	calMech := calendarMechanics{}

	_,
	err= calMech.convertAnyYearToAstronomicalYear(
		year, yearNumType, ePrefix)

	if err == nil {
		t.Error(ePrefix + "Error!\n" +
			"Expected an error return from calMech.convertAnyYearToAstronomicalYear()\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
		return
	}
}

func TestConvertAnyYearToAstronomicalYear06 (t *testing.T) {

	ePrefix := "TestConvertAnyYearToAstronomicalYear06() "

	var err error

	year := int64(0)
	yearNumType := CalYearType.BCE()

	calMech := calendarMechanics{}

	_,
	err= calMech.convertAnyYearToAstronomicalYear(
		year, yearNumType, ePrefix)

	if err == nil {
		t.Error(ePrefix + "Error!\n" +
			"Expected an error return from calMech.convertAnyYearToAstronomicalYear()\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
		return
	}
}

func TestConvertAnyYearToAstronomicalYear07 (t *testing.T) {

	ePrefix := "TestConvertAnyYearToAstronomicalYear07() "

	var err error

	year := int64(-4713)
	yearNumType := CalYearType.BCE()

	calMech := calendarMechanics{}

	_,
	err= calMech.convertAnyYearToAstronomicalYear(
		year, yearNumType, ePrefix)

	if err == nil {
		t.Error(ePrefix + "Error!\n" +
			"Expected an error return from calMech.convertAnyYearToAstronomicalYear()\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
		return
	}
}

func TestConvertAstronomicalYearToCommonEraYear01 (t *testing.T) {

	ePrefix := "TestConvertAstronomicalYearToCommonEraYear01() "


	astronomicalYear := int64(-4715)

	expectedYear := int64(4716)
	expectedYearNumType := CalYearType.BCE()

	calMech := calendarMechanics{}
	var commonEraYearValue int64
	var commonEraYearNumType CalendarYearNumType

	commonEraYearValue,
		commonEraYearNumType =
			calMech.convertAstronomicalYearToCommonEraYear(
		astronomicalYear)

	if commonEraYearValue != expectedYear {
		t.Errorf(ePrefix + "Result INVALID!\n" +
			"commonEraYearValue != expectedYear\n" +
			"commonEraYearValue='%v'\n" +
			"expectedYear='%v'\n",
			commonEraYearValue,
			expectedYear)
		return
	}

	if expectedYearNumType != commonEraYearNumType {
		t.Errorf(ePrefix + "Result INVALID!\n" +
			"expectedYearNumType != commonEraYearNumType\n" +
			"commonEraYearNumType='%v'\n" +
			"expectedYearNumType='%v'\n",
			commonEraYearNumType.String(),
			expectedYearNumType.String())
	}

	return
}

func TestConvertAstronomicalYearToCommonEraYear02 (t *testing.T) {

	ePrefix := "TestConvertAstronomicalYearToCommonEraYear02() "

	astronomicalYear := int64(2020)

	expectedYear := int64(2020)
	expectedYearNumType := CalYearType.CE()

	calMech := calendarMechanics{}
	var commonEraYearValue int64
	var commonEraYearNumType CalendarYearNumType

	commonEraYearValue,
		commonEraYearNumType =
			calMech.convertAstronomicalYearToCommonEraYear(
		astronomicalYear)

	if commonEraYearValue != expectedYear {
		t.Errorf(ePrefix + "Result INVALID!\n" +
			"commonEraYearValue != expectedYear\n" +
			"commonEraYearValue='%v'\n" +
			"expectedYear='%v'\n",
			commonEraYearValue,
			expectedYear)
		return
	}

	if expectedYearNumType != commonEraYearNumType {
		t.Errorf(ePrefix + "Result INVALID!\n" +
			"expectedYearNumType != commonEraYearNumType\n" +
			"commonEraYearNumType='%v'\n" +
			"expectedYearNumType='%v'\n",
			commonEraYearNumType.String(),
			expectedYearNumType.String())
	}

	return
}

func TestConvertAstronomicalYearToCommonEraYear03 (t *testing.T) {

	ePrefix := "TestConvertAstronomicalYearToCommonEraYear03() "


	astronomicalYear := int64(0)

	expectedYear := int64(1)
	expectedYearNumType := CalYearType.BCE()

	calMech := calendarMechanics{}
	var commonEraYearValue int64
	var commonEraYearNumType CalendarYearNumType

	commonEraYearValue,
		commonEraYearNumType =
		calMech.convertAstronomicalYearToCommonEraYear(
			astronomicalYear)

	if commonEraYearValue != expectedYear {
		t.Errorf(ePrefix + "Result INVALID!\n" +
			"commonEraYearValue != expectedYear\n" +
			"commonEraYearValue='%v'\n" +
			"expectedYear='%v'\n",
			commonEraYearValue,
			expectedYear)
		return
	}

	if expectedYearNumType != commonEraYearNumType {
		t.Errorf(ePrefix + "Result INVALID!\n" +
			"expectedYearNumType != commonEraYearNumType\n" +
			"commonEraYearNumType='%v'\n" +
			"expectedYearNumType='%v'\n",
			commonEraYearNumType.String(),
			expectedYearNumType.String())
	}

	return
}

