package save

import (
	"log"
)

func Save(fileType string) {
	if fileType != "json" && fileType != "csv" {
		log.Fatal("error: file type not supported, try json or csv")
		return
	}

	if fileType == "json" {
		saveInJSON()
	}

	if fileType == "csv" {
		saveInCSV()
	}

	log.Printf("File persisted in the Downloads folder in %s format", fileType)
}
