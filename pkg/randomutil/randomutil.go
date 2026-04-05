package randomutil

import "math/rand"

/* Returns a random item from a list */
func RandomItem(list []string) string {
	return list[rand.Intn(len(list))]
}
