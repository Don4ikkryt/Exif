package main

import (
	"log"
	"strconv"
	"strings"
)

func GetHeightFromAllPhotoes(filenames []string, path string) map[string]float64 {
	var heights map[string]float64 = make(map[string]float64)
	for _, filename := range filenames {
		file, height, errOfExif := RetriveHeight(OpenPhoto(filename, path), path, filename)
		file.Close()
		if errOfExif != 0 {
			continue
		}
		heights[filename] = fromStringToFloat64(height)

	}
	return heights
}
func fromStringToFloat64(line string) (floatNum float64) {
	separator := "/"
	newLine := strings.Replace(line, "\"", "", -1)
	substrings := strings.Split(newLine, separator)
	firstNumber, err := strconv.ParseFloat(substrings[0], 64)
	if err != nil {
		log.Fatal(err)
	}
	secondNumber, err := strconv.ParseFloat(substrings[1], 64)
	if err != nil {
		log.Fatal(err)
	}
	return firstNumber / secondNumber
}
