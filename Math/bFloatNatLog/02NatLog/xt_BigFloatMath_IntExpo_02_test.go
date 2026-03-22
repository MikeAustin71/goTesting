package naturalLogCalcs

// *************************************
// ** xt_BigFloatMath_Expo_02_test.go **
//      013Fix — IntExpoBigFloat tests
// *************************************
// These tests are designed to evaluate
// exponent calculation expected values
// to 1,000+ decimal places.

import (
	"math/big"
	"testing"
)

func Test_BigFloatMath_IntExpo_ManyDigits_01(t *testing.T) {

	ePrefix := "Test_BigFloatMath_IntExpo_ManyDigits_01"

	// baseStr to 57 decimal digits of accuracy
	baseStr := "3.716586431412345678901234567896665321489219683361042189632"

	exponentValue := int64(8)

	// Expected Result to 456 decimal digits of accuracy
	expectedResultStr := "36404.401562225975591741663072973189677987488996396044626112641492651053286615144796058628951198923359976623755686783423428496969116290496797290496395603454115934731971114958321403448741582304231510133842124829429416753443259489442443197222136207898977414246517466800582599164548256625911734795116932334044587952739317391815896874646908544842081850235941724431811478937457601741575847137432859848470036907221552484215297644767344006306721500401036591964855730176"

	bFloatHlpr := new(BigFloatHelper)

	var workingPrecisionBits uint

	var sourceCalculatedDecimalDigits, sourceBaseDecimalDigits, sourceExpectedResultDecimalDigits int

	_, sourceBaseDecimalDigits = bFloatHlpr.CountDigits(baseStr, '.')

	_, sourceExpectedResultDecimalDigits = bFloatHlpr.CountDigits(expectedResultStr, '.')

	sourceCalculatedDecimalDigits = sourceExpectedResultDecimalDigits + 3

	workingPrecisionBits = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

	baseBigFloat, isOk := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(workingPrecisionBits).
		SetString(baseStr)

	if !isOk {
		t.Errorf("%v\n"+
			"Failed to set base string as big.Float!\n"+
			"baseStr= '%v'\n\n",
			ePrefix, baseStr)
		return
	}

	actualResult, err := new(BigFloatMath).IntExpoBigFloat(baseBigFloat, exponentValue, workingPrecisionBits)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" actualResult, err := new(BigFloatMath).\n"+
			"IntExpoBigFloat(baseBigFloat, exponentValue, workingPrecisionBits)\n"+
			"multiplierStr= '%v'\n"+
			"baseBigFloat= '%v'\n"+
			"workingPrecisionBits= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseBigFloat.Text('f', sourceBaseDecimalDigits),
			exponentValue,
			workingPrecisionBits,
			err.Error())

		return
	}

	actualResultStr := actualResult.Text('f', sourceExpectedResultDecimalDigits)

	if expectedResultStr != actualResultStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != actualResultStr\n"+
			"Expected actualResultStr = '%v'\n"+
			"  Actual actualResultStr = '%v'\n\n",
			ePrefix, expectedResultStr, actualResultStr)
	}

	return
}

