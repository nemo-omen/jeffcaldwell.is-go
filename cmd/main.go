package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"jeffcaldwell.is/custommiddleware"
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
	isDevelopment := flag.Bool("dev", false, "Development Mode")
	flag.Parse()
	fmt.Printf("Running in dev mode?: %v", *isDevelopment)

	app := echo.New()
	app.Static("/public", "public")

	homeHandler := handler.HomeHandler{}
	blogHandler := handler.BlogHandler{}
	projectHandler := handler.ProjectHandler{}
	aboutHandler := handler.AboutHandler{}
	subscribeHandler := handler.SubscribeHandler{}
	feedHandler := handler.FeedHandler{}
	nowHandler := handler.NowHandler{}
	todoHandler := handler.TodoHandler{}
	themeHandler := handler.ThemeHandler{}
	statsHandler := handler.StatsHandler{}

	app.Use(custommiddleware.NewMiddlewareContextValue)
	app.Use(custommiddleware.SetRemoteAddr)
	app.Use(custommiddleware.SetCurrentPath)

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dev", *isDevelopment)
			return next(c)
		}
	})

	// app.Use(middleware.CORS())

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://jeffcaldwell.is", "http://localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	app.GET("/", homeHandler.HandleHomeIndex)
	app.GET("/blog", blogHandler.HandleBlogIndex)
	app.GET("/blog/:slug", blogHandler.HandleGetBlogPost)
	app.GET("/blog/tags", blogHandler.HandleGetTags)
	app.GET("/blog/tags/:tag", blogHandler.HandleGetPostsByTag)
	app.GET("/blog/blogroll", blogHandler.HandleGetBlogroll)
	app.GET("/blog/calendar", blogHandler.HandleGetCalendar)
	app.GET("/about", aboutHandler.HandleAboutIndex)
	app.GET("/subscribe", subscribeHandler.HandleSubscribeIndex)
	app.GET("/feed", subscribeHandler.HandleSubscribeIndex)
	app.GET("/feed/atom", feedHandler.HandleGetAtomFeed)
	app.GET("/feed/rss", feedHandler.HandleGetRssFeed)
	app.GET("feed/json", feedHandler.HandleGetJsonFeed)
	app.GET("/now", nowHandler.HandleGetNowIndex)
	app.GET("/todo", todoHandler.HandleGetTodoIndex)
	app.GET("/projects", projectHandler.HandleGetProjectIndex)
	app.GET("/projects/:slug", projectHandler.HandleGetProject)
	app.GET("/projects/experiments", projectHandler.HandleGetExperiments)
	app.GET("/projects/challenges", projectHandler.HandleGetChallenges)
	app.GET("/theme/:themeName", themeHandler.HandleGetTheme)
	app.GET("/blog/stats", statsHandler.HandleGetPostStats)

	app.Logger.Fatal(app.Start("0.0.0.0:1234"))
}
