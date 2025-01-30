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

	fmt.Println("Trying &Dog")
	testNil(&Dog{})

	fmt.Println("Trying &Robot")
	testNil(&Robot{})

	fmt.Println("Trying Cat")
	testNil(Cat{})

	fmt.Println("Trying Human")
	testNil(Human{})

	fmt.Println("Trying Owl")
	testNil(Owl{})

	fmt.Println("Trying Bird")
	testNil(Bird{})

	fmt.Println("Trying 'nil'")
	testNil(nil)

}

func testNil(thing IThing) {

	if thing == nil {
		fmt.Println("'thing' is 'nil'!")

		fmt.Println()

		return
	}

	fmt.Println(thing.GetName())

	fmt.Println()

}

/* Print Out
$ go run main.go
Trying &Dog
Dog

Trying &Robot
Robot

Trying Cat
Cat

Trying Human
Human

Trying Owl
Owl

Trying Bird
Bird

Trying 'nil'
'thing' is 'nil'!

*/
