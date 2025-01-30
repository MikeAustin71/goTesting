package main

import "fmt"

type stringDto struct {

 str1 string
 defaultStr2 string

}

func main() {
	sDto := stringDto{}
	sDto.defaultStr2 = "something"
	fmt.Println("sDto.defaultStr= ", sDto.defaultStr2)
}
