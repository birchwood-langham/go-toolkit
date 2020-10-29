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

func TestToCsv(t *testing.T) {
	type args struct {
		sep rune
		in []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one-item-comma-separated",
			args: args{
				sep: ',',
				in: []string{"Test"},
			},
			want: "Test",
		},
		{
			name: "multiple-items-comma-separated",
			args: args{
				sep: ',',
				in: []string{"This", "is", "a", "test"},
			},
			want: "This,is,a,test",
		},
		{
			name: "multiple-items-tab-separated",
			args: args{
				sep: '\t',
				in: []string{"This", "is", "a", "test"},
			},
			want: "This\tis\ta\ttest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.ToCsv(tt.args.sep, tt.args.in...); got != tt.want {
				t.Errorf("ToCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToQuotedCsv(t *testing.T) {
	type args struct {
		sep   rune
		quote strings.QuoteType
		in    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single-item-comma-separated-double-quoted",
			args: args{
				sep:   ',',
				quote: strings.DoubleQuote,
				in:    []string{"Test"},
			},
			want: `"Test"`,
		},
		{
			name: "multiple-items-comma-separated-double-quoted",
			args: args{
				sep:   ',',
				quote: strings.DoubleQuote,
				in:    []string{"This", "is", "a", "test"},
			},
			want: `"This","is","a","test"`,
		},
		{
			name: "single-item-comma-separated-single-quoted",
			args: args{
				sep:   ',',
				quote: strings.SingleQuote,
				in:    []string{"Test"},
			},
			want: `'Test'`,
		},
		{
			name: "multiple-items-comma-separated-single-quoted",
			args: args{
				sep:   ',',
				quote: strings.SingleQuote,
				in:    []string{"This", "is", "a", "test"},
			},
			want: `'This','is','a','test'`,
		},
		{
			name: "single-item-tab-separated-double-quoted",
			args: args{
				sep:   '\t',
				quote: strings.DoubleQuote,
				in:    []string{"Test"},
			},
			want: `"Test"`,
		},
		{
			name: "multiple-items-tab-separated-double-quoted",
			args: args{
				sep:   '\t',
				quote: strings.DoubleQuote,
				in:    []string{"This", "is", "a", "test"},
			},
			want: "\"This\"\t\"is\"\t\"a\"\t\"test\"",
		},
		{
			name: "single-item-tab-separated-single-quoted",
			args: args{
				sep:   '\t',
				quote: strings.SingleQuote,
				in:    []string{"Test"},
			},
			want: `'Test'`,
		},
		{
			name: "multiple-items-tab-separated-single-quoted",
			args: args{
				sep:   '\t',
				quote: strings.SingleQuote,
				in:    []string{"This", "is", "a", "test"},
			},
			want: "'This'\t'is'\t'a'\t'test'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.ToQuotedCsv(tt.args.sep, tt.args.quote, tt.args.in...); got != tt.want {
				t.Errorf("ToQuotedCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}