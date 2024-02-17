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
	isDev := c.Get("dev").(bool)
	contentService := service.NewPostService("./content/blog", isDev)

	posts, err := contentService.GetAllContent()

	if err != nil {
		fmt.Printf("error reading files at %s", "./content/blog")
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].PubDate.After(posts[j].PubDate)
	})

	return render(c, blog.Index(posts))
}

func (h BlogHandler) HandleGetBlogPost(c echo.Context) error {
	isDev := c.Get("dev").(bool)
	slug := c.Param("slug")
	contentService := service.NewPostService("./content/blog", isDev)
	post, err := contentService.GetPostBySlug(slug)

	if err != nil {
		// TODO: We need a good error page
		fmt.Printf("error getting post by slug: %v\n", err)
	}
	return render(c, blog.Post(post))
}

func (h BlogHandler) HandleGetBlogroll(c echo.Context) error {
	return render(c, blog.BlogRoll())
}

func (h BlogHandler) HandleGetTags(c echo.Context) error {
	// TODO: Filter for published posts only
	isDev := c.Get("dev").(bool)
	contentService := service.NewPostService("./content/blog", isDev)

	tags, err := contentService.GetAllTags()

	if err != nil {
		fmt.Printf("error getting all tags: %v\n", err)
	}

	return render(c, blog.Tags(tags))
}

func (h BlogHandler) HandleGetPostsByTag(c echo.Context) error {
	isDev := c.Get("dev").(bool)

	tag := c.Param("tag")
	contentService := service.NewPostService("./content/blog", isDev)
	posts, err := contentService.GetPostsByTag(tag)

	if err != nil {
		fmt.Printf("Error getting posts by tag %s: %v\n", tag, err)
	}

	return render(c, blog.PostList(tag, posts))
}

func (h BlogHandler) HandleGetCalendar(c echo.Context) error {
	return render(c, blog.Calendar())
}

func (h BlogHandler) HandleGetSubscribe(c echo.Context) error {
	return render(c, blog.Subscribe())
}
