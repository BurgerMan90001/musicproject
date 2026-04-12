package randutil

import (
	"crypto/rand"
	"hash/fnv"

	"encoding/base64"
	mathrand "math/rand"
)

// Returns a random item from a list
func RandItem(list []string) string {
	return list[mathrand.Intn(len(list))]
}

func FNV(s string) uint64 {
	hash := fnv.New64a()
	hash.Write([]byte(s))
	return hash.Sum64()
}

// Generate random strings
type String interface {
	New(length int) string
}

type RandRead struct {
}

func (s *RandRead) New(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
