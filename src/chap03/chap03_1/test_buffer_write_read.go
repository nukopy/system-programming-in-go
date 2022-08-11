package main

import (
	"bytes"
	"fmt"
)

func testBufferWriteRead() {
	buf1 := []byte("0123456789")
	var buf2 bytes.Buffer // io.Writer

	// buf1 を buf2 へ書き込む
	fmt.Printf("buf1 : %s\n", string(buf1))
	fmt.Printf("buf2  : %s\n", buf2.String())
	fmt.Printf("Write byte array of buf1 to buf2...\n")
	n, err := buf2.Write(buf1) // buf2 に対して buf1 の中身を書き込む
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d bytes were written to buf2.\n", n)
	fmt.Printf("buf1  : %s\n", string(buf1)) // Write は buf1 には影響しない
	fmt.Printf("buf2 : %s\n", buf2.String())
	fmt.Println()

	// buf2 から buf3 へ読み込む
	buf3 := make([]byte, 4)
	fmt.Printf("buf2 : %s\n", buf2.String())
	fmt.Printf("buf3 : %s\n", string(buf3))
	fmt.Printf("Read byte array of buf2 to buf3...\n")

	// buf2 の中身を buf3 へ読み込む。
	// buf2 は 10 byte のバイト列、buf3 のバッファ領域（メモリ確保領域）は 4 byte なので 10 - 4 = 6 byte 分が buf2 に残り、4 byte 分が buf3 へ書き込まれる。
	n2, err2 := buf2.Read(buf3)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("%d bytes were read from buf2 to buf3.\n", n2)
	fmt.Printf("buf1 : %s\n", buf2.String())
	fmt.Printf("buf3 : %s\n", string(buf3))
	fmt.Println()

	// buf2 から buf4 へ読み込む
	// buf2 が残り保持しているバイト列は 6 byte、buf4 は 10 byte なので、buf2 のバイト列は全ては buf4 へ移される
	buf4 := []byte("XXXXXXXXXX")
	fmt.Printf("buf2 : %s\n", buf2.String())
	fmt.Printf("buf4 : %s\n", string(buf4))
	fmt.Printf("Read byte array of buf2 to buf4...\n")
	n3, err3 := buf2.Read(buf4) // buf2 の中身を buf4 へ読み込む。buf2 の中身は消費される。
	// bytes.buffer.Read は
	if err3 != nil {
		panic(err3)
	}
	fmt.Printf("%d bytes were read from buf2 to buf4.\n", n3)
	fmt.Printf("buf2 : %s\n", buf2.String())
	fmt.Printf("buf4 : %s\n", string(buf4))
	fmt.Println()

	// まとめ
	fmt.Println("Summary")
	fmt.Printf("buf1 : %s\n", string(buf1))
	fmt.Printf("buf2 : %s\n", buf2.String())
	fmt.Printf("buf3 : %s\n", string(buf3))
	fmt.Printf("buf4 : %s\n", string(buf4))
}

/* 補足: bytes.Buffer.Read(p byte[]) の中身は buffer から buffer へのコピーを行っている

Buffer.buf はプライベートメンバで、bytes.Buffer のバッファ領域の実態。

```go
type Buffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.
}

...中略

func (b *Buffer) Read(p []byte) (n int, err error) {
	b.lastRead = opInvalid
	if b.empty() {
		// Buffer is empty, reset to recover space.
		b.Reset()
		if len(p) == 0 {
			return 0, nil
		}
		return 0, io.EOF
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	if n > 0 {
		b.lastRead = opRead
	}
	return n, nil
}
```

bytes.Buffer.buf の中身を p へコピーする https://pkg.go.dev/builtin#copy

```go
n = copy(p, b.buf[b.off:])  // func copy(dst []Type, src []Type) int
```
*/