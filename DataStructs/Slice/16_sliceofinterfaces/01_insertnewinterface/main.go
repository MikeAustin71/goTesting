package main

import "fmt"

type IThing interface {
	GetName() string
}

type Bird struct {
	Name string
}

func (thingBird Bird) GetName() string {
	return "Bird"
}

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

type Owl struct {
	Name string
}

func (thingOwl Owl) GetName() string {
	return "Owl"
}

type Human struct {
	Name string
}

func (thingHuman Human) GetName() string {
	return "Human"
}

type Robot struct {
	Name string
}

func (thingRobot *Robot) GetName() string {
	return "Robot"
}

func main() {

	//var things = []IThing {
	//	Bird{},
	//	Owl{},
	//	Cat{},
	//	Human{},
	//	Robot{},
	//
	//}

	var things = make([]IThing, 5)

	things[0] = Bird{}
	things[1] = Owl{}
	things[2] = Cat{}
	things[3] = Human{}
	things[4] = &Robot{}

	var err error

	err = PrintThings(things)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	err = ReplaceTheCat(things)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	fmt.Printf("Things After Replace The Cat!\n")
	err = PrintThings(things)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	things,
		err = DeleteTheCat(things)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	lenThings := len(things)

	fmt.Printf("main()\n"+
		"Length Of 'things' after deletion='%v'\n",
		lenThings)

	err = PrintThings(things)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	fmt.Printf("\nMain\n" +
		"       SUCCESSFUL COMPLETION!\n")
}

func PrintThings(realThings []IThing) error {

	ePrefix := "PrintThings()"
	var err error
	if len(realThings) == 0 {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'realThings' is an empty slice!\n",
			ePrefix)
		return err
	}

	fmt.Println(ePrefix)

	lenRealThings := len(realThings)

	for i := 0; i < lenRealThings; i++ {

		fmt.Printf("%v. Thing = '%v'\n",
			i+1,
			realThings[i].GetName())

	}

	fmt.Println()

	return nil
}

func ReplaceTheCat(realThings []IThing) error {

	ePrefix := "ReplaceTheCat()"
	var err error

	if len(realThings) == 0 {
		err = fmt.Errorf("%v\n"+
			"ERROR Input parameter 'realThings' is an empty slice!\n",
			ePrefix)

		return err
	}

	targetIdx := -1

	for i := 0; i < len(realThings); i++ {

		if realThings[i].GetName() == "Cat" {
			targetIdx = i
			break
		}

	}

	if targetIdx == -1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Could NOT locate 'Cat' in realThings slice.\n",
			ePrefix)

		return err
	}

	newDog := Dog{}

	var replacementThing = []IThing{
		newDog, // newDog and &newDog both work
	}

	// this works
	itemsCopied := copy(realThings[targetIdx:targetIdx+1], replacementThing[:])

	fmt.Printf("%v\n"+
		"Items Copied: %v\n",
		ePrefix,
		itemsCopied)

	return nil
}

func DeleteTheCat(realThings []IThing) ([]IThing, error) {

	ePrefix := "DeleteTheCat()"
	var err error
	thingToDelete := "Human"
	lenRealThings := len(realThings)

	if lenRealThings == 0 {
		err = fmt.Errorf("%v\n"+
			"ERROR Input parameter 'realThings' is an empty slice!\n",
			ePrefix)

		return []IThing{}, err
	}

	fmt.Printf("\n%v\n"+
		"Starting lenRealThings= %v\n",
		ePrefix,
		lenRealThings)

	var targetIdx = -1

	for i := 0; i < lenRealThings; i++ {

		if realThings[i].GetName() == thingToDelete {
			targetIdx = i
			break
		}

	}

	if targetIdx == -1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Could NOT locate '%v' in Slice 'realThings'!\n",
			ePrefix,
			thingToDelete)

		return []IThing{}, err
	}

	var result = make([]IThing, lenRealThings-1)

	copy(result[0:], realThings[0:targetIdx])

	if targetIdx < lenRealThings-1 {
		copy(result[targetIdx:], realThings[targetIdx+1:])
	}

	lenRealThings = len(result)

	fmt.Printf("\n%v\n"+
		"Ending Length Result= %v\n",
		ePrefix,
		lenRealThings)

	return result, nil
}
