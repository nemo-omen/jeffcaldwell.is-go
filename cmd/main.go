package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/handler"
)

func GetRemoteAddr(c echo.Context) string {
	var remoteAddr string = c.Get("remoteAddr").(string)
	if remoteAddr != "" {
		return remoteAddr
	}
	return ""
}

func main() {
	isDevelopment := flag.Bool("dev", true, "Development Mode")
	flag.Parse()
	fmt.Printf("Running in dev mode?: %v", *isDevelopment)

	app := echo.New()
	app.Static("/public", "public")

	homeHandler := handler.HomeHandler{}
	blogHandler := handler.BlogHandler{}
	aboutHandler := handler.AboutHandler{}
	subscribeHandler := handler.SubscribeHandler{}
	feedHandler := handler.FeedHandler{}

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// fmt.Printf("Current request remote address: %v\n", c.Request().RemoteAddr)
			c.Set("remoteAddr", c.Request().RemoteAddr)
			return next(c)
		}
	})

	app.GET("/", homeHandler.HandleHomeIndex)
	app.GET("/blog", blogHandler.HandleBlogIndex)
	app.GET("/blog/:slug", blogHandler.HandleGetBlogPost)
	app.GET("/blog/tags", blogHandler.HandleGetTags)
	app.GET("/blog/tags/:tag", blogHandler.HandleGetPostsByTag)
	app.GET("/blog/blogroll", blogHandler.HandleGetBlogroll)
	app.GET("/about", aboutHandler.HandleAboutIndex)
	app.GET("/subscribe", subscribeHandler.HandleSubscribeIndex)
	app.GET("/feed", subscribeHandler.HandleSubscribeIndex)
	app.GET("/feed/atom", feedHandler.HandleGetAtomFeed)
	app.GET("/feed/rss", feedHandler.HandleGetRssFeed)
	app.GET("feed/json", feedHandler.HandleGetJsonFeed)

	app.Logger.Fatal(app.Start(":1234"))
}
