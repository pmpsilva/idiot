package main

import (
	"errors"
	"strings"
)

// nibbleSwap swaps each adjacent pair of characters in s.
// The operation is its own inverse — applying it twice returns the original.
// If s has an odd length the last character is left in place.
func nibbleSwap(s string) string {
	b := []byte(s)
	for i := 0; i+1 < len(b); i += 2 {
		b[i], b[i+1] = b[i+1], b[i]
	}
	return string(b)
}

// swimsi converts a 15-digit IMSI to its telco nibble-swapped form, or
// converts an 18-digit swapped value back to the original IMSI.
//
// Forward  (15 digits → 18 digits): prepend "9", nibble-swap, prepend "08"
// Reverse  (18 digits → 15 digits): strip "08", nibble-swap, strip leading "9"
//
// Mirrors the Java helper getImsiCompleteAndSwapped / nibbleSwap used in
// the telco services.
func swimsi(input string) (string, error) {
	input = strings.TrimSpace(input)
	switch len(input) {
	case 15:
		if !isDigits(input) {
			return "", errors.New("IMSI must be digits only")
		}
		return "08" + nibbleSwap("9"+input), nil

	case 18:
		if !isDigits(input) {
			return "", errors.New("swapped IMSI must be digits only")
		}
		if !strings.HasPrefix(input, "08") {
			return "", errors.New("swapped IMSI must start with '08'")
		}
		body := nibbleSwap(input[2:]) // strip "08", nibble-swap back
		if len(body) == 0 || body[0] != '9' {
			return "", errors.New("invalid swapped IMSI: expected leading '9' after nibble-swap")
		}
		return body[1:], nil // strip the sentinel "9"

	default:
		return "", errors.New("input must be a 15-digit IMSI or an 18-digit swapped IMSI")
	}
}
