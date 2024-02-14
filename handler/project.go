package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/service"
	"jeffcaldwell.is/view/project"
)

type ProjectHandler struct{}

func (h ProjectHandler) HandleGetProjectIndex(c echo.Context) error {
	current := c.Request().URL.Path
	remoteAddr := c.Get("remoteAddr").(string)
	isDev := c.Get("dev").(bool)
	projectService := service.NewProjectService("./content/projects", isDev)

	projects, err := projectService.GetAllProjects()

	if err != nil {
		return echo.NewHTTPError(501, fmt.Sprintf("error getting projects: %v", err))
	}

	return render(c, project.Index(current, remoteAddr, isDev, projects))
}

func (h ProjectHandler) HandleGetProject(c echo.Context) error {
	current := c.Request().URL.Path
	remoteAddr := c.Get("remoteAddr").(string)
	isDev := c.Get("dev").(bool)
	slug := c.Param("slug")
	projectService := service.NewProjectService("./content/projects", isDev)
	proj, err := projectService.GetProjectBySlug(slug)

	if err != nil {
		// TODO: We need a good error page
		fmt.Printf("error getting post by slug: %v\n", err)
	}
	return render(c, project.Project(current, remoteAddr, proj))
}
