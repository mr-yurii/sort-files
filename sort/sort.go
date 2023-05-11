package sort

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func getDirPath() (string, error) {
	scriptPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("error getting directory path: %w", err)
	}
	dirPath := path.Dir(scriptPath)
	return dirPath, nil
}

func GetFilesInDir() ([]string, error) {
	dirPath, err := getDirPath()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	dir, err := os.ReadDir(dirPath)
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
	dirPath, err := getDirPath()
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	numMoved := 0
	for _, filename := range files {
		folderName := GetFolderName(filename)
		if folderName == "" {
			continue
		}
		sourcePath := filepath.Join(dirPath, filename)

		destDir := filepath.Join(dirPath, folderName)
		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			if err := os.Mkdir(destDir, 0755); err != nil {
				return numMoved, fmt.Errorf("error creating directory %s: %w", destDir, err)
			}
		}
		destPath := filepath.Join(dirPath, folderName, filename)

		if err := os.Rename(sourcePath, destPath); err != nil {
			return numMoved, fmt.Errorf("error moving file %s to %s: %w", filename, destPath, err)
		}
		numMoved++
	}
	return numMoved, nil
}
