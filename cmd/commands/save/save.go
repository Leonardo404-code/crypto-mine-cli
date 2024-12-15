package save

import (
	"errors"
	"log"
)

func Save(fileType string) error {
	if fileType != "json" && fileType != "csv" {
		return errors.New("file type not supported, try json or csv")
	}

	if fileType == "json" {
		if err := saveInJSON(); err != nil {
			return err
		}
	}

	if fileType == "csv" {
		if err := saveInCSV(); err != nil {
			return err
		}
	}

	log.Printf("File persisted in the Downloads folder in %s format", fileType)

	return nil
}
