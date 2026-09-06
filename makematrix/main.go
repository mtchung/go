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

type colRowKey struct {
	User  string
	Group string
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
	fmt.Printf("DEBUG Read from to: %s row:%d col:%d del:%s\n", *inputFile, *row, *col, *del)

	// 1. Open the delimited file
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// 2. Initialize the CSV reader
	reader := csv.NewReader(file)

	// 3. Define your custom delimiter (e.g., Use '\t' for TSV, ';' for semicolon)
	// Convert the first character of the string to a rune
	if len(*del) > 0 {
		reader.Comma = []rune(*del)[0]
	} else {
		reader.Comma = ',' // Fallback default if string is empty
	}

	// 4. Loop through the file line-by-line
	colRow := make(map[colRowKey]string)
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

		user := record[*row]
		group := record[*col]

		uniqueCol[user] = "Y"
		uniqueRow[group] = "Y"
		colRow[colRowKey{User: user, Group: group}] = "Y"
	}

	sortedGroups := make([]string, 0, len(uniqueRow))

	for g := range uniqueRow {
		sortedGroups = append(sortedGroups, g)
	}

	sortedUsers := make([]string, 0, len(uniqueCol))

	for u := range uniqueCol {
		sortedUsers = append(sortedUsers, u)
	}

	slices.Sort(sortedUsers)
	slices.Sort(sortedGroups)

	fmt.Printf("Name\t")
	for _, groupName := range sortedGroups {
		fmt.Printf("%s\t", groupName)
	}
	fmt.Printf("\n")

	for _, userVal := range sortedUsers {

		fmt.Printf("%s\t", userVal)
		for _, groupVal := range sortedGroups {
			keyToFind := colRowKey{User: userVal, Group: groupVal}

			_, exists := colRow[keyToFind]
			if exists {
				fmt.Printf("%s\t", "Y")
			} else {
				fmt.Printf("%s\t", ".")
			}
		}
		fmt.Printf("\n")
	}

}
