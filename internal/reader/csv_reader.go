package reader

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading all records:", err)
		return nil, err
	}
	for _, record := range records {
		fmt.Println(record)

	}

	return records, nil
}
