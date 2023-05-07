package main

import (
	"encoding/csv"
	"log"
	"os"

	"web-scraping-colly/scraping"
)

func main() {
	fileName := "cryptocoin.csv"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fileName, err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	scraping.Scraping(writer)

	log.Printf("Scraping finished, check file %s for results", fileName)
}
