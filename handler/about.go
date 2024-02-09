package handler

import (
	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/view/about"
)

type AboutHandler struct{}

func (h AboutHandler) HandleAboutIndex(c echo.Context) error {
	current := c.Request().URL.Path
	remoteAddr := c.Get("remoteAddr").(string)
	return render(c, about.Index(current, remoteAddr))
}
