package repository

import (
	"log"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/totoro081295/daily-report-api/refreshtoken"
	"github.com/totoro081295/daily-report-api/status"
)

type refreshTokenRepository struct {
	Conn *gorm.DB
}

// NewRefreshTokenRepository mount refresh_token repository
func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{
		Conn: db,
	}
}

// RefreshTokenRepository repository interface
type RefreshTokenRepository interface {
	GetByAccountID(id uuid.UUID) (*refreshtoken.RefreshToken, error)
}

func (m *refreshTokenRepository) GetByAccountID(id uuid.UUID) (*refreshtoken.RefreshToken, error) {
	var token refreshtoken.RefreshToken
	err := m.Conn.Model(&token).Where("account_id = ?", id).Find(&token).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &token, nil
}
