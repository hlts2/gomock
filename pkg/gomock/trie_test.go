package gomock

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		input    map[string]Response
		path     string
		expected struct {
			resp Response
			ok   bool
		}
	}{
		{
			input: map[string]Response{
				"/": {
					Code: 200,
					Body: "hello",
				},
			},
			path: "/",
			expected: struct {
				resp Response
				ok   bool
			}{
				resp: Response{
					Code: 200,
					Body: "hello",
				},
				ok: true,
			},
		},
		{
			input: map[string]Response{
				"/": {
					Code: 200,
					Body: "hello",
				},
				"/abc": {
					Code: 500,
					Body: "tiny",
				},
			},
			path: "/abc",
			expected: struct {
				resp Response
				ok   bool
			}{
				resp: Response{
					Code: 500,
					Body: "tiny",
				},
				ok: true,
			},
		},
		{
			input: map[string]Response{
				"/": {
					Code: 200,
					Body: "hello",
				},
				"/abc?id=111": {
					Code: 500,
					Body: "tiny",
				},
			},
			path: "/abc?id=111",
			expected: struct {
				resp Response
				ok   bool
			}{
				resp: Response{
					Code: 500,
					Body: "tiny",
				},
				ok: true,
			},
		},
		{
			input: map[string]Response{
				"/": {
					Code: 200,
					Body: "little",
				},
				"/abc?id=*": {
					Code: 500,
					Body: "tiny",
				},
			},
			path: "/abc?id=1111",
			expected: struct {
				resp Response
				ok   bool
			}{
				resp: Response{
					Code: 500,
					Body: "tiny",
				},
				ok: true,
			},
		},
	}

	for i, test := range tests {
		tree := NewTrie()

		for path, resp := range test.input {
			tree.Insert(path, resp)
		}

		resp, ok := tree.Search(test.path)

		if test.expected.ok != ok {
			t.Errorf("tests[%d] - Search ok is wrong. expected: %v, got: %v", i, test.expected.ok, ok)
		}

		if !reflect.DeepEqual(test.expected.resp, resp) {
			t.Errorf("tests[%d] - Search resp is wrong. expected: %v, got: %v", i, test.expected.resp, resp)
		}
	}
}

func TestStringToBytes(t *testing.T) {
	tests := []struct {
		str      string
		expected []byte
	}{
		{
			str:      "hello world",
			expected: []byte("hello world"),
		},
		{
			str:      "/a/b/c",
			expected: []byte("/a/b/c"),
		},
	}

	for i, test := range tests {
		got := stringToBytes(test.str)

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("tests[%d] - stringToBytes is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}

func TestBytesToString(t *testing.T) {
	tests := []struct {
		bytes    []byte
		startPos int
		endPos   int
		expected struct {
			hasPanic bool
			str      string
		}
	}{
		{
			bytes:    []byte("hello world"),
			startPos: 0,
			endPos:   11,
			expected: struct {
				hasPanic bool
				str      string
			}{
				hasPanic: false,
				str:      "hello world",
			},
		},
		{
			bytes:    []byte("hello world"),
			startPos: 0,
			endPos:   1,
			expected: struct {
				hasPanic bool
				str      string
			}{
				hasPanic: false,
				str:      "h",
			},
		},
		{
			bytes:    []byte("hello world"),
			startPos: 10,
			endPos:   11,
			expected: struct {
				hasPanic bool
				str      string
			}{
				hasPanic: false,
				str:      "d",
			},
		},
		{
			bytes:    []byte("hello world"),
			startPos: -10,
			endPos:   11,
			expected: struct {
				hasPanic bool
				str      string
			}{
				hasPanic: true,
				str:      "",
			},
		},
		{
			bytes:    []byte("hello world"),
			startPos: 120,
			endPos:   140,
			expected: struct {
				hasPanic bool
				str      string
			}{
				hasPanic: true,
				str:      "",
			},
		},
	}

	for i, test := range tests {
		func() {
			defer func() {
				hasPanic := !(recover() == nil)

				if test.expected.hasPanic != hasPanic {
					t.Errorf("tests[%d] - bytesToString hasPanic is wrong. expected: %v, got: %v", i, test.expected.hasPanic, hasPanic)
				}
			}()
			got := bytesToString(test.bytes, test.startPos, test.endPos)

			if test.expected.str != got {
				t.Errorf("tests[%d] - bytesToString is wrong. expected: %v, got: %v", i, test.expected.str, got)
			}
		}()
	}
}
