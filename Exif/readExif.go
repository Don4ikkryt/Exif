package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/rwcarlsen/goexif/tiff"
)

func decodeImg(file *os.File) (exifFile *exif.Exif) {
	exif.RegisterParsers(mknote.All...)
	exifFile, err := exif.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return
}
func getGPSAltitude(file *os.File) (tiffFile *tiff.Tag) {
	tiffFile, err := decodeImg(file).Get(exif.GPSAltitude)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}
	return
}
func RetriveHeight(file *os.File) {
	height := getGPSAltitude(file).Format()

	fmt.Println(height)
}
