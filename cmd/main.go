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

	app.GET("/", homeHandler.HandleHomeIndex)
	app.GET("/blog", blogHandler.HandleBlogIndex)
	app.GET("/blog/:slug", blogHandler.HandleBlogPost)
	app.GET("/about", aboutHandler.HandleAboutIndex)

	app.Logger.Fatal(app.Start(":1234"))
}
