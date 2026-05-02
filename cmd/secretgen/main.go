package main

import (
	"context"
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
	fmt.Printf("Short: %s\n", crand.NewShort())
}
