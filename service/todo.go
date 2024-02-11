package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"jeffcaldwell.is/model"
)

type TodoService struct {
	filePath string
}

func NewTodoService(directoryPath string) TodoService {
	return TodoService{
		filePath: "./content/todo/todo.md",
	}
}

func (s TodoService) GetTodos() (model.TodoPageData, error) {
	fileService := FileService{}
	markdownService := MarkdownService{}

	pageData := model.TodoPageData{}
	fileText, err := fileService.GetFileText(s.filePath)
	if err != nil {
		return pageData, fmt.Errorf("error getting file content: %v", err)
	}

	var matter struct {
		Title   string `yaml:"title"`
		Updated string `yaml:"updated"`
	}

	rest, err := frontmatter.Parse(strings.NewReader(fileText), &matter)

	if err != nil {
		return pageData, fmt.Errorf("error parsing frontmatter: %v", err)
	}

	loc, _ := time.LoadLocation("US/Central")

	updated, err := time.ParseInLocation(time.RFC3339, matter.Updated, loc)

	if err != nil {
		return pageData, fmt.Errorf("error parsing time: %v", err)
	}

	htmlContent, err := markdownService.ParseMarkdownContent(rest)
	if err != nil {
		return pageData, fmt.Errorf("error parsing markdown: %v", err)
	}

	pageData.Title = matter.Title
	pageData.Updated = updated
	pageData.PageContent = htmlContent

	return pageData, nil
}
