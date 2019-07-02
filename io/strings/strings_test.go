package strings_test

import (
	"testing"

	"gitlab.com/bl-go/toolkit.git/io/strings"
)

func TestStripMargin(t *testing.T) {
	testCases := []struct {
		input  string
		margin string
		want   string
	}{
		{
			input: `This
						 |Is
						 |A
						 |Test`, margin: "|", want: "This\nIs\nA\nTest",
		},
		{
			input: `
			| This
			| Is
			| A
			| Test
			`,
			margin: "| ", want: "This\nIs\nA\nTest",
		},
		{
			input: `This
						 #Is
						 #A
						 #Test`, margin: "#", want: "This\nIs\nA\nTest",
		},
		{
			input: `	This
						 #	Is
						 #	A
						 #	Test`, margin: "#\t", want: "This\nIs\nA\nTest",
		},
		{
			input: ` This
						 # Is
						 # A
						 # Test`, margin: "# ", want: "This\nIs\nA\nTest",
		},
	}

	for _, tc := range testCases {
		got := strings.StripMargin(tc.input, tc.margin)
		if tc.want != got {
			t.Errorf("StripMargin - want: %q but got: %q", tc.want, got)
		}
	}
}
