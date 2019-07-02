package strings

import gs "strings"

// StripMargin is a utility that replicates the Scala stripMargin function
func StripMargin(input string, margin rune) (string, error) {
	sb := gs.Builder{}
	nl := '\n'

	strip := false

	input = gs.TrimSpace(gs.Trim(input, "\n"))

	for _, r := range input {
		if r == margin {
			strip = false
			continue
		}

		if r == nl {
			strip = true
			if _, err := sb.WriteRune(r); err != nil {
				return "", err
			}
			continue
		}

		if strip {
			continue
		}

		if _, err := sb.WriteRune(r); err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}
