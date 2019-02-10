package toolkit

import "strings"

// IsTrue takes a string representing a boolean value and converts it to a boolean, if the string is true, or yes (regardless of case)
// it will be evaluated as true, anything else will be regarded as false
func IsTrue(value string) bool {
	if len(value) < 1 || len(value) > 5 {
		return false
	} // false is the longest

	switch strings.ToUpper(value) {
	case "TRUE", "YES":
		return true
	default:
		return false
	}
}

// MaxPort returns the maximum TCP/IP port number
func MaxPort() int {
	return 65535
}
