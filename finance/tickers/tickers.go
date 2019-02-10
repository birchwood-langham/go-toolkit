package tickers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// CalculateIsinCheckDigit calculates the check digit for a given Isin
func CalculateIsinCheckDigit(isin string) (uint8, error) {
	checkDigit := uint8(0)

	var b strings.Builder

	for _, d := range strings.ToUpper(isin)[0:11] {
		switch {
		case 'A' <= d && d <= 'Z':
			_, err := fmt.Fprintf(&b, "%d", uint8(d-55))
			if err != nil {
				return uint8(0), err
			}
		case '0' <= d && d <= '9':
			_, err := fmt.Fprintf(&b, "%d", uint8(d-48))
			if err != nil {
				return uint8(0), err
			}
		default:
			// this is an error because isins are only made up of capital letters and numbers
			return 0, errors.New("invalid character in ISIN")
		}
	}

	ldigits := make([]int, 0)
	rdigits := make([]int, 0)

	digits := b.String()

	for i, d := range digits {
		if i%2 == 0 {
			v, _ := strconv.Atoi(string(d))
			ldigits = append(ldigits, v)
			continue
		}

		if (i+1)%2 == 0 {
			v, _ := strconv.Atoi(string(d))
			rdigits = append(rdigits, v)
			continue
		}
	}

	if len(digits)%2 == 1 {
		ldigits = double(&ldigits)
	} else {
		rdigits = double(&rdigits)
	}

	b.Reset()

	for _, d := range ldigits {
		_, _ = fmt.Fprintf(&b, "%d", d)
	}

	for _, d := range rdigits {
		_, _ = fmt.Fprintf(&b, "%d", d)
	}

	sum := 0

	for _, v := range b.String() {
		i, _ := strconv.Atoi(string(v))
		sum += i
	}

	checkDigit = uint8((10 - (sum % 10)) % 10)

	return checkDigit, nil
}

func double(digits *[]int) []int {
	double := make([]int, len(*digits))
	for i, v := range *digits {
		double[i] = v * 2
	}

	return double
}

// CalculateCusipCheckDigit calculates the check digit for a given cusip
func CalculateCusipCheckDigit(cusip string) (uint8, error) {
	checkDigit := uint8(0)
	var sum int

	for i, d := range strings.ToUpper(cusip)[0:7] {
		v := 0
		switch {
		case 'A' <= d && d <= 'Z':
			v = int(d) - 55
		case '0' <= d && d <= '9':
			v = int(d) - 48
		case '*' == d:
			v = 36
		case '@' == d:
			v = 37
		case '#' == d:
			v = 38
		default:
			return checkDigit, errors.New("invalid character in CUSIP")
		}

		// if i is even, multiply v by 2
		if (i+1)%2 == 0 {
			v *= 2
		}

		sum += (v / 10) + (v % 10)
	}

	checkDigit = uint8((10 - (sum % 10)) % 10)

	return checkDigit, nil
}

// CalculateSedolCheckDigit calculates the check digit for a given Sedol
func CalculateSedolCheckDigit(sedol string) (uint8, error) {
	checkDigit := uint8(0)
	weights := []int{1, 3, 1, 7, 3, 9}

	sum := 0

	for i, d := range strings.ToUpper(sedol)[0:6] {
		switch {
		case '0' <= d && d <= '9':
			sum += (int(d) - 48) * weights[i]
		case d == 'A' || d == 'E' || d == 'I' || d == 'O' || d == 'U':
			return checkDigit, errors.New("invalid character in SEDOL")
		case 'B' <= d && d <= 'Z':
			sum += (int(d) - 55) * weights[i]
		default:
			return checkDigit, errors.New("invalid character in SEDOL")
		}
	}

	checkDigit = uint8((10 - (sum % 10)) % 10)

	return checkDigit, nil
}

