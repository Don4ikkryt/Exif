package main

import "fmt"

func main() {

	SetHeight, SetPath := CommandLine()
	fmt.Println(*SetHeight)
	fmt.Println(*SetPath)
	Filenames := ReadPhotoesFromDirectory(*SetPath)
	//fmt.Println(Filenames)
	Heights := GetHeightFromAllPhotoes(Filenames, *SetPath)
	//fmt.Println(Heights)
	CompareHeight(110, Heights, Filenames, *SetPath)

}
