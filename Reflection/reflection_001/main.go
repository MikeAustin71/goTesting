package main

import "fmt"

type person struct {
	name string
	age  int
}

type puppyDog struct {
	name string
	age  int
}

func main() {

	x := person{
		name: "Mike",
		age:  72,
	}
	typeSort01(x)

	z := puppyDog{
		name: "Jordy",
		age:  10,
	}

	typeSort02(z)

	typeSort03(x)

}

func typeSort01(x interface{}) {

	var typeName string

	switch x.(type) {
	case bool:
		typeName = "bool"
	case int:
		typeName = "int"
	case person:
		var j person
		j = x.(person)

		typeName = fmt.Sprintf("Type: 'person'\nName: %v Age: %v\n",
			j.name,
			j.age)

	}

	fmt.Println("typeSort01")
	fmt.Println(typeName)
	fmt.Println()
}

func typeSort02(x interface{}) {

	myPuppyDog, ok := x.(puppyDog)

	fmt.Println("typeSort02")
	if !ok {
		fmt.Printf("Hoped for 'puppyDog', but got %T\n",
			x)
		fmt.Println()
		return
	}

	fmt.Printf("My Puppy Dog! Name: %v Age: %v\n",
		myPuppyDog.name, myPuppyDog.age)

	fmt.Println()
}

func typeSort03(x interface{}) {

	myPuppyDog, ok := x.(puppyDog)

	fmt.Println("typeSort03")
	if !ok {
		fmt.Printf("Hoped for 'puppyDog', but got %T\n",
			x)
		fmt.Println()
		return
	}

	fmt.Printf("My Puppy Dog! Name: %v Age: %v",
		myPuppyDog.name, myPuppyDog.age)
	fmt.Println()

}
