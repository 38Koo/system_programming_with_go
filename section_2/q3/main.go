package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// favicon.icoのリクエストも含めて2回リクエストする場合の処理
	// if r.URL.Path == "/favicon.ico" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	jsonData, err := json.Marshal(source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gzipWriter := gzip.NewWriter(w)
	gzipWriter.Header.Name = "data.json"
	
	writer := io.MultiWriter(gzipWriter, os.Stdout)

	io.WriteString(writer, string(jsonData))
	gzipWriter.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
}
