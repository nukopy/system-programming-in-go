package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// s3_4_1() // `go run main.go < main.go` で 5 byte ごとにこのソースコード自体が出力される
	// s3_4_2()
	// s3_4_3_rawHttpResponse()
	// s3_4_3_parseHttpResponse()
	s3_4_4_buffer()
}

func open(filename string) (io.ReadCloser, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return r, nil
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
	// open file
	filename := "main.go"
	src, err := open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// define buf
	var buf bytes.Buffer
	// dst := io.MultiWriter(os.Stdin, &buf)
	dst := &buf

	// dst に src の中身をすべて書き込む
	fmt.Println("===== io.Copy =====")
	writeSize, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("size=%d, buffer=\n%s\n", writeSize, buf.String())

	// dst に srcFile の中身を 10 byte だけコピーする
	/*
	fmt.Println("===== io.CopyN =====")
	var size int64 = 10 // byte
	writeSize, err = io.CopyN(dst, src, size)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("size=%d, buffer=\n%s\n", writeSize, buf.String())
	*/
}

func s3_4_3_rawHttpResponse() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	httpRequest := "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"
	conn.Write([]byte(httpRequest))
	io.Copy(os.Stdout, conn)
}

func s3_4_3_parseHttpResponse() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	// http request
	httpRequest := "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"
	conn.Write([]byte(httpRequest))

	// parse HTTP response
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// output
	// fmt.Println(res.Proto, res.Status)
	// fmt.Println(res.Header)
	io.Copy(os.Stdout, res.Body)
}

func s3_4_4_buffer() {
	// メモリに蓄えた内容を io.Reader として読み出すバッファ

	// 空のバッファ
	// CAUTION: これはポインタでなく実体として初期化される
	var buffer1 bytes.Buffer

	// バイト列で初期化
	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})

	// 文字列で初期化
	buffer3 := bytes.NewBufferString("初期文字列")

	fmt.Println(buffer1)
	fmt.Println(buffer2)
	fmt.Println(buffer3)
}