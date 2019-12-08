package usecase

import (
	"github.com/k0kubun/pp"
	accountRepo "github.com/totoro081295/daily-report-api/account/repository"
	"github.com/totoro081295/daily-report-api/auth"
	rTokenRepo "github.com/totoro081295/daily-report-api/refreshtoken/repository"
)

type authUsecase struct {
	accountRepo accountRepo.AccountRepository
	rTokenRepo  rTokenRepo.RefreshTokenRepository
}

// NewAuthUsecase mount auth usecase
func NewAuthUsecase(
	accoount accountRepo.AccountRepository,
	rToken rTokenRepo.RefreshTokenRepository,
) AuthUsecase {
	return &authUsecase{
		accountRepo: accoount,
		rTokenRepo:  rToken,
	}
}

// AuthUsecase usecase interface
type AuthUsecase interface{}

func (a *authUsecase) Login(l auth.Login) (*auth.Token, error) {
	account, err := a.accountRepo.GetByEmail(l.Email)
	if err != nil {
		return nil, err
	}
	refresh, err := a.rTokenRepo.GetByAccountID(account.ID)
	if err != nil {
		return nil, err
	}
	pp.Println(refresh)
	return nil, nil
}
