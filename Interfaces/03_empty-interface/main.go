package main

import "fmt"

type Cat struct {
	Name string
}

func (thingCat Cat) GetName() string {
	return "Cat"
}

type Dog struct {
	Name     string
	Location string
	CollarId int
}

func (thingDog Dog) GetName() string {
	return "Dog"
}

func main() {

	testCompareEmptyInterfaces(Cat{}, Dog{})
}

func mainTestEmptyInterface() {

	testEmptyInterface(nil)

	testEmptyInterface(Cat{})

	testEmptyInterface("string says 'hello'!")

	var twoDArray = make([][2]string, 3)

	twoDArray[0][0] = "Xray00"
	twoDArray[0][1] = "Xray01"
	twoDArray[1][0] = "Alpha10"
	twoDArray[1][1] = "Alpha11"
	twoDArray[2][0] = "Delta20"
	twoDArray[2][1] = "Delta21"

	testEmptyInterface(twoDArray)

	testEmptyInterface(Dog{})

	testEmptyInterface(&Dog{})

}

func testEmptyInterface(unknownThing interface{}) {

	if unknownThing == nil {
		fmt.Println("'unknownThing' is nil")
		fmt.Println()
		return
	}

	var ok bool
	var cat Cat

	cat,
		ok = unknownThing.(Cat)

	if ok {
		fmt.Println("I am a ", cat.GetName())
		fmt.Println()
		return
	}

	var str string

	str,
		ok = unknownThing.(string)

	if ok {
		fmt.Println("This is a string: ", str)
		fmt.Println()
		return
	}

	var twoDArray [][2]string

	twoDArray,
		ok = unknownThing.([][2]string)

	if ok {
		lenTwoDArray := len(twoDArray)
		fmt.Println("twoDArray -- ")
		for i := 0; i < lenTwoDArray; i++ {
			fmt.Printf("twoDArray[i][0]='%v'  "+
				"twoDArray[i][1]='%v'\n",
				twoDArray[i][0],
				twoDArray[i][1])
		}
		fmt.Println()
		return
	}

	var dog Dog

	dog,
		ok = unknownThing.(Dog)

	if ok {
		fmt.Println("I am a ", dog.GetName())
		fmt.Println()
		return
	}

	var dogPtr *Dog

	dogPtr,
		ok = unknownThing.(*Dog)

	if ok {
		fmt.Println("*Dog says, \"I am a ", dogPtr.GetName(), "\"")
		fmt.Println()
		return
	}

}

func testCompareEmptyInterfaces(
	unknownThingOne interface{},
	unknownThingTwo interface{}) {

	funcName := "testCompareEmptyInterface"

	if unknownThingOne == nil &&
		unknownThingTwo != nil {

		fmt.Printf("%v\n"+
			"unknownThingOne is nil and\n"+
			"unknownThingTwo is NOT nil!\n"+
			"The two are NOT EQUAL!\n",
			funcName)

		return
	}

	if unknownThingOne != nil &&
		unknownThingTwo == nil {

		fmt.Printf("%v\n"+
			"unknownThingOne is NOT nil and\n"+
			"unknownThingTwo is nil!\n"+
			"The two are NOT EQUAL!\n",
			funcName)

		return
	}

	if unknownThingOne == nil &&
		unknownThingTwo == nil {

		fmt.Printf("%v\n"+
			"unknownThingOne is nil and\n"+
			"unknownThingTwo is nil!\n"+
			"The two are EQUAL!\n",
			funcName)

		return
	}

	if unknownThingOne == unknownThingTwo {

		fmt.Printf("%v\n"+
			"unknownThingOne is EQUAL TO unknownThingTwo!\n"+
			"The two are EQUAL!\n",
			funcName)

		return
	}

	if unknownThingOne != unknownThingTwo {

		fmt.Printf("%v\n"+
			"unknownThingOne is NOT EQUAL TO unknownThingTwo!\n"+
			"The two are NOT EQUAL!\n",
			funcName)

		return
	}

	fmt.Printf("%v\n"+
		"Equivalency status is UNDETERMINED!\n",
		funcName)

	return
}
