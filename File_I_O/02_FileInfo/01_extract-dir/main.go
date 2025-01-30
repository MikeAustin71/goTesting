package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	var inputFile = "./SampleFile.txt"

	fmt.Println("Initial Path and File Name: ", inputFile)

	absPath, _ := filepath.Abs(inputFile)

	fmt.Println("Absolute Path: ", absPath)

	fmt.Println("\nDir(): ", filepath.Dir(absPath))

	ext := filepath.Ext(absPath)

	fmt.Println("Ext(): ", ext)

	base := filepath.Base(absPath)

	fmt.Println("Base(): ", base)

	fmt.Println("File Name:", strings.TrimRight(base, ext))

	fmt.Println("VolName() ", filepath.VolumeName(absPath))

}

/*	Output - running on a Windows Machine
$ go run main.go
Initial Path and File Name:  ./SampleFile.txt
Absolute Path:  D:\go\work\src\bitbucket.org\AmarilloMike\GolangMikeSamples\File_I_O\02_FileInfo\01_extract-dir\SampleFile.txt

Dir():  D:\go\work\src\bitbucket.org\AmarilloMike\GolangMikeSamples\File_I_O\02_FileInfo\01_extract-dir
Ext():  .txt
Base():  SampleFile.txt
File Name: SampleFile
VolName()  D:

*/
