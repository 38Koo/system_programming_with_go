package main

import (
	"io"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	targetDir := currentDir + "/section_3/q1/"

	file, err := os.Open(targetDir + "file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create(targetDir + "copyFile.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	io.Copy(newFile, file)
}