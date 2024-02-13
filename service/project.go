package service

import (
	"errors"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"jeffcaldwell.is/model"
)

type ProjectService struct {
	DirectoryPath string
	IsDev         bool
}

func NewProjectService(contentDir string, isDev bool) ProjectService {
	return ProjectService{DirectoryPath: contentDir, IsDev: isDev}
}

func (s ProjectService) GetAllProjects() ([]*model.Project, error) {
	fileService := FileService{}
	markdownService := MarkdownService{}
	projects := []*model.Project{}

	contentPaths, err := fileService.GetDirectoryFilePaths(s.DirectoryPath)

	if err != nil {
		return projects, err
	}

	for _, path := range contentPaths {
		slug := strings.TrimSuffix(path, filepath.Ext(path))
		fileText, err := fileService.GetFileText(path)
		if err != nil {
			return projects, err
		}

		frontmatterMap, err := markdownService.GetFrontmatter([]byte(fileText))
		if err != nil {
			return projects, err
		}

		frontmatter := model.ProjectFrontmatter{}
		projectImage := model.Image{}

		mapstructure.Decode(frontmatterMap, &frontmatter)
		mapstructure.Decode(frontmatter.Image, &projectImage)

		htmlContent, err := markdownService.ParseMarkdownContent([]byte(fileText))
		if err != nil {
			return projects, err
		}

		loc, _ := time.LoadLocation("US/Central")
		startDate, err := time.ParseInLocation(time.RFC3339, frontmatter.StartDate, loc)
		if err != nil {
			return projects, err
		}

		endDate, err := time.ParseInLocation(time.RFC3339, frontmatter.EndDate, loc)
		if err != nil {
			endDate = startDate
		}

		project := model.Project{
			Name:          frontmatter.Name,
			Link:          frontmatter.Link,
			Summary:       frontmatter.Summary,
			StartDate:     startDate,
			EndDate:       endDate,
			FeaturedImage: &projectImage,
			Content:       htmlContent,
			Slug:          slug,
		}

		projects = append(projects, &project)
	}

	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].EndDate.After(projects[j].EndDate)
	})

	return projects, nil
}

func (s ProjectService) GetProjectBySlug(slug string) (*model.Project, error) {
	return &model.Project{}, errors.New("not implemented")
}
