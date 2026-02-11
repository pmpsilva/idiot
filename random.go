package main

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func generatePassword(length int, includeLetters, includeDigits, includeSpecial bool) (string, error) {
	if length <= 0 {
		return "", errors.New("password length must be greater than zero")
	}

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>/?"

	var charset string
	if includeLetters {
		charset += letters
	}
	if includeDigits {
		charset += digits
	}
	if includeSpecial {
		charset += special
	}

	if charset == "" {
		return "", errors.New("at least one character type must be selected")
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[idx.Int64()]
	}

	return string(password), nil
}
