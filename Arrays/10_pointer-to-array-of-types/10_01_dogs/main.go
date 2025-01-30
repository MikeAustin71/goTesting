package main

import "fmt"

type IThing interface {
	Empty()
	GetName() *string
	IsEmpty() bool
	SetName()
}

type Bird struct {
	Name *string
}

func (bird *Bird) Empty() {
	bird.Name = nil
}

func (bird *Bird) GetName() *string {
	return bird.Name
}

func (bird *Bird) IsEmpty() bool {
	return bird.Name == nil
}

func (bird Bird) New() (newBird *Bird) {

	newBird = &Bird{}
	newBird.SetName()

	return newBird
}

func (bird *Bird) SetName() {

	name := "Bird"
	bird.Name = &name
}

type Cat struct {
	Name *string
}

func (cat *Cat) Empty() {
	cat.Name = nil
}

func (cat *Cat) GetName() *string {

	return cat.Name
}

func (cat *Cat) IsEmpty() bool {
	return cat.Name == nil
}

func (cat Cat) New() (newCat *Cat) {

	newCat = &Cat{}
	newCat.SetName()

	return newCat
}

func (cat *Cat) SetName() {
	name := "Cat"
	cat.Name = &name
}

type Dog struct {
	Name *string
}

func (dog *Dog) Empty() {
	dog.Name = nil
}

func (dog *Dog) GetName() *string {
	return dog.Name
}

func (dog *Dog) IsEmpty() bool {
	return dog.Name == nil
}

func (dog Dog) New() (newDog *Dog) {

	newDog = &Dog{}
	newDog.SetName()

	return newDog
}

func (dog *Dog) SetName() {

	name := "Dog"
	dog.Name = &name
}

type Human struct {
	Name *string
}

func (human *Human) Empty() {
	human.Name = nil
}

func (human *Human) GetName() *string {
	return human.Name
}

func (human *Human) IsEmpty() bool {
	return human.Name == nil
}

func (human Human) New() (newHuman *Human) {

	newHuman = &Human{}
	newHuman.SetName()

	return newHuman
}

func (human *Human) SetName() {
	name := "Human"
	human.Name = &name
}

type Robot struct {
	Name *string
}

func (robot *Robot) Empty() {
	robot.Name = nil
}

func (robot *Robot) GetName() *string {

	return robot.Name
}

func (robot *Robot) IsEmpty() bool {
	return robot.Name == nil
}

func (robot Robot) New() (newRobot *Robot) {

	newRobot = &Robot{}
	newRobot.SetName()

	return newRobot
}

func (robot *Robot) SetName() {
	name := "Robot"
	robot.Name = &name
}

func main() {

	runTestEmptyArray()
}

func runTestEmptyArray() {

	things := make([]IThing, 5)
	things[0] = Dog{}.New()
	things[1] = Cat{}.New()
	things[2] = Bird{}.New()
	things[3] = Human{}.New()
	things[4] = Robot{}.New()

	ptrThing := &things

	fmt.Printf("----------------------------------\n")
	fmt.Printf("Running Initial Ptr Array Listing \n")
	fmt.Printf("----------------------------------\n")
	testPtrArray(ptrThing)
	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf("Running Array Empty Operation \n")
	fmt.Printf("----------------------------------\n")
	testEmptyPtrArray02(ptrThing)
	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf("Running 2nd Ptr Array Listing \n")
	fmt.Printf("----------------------------------\n")
	testPtrArray(ptrThing)
	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf("Running 2nd Concrete Array Listing\n")
	fmt.Printf("----------------------------------\n")
	testArray(things)
	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf("Setting Thing Elements to nil     \n")
	fmt.Printf("----------------------------------\n")
	testSetPtrArrayElementsToNil(ptrThing)
	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf("Running Concrete Array Listing\n")
	fmt.Printf("----------------------------------\n")
	testArray(things)
	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf("Setting Ptr Array to Nil\n")
	fmt.Printf("----------------------------------\n")
	testSetPtrArrayToNil(ptrThing)
	fmt.Printf("------------------------------------\n")
	fmt.Printf("Running Final Concrete Array Listing\n")
	fmt.Printf("------------------------------------\n")
	testArray(things)

}

