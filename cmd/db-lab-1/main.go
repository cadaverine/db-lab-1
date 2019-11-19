package main

import (
	"encoding/csv"
	"fmt"
	g "github.com/cadaverine/db-lab-1/db-generators"
	"log"
	"os"
)

func writeToCSV(path string, values [][]string) {
	file, err := os.Create(path)
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

	institutionsData := g.GenerateInstitutionsData(institutionsNum)

	writeToCSV("test-data/institutions.csv", institutionsData)
}
