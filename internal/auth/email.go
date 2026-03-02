package auth

import "regexp"

func ValidateEmail(email string) (bool, error) {
	return regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
}
