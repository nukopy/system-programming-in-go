package main

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main()  {
	q2_1()
	q2_2()
	q2_3()
}

func q2_1() {
	var a int = 31415
	var b float64 = 3.1415
	var c string = "Hello, world!"

	fmt.Fprintf(os.Stdout, "a: %d\nb: %.3f\nc: %s\n", a, b, c)
	fmt.Fprintf(os.Stdout, "a: %d\nb: %.10f\nc: %s\n", a, b, c) // %f で表示桁数を調整すると四捨五入される
}

func q2_2() {
	filename := "test.csv"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	header := []string{
		"Id",
		"Name",
		"y/o",
	}
	records := [][]string{
		header,
		{
			"1",
			"Bob",
			strconv.Itoa(15),
		},
		{
			"2",
			"Alice",
			strconv.Itoa(13),
		},
	}
	for _, r := range records {
		writer.Write(r)
	}
	writer.Flush() // gzip のときは Flush いらなかったけどなぜ必要？
}

func q2_3() {
	http.HandleFunc("/", q2_3_handler)
	http.ListenAndServe(":8080", nil)
}

var cnt int

func q2_3_handler(w http.ResponseWriter, r *http.Request) {
	// HTTP レスポンスヘッダの設定
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	// 標準出力の Writer に gzip 圧縮前、HTTP レスポンスの Writer に gzip 圧縮後を Write する
	stdoutWriter := os.Stdout
	gzipWriter := gzip.NewWriter(w)
	multiWriter := io.MultiWriter(stdoutWriter, gzipWriter)

	// レスポンスとして返す JSON の元データ
	resJson := map[string]string{
		"Language": "Go",
		"Hello": "World",
		"My Name is": "Bob",
		"cnt": strconv.Itoa(cnt),
		"time": time.Now().Format("2006-01-02 15:04:05"),
	}

	jsonEncoder := json.NewEncoder(multiWriter)
	jsonEncoder.Encode(resJson) // ここで Writer が Write する

	gzipWriter.Flush()

	cnt++
}