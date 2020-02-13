package auth

import (
	"errors"
	"regexp"

	"github.com/gofrs/uuid"
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

// LogoutValidate validate logout
func LogoutValidate(logout *Logout) error {
	// 必須チェック
	if logout.ID == uuid.FromStringOrNil("") {
		return errors.New("id is required")
	}
	return nil
}
