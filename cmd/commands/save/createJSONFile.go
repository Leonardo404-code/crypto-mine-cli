package save

import (
	"fmt"
	"os"
	"path/filepath"
)

func createJSONFile() (*os.File, error) {
	downloadFolderPath, err := getDownloadFolderPath()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(downloadFolderPath, JSONFileName)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot create file %q: %s\n", JSONFileName, err)
	}

	return file, nil
}
