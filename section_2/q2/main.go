package main

import (
	"encoding/csv"
	"os"
)

func main() {
	records := [][]string{
		{"1", "A", "B", "C"},
		{"2", "D", "E", "F"},
		{"3", "G", "H", "I"},
	}

	// file, err := os.Create("section_2/q2/test.csv")
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// writer := csv.NewWriter(file)
	writer := csv.NewWriter(os.Stdout)
	writer.WriteAll(records)
	writer.Flush()
}
