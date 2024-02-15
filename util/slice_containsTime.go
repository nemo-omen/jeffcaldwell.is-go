package util

import "time"

func ContainsTime(s []time.Time, t time.Time) bool {
	for _, tt := range s {
		if tt.Equal(t) {
			return true
		}
	}
	return false
}
