package service

import (
	"crypto/md5"
	"fmt"

	"jeffcaldwell.is/model"
)

type AnalyticsService struct {
	DataDir   string
	Pageviews *[]model.Pageview
}

func NewAnalyticsService(dataDir string) *AnalyticsService {
	return &AnalyticsService{
		DataDir: dataDir,
	}
}

func (s *AnalyticsService) SavePageview(remoteAddr, path string) {
	hash := md5.New()
	addrHash := hash.Sum([]byte(remoteAddr))
	pv := model.NewPageview(string(addrHash), path)
	// welp, need to learn more about creating
	// and storing time-series data in go
	fmt.Println(pv)
}
