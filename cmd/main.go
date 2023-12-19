package main

import (
	"log"

	"web-scraping-colly/scraping"
)

func main() {
	scraping.InitScraping()
	log.Printf("Scraping finished, check file %s for results", scraping.FileName)
}
