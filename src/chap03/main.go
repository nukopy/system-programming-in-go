package main

import (
	"bytes"
	"fmt"
)

func main()  {
	testBufferWriteRead()
}

func testBufferWriteRead() {
	var buf1 bytes.Buffer
	buf2 := make([]byte, 4)

	// by1 を buf1 へ書き込む
	by1 := []byte("0123456789")
	fmt.Printf("by1  : %s\n", string(by1))
	fmt.Printf("buf1 : %s\n", buf1.String())
	fmt.Printf("Write by1 to buf1...\n")
	n, err := buf1.Write(by1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes were written to buf1.\n", n)
	fmt.Printf("by1  : %s\n", string(by1)) // Write は by1 には影響しない
	fmt.Printf("buf1 : %s\n", buf1.String())
	fmt.Println()

	// buf1 から buf2 へ読み込む
	fmt.Printf("buf1 : %s\n", buf1.String())
	fmt.Printf("buf2 : %s\n", string(buf2))
	fmt.Printf("Read from buf1 to buf2...\n")
	n2, err2 := buf1.Read(buf2) // buf1 から buf2 へ文字列を読み込む
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("%d bytes were read from buf1 to buf2.\n", n2)
	fmt.Printf("buf1 : %s\n", buf1.String())
	fmt.Printf("buf2 : %s\n", string(buf2))
	fmt.Println()

	// buf1 から by2 へ読み込む
	by2 := []byte("XXXXXXXXXX")
	fmt.Printf("buf1 : %s\n", buf1.String())
	fmt.Printf("by2   : %s\n", string(by2))
	fmt.Printf("Read buf1 to by2...\n")
	n3, err3 := buf1.Read(by2)
	if err3 != nil {
		panic(err3)
	}
	fmt.Printf("%d bytes were read from buf1 to by2.\n", n3)
	fmt.Printf("buf1 : %s\n", buf1.String()) // buf1 が保持しているデータが全て by2 に移されているため、buf1 は空になる
	fmt.Printf("by2   : %s\n", string(by2))
}

/* 出力
buf1 : 0123456789
buf2 :
Read buf1 to buf2...
buf1 : 456789
buf2 : 0123

buf1 : 456789
by2   : XXXXXXXXXX
Read buf1 to by2...
buf1 :
by2   : 456789XXXX
*/