package main

func main() {

	Filenames := ReadPhotoesFromDirectory()
	//fmt.Println(Filenames)
	Heights := GetHeightFromAllPhotoes(Filenames)
	//fmt.Println(Heights)
	CompareHeight(110, Heights, Filenames)

}
