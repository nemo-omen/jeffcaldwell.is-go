package service

import "jeffcaldwell.is/model"

func filterPosts(pp []*model.Post, test func(*model.Post) bool) (ret []*model.Post) {
	for _, p := range pp {
		if test(p) {
			ret = append(ret, p)
		}
	}
	return
}
