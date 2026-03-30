package main

import (
	"log"

	"musicproject.com/internal/services/auth"
)

func main() {
	//log.Println(os.Getwd())
	p, err := auth.HashPassword("Dirtycash@123!")
	log.Println(p)
	log.Println(err)
	log.Println(auth.ComparePassword("Dirtycash@123!", "$2a$14$BnvtnMBd9iRf3/Y0B5uxUudj1dHseBePvjl0nw8Hb0qRgjyaGdhFu"))

	t := []string{"1231231", "asdadsasd", "asdadsasdasdf"}
	log.Println(t[1:])
	// fmt.Println(cfg.Oauth.Google.Scopes)
	//fmt.Println(time.Now().Unix())
	// fmt.Println(oauthCfg.ClientID)
	// fmt.Println(oauthCfg.ClientSecret)
	// fmt.Println(oauthCfg.RedirectURL)
	// fmt.Println(oauthCfg.Scopes)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
	// 	ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * -10)),
	// })
	// fmt.Println(token.Valid)
	//jwt.NewNumericDate()
}
