package main

import (
		"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadPhotoesFromDirectory() []string {
	Filenames := make([]string, 0)

	files, err := ioutil.ReadDir("C:\\Users\\don4i\\Desktop\\Go\\Exif")
	if err != nil {
		log.Fatal(err)
	}
	writeDirectories(&files, &Filenames)
	return Filenames
}

func writeDirectories(files *[]os.FileInfo, Filenames *[]string) {
	FileExtension1 := ".jpg"
	FileExtension2 := ".JPG"
	for _, file := range *files {
		if strings.HasSuffix(file.Name(), FileExtension1) || strings.HasSuffix(file.Name(), FileExtension2) {
			*Filenames = append(*Filenames, file.Name())
		}

	}
}
