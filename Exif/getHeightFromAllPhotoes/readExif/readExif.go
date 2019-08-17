package readExif

import (
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
	exifFile := decodeImg(file)
	tiffFile, err := exifFile.Get(exif.GPSAltitude)
	if err != nil {

		log.Fatal(err)
	}

	return
}
func RetriveHeight(file *os.File) (toCloseFIle *os.File, stringValueOfHeight string) {
	stringValueOfHeight = getGPSAltitude(file).String()
	toCloseFIle = file
	return
}
