package card

import (
	"testing"
)

func TestWrapDescription(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "This is a short description",
			expected: []string{"This is a short description "},
		},
		{
			input: "This is a long description that is longer than 55 characters and should be wrapped",
			expected: []string{
				"This is a long description that is longer than 55 characters ",
				"and should be wrapped ",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := wrapDescription(test.input)
			if len(actual) != len(test.expected) {
				t.Fatalf("expected %d lines, actual %d lines", len(test.expected), len(actual))
			}
			for i := range actual {
				if actual[i] != test.expected[i] {
					t.Errorf("expected %q, actual %q", test.expected[i], actual[i])
				}
			}
		})
	}
}

func TestFormatCount(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{input: 0, expected: "0"},
		{input: 1, expected: "1"},
		{input: 950, expected: "0.9k"},
		{input: 1000, expected: "1.0k"},
		{input: 10000, expected: "10.0k"},
		{input: 100000, expected: "100k"},
		{input: 1000000, expected: "1000k"},
	}
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			actual := formatCount(test.input)
			if actual != test.expected {
				t.Errorf("expected %q, actual %q", test.expected, actual)
			}
		})
	}
}
