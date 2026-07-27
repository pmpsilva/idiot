package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// eidLength is the required length of a GSMA EID (eUICC Identifier).
const eidLength = 32

// eidBodyLength is the number of digits before the 2 check digits.
const eidBodyLength = eidLength - 2

// generateEID returns a random 32-digit EID with valid GSMA SGP.02 check
// digits. The optional prefix pins the leading digits (digits only, at most
// 30 chars). If empty, "89" (telecom MII) is used. The output is intended
// for testing/development only — it is not a real, provisionable EID.
func generateEID(prefix string) (string, error) {
	if prefix == "" {
		prefix = "89"
	}
	if len(prefix) > eidBodyLength {
		return "", errors.New("eid prefix must be at most 30 digits")
	}
	if !isDigits(prefix) {
		return "", errors.New("eid prefix must be digits only")
	}

	body := make([]byte, eidBodyLength)
	copy(body, prefix)
	for i := len(prefix); i < eidBodyLength; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		body[i] = byte('0' + n.Int64())
	}

	candidate := string(body) + "00"
	cd := calculateEIDCheckDigits(candidate)
	return string(body) + padTwoDigits(cd), nil
}

func padTwoDigits(n int) string {
	return string([]byte{byte('0' + n/10), byte('0' + n%10)})
}

// calculateEIDCheckDigits computes the two check digits for a 32-digit EID
// per GSMA SGP.02: replace the last two digits with "00", then the check
// digits are 98 - (n mod 97).
func calculateEIDCheckDigits(number string) int {
	base := number[:len(number)-2] + "00"
	n := new(big.Int)
	n.SetString(base, 10)
	mod := new(big.Int).Mod(n, big.NewInt(97)).Int64()
	return int(98 - mod)
}

// validateEID reports whether s is a well-formed EID with valid check digits.
func validateEID(s string) bool {
	if len(s) != eidLength || !isDigits(s) {
		return false
	}
	got := int((s[30]-'0')*10 + (s[31] - '0'))
	return calculateEIDCheckDigits(s) == got
}
