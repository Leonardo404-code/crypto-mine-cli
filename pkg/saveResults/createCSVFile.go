package saveResults

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func createCSVFile() (*os.File, *csv.Writer, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, nil, fmt.Errorf("error getting home directory: %v", err)
	}

	var downloadFolder string
	if os.Getenv("OS") == "Windows_NT" {
		// Windows path
		downloadFolder = filepath.Join(homeDir, "Downloads")
	} else {
		// Linux/macOS path
		downloadFolder = filepath.Join(homeDir, "Downloads")
	}

	filePath := filepath.Join(downloadFolder, FileName)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot create file %q: %s\n", FileName, err)
	}

	writer := csv.NewWriter(file)

	return file, writer, nil
}
