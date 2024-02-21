package service

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"jeffcaldwell.is/model"
	"jeffcaldwell.is/util"
)

type ProjectService struct {
	DirectoryPath string
	IsDev         bool
}

func NewProjectService(contentDir string, isDev bool) ProjectService {
	return ProjectService{DirectoryPath: contentDir, IsDev: isDev}
}

func (s ProjectService) GetProjectBySlug(slug string) (*model.Project, error) {
	project := model.Project{}
	projects, err := s.GetAllProjects()

	if err != nil {
		return &project, err
	}

	filtered := util.Filter(projects, func(p *model.Project) bool {
		return p.Slug == slug
	})

	if len(filtered) > 1 {
		return nil, fmt.Errorf("more than one project matches the given slug")
	}

	if len(filtered) == 0 {
		return nil, nil
	}

	return filtered[0], nil
}

func (s ProjectService) GetLatestNProjects(n int) ([]*model.Project, error) {
	projects, err := s.GetAllProjects()

	if err != nil {
		return nil, err
	}

	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].StartDate.After(projects[j].StartDate)
	})

	if len(projects) <= n {
		return projects, nil
	}

	return projects[0:n], nil
}

func (s ProjectService) GetAllProjects() ([]*model.Project, error) {
	fileService := FileService{}
	markdownService := MarkdownService{}
	projects := []*model.Project{}

	contentPaths, err := fileService.GetDirectoryFilePaths(s.DirectoryPath)

	if err != nil {
		return projects, err
	}

	if len(contentPaths) < 1 {
		return projects, nil
	}

	for _, path := range contentPaths {
		_, file := filepath.Split(path)
		slug := strings.TrimSuffix(file, filepath.Ext(file))
		fileText, err := fileService.GetFileText(path)
		if err != nil {
			return projects, err
		}

		frontmatterMap, err := markdownService.GetFrontmatter([]byte(fileText))
		if err != nil {
			return projects, err
		}

		frontmatter := model.ProjectFrontmatter{}
		projectImages := []model.Image{}

		mapstructure.Decode(frontmatterMap, &frontmatter)
		mapstructure.Decode(frontmatter.Images, &projectImages)

		htmlContent, err := markdownService.ParseMarkdownContent([]byte(fileText))
		if err != nil {
			return projects, err
		}

		layout := "2006-01-02"

		loc, _ := time.LoadLocation("US/Central")
		startDate, err := time.ParseInLocation(layout, frontmatter.StartDate, loc)
		if err != nil {
			return projects, err
		}

		endDate, err := time.ParseInLocation(layout, frontmatter.EndDate, loc)
		if err != nil {
			endDate = startDate
		}

		project := model.Project{
			Name:      frontmatter.Name,
			Link:      frontmatter.Link,
			Summary:   frontmatter.Summary,
			StartDate: startDate,
			EndDate:   endDate,
			Content:   htmlContent,
			Slug:      slug,
			Images:    projectImages,
		}

		projects = append(projects, &project)
	}

	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].EndDate.After(projects[j].EndDate)
	})

	return projects, nil
}
