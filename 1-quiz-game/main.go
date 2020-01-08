package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadCsvRecords(filename string) [][]string {
	// try to open the file, and defer it's closure
	// https://golang.org/pkg/os
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create a csv reader pointer (*Reader) and configure it
	// to expect two fields per csv line (record)
	// https://golang.org/pkg/encoding/csv
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = 2

	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

func main() {
	filename := "problems.csv"

	records := loadCsvRecords(filename)

	fmt.Println(records)
}
