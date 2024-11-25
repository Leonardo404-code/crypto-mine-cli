package save

import (
	"fmt"
	"os"
	"path/filepath"

	"crypto-mine-cli/cmd/commands"
)

func createJSONFile() (*os.File, error) {
	downloadFolderPath, err := getDownloadFolderPath()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(downloadFolderPath, commands.JSONFileName)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot create file %q: %s\n", commands.JSONFileName, err)
	}

	return file, nil
}
