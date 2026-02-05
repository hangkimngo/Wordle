package store

import (
	"encoding/csv"
	"fmt"
	"os"
)

func StatFile() [][]string {
	dbFile := "stats.csv"
	file, err := os.OpenFile(dbFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(dbFile)
			if err != nil {
				fmt.Println("Error creating file:", err)
			}
		} else {
			fmt.Println("Error opening file:", err)
		}
	}
	defer file.Close()
	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}
	return records
}
