package util

import (
	"testing"
)

func TestFilterStrings(t *testing.T) {
	type filterTest struct {
		description string
		source      []string
		testFunc    func(s string) bool
		expected    []string
	}

	tests := []filterTest{
		{
			"Test no matching val",
			[]string{"b", "c", "d"},
			func(s string) bool {
				return s == "a"
			},
			[]string{},
		},
		{
			"Test slice of same val",
			[]string{"a", "a", "a", "a"},
			func(s string) bool {
				return s == "a"
			},
			[]string{"a", "a", "a", "a"},
		},
		{
			"Test slice with single matching val",
			[]string{"a", "b", "c", "d"},
			func(s string) bool {
				return s == "a"
			},
			[]string{"a"},
		},
	}

	for _, tt := range tests {
		out := Filter(tt.source, tt.testFunc)

		if !ComparableSliceEqual(out, tt.expected) {
			t.Errorf("Expected %v got %v", tt.expected, out)
		}
	}
}

func TestFilterStructs(t *testing.T) {
	type person struct {
		First string
		Last  string
		Age   int
	}

	type filterTest struct {
		description string
		source      []person
		testFunc    func(p person) bool
		expected    []person
	}

	testVals := []person{
		{"James", "Caldwell", 23},
		{"Juan Carlos", "Caldwell", 39},
		{"Jeff", "Caldwell", 45},
		{"Maria", "Caldwell", 69},
		{"Wayne", "Caldwell", 72},
	}

	compVals := []person{
		{"James", "Caldwell", 23},
		{"Juan Carlos", "Caldwell", 39},
		{"Jeff", "Caldwell", 45},
		{"Maria", "Caldwell", 69},
		{"Wayne", "Caldwell", 72},
	}

	tests := []filterTest{
		{
			"Test no matching person",
			testVals,
			func(p person) bool {
				return p.Age == 19
			},
			[]person{},
		},
		{
			"Test one matching person",
			testVals,
			func(p person) bool {
				return p.Age == 45
			},
			[]person{
				{"Jeff", "Caldwell", 45},
			},
		},
		{
			"Test all elements match",
			testVals,
			func(p person) bool {
				return p.Last == "Caldwell"
			},
			compVals,
		},
	}

	for _, tt := range tests {
		res := Filter(tt.source, tt.testFunc)

		if !NonComparableSliceEqual(res, tt.expected) {
			t.Errorf("Expected %v got %v", tt.expected, res)
		}
	}
}
