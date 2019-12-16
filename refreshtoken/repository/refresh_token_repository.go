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
	Create(token *refreshtoken.RefreshToken) error
	Delete(token string) error
}

func (m *refreshTokenRepository) GetByAccountID(id uuid.UUID) (*refreshtoken.RefreshToken, error) {
	var rToken refreshtoken.RefreshToken
	err := m.Conn.Model(&rToken).Where("account_id = ?", id).Find(&rToken).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &rToken, nil
}

func (m *refreshTokenRepository) Create(token *refreshtoken.RefreshToken) error {
	err := m.Conn.Create(token).Error
	if err != nil {
		log.Println(err)
		return errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return nil
}

func (m *refreshTokenRepository) Delete(token string) error {
	var rToken refreshtoken.RefreshToken
	err := m.Conn.Model(&rToken).Where("refresh_token = ?", token).Delete(&rToken).Error
	if err != nil {
		log.Println(err)
		return errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return nil
}
