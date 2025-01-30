package main

import (
	"bitbucket.org/AmarilloMike/GolangMikeSamples/TypeArchitecture/03_TypeNewPtr/app/appTest"
	"fmt"
)

func main() {
	appTest01()
	appTest02()

}

func appTest01() {

	// This works
	num1 := 1
	num2 := 2

	sum := appTest.NumMgr{}.NewPtr().AddNums(num1,num2)

	fmt.Println("---- appTest01() ----")
	fmt.Printf("Sum of %v + %v = %v\n",
		num1, num2, sum)
	fmt.Println("---------------------")
	fmt.Println()

}

func appTest02() {
	// This Works

	num1 := 1
	num2 := 2

	sum := appTest.NumMgr{}.This().AddNums(num1,num2)

	fmt.Println("---- appTest02() ----")
	fmt.Printf("Sum of %v + %v = %v\n",
		num1, num2, sum)
	fmt.Println("---------------------")
	fmt.Println()

}