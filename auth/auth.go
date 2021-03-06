package auth

import (
	"github.com/gofrs/uuid"
)

// Login login struct
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token token struct
type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// Logout logout struct
type Logout struct {
	ID uuid.UUID `json:"accountId"`
}

// Refresh refresh struct
type Refresh struct {
	AccountID    uuid.UUID `json:"accountId"`
	RefreshToken string    `json:"refreshToken"`
}
