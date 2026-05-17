package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ReplaceAll("\"name\":\"Cool music\",\"artists\":[\"rockguy\",\"rockguy2\"],\"genres\":[\"Rock\",\"Pop\"],\"creationDate\":\"2026-05-16\",\"audio\":\"mysong.mp3\"}", "\\", ""))
}