func runPtrTestArray() {

	things := make([]IThing, 5)
	things[0] = Dog{}.New()
	things[1] = Cat{}.New()
	things[2] = Bird{}.New()
	things[3] = Human{}.New()
	things[4] = Robot{}.New()

	testPtrArray(&things)
}

func runTestArray() {

	things := make([]IThing, 6)

	things[0] = &Dog{}
	things[1] = &Cat{}
	things[2] = &Bird{}
	things[3] = &Human{}
	things[4] = &Robot{}
	things[5] = nil

	for i := 0; i < 5; i++ {
		things[i].SetName()
	}

	testArray(things)
}

func runTestNil() {

	fmt.Println("Trying &Dog")
	testNil(&Dog{})

	fmt.Println("Trying &Robot")
	testNil(&Robot{})

	fmt.Println("Trying Cat")
	testNil(&Cat{})

	fmt.Println("Trying Human")
	testNil(&Human{})

	fmt.Println("Trying Bird")
	testNil(&Bird{})

	fmt.Println("Trying 'nil'")
	testNil(nil)

}

func testPtrArray(things *[]IThing) {

	funcName := "testPtrArray(things *[]IThing)"

	if things == nil {
		fmt.Printf("%v - ERROR\n"+
			"Input parameter '*things' is nil!\n",
			funcName)

		return
	}

	emptyCnt := 0
	fullCnt := 0
	nilCnt := 0

	for idx, val := range *things {

		if val == nil {

			fmt.Printf("Index= '%v' - 'thing' is 'nil'!\n",
				idx)

			nilCnt++

		} else if val.IsEmpty() {

			fmt.Printf("Index= '%v' - 'thing' is 'Empty' - Type= %T\n",
				idx,
				val)

			emptyCnt++

		} else {

			fmt.Printf("Index= '%v' - 'thing' is a %v\n",
				idx,
				*val.GetName())
			fullCnt++
		}
	}

	fmt.Printf("\n"+
		"Results: %v \n"+
		"Full Valid Elements: %v\n"+
		"     Empty Elements: %v\n"+
		"       Nil Elements; %v\n"+
		"-------------------------------\n"+
		"     Total Elements: %v\n",
		funcName,
		fullCnt,
		emptyCnt,
		nilCnt,
		fullCnt+emptyCnt+nilCnt)

	fmt.Println()

}

func testSetPtrArrayToNil(things *[]IThing) {

	*things = nil
}

// This works! All individual array elements
// are set to nil in the underlying concrete
// array.
func testSetPtrArrayElementsToNil(things *[]IThing) {

	funcName := "testEmptyPtrArray(things *[]IThing)"

	if things == nil {
		fmt.Printf("\n%v - ERROR\n"+
			"Input parameter '*things' is nil!\n\n",
			funcName)

		return
	}

	concreteThings := *things

	lenThings := len(concreteThings)
	nilCnt := 0
	fullCnt := 0
	emptyCnt := 0

	for i := 0; i < lenThings; i++ {

		if concreteThings[i] == nil {

			fmt.Printf("Index= '%v' - 'thing' is 'nil'!\n",
				i)

			nilCnt++

			continue

		} else {

			if concreteThings[i].IsEmpty() {

				fmt.Printf("Index= '%v' - Type= %T - "+
					"'thing' is 'Empty'!\n",
					i,
					concreteThings[i])

				emptyCnt++

			} else {
				// Must be Full and Valid

				fmt.Printf("Index= '%v' - 'thing' is a %v\n",
					i,
					*concreteThings[i].GetName())

				fullCnt++

			}

			concreteThings[i] = nil
		}

	}

	fmt.Printf("======================\n"+
		"%v - Results\n"+
		"         nil item count = %v\n"+
		"valid filled item count = %v\n"+
		"       empty item count = %v\n"+
		"            total items = %v\n"+
		"======================\n\n",
		funcName,
		nilCnt,
		fullCnt,
		emptyCnt,
		nilCnt+fullCnt+emptyCnt)

}

