package strings_test

import (
	"reflect"
	"testing"

	"github.com/birchwood-langham/go-toolkit/io/strings"
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

func TestSplitAndTrimSpace(t *testing.T) {
	type args struct {
		input string
		sep   string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []string
	}{
		{"Split comma, no spaces", args{"A,B,C,D", ","}, []string{"A", "B", "C", "D"}},
		{"Split comma, with spaces", args{"A, B, C, D", ","}, []string{"A", "B", "C", "D"}},
		{"Split tab, no spaces", args{"A	B	C	D", "\t"}, []string{"A", "B", "C", "D"}},
		{"Split tab, with spaces", args{"A	 B	 C	 D", "\t"}, []string{"A", "B", "C", "D"}},
		{"Split empty string using comma", args{"", ","}, []string{}},
		{"Split empty string using tab", args{"", "\t"}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput := strings.SplitAndTrimSpace(tt.args.input, tt.args.sep)

			if len(gotOutput) == 0 && len(tt.wantOutput) == 0 {
				return
			}

			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("SplitAndTrimSpace() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
