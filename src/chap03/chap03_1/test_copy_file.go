package main

import (
	"fmt"
	"io"
	"os"
)

func testCopyFile() {
	// reader, _ := os.Open(os.Args[1])
	input := "read.txt"
	output := "output.txt"
	reader, _ := os.Open(input)
	writer, _ := os.OpenFile(output, os.O_WRONLY | os.O_CREATE, 0644)
	copy(reader, writer)
}

func copy(src io.Reader, dst io.Writer) {
	// バッファの定義
	buf := make([]byte, 5)

	for {
		n, _ := src.Read(buf) // reader -> buf へ読み込み
		fmt.Println(string(buf[:n]), n)
		if n == 0 {
			break
		}
		dst.Write(buf[:n]) // buf -> writer へ書き込み
	}
}