// This doesn't work. Setting 'val' to nil doesn't
// set the underlying array element to nil.
func testEmptyPtrArray(things *[]IThing) {

	if things == nil {
		fmt.Printf("\n%v - ERROR\n"+
			"Input parameter '*things' is nil!\n\n",
			"testEmptyPtrArray(things *[]IThing)")

		return
	}

	for idx, val := range *things {

		if val == nil {
			fmt.Printf("Index= '%v' - 'thing' is 'nil'!\n",
				idx)

		} else {
			fmt.Printf("Index= '%v' - Set 'thing' %v to nil.\n",
				idx,
				*val.GetName())

			val = nil
		}
	}

	return
}

// Call Empty On All Array Members
func testEmptyPtrArray02(things *[]IThing) {

	funcName := "testEmptyPtrArray02(things *[]IThing)"

	if things == nil {
		fmt.Printf("\n%v - ERROR\n"+
			"Input parameter '*things' is nil!\n\n",
			funcName)

		return
	}

	concreteThings := *things

	lenThings := len(concreteThings)

	emptyCnt := 0
	fullCnt := 0
	nilCnt := 0

	for i := 0; i < lenThings; i++ {

		if concreteThings[i] == nil {
			fmt.Printf("Index= '%v' - 'thing' is "+
				"'nil'!\n",
				i)

			nilCnt++

		} else if concreteThings[i].IsEmpty() {

			fmt.Printf("Index= '%v' - 'thing' is "+
				"'Empty' - Type= %T\n",
				i,
				concreteThings[i])

			emptyCnt++

		} else {
			fmt.Printf("Index= '%v' - Set 'thing' %v "+
				"to Empty.\n",
				i,
				*concreteThings[i].GetName())

			concreteThings[i].Empty()

			fullCnt++
		}

	}

	fmt.Printf("\n"+
		"Results: %v \n"+
		"Full Valid Elements: %v\n"+
		"     Empty Elements: %v\n"+
		"       Nil Elements; %v\n"+
		"-------------------------------\n"+
		"     Total Elements: %v\n",
		funcName,
		fullCnt,
		emptyCnt,
		nilCnt,
		fullCnt+emptyCnt+nilCnt)

	fmt.Println()

	return
}

func testArray(things []IThing) {

	funcName := "testArray(things []IThing)"

	if things == nil {
		fmt.Printf("%v - ERROR\n"+
			"Input parameter 'things' is nil!\n",
			funcName)

		return
	}

	nilCnt := 0
	fullCnt := 0
	emptyCnt := 0

	for i := 0; i < len(things); i++ {

		if things[i] == nil {

			fmt.Printf("Index= '%v' - 'thing' is 'nil'!\n",
				i)

			nilCnt++

		} else if things[i].IsEmpty() {

			fmt.Printf("Index= '%v' - Type= %T - "+
				"'thing' is 'Empty'!\n",
				i,
				things[i])

			emptyCnt++

		} else {

			fmt.Printf("Index= '%v' - 'thing' is a %v\n",
				i,
				*things[i].GetName())

			fullCnt++

		}

		fmt.Println()
	}

	fmt.Printf("======================\n"+
		"%v - Results\n"+
		"         nil item count = %v\n"+
		"valid filled item count = %v\n"+
		"       empty item count = %v\n"+
		"            total items = %v\n"+
		"======================\n\n",
		funcName,
		nilCnt,
		fullCnt,
		emptyCnt,
		nilCnt+fullCnt+emptyCnt)
}

func testNil(thing IThing) {

	if thing == nil {
		fmt.Println("'thing' is 'nil'!")

		fmt.Println()

		return
	}

	thing.SetName()

	var output *string

	output = thing.GetName()

	fmt.Printf("thing is a %v\n",
		*output)

	fmt.Println()

}
