package main

import (
	"io"
	"os"
)

func copyNClone(dest io.Writer, src io.Reader, length int) {
	// srcを任意のlengthで区切る
	size := io.LimitReader(src, int64(length))

	io.Copy(dest, size)
}


func main() {
	copyNClone(os.Stdout, os.Stdin, 5)
}
