package main

import (
	"os"
)

func CompareHeight(height float64, heights map[string]float64, filenames []string) {
	path := "C:\\Users\\don4i\\Desktop\\Go\\Exif\\Nonvalid"
	createNonvalidDir(path)
	for _, filename := range filenames {

		if heights[filename] < height {
			os.Rename("C:\\Users\\don4i\\Desktop\\Go\\Exif\\"+filename, "C:\\Users\\don4i\\Desktop\\Go\\Exif\\Nonvalid\\"+filename)
		}

	}

}
func createNonvalidDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}
