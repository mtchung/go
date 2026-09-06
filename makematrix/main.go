package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"

	"github.com/pborman/getopt/v2"
)

type rowColKey struct {
	Row string
	Col string
}

func main() {
	// 1. Declare flags. Getopt returns pointers to the parsed values.
	// Arguments: (name, short character, default value, description)
	helpFlag := getopt.BoolLong("help", 'h', "Display help menu")
	versionFlag := getopt.BoolLong("version", 'V', "Display version info")

	// Example of a flag that takes a string argument
	inputFile := getopt.StringLong("inputCSV", 'i', "input.csv", "Path to the output file")
	del := getopt.StringLong("delimiter", 'd', ",", "Delimiter Value Default to comma(,)")
	row := getopt.IntLong("row", 'r', 0, "Column Field1 (row)")
	col := getopt.IntLong("column", 'c', 1, "Column Field2 (column)")

	// 2. Parse the command-line arguments
	getopt.Parse()

	// 3. Handle help menu manually if requested
	if *helpFlag {
		getopt.Usage()
		return
	}

	if *versionFlag {
		fmt.Println("App Version: 0.0.1")
		return
	}

	// 4. Access parsed values
	//fmt.Printf("DEBUG Read from to: %s row:%d col:%d del:%s\n", *inputFile, *row, *col, *del)

	// 5. Open the delimited file
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// 6. Initialize the CSV reader
	reader := csv.NewReader(file)

	// 7. Define your custom delimiter (e.g., Use '\t' for TSV, ';' for semicolon)
	// Convert the first character of the string to a rune
	if len(*del) > 0 {
		reader.Comma = []rune(*del)[0]
	} else {
		reader.Comma = ',' // Fallback default if string is empty
	}

	// 8. Loop through the file line-by-line
	rowCol := make(map[rowColKey]string)
	uniqueCol := make(map[string]string)
	uniqueRow := make(map[string]string)

	for {
		record, err := reader.Read()

		// Break out of the loop when hitting the End Of File (EOF)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("Error reading record: %s", err)
		}

		theRow := record[*row]
		theCol := record[*col]

		// Save records into row and cols hashmap
		uniqueCol[theCol] = "Y"
		uniqueRow[theRow] = "Y"
		rowCol[rowColKey{Row: theRow, Col: theCol}] = "Y"
	}

	// Make sorted column slice
	sortedCols := make([]string, 0, len(uniqueCol))

	for g := range uniqueCol {
		sortedCols = append(sortedCols, g)
	}

	// Make sorted rows slice
	sortedRows := make([]string, 0, len(uniqueRow))

	for u := range uniqueRow {
		sortedRows = append(sortedRows, u)
	}

	// Make sorted rows and columns
	slices.Sort(sortedRows)
	slices.Sort(sortedCols)

	fmt.Printf("\t")

	// Print columns
	for _, colName := range sortedCols {
		fmt.Printf("%s\t", colName)
	}
	fmt.Printf("\n")

	for _, rowVal := range sortedRows {

		// Print columns
		fmt.Printf("%s\t", rowVal)
		for _, colVal := range sortedCols {
			keyToFind := rowColKey{Row: rowVal, Col: colVal}

			_, exists := rowCol[keyToFind]
			if exists {
				fmt.Printf("%s\t", "Y")
			} else {
				fmt.Printf("%s\t", ".")
			}
		}
		fmt.Printf("\n")
	}

}
