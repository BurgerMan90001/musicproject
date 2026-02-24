package main

import (
	"fmt"
	"log"
	"net/http"

	"movieexample.com/internal/controller/auth"
	"movieexample.com/internal/controller/user"
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

	authController := auth.New(repo, []byte(cfg.APIConfig.JWTKey))
	userController := user.New(repo)

	handler := handler.New(authController, userController)

	mux := http.NewServeMux()

	handler.Register(mux)
	// start server
	log.Printf("Server listening at localhost:%d", port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), mux); err != nil {
		panic(err)
	}
}
