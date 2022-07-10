package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// testBufferWriteRead()

	// s3_4_1() // `go run main.go < main.go` で 5 byte ごとにこのソースコード自体が出力される
	s3_4_2()
}

func s3_4_1() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer) // read stdin to buffer
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("size=%d, input='%s'\n", size, string(buffer))
	}
}

func s3_4_2() {
	srcFile, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	var buffer bytes.Buffer
	dst := io.MultiWriter(os.Stdin, &buffer)

	// dst に srcFile の中身をすべてコピーする
	fmt.Println("===== io.Copy =====")
	writeSize, err := io.Copy(dst, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("size=%d, buffer=\n%s\n", writeSize, buffer.String())

	// dst に srcFile の中身を 10 byte だけコピーする
	fmt.Println("===== io.CopyN =====")
	var size int64 = 10 // byte
	writeSize, err = io.CopyN(dst, srcFile, size)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("size=%d, buffer=\n%s\n", writeSize, buffer.String())
}
