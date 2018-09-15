package gomock

import (
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		queries  Queries
		targets  Queries
		expected bool
	}{
		{
			queries: Queries{
				{
					key:   "id",
					value: "123",
				},
			},
			targets: Queries{
				{
					key:   "id",
					value: "123",
				},
			},
			expected: true,
		},
		{
			queries: Queries{
				{
					key:   "id",
					value: "*",
				},
			},
			targets: Queries{
				{
					key:   "id",
					value: "3333",
				},
			},
			expected: true,
		},
		{
			queries: Queries{
				{
					key:   "id",
					value: "*",
				},
			},
			targets: Queries{
				{
					key:   "name",
					value: "little",
				},
			},
			expected: false,
		},
		{
			queries: Queries{
				{
					key:   "id",
					value: "*",
				},
				{
					key:   "name",
					value: "*",
				},
			},
			targets: Queries{
				{
					key:   "name",
					value: "little",
				},
			},
			expected: false,
		},
	}

	for i, test := range tests {
		got := test.queries.Match(test.targets)

		if test.expected != got {
			t.Errorf("tests[%d] - Match is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestParseQuery(t *testing.T) {
	tests := []struct {
		path     string
		expected Queries
	}{
		{
			path: "/a/b/c?id=1111",
			expected: Queries{
				{
					key:   "id",
					value: "1111",
				},
			},
		},
		{
			path: "id=1111",
			expected: Queries{
				{
					key:   "id",
					value: "1111",
				},
			},
		},
		{
			path: "id=1111&name=little",
			expected: Queries{
				{
					key:   "id",
					value: "1111",
				},
				{
					key:   "name",
					value: "little",
				},
			},
		},
		{
			path:     "id=1111&name==little",
			expected: Queries{},
		},
		{
			path:     "?",
			expected: Queries{},
		},
		{
			path:     "",
			expected: Queries{},
		},
	}

	for i, test := range tests {
		got := parseQuery(test.path)

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("tests[%d] - parseQuery is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
