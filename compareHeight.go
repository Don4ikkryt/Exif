package main
import (
	"os"
)

func CompareHeight(height float64, heights map[string]float64, filenames []string, path string) {
	pathOfNonvalidFolder := path + "\\Nonvalid"
	createNonvalidDir(pathOfNonvalidFolder)
	for _, filename := range filenames {

		if heights[filename] < height {
			oldPath := path + "\\" + filename
			newPath := pathOfNonvalidFolder + "\\" + filename
			os.Rename(oldPath, newPath)
		}

	}

}
func createNonvalidDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}
