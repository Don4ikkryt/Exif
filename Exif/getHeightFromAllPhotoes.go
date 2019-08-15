package main

import (
	"log"
	"strconv"
	"strings"
)

func GetHeightFromAllPhotoes(filenames []string) (heights []float64) {

	Path := "C:\\Users\\don4i\\Desktop\\Go\\Exif\\"
	for _, filename := range filenames {
		file, height := RetriveHeight(openPhoto(string(Path + filename)))
		file.Close()
		heights = append(heights, fromStringToFloat64(height))
	}
	return
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