// CalculateFigiCheckDigit calculates the check digit for a given FIGI
func CalculateFigiCheckDigit(figi string) (uint8, error) {
	checkDigit := uint8(0)

	var sum int
	var b strings.Builder

	for i, d := range strings.ToUpper(figi)[0:11] {
		v := 0
		switch {
		case 'A' <= d && d <= 'Z':
			v = int(d) - 55
		case '0' <= d && d <= '9':
			v = int(d) - 48
		default:
			return checkDigit, errors.New("invalid character in FIGI")
		}

		// if i is even, multiply v by 2
		if (i+1)%2 == 0 {
			v *= 2
		}

		_, _ = fmt.Fprintf(&b, "%d", v)
	}

	for _, v := range b.String() {
		sum += int(v) - 48
	}

	checkDigit = uint8((10 - (sum % 10)) % 10)

	return checkDigit, nil
}

// IsIsinValid checks if a given Isin is valid
func IsIsinValid(isin string) bool {
	if len(isin) != 12 {
		return false
	}

	cd, err := CalculateIsinCheckDigit(isin)

	if err != nil {
		return false
	}

	return cd == uint8(isin[11]-48)
}

// IsCusipValid checks if the given Cusip is valid
func IsCusipValid(cusip string) bool {
	if len(cusip) != 9 {
		return false
	}

	cd, err := CalculateCusipCheckDigit(cusip)

	if err != nil {
		return false
	}

	return cd == uint8(cusip[8]-48)
}

// IsSedolValid checks if the given Sedol is valid
func IsSedolValid(sedol string) bool {
	if len(sedol) != 7 {
		return false
	}

	cd, err := CalculateSedolCheckDigit(sedol)

	if err != nil {
		return false
	}

	return cd == uint8(sedol[6]-48)
}

// IsFigiValid checks if the given Financial Instrument Global Identifier (Formerly a Bloomberg UID) is valid
func IsFigiValid(figi string) bool {
	if len(figi) != 12 {
		return false
	}

	cd, err := CalculateFigiCheckDigit(figi)

	if err != nil {
		return false
	}

	return cd == uint8(figi[11]-48)
}

// SedolToGbIsin convers a Sedol identifier to a GB Isin identifier
func SedolToGbIsin(sedol string) (string, error) {
	if !IsSedolValid(sedol) {
		return "", fmt.Errorf("%s is not a valid sedol", sedol)
	}

	prefix := fmt.Sprintf("GB00%s", sedol)
	cd, err := CalculateIsinCheckDigit(prefix)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%d", prefix, cd), nil
}

// SedolToIeIsin converts a Sedol identifier into an IE Isin identifier
func SedolToIeIsin(sedol string) (string, error) {
	if !IsSedolValid(sedol) {
		return "", fmt.Errorf("%s is not a valid sedol", sedol)
	}

	prefix := fmt.Sprintf("IE00%s", sedol)
	cd, err := CalculateIsinCheckDigit(prefix)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%d", prefix, cd), nil
}

// SedolToIsin converts a Sedol to a GB or IE Isin identifier
func SedolToIsin(sedol string, isGb bool) (string, error) {
	if isGb {
		return SedolToGbIsin(sedol)
	} else {
		return SedolToIeIsin(sedol)
	}
}

// CusipToUsIsin takes a cusip identifier and converts it to a US Isin
func CusipToUsIsin(cusip string) (string, error) {
	if !IsCusipValid(cusip) {
		return "", fmt.Errorf("%s is not a valid cusip", cusip)
	}
	prefix := fmt.Sprintf("US%s", cusip)
	cd, err := CalculateIsinCheckDigit(prefix)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%d", prefix, cd), nil
}

// CusipToCaIsin takes a Cusip and converts it to a CA Isin identifier
func CusipToCaIsin(cusip string) (string, error) {
	if !IsCusipValid(cusip) {
		return "", fmt.Errorf("%s is not a valid cusip", cusip)
	}
	prefix := fmt.Sprintf("CA%s", cusip)
	cd, err := CalculateIsinCheckDigit(prefix)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%d", prefix, cd), nil
}

// CusipToIsin takes a Cusip identifier and converts it to a US or CA Isin identifier
func CusipToIsin(cusip string, isUs bool) (string, error) {
	if isUs {
		return CusipToUsIsin(cusip)
	} else {
		return CusipToCaIsin(cusip)
	}
}
