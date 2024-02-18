package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/snabb/sitemap"
	"jeffcaldwell.is/service"
)

type SitemapHandler struct{}

func (h SitemapHandler) HandleGetSitemap(c echo.Context) error {
	sm := sitemap.New()
	rootLoc := "https://jeffcaldwell.is"
	staticPaths := []string{
		"/",
		"/blog",
		"/projects",
		"/about",
	}

	for _, path := range staticPaths {
		urlLoc := rootLoc + path
		t := time.Now()

		sm.Add(&sitemap.URL{
			Loc:        urlLoc,
			LastMod:    &t,
			ChangeFreq: sitemap.Weekly,
			Priority:   0.5,
		})
	}

	// false = not dev
	postService := service.NewPostService("./content/blog", false)
	projectService := service.NewProjectService("./content/projects/projects", false)

	posts, err := postService.GetAllContent()
	if err != nil {
		fmt.Println("error getting posts: ", err.Error())
		return echo.ErrInternalServerError
	}

	for _, post := range posts {
		t := post.Updated
		urlLoc := rootLoc + "/blog/" + post.Slug

		sm.Add(&sitemap.URL{
			Loc:        urlLoc,
			LastMod:    &t,
			ChangeFreq: sitemap.Weekly,
			Priority:   0.5,
		})
	}

	projects, err := projectService.GetAllProjects()
	if err != nil {
		fmt.Println("error getting projects: ", err.Error())
		return echo.ErrInternalServerError
	}

	for _, project := range projects {
		t := project.EndDate
		urlLoc := rootLoc + "/projects/" + project.Slug

		sm.Add(&sitemap.URL{
			Loc:        urlLoc,
			LastMod:    &t,
			ChangeFreq: sitemap.Monthly,
			Priority:   0.5,
		})
	}

	return c.XMLPretty(http.StatusOK, sm, "    ")
}
