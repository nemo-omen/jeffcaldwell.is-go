package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/service"
)

type AnalyticsHandler struct{}

func (h AnalyticsHandler) PostPageview(c echo.Context) error {
	analyticsService := service.NewAnalyticsService("./data")
	data := c.Request().Header.Get("remoteaddr")
	analyticsService.SavePageview(data, c.Request().URL.Path)
	return c.String(http.StatusOK, "Yep!")
}
