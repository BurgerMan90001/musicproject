package random

import "math/rand"

/* Returns a random item from a list */
func RandItem(list []string) string {
	return list[rand.Intn(len(list))]
}
