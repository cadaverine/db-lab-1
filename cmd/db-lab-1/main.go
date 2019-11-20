package main

import (
	"encoding/csv"
	"fmt"
	g "github.com/cadaverine/db-lab-1/db-generators"
	"log"
	"os"
	"sync"
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
	var institutionsNum, studentsNum, teachersNum, coursesNum int

	fmt.Print("Type institutions num: ")
	fmt.Scanln(&institutionsNum)

	fmt.Print("Type students num: ")
	fmt.Scanln(&studentsNum)

	fmt.Print("Type teachers num: ")
	fmt.Scanln(&teachersNum)

	fmt.Print("Type courses num: ")
	fmt.Scanln(&coursesNum)

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		institutionsData := g.GenerateInstitutionsData(institutionsNum)
		writeToCSV("test-data/institutions.csv", institutionsData)
		wg.Done()
	}()

	go func() {
		studentsData := g.GenerateStudentsData(studentsNum)
		writeToCSV("test-data/students.csv", studentsData)
		wg.Done()
	}()

	go func() {
		teachersData := g.GenerateTeachersData(teachersNum)
		writeToCSV("test-data/teachers.csv", teachersData)
		wg.Done()
	}()

	wg.Wait()

	coursesData := g.GenerateCoursesData(coursesNum, teachersNum, institutionsNum)
	writeToCSV("test-data/courses.csv", coursesData)
}
