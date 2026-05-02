package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	s, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	fmt.Printf("uuid: %v\n", uuid.New().String())
	fmt.Printf("uuid v7: %v\n", s.String())
}
