package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

const pngSign = "\x89PNG\r\n\x1a\n"

type PngChunk struct {
	Size uint32 // データ長 (4 byte)
	Type []byte // チャンクタイプ (4 byte)
	Data []byte // データ本体 (可変長)
	CRC uint32 // CRC (4 byte)
}

func (p *PngChunk) InitDataBySize(size uint32) {
	p.Data = make([]byte, size)
}

func (p *PngChunk) SetSize(s uint32) {
	p.Size = s
}

func (p *PngChunk) SetType(t []byte) {
	p.Type = t
}

func (p *PngChunk) SetData(d []byte) {
	p.Data = d
}

func (p *PngChunk) SetCRC(c uint32) {
	p.CRC = c
}

func NewPngChunk() *PngChunk {
	return &PngChunk{
		Size: 0,
		Type: make([]byte, 4),
		Data: []byte{},
		CRC: 0,
	}
}

func analyzePngFile(pngFile *os.File) ([]*PngChunk, error) {
	// PNG シグネチャの読み込み
	lengthPngSign := 8
	buffer := make([]byte, lengthPngSign)

	pngFile.Read(buffer) // 8 byte 分読み込む
	if !bytes.Equal([]byte(pngSign), buffer) {
		return nil, fmt.Errorf("this file format is not .png")
	}

	// chunk の分割
	// PNG のチャンク = データ長（4 byte） + 種類（4 byte） + データ（xxx byte） + CRC（4 byte）
	var chunks []*PngChunk
	var rChunks []io.Reader
	var offset int64 = int64(lengthPngSign)
	for {
		chunk := NewPngChunk()

		// Size（4 byte 分読み込む）
		err := binary.Read(pngFile, binary.BigEndian, &chunk.Size) // file(io.Reader) から 4 byte 分読み込んで int32 の変数に格納する
		if err == io.EOF {
			break
		}

		// Type（4 byte 分読み込む）
		err = binary.Read(pngFile, binary.BigEndian, &chunk.Type)
		if err == io.EOF {
			break
		}

		// Data（chunk.Size byte 分読み込む）
		chunk.InitDataBySize(chunk.Size)
		err = binary.Read(pngFile, binary.BigEndian, &chunk.Data)
		if err == io.EOF {
			break
		}

		// CRC（4 byte 分読み込む）
		err = binary.Read(pngFile, binary.BigEndian, &chunk.CRC)
		if err == io.EOF {
			break
		}

		// append
		chunks = append(chunks, chunk)
		chunkSize := 4 + 4 + chunk.Size + 4 // 現在読込中のチャンクサイズ
		rChunks = append(rChunks, io.NewSectionReader(pngFile, offset, int64(chunkSize)))

		offset += int64(chunkSize) // Seek 自体は既に binary.Read で進んでいるため offset の更新のみで良い。file.Seek は不要。
	}
	fmt.Printf("%d chunks gotten!\n", len(chunks))

	return chunks, nil
}

func dumpPngChunk(chunk *PngChunk) {
	// テキストチャンクのときのみデータの中身を出力
	if bytes.Equal([]byte("tExt"), chunk.Type) {
		fmt.Printf("chunk '%v' (%d bytes) %s\n", string(chunk.Type), chunk.Size, string(chunk.Data))
		return
	}
	fmt.Printf("chunk '%v' (%d bytes)\n", string(chunk.Type), chunk.Size)
}

func createTextChunk(text string) *PngChunk {
	textChunk := NewPngChunk()
	byteData := []byte(text)

	// Size の計算 & セット
	size := uint32(len(byteData))
	textChunk.SetSize(size)

	// Type のセット
	chunkType := "tExt"
	textChunk.SetType([]byte(chunkType))

	// Data のセット
	textChunk.SetData(byteData)

	// CRC の計算 & セット
	crc := crc32.NewIEEE()
	crc.Write([]byte(chunkType))
	crc.Write(byteData)
	textChunk.SetCRC(crc.Sum32())

	return textChunk
}

func createPngFileFromChunks(chunks []*PngChunk, pngFile io.Writer) {
	_, err := pngFile.Write([]byte(pngSign))
	if err != nil {
		panic(err)
	}

	// チャンクの書き込み
	for _, chunk := range chunks {
		// Size
		binary.Write(pngFile, binary.BigEndian, chunk.Size)

		// Type
		binary.Write(pngFile, binary.BigEndian, chunk.Type)

		// Data
		binary.Write(pngFile, binary.BigEndian, chunk.Data)

		// CRC
		binary.Write(pngFile, binary.BigEndian, chunk.CRC)
	}
}

func createPngFileWithTextChunk(chunks []*PngChunk, pngFile io.Writer, text string) {
	// テキストチャンクの作成
	textChunk := createTextChunk(text)

	// テキストチャンクの挿入: IHDR チャンクの直後にテキストチャンク（tExt チャンク）を挿入する
	insertIdx := 1
	chunks = append(chunks[:insertIdx+1], chunks[insertIdx:]...)
	chunks[insertIdx] = textChunk

	// テキストチャンク入りの PNG ファイルを作成
	fmt.Println("Writing to new PNG file...")

	// PNG シグネチャの書き込み
	createPngFileFromChunks(chunks, pngFile)
}