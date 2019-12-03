package account

import "github.com/gofrs/uuid"

// Account account struct
type Account struct {
	ID    uuid.UUID `json:"id" gorm:"primary_key" sql:"type:uuid" name:"id"`
	Email string    `json:"email" name:"メールアドレス"`
}
