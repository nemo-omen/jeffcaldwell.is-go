package model

import "time"

type Pageview struct {
	IpHash string
	Path   string
	Date   string
}

func NewPageview(ipHash, path string) *Pageview {
	now := time.Now()
	date := now.Format("2006-01-02")

	return &Pageview{
		IpHash: ipHash,
		Path:   path,
		Date:   date,
	}
}
