package gomock

import (
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

	index := endpoints.GetMachingEndpointIndex("GET", "/list/1/")
	if index == -1 {
		t.Errorf("GetMachingEndpointIndex(method, path) expected: %v, got: %v", 0, index)
	}

	index = endpoints.GetMachingEndpointIndex("GET", "/list/{:id}/name/")
	if index == -1 {
		t.Errorf("GetMachingEndpointIndex(method, path) expected: %v, got: %v", 1, index)
	}

	index = endpoints.GetMachingEndpointIndex("GET", "/list?id={:id}")
	if index == -1 {
		t.Errorf("GetMachingEndpointIndex(method, path) expected: %v, got: %v", 2, index)
	}

	index = endpoints.GetMachingEndpointIndex("GET", "/search?ei={:ei}&q={:q}")
	if index == -1 {
		t.Errorf("GetMachingEndpointIndex(method, path) expected: %v, got: %v", 3, index)
	}

	index = endpoints.GetMachingEndpointIndex("GET", "/search?ei={:ei}&q={:q}&hoge=11")
	if index != -1 {
		t.Errorf("GetMachingEndpointIndex(method, path) expected: %v, got: %v", -1, index)
	}
}
