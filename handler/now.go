package handler

import (
	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/view/now"
)

type NowHandler struct{}

func (h NowHandler) HandleGetNowIndex(c echo.Context) error {
	current := c.Request().URL.Path
	remoteAddr := c.Request().RemoteAddr

	return render(c, now.Index(current, remoteAddr))
}
