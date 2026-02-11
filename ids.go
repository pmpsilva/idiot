package main

import (
	"crypto/rand"
	"strings"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

func generateUUID() string {
	return uuid.NewString()
}

func generateULID(prefix string) string {
	id := ulid.MustNew(ulid.Now(), rand.Reader).String()
	if p := strings.TrimSpace(prefix); p != "" {
		return p + id
	}
	return id
}
