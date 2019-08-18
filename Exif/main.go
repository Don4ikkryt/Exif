package main

import (
	"./cmd"
	"./compareheight"
	"./getheightfromallphotoes"
	"./readphotoes"
)

func main() {

	SetHeight, SetPath := cmd.CommandLine()
	Filenames := readphotoes.ReadPhotoesFromDirectory(*SetPath)
	Heights := getheightfromallphotoes.GetHeightFromAllPhotoes(Filenames, *SetPath)
	compareheight.CompareHeight(*SetHeight, Heights, Filenames, *SetPath)

}
