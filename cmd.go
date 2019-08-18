package main

import (
	"flag"
)

func CommandLine() (setHeight *float64, setPath *string) {
	setHeight = flag.Float64("Height", 0, "Height of dron")
	setPath = flag.String("Path", "", "Path to the folder with photoes")
	flag.Parse()
	return
}
