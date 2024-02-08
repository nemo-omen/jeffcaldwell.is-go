package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"jeffcaldwell.is/model"
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

	return fileList, nil
}

func (s ContentService) GetFiles() ([]map[string]string, error) {
	filePaths, err := s.GetContentFilePaths()
	if err != nil {
		return nil, err
	}

	fileContents := []map[string]string{}

	for _, path := range filePaths {
		content, err := os.ReadFile(path)
		dir, file := filepath.Split(path)

		if err != nil {
			return nil, fmt.Errorf("error reading file at %s, %v", path, err)
		}

		fileContents = append(
			fileContents,
			map[string]string{
				"dir":      dir,
				"filename": file,
				"slug":     strings.TrimSuffix(file, filepath.Ext(file)),
				"content":  string(content),
			},
		)
	}

	return fileContents, nil
}

func (s ContentService) GetAllContent() ([]model.Post, error) {
	posts := []model.Post{}
	fileContents, err := s.GetFiles()

	if err != nil {
		return nil, fmt.Errorf("error getting file contents %v", err)
	}

	for _, filedata := range fileContents {
		content := filedata["content"]
		slug := filedata["slug"]
		postmeta, err := s.GetContentFrontmatter(content)

		if err != nil {
			return []model.Post{}, fmt.Errorf("error parsing frontmatter: %v", err)
		}

		renderedContent, err := s.ParseMarkdownContent(postmeta.RawContent)

		if err != nil {
			return []model.Post{}, fmt.Errorf("could not render content: %v", err)
		}

		fmt.Printf("%+v\n", renderedContent)
		props := model.PostProps{
			Subtitle: postmeta.Subtitle,
			Summary:  postmeta.Summary,
			Updated:  postmeta.Updated,
			Tags:     postmeta.Tags,
		}
		post := model.NewPost(slug, postmeta.Title, postmeta.Published, renderedContent, props)
		posts = append(posts, *post)
	}

	return posts, fmt.Errorf("not implemented")
}

func (s ContentService) ParseMarkdownContent(content []byte) (string, error) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(content)
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.LazyLoadImages
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer)), nil
}

func (s ContentService) GetContentFrontmatter(content string) (model.Frontmatter, error) {
	var matter struct {
		Title     string   `yaml:"title"`
		Subtitle  string   `yaml:"subtitle"`
		Author    string   `yaml:"author"`
		Summary   string   `yaml:"summary"`
		Published string   `yaml:"published"`
		Updated   string   `yaml:"updated"`
		Draft     bool     `yaml:"draft"`
		Tags      []string `yaml:"tags"`
	}

	rest, err := frontmatter.Parse(strings.NewReader(content), &matter)

	if err != nil {
		return model.Frontmatter{}, fmt.Errorf("Error reading frontmatter %v", err)
	}

	frontmatter := model.Frontmatter{
		Title:      matter.Title,
		Subtitle:   matter.Subtitle,
		Summary:    matter.Summary,
		Published:  matter.Published,
		Draft:      matter.Draft,
		Tags:       matter.Tags,
		RawContent: rest,
	}

	return frontmatter, nil
}
