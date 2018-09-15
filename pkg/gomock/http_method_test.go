package gomock

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		method   HTTPMethod
		expected string
	}{
		{
			method:   Get,
			expected: "GET",
		},
		{
			method:   Head,
			expected: "HEAD",
		},
		{
			method:   Post,
			expected: "POST",
		},
		{
			method:   Put,
			expected: "PUT",
		},
		{
			method:   Delete,
			expected: "DELETE",
		},
		{
			method:   Options,
			expected: "OPTIONS",
		},
		{
			method:   Patch,
			expected: "PATCH",
		},
		{
			method:   HTTPMethod(7),
			expected: "",
		},
		{
			method:   HTTPMethod(-1),
			expected: "",
		},
	}

	for i, test := range tests {
		got := test.method.String()

		if test.expected != got {
			t.Errorf("tests[%d] - String is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
