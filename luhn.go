package main

import "strings"

func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func calculateLuhnCheckDigit(number string) int {
	sum := 0
	double := true
	for i := len(number) - 1; i >= 0; i-- {
		d := int(number[i] - '0')
		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return (10 - (sum % 10)) % 10
}

func validateLuhn(number string) bool {
	if len(number) < 2 {
		return false
	}
	body := number[:len(number)-1]
	checkDigit := int(number[len(number)-1] - '0')
	return calculateLuhnCheckDigit(body) == checkDigit
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}
