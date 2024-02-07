package handler

import (
	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/view/blog"
)

type BlogHandler struct{}

// Okay, here's the question. Do I implement a separate BlogService to
// read/write Post files and initialize Post structs? Seems like the most
// flexible way to go. That way, if I wanted to create my own private
// editor, I could do so without the handlers getting to hairy.
func (h BlogHandler) HandleBlogIndex(c echo.Context) error {
	// contentService := service.NewContentService("./content")

	current := c.Request().URL.Path

	// contentFiles, err := contentService.GetContentFilePaths()
	// if err != nil {
	// 	fmt.Printf("error reading files at %s", "./content")
	// }

	return render(c, blog.Index(current))
}

func (h BlogHandler) HandleBlogPost(c echo.Context) error {
	current := c.Request().URL.Path
	slug := c.Param("slug")
	frontmatter := map[string]string{"title": slug, "summary": "Just testing a very simple render"}
	return render(c, blog.Post(frontmatter, current))
}
