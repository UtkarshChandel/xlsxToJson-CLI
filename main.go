package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var inputFile string

	for {
		fmt.Print("Enter the path of the input Excel file: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if fileExists(input) && filepath.Ext(input) == ".xlsx" {
			inputFile = input
			break
		} else {
			fmt.Println("Invalid file path or the file is not an .xlsx file. Please try again.")
		}
	}

	fmt.Print("Enter the path for the output JSON file: ")
	outputFile, _ := reader.ReadString('\n')
	outputFile = strings.TrimSpace(outputFile)
	f, err := excelize.OpenFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	rows := f.GetRows(f.GetSheetName(1))

	var dataJsonArray []map[string]interface{}

	// Assuming the first row contains headers
	headers := rows[0] // First row is headers

	// Start from the second row
	for _, row := range rows[1:] {
		item := make(map[string]interface{})
		for j, colCell := range row {
			// Use the header for the key, converting to upper case
			headerKey := strings.ToUpper(headers[j])
			item[headerKey] = colCell
		}
		dataJsonArray = append(dataJsonArray, item)
	}

	jsonData, err := json.MarshalIndent(dataJsonArray, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON file written successfully to", outputFile)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return !os.IsNotExist(err) && !info.IsDir()
}
