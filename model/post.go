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
	RawContent []byte
}

type PostProps struct {
	Subtitle string
	Author   string
	Summary  string
	Updated  string
	Tags     []string
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
}

func NewPost(slug, title, pubDate, content string, props PostProps) *Post {
	shaBytes := sha256.Sum256([]byte(slug))
	pubTime, error := time.Parse(time.RFC3339, pubDate)

	if error != nil {
		panic(fmt.Sprintf("NewPost pubDate argument: %v", error))
	}

	if props.Updated == "" {
		props.Updated = pubTime.String()
	}

	updateTime, error := time.Parse(time.RFC3339, props.Updated)

	if error != nil {
		fmt.Printf("NewPost error parsing updated date %v\n", error)
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
	}
}
