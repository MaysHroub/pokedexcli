package test

import "testing"

const (
	RED   = "\033[31m"
	GREEN = "\033[32m"
	RESET = "\033[0m"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world        ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "my NAME             is    mAys",
			expected: []string{"my", "name", "is", "mays"},
		},
	}

	// loop through the test cases
	for _, c := range cases {
		output := cleanInput(c.input)

		if len(output) != len(c.expected) {
			t.Errorf(RED+"Length mismatch for input \"%v\"; got %v, want %v"+RESET, c.input, len(output), len(c.expected))
			continue
		}

		for i := range output {
			word := output[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(RED+"Word mismatch at index %v; got \"%v\", want \"%v\""+RESET, i, word, expectedWord)
			}
		}
	}

}
