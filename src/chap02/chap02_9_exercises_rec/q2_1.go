package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func q2_1() {
	os.Stdout.Write([]byte("===== q2.1 =====\n"))

	// file object
	f, err := os.Create("q2_1_test.txt")
	if err != nil {
		log.Fatal("File creation error.")
	}
	defer f.Close()

	writer := io.MultiWriter(os.Stdout, f)

	// write to "test.txt"
	strValue := "Golang"
	intValue := 1
	floatValue := 1.19
	fmt.Fprintf(writer, "string : %s\n", strValue)
	fmt.Fprintf(writer, "int    : %d\n", intValue)
	fmt.Fprintf(writer, "float  : %f\n", floatValue)
	fmt.Fprintf(writer, "float  : %.f\n", floatValue)
	fmt.Fprintf(writer, "float  : %.2f\n", floatValue)
	fmt.Fprintf(writer, "float  : %.10f\n", floatValue)
	fmt.Println()
}
