package main

import (
	"fmt"
	"log"
	"net/http"

	"movieexample.com/internal/user"
	"movieexample.com/internal/handler"
	"movieexample.com/internal/util/fileutil"
)

func main() {

	cfg, err := fileutil.ReadYAML[serviceConfig]("base.yml")
	if err != nil {
		panic(err)
	}
	port := cfg.APIConfig.Port

	repo := newRepository(cfg.RepositoryConfig)

	//authController := auth.New(repo, []byte(cfg.APIConfig.JWTKey))
	userController := user.New(repo)

	handler := handler.New(nil, userController)

	mux := http.NewServeMux()

	handler.Register(mux)

	host := "localhost"

	// start server
	log.Printf("Server listening at %v:%d", host, port)
	if err := http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), mux); err != nil {
		panic(err)
	}

}
