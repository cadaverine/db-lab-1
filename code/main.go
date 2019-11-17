package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func writeToCSV(path string, values [][]string) {
	file, err := os.Create("../test-data/institutions.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.WriteAll(values)
	writer.Flush()
}

func main() {
	var institutionsNum int

	fmt.Print("Type institutions num: ")
	fmt.Scanln(&institutionsNum)

	institutionsData := generateInstitutionsData(institutionsNum)

	writeToCSV("../test-data/institutions.csv", institutionsData)
}
