package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/service"
)

type FeedHandler struct{}

func (h FeedHandler) HandleGetAtomFeed(c echo.Context) error {
	feedService := service.FeedService{}
	contentService := service.NewContentService("./content/blog")
	posts, err := contentService.GetAllContent()

	if err != nil {
		fmt.Printf("Error getting posts for feed: %+v\n", err)
	}
	feed, err := feedService.GetAtomFeed(posts)

	if err != nil {
		return echo.NewHTTPError(501, fmt.Sprintf("Error getting feed %v\n", err))
	}

	return c.XMLPretty(http.StatusOK, feed, "    ")
}

func (h FeedHandler) HandleGetRssFeed(c echo.Context) error {
	feedService := service.FeedService{}
	contentService := service.NewContentService("./content/blog")
	posts, err := contentService.GetAllContent()

	if err != nil {
		fmt.Printf("Error getting posts for feed: %+v\n", err)
	}
	feed, err := feedService.GetRSSFeed(posts)

	if err != nil {
		return echo.NewHTTPError(501, fmt.Sprintf("Error getting feed %v\n", err))
	}
	return c.XMLPretty(http.StatusOK, feed, "    ")
}

func (h FeedHandler) HandleGetJsonFeed(c echo.Context) error {
	feedService := service.FeedService{}
	contentService := service.NewContentService("./content/blog")
	posts, err := contentService.GetAllContent()

	if err != nil {
		fmt.Printf("Error getting posts for feed: %+v\n", err)
	}
	feed, err := feedService.GetJSONFeed(posts)

	if err != nil {
		return echo.NewHTTPError(501, fmt.Sprintf("Error getting feed %v\n", err))
	}
	return c.JSONPretty(http.StatusOK, feed, "  ")
}
