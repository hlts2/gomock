package gomock

import (
	"reflect"
	"regexp"
	"testing"
)

func TestGetMachingEndpointIndex(t *testing.T) {
	endpoints := Endpoints{
		{
			Request: Request{
				RegexRoute: regexp.MustCompile("GET/list/1/$"),
			},
		},
		{
			Request: Request{
				RegexRoute: regexp.MustCompile("GET/list/[^/&]+?/name/$"),
			},
		},
		{
			Request: Request{
				RegexRoute: regexp.MustCompile("GET/list\\?id=[^/&]+?$"),
			},
		},
		{
			Request: Request{
				RegexRoute: regexp.MustCompile("GET/search\\?ei=[^/&]+?&q=[^/&]+?$"),
			},
		},
	}

	e, ok := endpoints.MatchedEndpoint("GET", "/list/1/")
	if !ok {
		t.Errorf("MatchedEndpoint ok is wrong. expected: %v, got: %v", true, false)
	}

	if !reflect.DeepEqual(e, endpoints[0]) {
		t.Errorf("MatchedEndpoint e is wrong. expected: %v, got: %v", endpoints[0], e)
	}

	e, ok = endpoints.MatchedEndpoint("GET", "/list/{:id}/name/")
	if !ok {
		t.Errorf("MatchedEndpoint ok is wrong. expected: %v, got: %v", true, false)
	}

	if !reflect.DeepEqual(e, endpoints[1]) {
		t.Errorf("MatchedEndpoint e is wrong. expected: %v, got: %v", endpoints[1], e)
	}

	e, ok = endpoints.MatchedEndpoint("GET", "/list?id={:id}")
	if !ok {
		t.Errorf("MatchedEndpoint ok is wrong. expected: %v, got: %v", true, false)
	}

	if !reflect.DeepEqual(e, endpoints[2]) {
		t.Errorf("MatchedEndpoint e is wrong. expected: %v, got: %v", endpoints[2], e)
	}

	e, ok = endpoints.MatchedEndpoint("GET", "/search?ei={:ei}&q={:q}")
	if !ok {
		t.Errorf("MatchedEndpoint ok is wrong. expected: %v, got: %v", true, false)
	}

	if !reflect.DeepEqual(e, endpoints[3]) {
		t.Errorf("MatchedEndpoint e is wrong. expected: %v, got: %v", endpoints[2], e)
	}

	e, ok = endpoints.MatchedEndpoint("GET", "/search?ei={:ei}&q={:q}&hoge=11")
	if ok {
		t.Errorf("MatchedEndpoint is wrong. expected: %v, got: %v", false, true)
	}

	if !reflect.DeepEqual(e, Endpoint{}) {
		t.Errorf("MatchedEndpoint e is wrong. expected: %v, got: %v", Endpoint{}, e)
	}
}
