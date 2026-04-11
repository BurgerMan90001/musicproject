package randutil

import (
	"crypto/rand"

	"encoding/base64"
	mathrand "math/rand"
)

// Returns a random item from a list
func RandItem(list []string) string {
	return list[mathrand.Intn(len(list))]
}

// Generate random strings
type String interface {
	New(length int) string
}

type RandRead struct {
}

func (s *RandRead) New() string {
	b := make([]byte, 128)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
