package usecase

import (
	"github.com/gofrs/uuid"
	model "github.com/totoro081295/daily-report-api/account"
	"github.com/totoro081295/daily-report-api/account/repository"
)

type accountUsecase struct {
	accountRepo repository.AccountRepository
}

// NewAccountUsecase mount account usecase
func NewAccountUsecase(account repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepo: account,
	}
}

// AccountUsecase usecase interface
type AccountUsecase interface {
	Get(id uuid.UUID) (*model.Response, error)
}
