package main

import (
	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/handler"
)

func main() {
	app := echo.New()
	app.Static("/public", "public")

	homeHandler := handler.HomeHandler{}
	blogHandler := handler.BlogHandler{}
	aboutHandler := handler.AboutHandler{}
	subscribeHandler := handler.SubscribeHandler{}
	feedHandler := handler.FeedHandler{}

	app.GET("/", homeHandler.HandleHomeIndex)
	app.GET("/blog", blogHandler.HandleBlogIndex)
	app.GET("/blog/:slug", blogHandler.HandleBlogPost)
	app.GET("/about", aboutHandler.HandleAboutIndex)
	app.GET("/subscribe", subscribeHandler.HandleSubscribeIndex)
	app.GET("/feed", subscribeHandler.HandleSubscribeIndex)
	app.GET("/feed/atom", feedHandler.HandleAtomFeed)
	app.GET("/feed/rss", feedHandler.HandleRssFeed)
	app.GET("feed/json", feedHandler.HandleJsonFeed)

	app.Logger.Fatal(app.Start(":1234"))
}
