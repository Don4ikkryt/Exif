package readPhotoes

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadPhotoesFromDirectory(path string) []string {
	Filenames := make([]string, 0)

	files, err := ioutil.ReadDir(path)
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
