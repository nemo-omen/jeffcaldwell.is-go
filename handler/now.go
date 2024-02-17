package handler

import (
	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/view/now"
)

type NowHandler struct{}

func (h NowHandler) HandleGetNowIndex(c echo.Context) error {
	return render(c, now.Index())
}
