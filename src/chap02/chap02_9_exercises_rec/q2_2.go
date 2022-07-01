package main

import (
	"encoding/csv"
	"log"
	"os"
)

func q2_2() {
	os.Stdout.Write([]byte("===== q2.2 =====\n"))

	// file object
	f, err := os.Create("q2_2_test.csv")
	if err != nil {
		log.Fatal("File creation error.")
	}
	defer f.Close()

	// wrap f with csv writer
	csvWriter := csv.NewWriter(f)

	header := []string{"id", "name", "yo"}
	record1 := []string{"001", "John", "19"}
	record2_3 := [][]string{
		{"002", "Mary", "22"},
		{"003", "Bob", "50"},
	}
	csvWriter.Write(header)
	csvWriter.Write(record1)
	csvWriter.WriteAll(record2_3)

	csvWriter.Flush()

	os.Stdout.Write([]byte("Done!\n"))
}