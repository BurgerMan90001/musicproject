package crand

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"hash/fnv"
	mathrand "math/rand"
	"strings"
)

var escaper = strings.NewReplacer("9", "99", "-", "90", "_", "91")

// Returns a random item from a list
func RandItem(list []string) string {
	return list[mathrand.Intn(len(list))]
}

func NewFNV(s string) uint64 {
	hash := fnv.New64a()
	hash.Write([]byte(s))
	return hash.Sum64()
}

func NewHex(n int) string {
	// n * 2 characters
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Random string that is 8 characters long
func NewShort() string {
	b := make([]byte, 4)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewB64(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return escaper.Replace(base64.StdEncoding.WithPadding('d').EncodeToString(b))
}

// Usually used for urls and filenames
func NewB64Url(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return escaper.Replace(base64.RawURLEncoding.EncodeToString(b))
}
