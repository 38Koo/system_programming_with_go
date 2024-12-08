package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)

	zipWriter := zip.NewWriter(buf)
	defer zipWriter.Close()

	header := &zip.FileHeader{
		Name:  "sample.txt",
		Method: zip.Deflate,
	}

	header.Modified = time.Now()
	header.SetMode(0644)

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content := "A text file in a zip file"
	_, err = writer.Write([]byte(content))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = zipWriter.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=sample.zip")
	w.Header().Set("Content-Length", string((buf.Bytes())))

	_, err = w.Write(buf.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
