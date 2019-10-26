package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	f := gofpdf.New("P", "mm", "A4", "")
	f.AddPage()
	f.SetFillColor(255, 0, 0)
	f.SetLineWidth(0.1)
	f.Circle(30, 40, 2, "F")
	file := "C:\\Users\\don4i\\Desktop\\Go\\PDF\\drd"
	err := f.OutputFileAndClose(file)
	if err != nil {
		fmt.Println("err")
	}
}
