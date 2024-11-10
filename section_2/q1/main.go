package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var file *os.File
	var err error

  dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory")
		panic(err)
	}
	targetDir := dir + "/section_2/q1/"
	targetFile := targetDir + "output.txt"
	_, err = os.Stat(targetFile)
	if os.IsNotExist(err) {
		file, err = os.Create(targetFile)
		if err != nil {
			fmt.Println("Error creating file")
			panic(err)
		}
	} else {
		file, err = os.OpenFile(targetFile, os.O_RDWR | os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Error opening file")
			panic(err)
		}
	}
	
	defer file.Close()

	fmt.Fprintln(file, "Hello, World!")
	fmt.Fprintf(file, "it's %v \n", time.Now())
	fmt.Fprintf(file, "%d \n", 128)
	fmt.Fprintf(file, "%f \n", 1.14)
}
