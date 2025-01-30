package main

import (
	"bitbucket.org/AmarilloMike/GolangMikeSamples/DataStructs/Slice/05_slice-as-parameters/common"
	"fmt"
)

func main() {

	dInfo := common.DirInfo{Name: "Test001"}

	dInfo.DirList = make([]string, 0, 150)
	dInfo.FileList = make([]string, 0, 150)
	dInfo.ErrList = make([]string, 0, 150)
	common.TestSlice001(&dInfo)

	l := len(dInfo.FileList)
	fmt.Println("Length Of FileList:", l)

	if l > 0 {
		for i, f := range dInfo.FileList {
			fmt.Println("FileList:", i, " ", f)
		}
	}

}
