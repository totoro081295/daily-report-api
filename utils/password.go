package utils

import "golang.org/x/crypto/bcrypt"

// CompareHashPassword パスワードの照合
func CompareHashPassword(hash string, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return nil
	}
	return err
}
