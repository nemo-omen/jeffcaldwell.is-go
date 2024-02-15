package model

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Frontmatter struct {
	Title      string
	Subtitle   string
	Author     string
	Summary    string
	Published  string
	Updated    string
	Draft      bool
	Tags       []string
	Scripts    []string
	RawContent []byte
}

type PostProps struct {
	Subtitle string
	Author   string
	Summary  string
	Updated  string
	Tags     []string
	Scripts  []string
	Draft    bool
}

type Post struct {
	Slug     string
	Id       string
	Title    string
	Subtitle string
	Author   string
	Summary  string
	PubDate  time.Time
	Updated  time.Time
	Tags     []string
	Content  string
	Scripts  []string
	Draft    bool
}

func NewPost(slug, title, pubDate, content string, props PostProps) *Post {
	shaBytes := sha256.Sum256([]byte(slug))
	layout := "2006-01-02"
	loc, _ := time.LoadLocation("US/Central")
	pubTime, error := time.ParseInLocation(layout, pubDate, loc)

	if error != nil {
		panic(fmt.Sprintf("NewPost pubDate argument: %v", error))
	}

	if props.Updated == "" {
		props.Updated = pubTime.String()
	}

	updateTime, error := time.ParseInLocation(layout, props.Updated, loc)

	if error != nil {
		// fmt.Printf("NewPost error parsing updated date %v\n", error)
		updateTime = pubTime
	}

	return &Post{
		Slug:     slug,
		Id:       hex.EncodeToString(shaBytes[:]),
		Title:    title,
		Subtitle: props.Subtitle,
		Author:   props.Author,
		Summary:  props.Summary,
		PubDate:  pubTime,
		Updated:  updateTime,
		Tags:     props.Tags,
		Content:  content,
		Scripts:  props.Scripts,
		Draft:    props.Draft,
	}
}
