package handler

import (
	"fmt"
	"net/http"
	"slices"
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
		loc, _ := time.LoadLocation("US/Central")
		y, m, d := time.Time(p.PubDate).Date()
		zeroDate := time.Date(y, m, d, 0, 0, 0, 0, loc)
		return zeroDate
	})

	start := filtered[len(filtered)-1].PubDate
	end := filtered[0].PubDate
	years := []int{}

	for d := start; !d.After(end); d = d.AddDate(1, 0, 0) {
		years = append(years, d.Year())
	}

	type calendarDay map[time.Time]bool
	type calendarMonth map[time.Month][]calendarDay
	type calendarYear map[int][]calendarMonth

	calendar := calendarYear{}

	for _, y := range years {
		calendar[y] = []calendarMonth{}
		yearMonths := []time.Month{}
		loc, _ := time.LoadLocation("US/Central")
		jan := time.Date(y, 1, 1, 0, 0, 0, 0, loc)
		dec := time.Date(y, 12, 1, 0, 0, 0, 0, loc)
		for d := jan; !d.After(dec); d = d.AddDate(0, 1, 0) {
			yearMonths = append(yearMonths, d.Month())

			for _, m := range yearMonths {
				month := calendarMonth{}
				first := time.Date(y, m, 1, 0, 0, 0, 0, loc)
				last := time.Date(y, m+1, 0, 0, 0, 0, 0, loc)
				for d := first; !d.After(last); d = d.AddDate(0, 0, 1) {
					didPost := false
					if slices.Contains(postDates, d) {
						didPost = true
					}
					cDay := calendarDay{d: didPost}
					month[m] = append(month[m], cDay)
				}
				calendar[y] = append(calendar[y], month)
			}
		}
	}

	fmt.Printf("%+v", calendar)

	return c.JSON(http.StatusOK, postDates)
}
