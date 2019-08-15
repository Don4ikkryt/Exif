package main

func main() {

	Filenames := ReadPhotoesFromDirectory()
	Heights := GetHeightFromAllPhotoes(Filenames)
	CompareHeight(128, Heights, Filenames)

}
