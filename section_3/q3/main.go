package main

import (
	"archive/zip"
	"os"
	"time"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	
	targetDir := currentDir + "/section_3/q3/"

	file, err := os.Create(targetDir + "output.zip")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	header := &zip.FileHeader{
		Name: "output.txt",
		Method: zip.Deflate,
	}
	header.Modified = time.Now()
	header.SetMode(0644)

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		panic(err)
	}

	content := "Hello, World!"
	_, err = writer.Write([]byte(content))
	if err != nil {
		panic(err)
	}

}
