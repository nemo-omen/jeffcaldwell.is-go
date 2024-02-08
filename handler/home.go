package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/service"
	"jeffcaldwell.is/view/home"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeIndex(c echo.Context) error {
	current := c.Request().URL.Path
	contentService := service.NewContentService("./content/blog")
	latestPosts, err := contentService.GetLatestNContent(2)

	if err != nil {
		fmt.Printf("Error getting the latest posts: %v\n", err)
	}

	return render(c, home.Index(current, latestPosts))
}
