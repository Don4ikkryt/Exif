package main

import (
	"log"
	"os"
)

func OpenPhoto(filename string, path string) *os.File {
	image, err := os.Open(path + "\\" + filename)
	if err != nil {
		log.Fatal(err)
	}

	return image
}
