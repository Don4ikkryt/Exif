package main

import (
	"./cmd"
	"./compareHeight"
	"./getHeightFromAllPhotoes"
	"./readPhotoes"
)

func main() {

	SetHeight, SetPath := cmd.CommandLine()
	Filenames := readPhotoes.ReadPhotoesFromDirectory(*SetPath)
	Heights := getHeightFromAllPhotoes.GetHeightFromAllPhotoes(Filenames, *SetPath)
	compareHeight.CompareHeight(*SetHeight, Heights, Filenames, *SetPath)

}
