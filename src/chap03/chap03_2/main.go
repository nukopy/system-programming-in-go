package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// s3_2_1()
	s3_2_2()
}

func s3_2_1() {
	input := "read.txt"
	reader, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	// すべて読み込む
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buffer))

	// 決まったバイト数のみ読み込む
	reader, err = os.Open(input)
	if err != nil {
		panic(err)
	}
	// buffer = make([]byte, 4) 指定したバッファサイズ分読み込める
	buffer = make([]byte, 14) // 指定したバッファサイズ分読み込める
	// buffer = make([]byte, 15) // 指定したバッファサイズ分読み込めないのでエラー（read.txt は 14 byte のテキストファイル）
	size, err := io.ReadFull(reader, buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buffer), size)
}

func s3_2_2() {
	input := "read.txt"
	reader, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
}