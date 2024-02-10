package service

import (
	"github.com/gorilla/feeds"
	"jeffcaldwell.is/model"
)

type FeedService struct{}

func (s FeedService) GetFeed(posts []*model.Post) (*feeds.Feed, error) {
	feed := &feeds.Feed{
		Id:          "https://jeffcaldwell.is/blog",
		Title:       "Jeff Caldwell — Blog",
		Link:        &feeds.Link{Href: "https://jeffcaldwell.is/blog"},
		Description: "A blog about programming, web development, and anything else I feel like writing.",
		Author:      &feeds.Author{Name: "Jeff Caldwell", Email: "jeff@jeffcaldwell.is"},
		Copyright:   "&copy; Jeff Caldwell",
	}

	feed.Items = []*feeds.Item{}

	for _, post := range posts {
		feed.Items = append(feed.Items, &feeds.Item{
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

	return feed, nil
}

func (s FeedService) GetAtomFeed(posts []*model.Post) (*feeds.AtomFeed, error) {
	feed, err := s.GetFeed(posts)
	if err != nil {
		return &feeds.AtomFeed{}, err
	}

	atomFeed := (&feeds.Atom{Feed: feed}).AtomFeed()
	atomFeed.Category = "Web Development"

	return atomFeed, nil
}

func (s FeedService) GetRSSFeed(posts []*model.Post) (*feeds.RssFeed, error) {
	feed, err := s.GetFeed(posts)
	if err != nil {
		return &feeds.RssFeed{}, err
	}

	rssFeed := (&feeds.Rss{Feed: feed}).RssFeed()
	rssFeed.Category = "Web Development"

	return rssFeed, nil
}

func (s FeedService) GetJSONFeed(posts []*model.Post) (*feeds.JSONFeed, error) {
	feed, err := s.GetFeed(posts)
	if err != nil {
		return &feeds.JSONFeed{}, err
	}

	jsonFeed := (&feeds.JSON{Feed: feed}).JSONFeed()

	return jsonFeed, nil
}
