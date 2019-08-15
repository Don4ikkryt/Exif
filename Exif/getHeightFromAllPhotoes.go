package main

func GetHeightFromAllPhotoes(filenames []string) (heights []string) {

	Path := "C:\\Users\\don4i\\Desktop\\Go\\Exif\\"
	for _, filename := range filenames {
		file, height := RetriveHeight(openPhoto(string(Path + filename)))
		file.Close()
		heights = append(heights, height)
	}
	return
}