func Test_BigFloatMath_IntExpo_ManyDigits_02(t *testing.T) {

	ePrefix := "Test_BigFloatMath_IntExpo_ManyDigits_02"

	// baseStr to 57 decimal digits of accuracy
	baseStr := "3.716586431412345678901234567896665321489219683361042189632"

	exponentValue := int64(15)

	// Expected Result to 855 decimal digits of accuracy
	expectedResultStr := "356585398.338275481921635377025649841709489427652437163336006441045956751959083186446657965399656758339102183981074408924482110146289650335707161101579361038419911303189387748794200028710982305980161327966211470989028012430633887254649707764402259951466293482386519184432416727925819630160217055108965310259018011089284739290288852255712263785936013202292671525309585386329005063437053702303769825524402845779965751533804001047586566924155469662176344535507201926013331932487894620777304342234981212167472452928971500569935671799753063140179626327502631906427117941860442312256326049309506643388505407994958670153319599761817199304982209039985393918610674539477112242038450573823179723535992601542467925504815566970253460502657305372039897498796266173927056925786610413259145722761922599379349272267963327162899890405324145867112755573042172602724463515239677165568"

	bFloatHlpr := new(BigFloatHelper)

	var workingPrecisionBits uint

	var sourceCalculatedDecimalDigits, sourceBaseDecimalDigits, sourceExpectedResultDecimalDigits int

	_, sourceBaseDecimalDigits = bFloatHlpr.CountDigits(baseStr, '.')

	_, sourceExpectedResultDecimalDigits = bFloatHlpr.CountDigits(expectedResultStr, '.')

	sourceCalculatedDecimalDigits = sourceExpectedResultDecimalDigits + 3

	workingPrecisionBits = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceCalculatedDecimalDigits), 1)

	baseBigFloat, isOk := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(workingPrecisionBits).
		SetString(baseStr)

	if !isOk {
		t.Errorf("%v\n"+
			"Failed to set base string as big.Float!\n"+
			"baseStr= '%v'\n\n",
			ePrefix, baseStr)
		return
	}

	actualResult, err := new(BigFloatMath).IntExpoBigFloat(baseBigFloat, exponentValue, workingPrecisionBits)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" actualResult, err := new(BigFloatMath).\n"+
			"IntExpoBigFloat(baseBigFloat, exponentValue, workingPrecisionBits)\n"+
			"multiplierStr= '%v'\n"+
			"baseBigFloat= '%v'\n"+
			"workingPrecisionBits= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseBigFloat.Text('f', sourceBaseDecimalDigits),
			exponentValue,
			workingPrecisionBits,
			err.Error())

		return
	}

	actualResultStr := actualResult.Text('f', sourceExpectedResultDecimalDigits)

	if expectedResultStr != actualResultStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != actualResultStr\n"+
			"Expected actualResultStr = '%v'\n"+
			"  Actual actualResultStr = '%v'\n\n",
			ePrefix, expectedResultStr, actualResultStr)
	}

	return
}

func Test_BigFloatMath_IntExpo_ManyDigits_03(t *testing.T) {

	ePrefix := "Test_BigFloatMath_IntExpo_ManyDigits_03"

	// baseStr to 57 decimal digits of accuracy
	baseStr := "3.716586431412345678901234567896665321489219683361042189632"

	exponentValue := int64(15)

	// Expected Result to 855 decimal digits of accuracy
	expectedResultStr := "356585398.338275481921635377025649841709489427652437163336006441045956751959083186446657965399656758339102183981074408924482110146289650335707161101579361038419911303189387748794200028710982305980161327966211470989028012430633887254649707764402259951466293482386519184432416727925819630160217055108965310259018011089284739290288852255712263785936013202292671525309585386329005063437053702303769825524402845779965751533804001047586566924155469662176344535507201926013331932487894620777304342234981212167472452928971500569935671799753063140179626327502631906427117941860442312256326049309506643388505407994958670153319599761817199304982209039985393918610674539477112242038450573823179723535992601542467925504815566970253460502657305372039897498796266173927056925786610413259145722761922599379349272267963327162899890405324145867112755573042172602724463515239677165568"

	bFloatHlpr := new(BigFloatHelper)

	var workingPrecisionBits uint

	var sourceBaseDecimalDigits, sourceExpectedResultDecimalDigits int

	_, sourceBaseDecimalDigits = bFloatHlpr.CountDigits(baseStr, '.')

	_, sourceExpectedResultDecimalDigits = bFloatHlpr.CountDigits(expectedResultStr, '.')

	workingPrecisionBits = bFloatHlpr.ComputeBigFloatPrecisionBits(uint(sourceExpectedResultDecimalDigits), 1)

	baseBigFloat, isOk := new(big.Float).
		SetMode(big.AwayFromZero).
		SetPrec(workingPrecisionBits).
		SetString(baseStr)

	if !isOk {
		t.Errorf("%v\n"+
			"Failed to set base string as big.Float!\n"+
			"baseStr= '%v'\n\n",
			ePrefix, baseStr)
		return
	}

	actualResult, err := new(BigFloatMath).IntExpoBigFloatDigits(baseBigFloat, exponentValue, uint(sourceExpectedResultDecimalDigits))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" actualResult, err := new(BigFloatMath).\n"+
			"IntExpoBigFloat(baseBigFloat, exponentValue, workingPrecisionBits)\n"+
			"multiplierStr= '%v'\n"+
			"baseBigFloat= '%v'\n"+
			"workingPrecisionBits= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseBigFloat.Text('f', sourceBaseDecimalDigits),
			exponentValue,
			workingPrecisionBits,
			err.Error())

		return
	}

	actualResultStr := actualResult.Text('f', sourceExpectedResultDecimalDigits)

	if expectedResultStr != actualResultStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != actualResultStr\n"+
			"Expected actualResultStr = '%v'\n"+
			"  Actual actualResultStr = '%v'\n\n",
			ePrefix, expectedResultStr, actualResultStr)
	}

	return
}
