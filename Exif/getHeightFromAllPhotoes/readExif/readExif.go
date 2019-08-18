package readexif

import (
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/rwcarlsen/goexif/tiff"
)

func decodeImg(file *os.File, path string, filename string) (exifFile *exif.Exif, errOfExif int) {
	errOfExif = 0
	exif.RegisterParsers(mknote.All...)
	exifFile, err := exif.Decode(file)
	if err != nil {
		errOfExif = -1
		exifFile = nil
		pathOfNonvalidFolder := path + "\\Nonvalid"
		createNonvalidDirNew(pathOfNonvalidFolder)
		moveToNonvalid(path, filename)
		return
	}
	return
}
func getGPSAltitude(file *os.File, path string, filename string) (tiffFile *tiff.Tag, errOfExif int) {
	exifFile, errOfExif := decodeImg(file, path, filename)
	if errOfExif != 0 {
		errOfExif = -1
		tiffFile = nil
		return
	}
	tiffFile, err := exifFile.Get(exif.GPSAltitude)
	if err != nil {
		pathOfNonvalidFolder := path + "\\Nonvalid"
		createNonvalidDirNew(pathOfNonvalidFolder)
		moveToNonvalid(path, filename)
		errOfExif = -1
		tiffFile = nil
		return
	}

	return
}
func RetriveHeight(file *os.File, path string, filename string) (toCloseFIle *os.File, stringValueOfHeight string, errOfExif int) {
	tiffFile, errOfExif := getGPSAltitude(file, path, filename)
	toCloseFIle = file
	if errOfExif != 0 {
		errOfExif = -1
		stringValueOfHeight = ""
		return
	}
	stringValueOfHeight = tiffFile.String()

	return
}

func createNonvalidDirNew(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}
func moveToNonvalid(oldPath string, filename string) {
	oldPlace := oldPath + "\\Nonvalid" + filename
	newPlace := oldPath + filename
	os.Rename(oldPlace, newPlace)
}
