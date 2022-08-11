package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// s3_5_1()
	// s3_5_2_endian()
	s3_5_3_analysis_png()
}

// io.LimitReader / io.SectionReader
func s3_5_1() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	offset := 14
	readBytes := 7
	sectionReader := io.NewSectionReader(reader, int64(offset), int64(readBytes))
	io.Copy(os.Stdout, sectionReader) // "Section"
}

func s3_5_2_endian() {
	// 10000 = 0x2710 (little endian)
	// dataLittleEndian := []byte{0x10, 0x27, 0x0, 0x0} // リトルエンディアン環境下（主流の CPU 環境下）で 10000
	dataBigEndian := []byte{0x0, 0x0, 0x27, 0x10} // ビッグエンディアン（ネットワークバイトオーダー）環境下（ネットワーク環境下）で 10000

	// エンディアンの変換
	var res1, res2 int32
	binary.Read(bytes.NewReader(dataBigEndian), binary.BigEndian, &res1)
	binary.Read(bytes.NewReader(dataBigEndian), binary.LittleEndian, &res2)
	fmt.Printf("res1: %d\nres2: %d\n", res1, res2)
}

func s3_5_3_analysis_png() {
	// --------------------------------------------------------------------------------
	// PNG ファイルの解析
	// --------------------------------------------------------------------------------

	// ファイルの読み込み
	filename := "Lenna.png"
	pngFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer pngFile.Close()

	// PNG ファイルの解析
	fmt.Printf("Analyzing .png file \"%s\"...\n", filename)
	chunks, err := analyzePngFile(pngFile)
	if err != nil {
		panic(err)
	}

	// chunk の概要を出力
	for _, chunk := range chunks {
		dumpPngChunk(chunk)
	}

	// --------------------------------------------------------------------------------
	// 既存の PNG ファイルにテキストチャンクを挿入した PNG ファイルを作成する
	// --------------------------------------------------------------------------------

	// 新規ファイルの作成
	filename = "Lenna_with_text_chunk.png"
	pngFile2, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	// テキストチャンクを挿入した PNG ファイルを作成
	text := "ASCII PROGRAMMING++"
	createPngFileWithTextChunk(chunks, pngFile2, text)
	pngFile2.Close() // 一旦閉じないとだめ。なぜかそのまま analyzePngFile しようとするとシグネチャの読み込み段階で落ちる。

	// PNG ファイルの解析
	pngFile2, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer pngFile2.Close()

	fmt.Printf("Analyzing .png file \"%s\"...\n", filename)
	chunks, err = analyzePngFile(pngFile2)
	if err != nil {
		panic(err)
	}

	// テキストチャンク入りの PNG ファイルの概要を出力
	for _, chunk := range chunks {
		dumpPngChunk(chunk)
	}
}
