package service

import (
	"fmt"
	"os"
	"path/filepath"
)

type ContentService struct {
	DirectoryPath string
}

func NewContentService(dir string) ContentService {
	return ContentService{DirectoryPath: dir}
}

func (s ContentService) GetContentFilePaths() ([]string, error) {
	root := s.DirectoryPath
	var fileList []string = []string{}
	if _, err := os.Stat(s.DirectoryPath); os.IsNotExist(err) {
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

	fmt.Println(fileList)

	return fileList, nil
}
