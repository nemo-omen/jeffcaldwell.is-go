package handler

import (
	"github.com/labstack/echo/v4"
)

type FeedHandler struct{}

func (h FeedHandler) HandleAtomFeed(c echo.Context) error {
	// TODO: Generate an Atom feed here
	return echo.NewHTTPError(501, "Not implemented... yet")
}

func (h FeedHandler) HandleRssFeed(c echo.Context) error {
	// TODO: Generate an RSS feed here
	return echo.NewHTTPError(501, "Not implemented... yet")
}

func (h FeedHandler) HandleJsonFeed(c echo.Context) error {
	// TODO: Generate a JSON feed here
	return echo.NewHTTPError(501, "Not implemented... yet")
}
