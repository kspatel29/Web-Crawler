package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
	}{
		{
			name:     "Return Normalized URL",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "Return Normalized URL",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "Return Normalized URL",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "HTTP without trailing slash",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
				name:     "No scheme with trailing slash",
				inputURL: "blog.boot.dev/path/",
				expected: "blog.boot.dev/path",
		},
		{
				name:     "No scheme without trailing slash",
				inputURL: "blog.boot.dev/path",
				expected: "blog.boot.dev/path",
		},
		{
				name:     "Mixed case scheme",
				inputURL: "HTTPS://blog.boot.dev/path/",
				expected: "blog.boot.dev/path",
		},
		
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}