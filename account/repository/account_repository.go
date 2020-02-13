package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/totoro081295/daily-report-api/account"
	"github.com/totoro081295/daily-report-api/status"
)

type accountRepository struct {
	Conn *gorm.DB
}

// NewAccountRepository mount account repository
func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		Conn: db,
	}
}

// AccountRepository repository interface
type AccountRepository interface {
	GetByEmail(email string) (*account.Account, error)
}

func (m *accountRepository) GetByEmail(email string) (*account.Account, error) {
	var a account.Account
	err := m.Conn.Model(&a).Where("email = ?", email).Find(&a).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &a, nil
}
