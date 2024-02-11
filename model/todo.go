package model

import "time"

type TodoPageData struct {
	Title       string
	Updated     time.Time
	PageContent string
}
