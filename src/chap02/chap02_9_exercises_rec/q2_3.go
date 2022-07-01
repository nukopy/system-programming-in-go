package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func q2_3() {
	os.Stdout.Write([]byte("===== q2.3 =====\n"))
	fmt.Println("HTTP server is running...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-My-Header", "Hi!")

	// create body
	body := map[string]string{
		"Hello": "Golang",
	}

	// gzip writer
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	writer := io.MultiWriter(os.Stdout, gzipWriter)
	gzipWriter.Flush()

	// json encoder
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "")
	err := encoder.Encode(body)
	if err != nil {
		panic(err)
	}

}