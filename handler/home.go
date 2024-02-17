package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/service"
	"jeffcaldwell.is/view/home"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeIndex(c echo.Context) error {
	isDev := c.Get("dev").(bool)
	contentService := service.NewPostService("./content/blog", isDev)
	latestPosts, err := contentService.GetLatestNContent(2)

	if err != nil {
		fmt.Printf("Error getting the latest posts: %v\n", err)
	}

	return render(c, home.Index(latestPosts))
}
