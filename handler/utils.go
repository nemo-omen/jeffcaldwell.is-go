package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	// currentPath := c.Request().URL.Path
	// ctx := context.WithValue(c.Request().Context(), "currentPath", currentPath)
	// return component.Render(ctx, c.Response())
	return component.Render(c.Request().Context(), c.Response())
}
