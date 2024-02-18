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
	feedHandler := handler.FeedHandler{}
	nowHandler := handler.NowHandler{}
	todoHandler := handler.TodoHandler{}
	themeHandler := handler.ThemeHandler{}
	statsHandler := handler.StatsHandler{}
	sitemapHandler := handler.SitemapHandler{}

	app.Use(custommiddleware.NewMiddlewareContextValue)
	app.Use(custommiddleware.SetRemoteAddr)
	app.Use(custommiddleware.SetCurrentPath)

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dev", *isDevelopment)
			return next(c)
		}
	})

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://jeffcaldwell.is", "http://localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	getRoutes := []struct {
		Route   string
		Handler echo.HandlerFunc
	}{
		{"/", homeHandler.HandleHomeIndex},
		{"/blog", blogHandler.HandleBlogIndex},
		{"/blog/:slug", blogHandler.HandleGetBlogPost},
		{"/blog/tags", blogHandler.HandleGetTags},
		{"/blog/tags/:tag", blogHandler.HandleGetPostsByTag},
		{"/blog/blogroll", blogHandler.HandleGetBlogroll},
		{"/blog/calendar", blogHandler.HandleGetCalendar},
		{"/blog/subscribe", blogHandler.HandleGetSubscribe},
		{"/about", aboutHandler.HandleAboutIndex},
		{"/feed/atom", feedHandler.HandleGetAtomFeed},
		{"/feed/rss", feedHandler.HandleGetRssFeed},
		{"feed/json", feedHandler.HandleGetJsonFeed},
		{"/now", nowHandler.HandleGetNowIndex},
		{"/todo", todoHandler.HandleGetTodoIndex},
		{"/projects", projectHandler.HandleGetProjectIndex},
		{"/projects/:slug", projectHandler.HandleGetProject},
		{"/projects/experiments", projectHandler.HandleGetExperiments},
		{"/projects/challenges", projectHandler.HandleGetChallenges},
		{"/theme/:themeName", themeHandler.HandleGetTheme},
		{"/blog/stats", statsHandler.HandleGetPostStats},
		{"/sitemap", sitemapHandler.HandleGetSitemap},
		{"/sitemap.xml", sitemapHandler.HandleGetSitemap},
		{"/robots.txt", func(c echo.Context) error {
			smTxt := "Sitemap: https://jeffcaldwell.is/sitemap.xml\nUser-agent: *\nDisallow:"
			return c.String(http.StatusOK, smTxt)
		}},
	}

	for _, r := range getRoutes {
		app.GET(r.Route, r.Handler)
	}

	app.Logger.Fatal(app.Start("0.0.0.0:1234"))
}
