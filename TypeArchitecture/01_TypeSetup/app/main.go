package main

import (
	"fmt"

	"github.com/mikeaustin71/TypeArchitecture/01_TypeSetup/common"
)

func main() {
	n := common.NumMgr{}.New()
	n.SetNum(92)

	fmt.Println("  n.SetNum(92): ", n.Num)

}

/*
func Test01() {

	n := common.NumMgr{}.New()

	fmt.Println("Initialized Value: ", n.Num)

}


func Test02() {

	n, err := common.NumMgr{}.New()

	if err != nil {
		panic(err)
	}

	fmt.Println("  n.Num Value: ", n.Num)

	n.SetNum(36)

	fmt.Println("n.Num Value#2: ", n.Num)

	n.AddNum(10, 10)

	fmt.Println("    Add 10+10: ", n.Num)

}

func Test03() {
	n:= common.NumMgr{}

	d3, err := n.SetNum(25)

	if err != nil {
		panic(err)
	}

	fmt.Println(" n.Num = ", n.Num)
	fmt.Println("d3.Num = ", d3.Num)

	// Change d3 pointer
	d3.SetNum(36)
	fmt.Println("After:")
	fmt.Println("    n.Num = ", n.Num)
	fmt.Println("   d3.Num = ", d3.Num)

	/*
	 		n.Num =  25
			d3.Num =  25
	After:
 			n.Num =  36
			d3.Num =  36


}
*/
