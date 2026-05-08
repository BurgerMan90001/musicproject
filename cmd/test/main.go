package main

import (
	"fmt"
	"net/url"
)

func main() {
	objectUrl, err := url.JoinPath("storage.songsled.com", "songsled", "myfile.png")
	if err != nil {
		panic(err)
	}

	fmt.Println(objectUrl)
}
