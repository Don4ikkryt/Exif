package main

 import(
	"log"
	"os"
 )
 func openPhoto(filename string) *os.File {
	image, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return image
}
/*func OpenAllPhotoes(filenames []string, string Path){
	
    for _, value = range filenames{
    
	}

}*/ 