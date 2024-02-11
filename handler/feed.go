package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/service"
)

type FeedHandler struct{}

func (h FeedHandler) HandleGetAtomFeed(c echo.Context) error {
	feedService := service.FeedService{}
	contentService := service.NewPostService("./content/blog")
	posts, err := contentService.GetAllContent()

	if err != nil {
		fmt.Printf("Error getting posts for feed: %+v\n", err)
	}
	feed, err := feedService.GetStyledAtomFeed(posts)

	if err != nil {
		return echo.NewHTTPError(501, fmt.Sprintf("Error getting feed %v\n", err))
	}

	feedString, err := xml.MarshalIndent(feed, "", "    ")

	if err != nil {
		return echo.NewHTTPError(501, fmt.Sprintf("Error converting feed to XML %v\n", err))
	}

	stylesheetString := `<?xml-stylesheet href="/public/style/feed.xsl" type="text/xsl"?>` + "\n"

	// feedString = []byte(xml.Header + string(feedString))
	feedString = []byte(xml.Header + stylesheetString + string(feedString))

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return c.String(http.StatusOK, string(feedString))
	// return c.XMLPretty(http.StatusOK, feed, "    ")
}

func (h FeedHandler) HandleGetRssFeed(c echo.Context) error {
	feedService := service.FeedService{}
	contentService := service.NewPostService("./content/blog")
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
	contentService := service.NewPostService("./content/blog")
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
