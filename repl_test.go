package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Mankey Machoke",
			expected: []string{"mankey", "machoke"},
		},
		{
			input:    "GaRdevoir LuCario",
			expected: []string{"gardevoir", "lucario"},
		},
		{
			input:    "Hojo   Jojo   ",
			expected: []string{"hojo", "jojo"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) < len(c.expected) || len(actual) > len(c.expected) {
			t.Errorf("size of slices do not match")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("the words in the slice do no match")
			}
		}
	}
}
