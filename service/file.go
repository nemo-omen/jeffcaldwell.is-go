package service

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileService struct{}

func (s FileService) GetFileText(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}

	return string(data), nil
}

func (s FileService) GetDirectoryFilePaths(dirPath string) ([]string, error) {
	root := dirPath
	var fileList []string = []string{}
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil, err
	}

	err := filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileList = append(fileList, file)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}
