package main

import (
	"fmt"
)

func main() {

	Filenames := ReadPhotoesFromDirectory()
	for _, value := range Filenames {
		fmt.Println(value)
	}

	RetriveHeight(openPhoto(string("C:\\Users\\don4i\\Desktop\\Go\\Exif\\" + Filenames[0])))

}
