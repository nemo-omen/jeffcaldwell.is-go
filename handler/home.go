package handler

import (
	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/view/home"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeIndex(c echo.Context) error {
	return render(c, home.Index())
}
