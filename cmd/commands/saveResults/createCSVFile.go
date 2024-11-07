package saveResults

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func createCSVFile() (*os.File, *csv.Writer, error) {
	downloadFolderPath, err := getDownloadFolderPath()
	if err != nil {
		return nil, nil, err
	}

	filePath := filepath.Join(downloadFolderPath, FileName)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot create file %q: %s\n", FileName, err)
	}

	writer := csv.NewWriter(file)

	return file, writer, nil
}
