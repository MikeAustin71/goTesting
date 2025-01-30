package main

import (
	"fmt"
	"math/big"
	"strings"
	"sync"
)

type IInputParamToStr interface {
	*big.Int | *big.Float | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 | ~string | ~bool
}

var paramToStrLock sync.Mutex

func ParameterToStr[T IInputParamToStr](param T) string {

	paramToStrLock.Lock()

	defer paramToStrLock.Unlock()

	str := fmt.Sprintf("%v",
		param)

	return str
}

type LittleHawk struct {
	sound  string
	repeat int
}

func (littleHawk LittleHawk) String() string {

	if littleHawk.repeat < 1 {
		littleHawk.repeat = 1
	}

	return strings.Repeat(littleHawk.sound, littleHawk.repeat)
}

type Hawk struct {
	sound  string
	repeat int
}

func (hawk *Hawk) SetSound(rawNewSound interface{}, repeat int) string {

	str := fmt.Sprintf("%s", rawNewSound)

	return strings.Repeat(str, repeat)
}

func (hawk *Hawk) CheckSound(rawNewSound interface{}) {

	switch rawNewSound.(type) { // the switch uses the type of the interface

	case uint, uint8, uint16, uint32, uint64:
		fmt.Printf("Type: Unsigned Integer\n")
	case int:
		fmt.Println("Type: Int")

	case big.Int:
		fmt.Printf("Type: big.Int\n")

	case big.Float:
		fmt.Printf("Type: big.Float\n")

	case fmt.Stringer:
		fmt.Printf("From 'case' Type: fmt.Stringer \n")
		iStringer,
			ok := rawNewSound.(fmt.Stringer)

		if !ok {
			fmt.Printf("Error: Could NOT convert empty interface to fmt.Stringer!!\n")
			return
		}

		fmt.Printf("Value of fmt.Stringer: %v\n",
			iStringer.String())

	case string:
		fmt.Println("Type: String")

	case nil:
		fmt.Printf("Type: nil\n")

	default:

		var iStr fmt.Stringer
		var ok bool

		iStr,
			ok = rawNewSound.(fmt.Stringer)

		if ok {
			fmt.Printf("From Default: Type Is fmt.Stringer\n")
			fmt.Printf("%v\n",
				iStr.String())
			return
		}

		fmt.Println("Other")
	}

}

func (hawk *Hawk) CheckSound2(rawNewSound interface{}) {

	var iStr fmt.Stringer
	var ok bool

	iStr,
		ok = rawNewSound.(fmt.Stringer)

	if ok {
		fmt.Printf("Is fmt.Stringer\n")
		fmt.Printf("%v\n",
			iStr.String())
		return
	}

	_,
		ok = rawNewSound.(big.Int)

	if ok {
		fmt.Printf("big.Int\n")
		return
	}

}
func (hawk *Hawk) ConvertSoundToStr(rawNewSound interface{}) (soundStr string) {

	var realString string
	var ok bool
	var iStringer fmt.Stringer

	switch rawNewSound.(type) { // the switch uses the type of the interface

	case uint, uint8, uint16, uint32, uint64:
		fmt.Printf("Type: Unsigned Integer\n")

		soundStr = fmt.Sprintf("%v",
			rawNewSound)

		return soundStr

	case int, int8, int16, int32, int64:
		fmt.Println("Type: Int")

		soundStr = fmt.Sprintf("%v",
			rawNewSound)

		return soundStr

	case bool:
		fmt.Println("Type: bool")

		soundStr = fmt.Sprintf("%v",
			rawNewSound)

		return soundStr

	case *big.Int:
		fmt.Printf("Type: big.Int\n")

		soundStr = fmt.Sprintf("%v",
			rawNewSound)

		return soundStr

	case *big.Float:
		fmt.Printf("Type: big.Float\n")

		soundStr = fmt.Sprintf("%v",
			rawNewSound)

		return soundStr

	case fmt.Stringer:
		fmt.Printf("Type: fmt.Stringer \n")
		iStringer,
			ok = rawNewSound.(fmt.Stringer)

		if !ok {
			fmt.Printf("Error: Could NOT convert empty interface to fmt.Stringer!!\n")
			return
		}

		soundStr = iStringer.String()

		return soundStr

	case string:
		fmt.Println("Type: String")

		realString,
			ok = rawNewSound.(string)

		if !ok {
			fmt.Printf("Error: Could NOT convert empty interface to string!!\n")
			return
		}

		soundStr = realString

		return soundStr

	case nil:
		fmt.Printf("Type: nil\n")
		soundStr = "nil"
		return soundStr

	default:

		fmt.Println("Other")
		soundStr = fmt.Sprintf("Could NOT convert rawNewSound to string!\n"+
			"interface{} = Type:  %T\n",
			rawNewSound)
	}

	return soundStr
}

