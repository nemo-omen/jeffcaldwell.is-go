package service

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/adrg/frontmatter"
	"jeffcaldwell.is/model"
	"jeffcaldwell.is/util"
)

type PostService struct {
	DirectoryPath string
	IsDev         bool
}

func NewPostService(dir string, isDev bool) PostService {
	return PostService{DirectoryPath: dir, IsDev: isDev}
}

func (s PostService) GetContentFilePaths() ([]string, error) {
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

func (s PostService) GetFileContent(path string) (map[string]string, error) {
	content, err := os.ReadFile(path)
	dir, file := filepath.Split(path)

	if err != nil {
		return nil, fmt.Errorf("error reading file at %s, %v", path, err)
	}

	return map[string]string{
		"dir":      dir,
		"filename": file,
		"slug":     strings.TrimSuffix(file, filepath.Ext(file)),
		"content":  string(content),
	}, nil
}

func (s PostService) GetFiles() ([]map[string]string, error) {
	filePaths, err := s.GetContentFilePaths()
	if err != nil {
		return nil, err
	}

	fileContents := []map[string]string{}

	for _, path := range filePaths {
		fileInfo, err := s.GetFileContent(path)

		if err != nil {
			return nil, err
		}

		fileContents = append(
			fileContents,
			fileInfo,
		)
	}

	return fileContents, nil
}

func (s PostService) MakePostFromFileData(fileData map[string]string) (*model.Post, error) {
	mdService := MarkdownService{}
	content := fileData["content"]
	slug := fileData["slug"]
	postmeta, err := s.GetContentFrontmatter(content)

	if err != nil {
		return &model.Post{}, fmt.Errorf("error parsing frontmatter: %v", err)
	}

	renderedContent, err := mdService.ParseMarkdownContent(postmeta.RawContent)

	if err != nil {
		return &model.Post{}, fmt.Errorf("could not render content: %v", err)
	}

	props := model.PostProps{
		Subtitle: postmeta.Subtitle,
		Summary:  postmeta.Summary,
		Updated:  postmeta.Updated,
		Tags:     postmeta.Tags,
		Scripts:  postmeta.Scripts,
		Draft:    postmeta.Draft,
	}
	post := model.NewPost(slug, postmeta.Title, postmeta.Published, renderedContent, props)
	return post, nil
}

func (s PostService) GetPostBySlug(slug string) (*model.Post, error) {
	posts, err := s.GetAllContent()

	if err != nil {
		return &model.Post{}, err
	}

	filtered := util.Filter(posts, func(p *model.Post) bool {
		return p.Slug == slug
	})

	if len(filtered) > 1 {
		return nil, fmt.Errorf("more than one post matches the given slug")
	}

	if len(filtered) == 0 {
		return nil, nil
	}

	return filtered[0], nil
}

func (s PostService) GetPostsByTag(tag string) ([]*model.Post, error) {
	posts, err := s.GetAllContent()

	if err != nil {
		return []*model.Post{}, err
	}

	filtered := util.Filter(posts, func(p *model.Post) bool {
		return slices.Contains(p.Tags, tag)
	})

	if !s.IsDev {
		filtered = util.Filter(filtered, func(p *model.Post) bool {
			return !p.Draft
		})
	}

	if len(filtered) == 0 {
		return nil, nil
	}

	return filtered, nil
}

func (s PostService) GetAllTags() ([]string, error) {
	posts, err := s.GetAllContent()
	tags := []string{}

	if err != nil {
		return tags, err
	}

	for _, post := range posts {
		for _, tag := range post.Tags {
			if !slices.Contains(tags, tag) {
				tags = append(tags, tag)
			}
		}
	}

	sort.Strings(tags)
	return tags, nil
}

func (s PostService) GetAllContent() ([]*model.Post, error) {
	posts := []*model.Post{}
	fileContents, err := s.GetFiles()

	if err != nil {
		return nil, fmt.Errorf("error getting file contents %v", err)
	}

	for _, fileData := range fileContents {
		post, err := s.MakePostFromFileData(fileData)
		if err != nil {
			return posts, fmt.Errorf("error initializing model.Post: %v", err)
		}
		posts = append(posts, post)
	}

	if !s.IsDev {
		posts = util.Filter(posts, func(p *model.Post) bool {
			return !p.Draft
		})
	}

	return posts, nil
}

func (s PostService) GetLatestNContent(n int) ([]*model.Post, error) {
	posts, err := s.GetAllContent()

	if err != nil {
		return nil, err
	}

	if !s.IsDev {
		posts = util.Filter(posts, func(p *model.Post) bool {
			return !p.Draft
		})
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].PubDate.After(posts[j].PubDate)
	})

	if len(posts) <= n {
		return posts, nil
	}

	return posts[0:n], nil
}

func (s PostService) GetContentFrontmatter(content string) (model.Frontmatter, error) {
	var matter struct {
		Title     string   `yaml:"title"`
		Subtitle  string   `yaml:"subtitle"`
		Author    string   `yaml:"author"`
		Summary   string   `yaml:"summary"`
		Published string   `yaml:"published"`
		Updated   string   `yaml:"updated"`
		Draft     bool     `yaml:"draft"`
		Tags      []string `yaml:"tags"`
		Scripts   []string `yaml:"scripts"`
	}

	rest, err := frontmatter.Parse(strings.NewReader(content), &matter)

	if err != nil {
		return model.Frontmatter{}, fmt.Errorf("error reading frontmatter %v", err)
	}

	frontmatter := model.Frontmatter{
		Title:      matter.Title,
		Subtitle:   matter.Subtitle,
		Summary:    matter.Summary,
		Published:  matter.Published,
		Draft:      matter.Draft,
		Tags:       matter.Tags,
		Scripts:    matter.Scripts,
		RawContent: rest,
	}

	return frontmatter, nil
}
