package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	targetDir := currentDir + "/section_3/q2/"
	
	file, err := os.Create(targetDir + "output.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	io.CopyN(file, rand.Reader, 1024)	
}
