package gomock

import (
	"strings"
)

// Query represetns URL query item
type Query struct {
	key   string
	value string
}

// Queries represetns Query slice
type Queries []Query

// Match returns true if queries and targets are the same, otherwise it returns false.
func (queries Queries) Match(targets Queries) bool {
	if len(queries) != len(targets) {
		return false
	}

	for _, target := range targets {
		for _, query := range queries {
			exist := false
			if target.key == query.key {
				exist = true
				if query.value != "*" && query.value != target.value {
					return false
				}
			}

			if !exist {
				return false
			}
		}
	}

	return true
}

func parseQuery(path string) Queries {
	pos := 0
	for i, b := range []byte(path) {
		if b == '?' {
			pos = i + 1
			break
		}
	}

	if pos >= len(path) {
		return Queries{}
	}

	var queries Queries
	for _, rawQuery := range strings.Split(path[pos:], "&") {
		items := strings.Split(rawQuery, "=")
		if len(items) != 2 {
			return Queries{}
		}

		query := Query{
			key:   items[0],
			value: items[1],
		}

		queries = append(queries, query)
	}

	return queries
}