func SetSubSound[T IInputParamToStr](newSound T, repeat int) string {

	str := fmt.Sprintf("%v", newSound)

	return strings.Repeat(str, repeat)

}

type Bird[T any] struct {
	sound  string
	repeat int
	param  T
}

func (bird *Bird[param]) SetSound(newSound param, repeat int) string {
	locStr := fmt.Sprintf(
		"%v",
		newSound)

	return strings.Repeat(locStr, repeat)
}

/*
https://tip.golang.org/ref/spec#Method_declarations

type Pair[A, B any] struct {
	a A
	b B
}

func (p Pair[A, B]) Swap() Pair[B, A]  { … }  // receiver declares A, B
func (p Pair[First, _]) First() First  { … }  // receiver declares First, corresponds to A in Pair

*/

func main() {

	hawk := Hawk{
		sound:  "awk",
		repeat: 1,
	}

	// Little Hawk
	fmt.Printf("Testing Little Hawk ConvertSoundToStr\n")

	littleHawk := LittleHawk{
		sound:  "EEEEE!",
		repeat: 1,
	}

	result := hawk.ConvertSoundToStr(littleHawk)
	fmt.Printf("Result: %v\n\n", result)

	// big.Int
	fmt.Printf("Testing big.Int ConvertSoundToStr\n")

	bigI := big.NewInt(0)
	var ok bool

	bigI,
		ok =
		bigI.SetString("123456789112345678921234567893", 10)

	if !ok {
		fmt.Printf("Error: \n" +
			"bigI.SetString(\"123456789112345678921234567893\", 10)\n" +
			"This String Parsing Function FAILED!\n")

		return
	}

	result = hawk.ConvertSoundToStr(bigI)
	fmt.Printf("Result: %v\n\n", result)

	// A Real String
	fmt.Printf("Testing Real String ConvertSoundToStr\n")

	result = hawk.ConvertSoundToStr("Hello World!")
	fmt.Printf("Result: %v\n\n", result)

	// boolean value
	fmt.Printf("Testing 'bool' ConvertSoundToStr\n")
	var isReal bool
	isReal = true
	result = hawk.ConvertSoundToStr(isReal)
	fmt.Printf("Result: %v\n\n", result)

	return
}

func MainSub01() {

	var isBird bool

	isBird = false

	myBird := Hawk{}

	result := myBird.SetSound(isBird, 2)

	fmt.Printf("Result: %v\n\n",
		result)

	littleHawk := LittleHawk{
		sound:  "awk",
		repeat: 1,
	}

	fmt.Printf("Test Little Hawk\n")
	result = myBird.SetSound(littleHawk, 1)

	fmt.Printf("Result: %v\n\n",
		result)

}

func MainSub02() {

	hawk := Hawk{
		sound:  "awk",
		repeat: 1,
	}

	fmt.Printf("Testing Little Hawk CheckSound\n")

	littleHawk := LittleHawk{
		sound:  "EEEEE!",
		repeat: 1,
	}

	hawk.CheckSound(littleHawk)

	// Testing big.Int

	fmt.Printf("Testing 'big.Int'\n")

	bigI := big.NewInt(0)
	var ok bool

	bigI,
		ok =
		bigI.SetString("123456789112345678921234567893", 10)

	if !ok {
		fmt.Printf("Error: \n" +
			"bigI.SetString(\"123456789112345678921234567893\", 10)\n" +
			"This String Parsing Function FAILED!\n")

		return
	}

	hawk.CheckSound(bigI)

	// -----------------------------------------
	// big.Float

	fmt.Printf("Testing 'big.Float'\n")

	bigF := big.NewFloat(0)

	bigF.SetPrec(100)
	bigF.SetMode(big.ToNearestAway)

	var actualBase int

	var err error

	bigF,
		actualBase,
		err =
		bigF.Parse("1.2345679112345678921234567893", 10)

	if err != nil {
		fmt.Printf("Error: \n"+
			"bigF.Parse(\"1.234567901234567890\", 10)\n"+
			"%v\n",
			err.Error())

		return
	}

	if actualBase != 10 {
		fmt.Printf("Error: \n"+
			"bigF.Parse(\"1.234567901234567890\", 10)\n"+
			"'actualBase' is NOT EQUAL to '10'\n"+
			"actualBase = %v\n",
			actualBase)

		return
	}

	hawk.CheckSound(bigF)

	// uint32 Unsigned Integer
	var uI uint32 = 57

	fmt.Printf("Testing Composite 'uint'\n")

	hawk.CheckSound(uI)

	// int32 Signed Integer
	var i32 int32 = 82

	fmt.Printf("Testing 'int32'\n")

	hawk.CheckSound(i32)

	// Test 'nil'

	fmt.Printf("Testing 'nil'\n")

	hawk.CheckSound(nil)

	// Test fmt.Stringer
	littleHawk = LittleHawk{
		sound:  "awk-awk",
		repeat: 1,
	}

	hawk.CheckSound(littleHawk)

	return

}
