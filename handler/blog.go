package handler

import (
	"fmt"
	"sort"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/service"
	"jeffcaldwell.is/view/blog"
)

type BlogHandler struct{}

// Okay, here's the question. Do I implement a separate BlogService to
// read/write Post files and initialize Post structs? Seems like the most
// flexible way to go. That way, if I wanted to create my own private
// editor, I could do so without the handlers getting to hairy.
func (h BlogHandler) HandleBlogIndex(c echo.Context) error {
	contentService := service.NewContentService("./content/blog")

	current := c.Request().URL.Path
	remoteAddr := c.Get("remoteAddr").(string)

	posts, err := contentService.GetAllContent()

	if err != nil {
		fmt.Printf("error reading files at %s", "./content/blog")
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].PubDate.After(posts[j].PubDate)
	})

	return render(c, blog.Index(current, remoteAddr, posts))
}

func (h BlogHandler) HandleBlogPost(c echo.Context) error {
	current := c.Request().URL.Path
	remoteAddr := c.Get("remoteAddr").(string)
	slug := c.Param("slug")
	contentService := service.NewContentService("./content/blog")
	post, err := contentService.GetPostBySlug(slug)

	if err != nil {
		// TODO: We need a good error page
		fmt.Printf("error getting post by slug: %v\n", err)
	}
	return render(c, blog.Post(post, current, remoteAddr))
}
