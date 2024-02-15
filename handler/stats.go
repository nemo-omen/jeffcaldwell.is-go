package handler

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"jeffcaldwell.is/model"
	"jeffcaldwell.is/service"
	"jeffcaldwell.is/util"
	"jeffcaldwell.is/view/blog"
)

type StatsHandler struct{}

func (s StatsHandler) HandleGetPostStats(c echo.Context) error {
	isDev := c.Get("dev").(bool)
	postService := service.NewPostService("./content/blog", isDev)
	allPosts, err := postService.GetAllContent()
	calendar := model.Calendar{}

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

	for _, y := range years {
		year := model.Year{
			YearNum:    y,
			YearString: strconv.Itoa(y),
		}

		yearMonths := []time.Month{}
		loc, _ := time.LoadLocation("US/Central")
		jan := time.Date(y, 1, 1, 0, 0, 0, 0, loc)
		dec := time.Date(y, 12, 1, 0, 0, 0, 0, loc)
		for d := jan; !d.After(dec); d = d.AddDate(0, 1, 0) {
			yearMonths = append(yearMonths, d.Month())
		}

		for i, m := range yearMonths {
			month := model.Month{
				MonthNum:  i + 1,
				MonthName: m.String(),
			}

			first := time.Date(y, m, 1, 0, 0, 0, 0, loc)
			last := time.Date(y, m+1, 0, 0, 0, 0, 0, loc)
			for d := first; !d.After(last); d = d.AddDate(0, 0, 1) {
				didPost := false
				if util.ContainsTime(postDates, d) {
					didPost = true
				}
				day := model.Day{
					DayNum:  d.Day(),
					DayName: d.Weekday().String(),
					Date:    d,
					DidPost: didPost,
				}

				if didPost {
					posts := util.Filter(allPosts, func(p *model.Post) bool {
						return p.PubDate.Equal(d)
					})

					if len(posts) > 0 {
						day.Slug = posts[0].Slug
						day.Title = posts[0].Title
					}
				}

				month.Days = append(month.Days, day)
			}
			year.Months = append(year.Months, month)
		}
		calendar.Years = append(calendar.Years, year)
	}

	sort.Slice(calendar.Years, func(i, j int) bool {
		return calendar.Years[i].YearNum > calendar.Years[j].YearNum
	})

	fmt.Println(calendar.Years)

	return render(c, blog.PostGrid(calendar))
}
