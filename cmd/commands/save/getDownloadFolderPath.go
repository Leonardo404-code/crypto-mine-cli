package save

import (
	"fmt"
	"os"
	"path/filepath"
)

func getDownloadFolderPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %v", err)
	}

	var downloadFolder string
	if os.Getenv("OS") == "Windows_NT" {
		downloadFolder = filepath.Join(homeDir, "Downloads")
	} else {
		downloadFolder = filepath.Join(homeDir, "Downloads")
	}

	return downloadFolder, nil
}
