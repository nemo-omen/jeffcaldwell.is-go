package util

import "reflect"

func ComparableSliceEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, el := range a {
		if el != b[i] {
			return false
		}
	}
	return true
}

func NonComparableSliceEqual[T any](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, el := range a {
		if !reflect.DeepEqual(el, b[i]) {
			return false
		}
	}
	return true
}
