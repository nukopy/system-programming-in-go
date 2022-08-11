package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

var p [1]byte

func readByte(r io.Reader) (rune, error) {
	n, err := r.Read(p[:])
	if n > 0 {
		return rune(p[0]), nil
	}

	return 0, err
}

func wordCount(r io.Reader) (int, error) {
	words := 0
	inword := false

	// word count
	for {
		rn, err := readByte(r)

		if unicode.IsSpace(rn) { // "hello world"
			if inword {
				words++
			}
			inword = false
		} else {
			inword = true
		}

		if err == io.EOF {
			log.Printf("Reached EOF\n")
			return words, nil
		}

		if err != nil {
			return -1, err
		}
	}
}

func main() {
	// filename := os.Args[1]
	filename := "MobyDickTextBook.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot open file %q: %v\n", filename, err)
	}
	defer f.Close()

	words, err := wordCount(f)
	if err != nil {
		log.Fatalf("read failed: %v\n", err)
	}
	fmt.Printf("%q: %d words\n", filename, words)
}