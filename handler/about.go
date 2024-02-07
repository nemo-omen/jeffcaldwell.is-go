package handler

import (
	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/view/about"
)

type AboutHandler struct{}

func (h AboutHandler) HandleAboutIndex(c echo.Context) error {
	return render(c, about.Index())
}
