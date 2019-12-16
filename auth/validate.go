package auth

import (
	"errors"
	"regexp"
)

// LoginValidate validate login
func LoginValidate(login *Login) error {
	// 必須チェック
	if login.Email == "" || login.Password == "" {
		return errors.New("email or password is required")
	}

	// formatチェック
	if login.Password != "" {
		var passRegexp = regexp.MustCompile(`^[0-9a-zA-Z]+$`)
		ok := passRegexp.MatchString(login.Password)
		if !ok {
			return errors.New("invalid password")
		}
		if len([]byte(login.Password)) < 8 {
			return errors.New("password must be 8 length")
		}
	}
	return nil
}
