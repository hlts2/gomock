package gomock

import (
	"regexp"
	"testing"
)

func TestNewRegexRoute(t *testing.T) {
	tests := []struct {
		route string
		want  string
	}{
		{
			route: "GET/list/1",
			want:  regexp.MustCompile("GET/list/1").String(),
		},
		{
			route: "GET/list/{:id}",
			want:  regexp.MustCompile("GET/list/.*?").String(),
		},
		{
			route: "GET/list/{:id}/{:name}",
			want:  regexp.MustCompile("GET/list/.*?/.*?").String(),
		},
		{
			route: "GET/list?id={:id}",
			want:  regexp.MustCompile("GET/list\\?id=.*?").String(),
		},
		{
			route: "GET/search?ei={:ei}&q={:q}",
			want:  regexp.MustCompile("GET/search\\?ei=.*?&q=.*?").String(),
		},
	}

	for i, test := range tests {
		got, err := newRegexRoute(test.route)

		if err != nil {
			t.Errorf("i = %d newRegexRoute(route) err is error: %v", i, err)
		}

		if test.want != got.String() {
			t.Errorf("i = %d newRegexRoute(route) want: %v, got: %v", i, test.want, got.String())
		}
	}
}
