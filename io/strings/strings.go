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
