package model

import "time"

type CalendarDay map[time.Time]bool
type CalendarMonth map[time.Month][]CalendarDay
type CalendarYear map[int][]CalendarMonth

type Calendar struct {
	Years []Year
}

type Year struct {
	YearNum    int
	YearString string
	Months     []Month
}

type Month struct {
	MonthNum  int
	MonthName string
	Days      []Day
}

type Day struct {
	DayNum  int
	DayName string
	Date    time.Time
	DidPost bool
	Slug    string
	Title   string
}
