package sort

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFilesInDir() ([]string, error) {
	dir, err := os.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	var files []string
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		files = append(files, file.Name())
	}

	return files, nil
}

func MoveFilesToDirs(files []string) (int, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return 0, fmt.Errorf("error getting current working directory: %w", err)
	}

	numMoved := 0
	for _, filename := range files {
		folderName := GetFolderName(filename)
		if folderName == "" {
			continue
		}
		sourcePath := filepath.Join(cwd, filename)

		destDir := filepath.Join(cwd, folderName)
		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			if err := os.Mkdir(destDir, 0755); err != nil {
				return numMoved, fmt.Errorf("error creating directory %s: %w", destDir, err)
			}
		}
		destPath := filepath.Join(cwd, folderName, filename)

		if err := os.Rename(sourcePath, destPath); err != nil {
			return numMoved, fmt.Errorf("error moving file %s to %s: %w", filename, destPath, err)
		}
		numMoved++
	}
	return numMoved, nil
}
