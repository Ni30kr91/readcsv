package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	input := "prince"
	records := readCsvFile("./read.csv")
	city := getcity(records, input)
	fmt.Println(city)
}

func getcity(records [][]string, input string) string {
	var city string
	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			if records[i][j] == input {
				//	fmt.Println(records[i][j+1])
				city = records[i][j+1]
				break
			}
		}
	}
	return city
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
