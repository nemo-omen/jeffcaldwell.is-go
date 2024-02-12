package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/view/project"
)

type ProjectHandler struct{}

func (h ProjectHandler) HandleGetProjectIndex(c echo.Context) error {
	current := c.Request().URL.Path
	remoteAddr := c.Get("remoteAddr").(string)
	isDev := c.Get("dev").(bool)
	fmt.Println("is dev?: ", isDev)
	return render(c, project.Index(current, remoteAddr))
}
