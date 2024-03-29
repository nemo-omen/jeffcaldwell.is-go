package model

import "time"

type ProjectFrontmatter struct {
	Name      string             `mapstructure:"name"`
	Link      string             `mapstructure:"link"`
	Summary   string             `mapstructure:"summary"`
	StartDate string             `mapstructure:"startDate"`
	EndDate   string             `mapstructure:"endDate"`
	Images    []ImageFrontmatter `mapstructure:"images"`
}

type ImageFrontmatter struct {
	Src string `mapstructure:"src"`
	Alt string `mapstructure:"alt"`
}

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
	Content   string
	Slug      string
	Images    []Image
}
