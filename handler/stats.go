package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/model"
	"jeffcaldwell.is/service"
	"jeffcaldwell.is/util"
)

type StatsHandler struct{}

func (s StatsHandler) HandleGetPostStats(c echo.Context) error {
	isDev := c.Get("dev").(bool)
	postService := service.NewPostService("./content/blog", isDev)

	allPosts, err := postService.GetAllContent()

	if err != nil {
		return echo.NewHTTPError(501, "Error getting blog posts")
	}

	filtered := util.Filter(allPosts, func(post *model.Post) bool {
		return !post.Draft
	})

	postDates := util.MapSlice[*model.Post, time.Time](filtered, func(p *model.Post) time.Time {
		return p.PubDate
	})

	return c.JSON(http.StatusOK, postDates)
}
