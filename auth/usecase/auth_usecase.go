package usecase

import (
	"github.com/k0kubun/pp"
	accountRepo "github.com/totoro081295/daily-report-api/account/repository"
	"github.com/totoro081295/daily-report-api/auth"
)

type authUsecase struct {
	accountRepo accountRepo.AccountRepository
}

// NewAuthUsecase mount auth usecase
func NewAuthUsecase(
	accoount accountRepo.AccountRepository,
) AuthUsecase {
	return &authUsecase{
		accountRepo: accoount,
	}
}

// AuthUsecase usecase interface
type AuthUsecase interface{}

func (a *authUsecase) Login(l auth.Login) (*auth.Token, error) {
	account, err := a.accountRepo.GetByEmail(l.Email)
	if err != nil {
		return nil, err
	}
	pp.Println(account)
	return nil, nil
}
