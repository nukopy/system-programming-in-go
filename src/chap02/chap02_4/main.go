package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

/* interface: io.Writer
Implementations must not retain p.

type Writer interface {
	Write(p []byte) (n int, err error)
}
*/

func main() {
	// 2.4.1 ファイル出力
	fileOutput()

	// 2.4.2 画面出力
	screenOutput()

	// 2.4.3 書かれた内容を記憶しておくバッファ（1）：bytes.Buffer
	bufferOutput1()

	// 2.4.4 書かれた内容を記憶しておくバッファ（2）：strings.Builder
	bufferOutput2()

	// 2.4.5 インターネットアクセスの送信
	internetAccess()
	httpServer()

	// 2.4.6 io.Writer のデコレータ
	gzipSample()
}

/* インタフェース io.Writer を満たす構造体の例
- os.File
- os.Stdout
- bytes.Buffer
- strings.Builder

*/

func fileOutput() {
	filename := "test.txt"
	file, err := os.Create(filename) // file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	n, err := file.Write([]byte("os.File example\n"))
	if err != nil {
		panic(err)
	}
	file.Close()

	fmt.Printf("%d bytes were written to file.\n", n) // 制御文字含めて 16 bytes
}

func screenOutput() {
	n, err := os.Stdout.Write([]byte("os.Stdout example\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d bytes were written to stdout.\n", n) // 制御文字含めて 18 bytes
}

func bufferOutput1() {
	var buffer bytes.Buffer
	n, err := buffer.Write([]byte("bytes.Buffer example\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", buffer.String())

	fmt.Printf("%d bytes were written to bytes.Buffer.\n", n) // 制御文字含めて 21 bytes
}

func bufferOutput2() {
	var builder strings.Builder
	n, err := builder.Write([]byte("strings.Builder example\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", builder.String())

	fmt.Printf("%d bytes were written to strings.Builder.\n", n) // 制御文字含めて 24 bytes
}

func internetAccess() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	httpRequest := "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"
	fmt.Printf("HTTP Request: \n%s", httpRequest)
	io.WriteString(conn, httpRequest)
	io.Copy(os.Stdout, conn)
}

func httpServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, http.ResponseWriter sample!")
}

func gzipSample() {
	filename := "test.txt.gz"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	writer := gzip.NewWriter(file)
	// writer.Header.Name = "test.txt" // なくても影響しないのだけど何をしている？
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}