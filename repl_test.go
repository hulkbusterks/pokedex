package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO   WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "abc abc abc abc ABC abc ABC",
			expected: []string{"abc", "abc", "abc", "abc", "abc", "abc", "abc"},
		}, // add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length of output is diff")
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words dont match")
				t.Fail()
			}
		}
	}
}
