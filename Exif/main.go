package main

import "fmt"

func main() {

	Filenames := ReadPhotoesFromDirectory()
	Heights := GetHeightFromAllPhotoes(Filenames)
	for _, height := range Heights {
		fmt.Println(height)
	}

}
