package main

import (
	"fmt"
)

func main()  {

	sSlice := make([] string, 0)
	sSlice = append(sSlice, "B1")
	sSlice = append(sSlice, "B2")
	sSlice = append(sSlice, "B3")
	sSlice = append(sSlice, "B4")
	sSlice = append(sSlice, "B5")

	subFunc(sSlice)

	fmt.Println("sSlice # 1 :", sSlice)
	// Prints Out
	// sSlice # 1 : [B1 B2 B3 B4 B5]

}

func subFunc( s1 []string) error {
	
	s2:= make([]string, 0, 10)
	s2 = append(s2,"H1")
	s2 = append(s2,"H2")
	s2 = append(s2,"H3")
	s2 = append(s2,"H4")
	s2 = append(s2,"H5")

	s1 = s2[0:]
	return nil
}
