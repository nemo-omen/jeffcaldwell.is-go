package handler

import (
	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/view/blog"
)

type BlogHandler struct{}

func (h BlogHandler) HandleBlogIndex(c echo.Context) error {
	return render(c, blog.Index())
}

func (h BlogHandler) HandleBlogPost(c echo.Context) error {
	slug := c.Param("slug")
	frontmatter := map[string]string{"title": slug, "summary": "Just testing a very simple render"}
	return render(c, blog.Post(frontmatter))
}
