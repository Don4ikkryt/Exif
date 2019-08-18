package main

func main() {

	SetHeight, SetPath := CommandLine()
	Filenames := ReadPhotoesFromDirectory(*SetPath)
	Heights := GetHeightFromAllPhotoes(Filenames, *SetPath)
	CompareHeight(*SetHeight, Heights, Filenames, *SetPath)

}
