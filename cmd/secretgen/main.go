package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"

	"songsled.com/internal/services/file"
	"songsled.com/pkg/crand"
)

func main() {
	var (
		envVar   string
		filename string
	)
	flag.StringVar(&envVar, "envVar", "", "update environment variable in .env file")
	flag.StringVar(&filename, "filename", "", "")
	flag.Parse()

	s := crand.NewB64Url(64)

	switch filename {
	case "":

		fmt.Println(s)
	default:
		fs := file.NewFileSystem()
		err := fs.CreateObject(context.Background(), "", filename, []byte(s), false, "")
		if err != nil {
			panic(err)
		}
	}
	// crand.NewB64()
	fmt.Printf("Short: %s\n")
	base64.URLEncoding.EncodeToString([]byte("file.png"))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("file.png")))
}
