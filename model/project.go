package model

import "time"

type Image struct {
	Src string
	Alt string
}

type Project struct {
	Name      string
	Link      string
	Summary   string
	StartDate time.Time
	EndDate   time.Time
	Images    []*Image
	Content   string
}
