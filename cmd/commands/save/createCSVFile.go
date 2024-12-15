package save

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"crypto-mine-cli/cmd/commands"
)

func createCSVFile() (*os.File, *csv.Writer, error) {
	downloadFolderPath, err := getDownloadFolderPath()
	if err != nil {
		return nil, nil, err
	}

	filePath := filepath.Join(downloadFolderPath, commands.CSVFileName)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot create file %q: %s\n", commands.CSVFileName, err)
	}

	writer := csv.NewWriter(file)

	return file, writer, nil
}
