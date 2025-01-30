package common

import "fmt"

type DirInfo struct{
	Name	string
	DirList [] string
	FileList [] string
	ErrList [] string
}

func TestSlice001(dInfo *DirInfo) {
	for i := 0; i < 50; i++ {
		createAString(dInfo, i)
	}
}

func createAString(dInfo *DirInfo, index int) {

	dInfo.FileList = append(dInfo.FileList, fmt.Sprintf("My String # %v", index))
}
