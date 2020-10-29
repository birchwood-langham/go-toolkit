package strings

import (
	"fmt"
	gs "strings"
)

func stripToRune(input string, margin rune) string {
	sb := gs.Builder{}
	nl := '\n'

	strip := false

	// in := gs.TrimSpace(input)

	for _, r := range input {
		if r == margin {
			sb.WriteRune(r)
			strip = false
			continue
		}

		if r == nl {
			strip = true
			sb.WriteRune(r)
			continue
		}

		if strip {
			continue
		}

		sb.WriteRune(r)
	}

	return sb.String()
}

/* StripMargin removes any leading spaces and the specified margin from multi-line strings
 * allowing you to write and aligh multi-line strings nicely in your code similar to the Scala
 * StripMarging
 */
func StripMargin(input string, margin string) string {
	mrs := []rune(margin)

	in := stripToRune(input, mrs[0])

	m := fmt.Sprintf("%c%s", '\n', margin)
	lines := gs.Split(in, m)

	sb := gs.Builder{}

	for _, l := range lines {
		sb.WriteString(l)
		sb.WriteRune('\n')
	}

	return gs.TrimSpace(gs.Trim(sb.String(), "\n"))
}

// SplitAndTrimSpace splits a string using the provided separators and removes any leading and trailing spaces around the results
func SplitAndTrimSpace(input string, sep string) (output []string) {
	if len(input) == 0 {
		return output
	}

	values := gs.Split(input, sep)

	output = make([]string, len(values))

	for i, v := range values {
		output[i] = gs.TrimSpace(v)
	}

	return
}

func ToCsv(sep rune, in ...string) string {
	b := gs.Builder{}

	for i, s := range in {
		if i > 0 {
			b.WriteRune(sep)
		}

		b.WriteString(s)
	}

	return b.String()
}

type QuoteType rune

const SingleQuote QuoteType = '\''
const DoubleQuote QuoteType = '"'

func ToQuotedCsv(sep rune, quote QuoteType, in ...string) string {
	b := gs.Builder{}

	for i, s := range in {
		if i > 0 {
			b.WriteRune(sep)
		}

		b.WriteRune(rune(quote))
		b.WriteString(s)
		b.WriteRune(rune(quote))
	}

	return b.String()
}