package main

import "fmt"

func main() {

	SetHeight, SetPath := CommandLine()
	fmt.Println(*SetHeight)
	fmt.Println(*SetPath)
	Filenames := ReadPhotoesFromDirectory()
	//fmt.Println(Filenames)
	Heights := GetHeightFromAllPhotoes(Filenames)
	//fmt.Println(Heights)
	CompareHeight(110, Heights, Filenames)

}
