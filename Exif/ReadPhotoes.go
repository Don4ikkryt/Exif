package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type filename string

func ReadPhotoesFromDirectory() []filename {
	Filenames := make([]filename, 0)

	files, err := ioutil.ReadDir("C:\\Users\\don4i\\Desktop\\Go\\Exif")
	if err != nil {
		log.Fatal(err)
	}
	writeDirectories(&files, &Filenames)
	return Filenames
}

func writeDirectories(files *[]os.FileInfo, Filenames *[]filename) {
	FileExtension1 := ".jpg"
	FileExtension2 := ".JPG"
	for _, file := range *files {
		fmt.Println(file.Name())
		if strings.HasSuffix(file.Name(), FileExtension1) || strings.HasSuffix(file.Name(), FileExtension2) {
			*Filenames = append(*Filenames, filename(file.Name()))
		}

	}
}
