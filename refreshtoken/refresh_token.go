package refreshtoken

import (
	"time"

	"github.com/gofrs/uuid"
)

// RefreshToken refresh token struct
type RefreshToken struct {
	ID           uuid.UUID `form:"id" json:"id" xml:"id" gorm:"primary_key" sql:"type:uuid" name:"ID"`
	AccountID    uuid.UUID `form:"accountId" xml:"accountId" json:"accountId" sql:"type:uuid" name:"アカウントID"`
	RefreshToken string    `form:"refreshToken" json:"refreshToken" name:"リフレッシュトークン"`
	Expired      time.Time `form:"expired" json:"expired" name:"expired"`
	CreatedAt    time.Time `json:"createdAt" name:"作成日"`
}
