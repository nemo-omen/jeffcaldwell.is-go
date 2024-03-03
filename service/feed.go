package service

import (
	"encoding/xml"
	"sort"

	"github.com/gorilla/feeds"
	"jeffcaldwell.is/model"
)

type FeedService struct{}

func (s FeedService) GetFeed(posts []*model.Post) (*feeds.Feed, error) {
	feedImg := feeds.Image{
		Url:    "https://jeffcaldwell.is/public/images/favicon.svg",
		Title:  "Jeff Caldwell",
		Link:   "https://jeffcaldwell.is",
		Width:  32,
		Height: 32,
	}
	feed := &feeds.Feed{
		Id:          "https://jeffcaldwell.is/blog",
		Title:       "Jeff Caldwell — Blog",
		Link:        &feeds.Link{Href: "https://jeffcaldwell.is/feed/atom", Rel: "self"},
		Description: "A blog about programming, web development, and anything else I feel like writing.",
		Author:      &feeds.Author{Name: "Jeff Caldwell", Email: "jeff@jeffcaldwell.is"},
		Copyright:   "© Jeff Caldwell",
		Image:       &feedImg,
		Updated:     posts[0].PubDate,
	}

	items := []*feeds.Item{}

	for _, post := range posts {
		if post.Updated.IsZero() {
			post.Updated = post.PubDate
		}

		items = append(items, &feeds.Item{
			Id:          "https://jeffcaldwell.is/blog/" + post.Slug,
			Title:       post.Title,
			Link:        &feeds.Link{Href: "https://jeffcaldwell.is/blog/" + post.Slug},
			Description: post.Summary,
			Author:      feed.Author,
			Created:     post.PubDate,
			Updated:     post.Updated,
			Content:     post.Content,
		})
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Created.After(items[j].Created)
	})

	feed.Items = items

	return feed, nil
}

func (s FeedService) GetAtomFeed(posts []*model.Post) (*feeds.AtomFeed, error) {
	feed, err := s.GetFeed(posts)
	if err != nil {
		return &feeds.AtomFeed{}, err
	}

	atomFeed := (&feeds.Atom{Feed: feed}).AtomFeed()
	atomFeed.Icon = feed.Image.Url
	atomFeed.Logo = feed.Image.Url

	for _, entry := range atomFeed.Entries {
		entry.Links = append(entry.Links, feeds.AtomLink{Href: "https://jeffcaldwell.is", Rel: "related", Type: "text/html"})
	}

	return atomFeed, nil
}

func (s FeedService) GetRSSFeed(posts []*model.Post) (*feeds.RssFeed, error) {
	feed, err := s.GetFeed(posts)
	if err != nil {
		return &feeds.RssFeed{}, err
	}

	rssFeed := (&feeds.Rss{Feed: feed}).RssFeed()
	// rssFeed.Category = "Web Development"

	return rssFeed, nil
}

func (s FeedService) GetJSONFeed(posts []*model.Post) (*feeds.JSONFeed, error) {
	feed, err := s.GetFeed(posts)
	if err != nil {
		return &feeds.JSONFeed{}, err
	}

	jsonFeed := (&feeds.JSON{Feed: feed}).JSONFeed()
	jsonFeed.HomePageUrl = "https://jeffcaldwell.is"
	jsonFeed.Favicon = feed.Image.Url
	jsonFeed.FeedUrl = feed.Link.Href

	return jsonFeed, nil
}

type StyleSheet struct {
	Href string `xml:"href,attr"`
	Type string `xml:"type,attr"`
}

type Category struct {
	XMLName xml.Name `xml:"category,omitempty"`
	Term    string   `xml:"term,attr"`
}

type StyledFeed struct {
	XMLName xml.Name `xml:"feed"`
	Xmlns   string   `xml:"xmlns,attr"`
	Title   string   `xml:"title"`   // required
	Id      string   `xml:"id"`      // required
	Updated string   `xml:"updated"` // required
	// Categories  string `xml:"category,omitempty"`
	// Categories  []Category
	Icon        string `xml:"icon,omitempty"`
	Logo        string `xml:"logo,omitempty"`
	Rights      string `xml:"rights,omitempty"` // copyright used
	Subtitle    string `xml:"subtitle,omitempty"`
	Link        *feeds.AtomLink
	Author      *feeds.AtomAuthor `xml:"author,omitempty"`
	Contributor *feeds.AtomContributor
	Entries     []*feeds.AtomEntry `xml:"entry"`
}

func (s FeedService) GetStyledAtomFeed(posts []*model.Post) (*StyledFeed, error) {
	atomFeed, err := s.GetAtomFeed(posts)
	if err != nil {
		return &StyledFeed{}, err
	}

	styledFeed := StyledFeed{
		XMLName:     atomFeed.XMLName,
		Xmlns:       atomFeed.Xmlns,
		Title:       atomFeed.Title,
		Id:          atomFeed.Id,
		Updated:     atomFeed.Updated,
		Icon:        atomFeed.Icon,
		Logo:        atomFeed.Logo,
		Rights:      atomFeed.Rights,
		Subtitle:    atomFeed.Subtitle,
		Link:        atomFeed.Link,
		Author:      atomFeed.Author,
		Contributor: atomFeed.Contributor,
		Entries:     atomFeed.Entries,
	}

	return &styledFeed, nil
}
