package usecase

import (
	"os"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	accountRepo "github.com/totoro081295/daily-report-api/account/repository"
	"github.com/totoro081295/daily-report-api/auth"
	"github.com/totoro081295/daily-report-api/refreshtoken"
	rTokenRepo "github.com/totoro081295/daily-report-api/refreshtoken/repository"
	"github.com/totoro081295/daily-report-api/token"
)

type authUsecase struct {
	accountRepo accountRepo.AccountRepository
	rTokenRepo  rTokenRepo.RefreshTokenRepository
	token       token.Handler
}

// NewAuthUsecase mount auth usecase
func NewAuthUsecase(
	accoount accountRepo.AccountRepository,
	rToken rTokenRepo.RefreshTokenRepository,
	tokenHandler token.Handler,
) AuthUsecase {
	return &authUsecase{
		accountRepo: accoount,
		rTokenRepo:  rToken,
		token:       tokenHandler,
	}
}

// AuthUsecase usecase interface
type AuthUsecase interface {
	Login(l *auth.Login) (*auth.Token, error)
	Logout(id uuid.UUID) error
}

func (a *authUsecase) Login(l *auth.Login) (*auth.Token, error) {
	account, err := a.accountRepo.GetByEmail(l.Email)
	if err != nil {
		return nil, err
	}
	refresh, err := a.rTokenRepo.GetByAccountID(account.ID)
	if refresh != nil {
		err := a.rTokenRepo.Delete(refresh.RefreshToken)
		if err != nil {
			return nil, err
		}
	}
	refreshToken := a.token.RandToken()
	// 文字列を数値に
	atoi, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRES_MIN"))
	insertID, _ := uuid.NewV4()
	insert := &refreshtoken.RefreshToken{
		ID:           insertID,
		AccountID:    account.ID,
		RefreshToken: refreshToken,
		Expired:      time.Now().Add(time.Minute * time.Duration(atoi)),
	}

	err = a.rTokenRepo.Create(insert)
	if err != nil {
		return nil, err
	}

	token, err := a.token.GenerateJWT(account.ID, false)
	if err != nil {
		return nil, err
	}

	res := auth.Token{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}
	return &res, nil
}

func (a *authUsecase) Logout(id uuid.UUID) error {
	refresh, err := a.rTokenRepo.GetByAccountID(id)
	if err != nil {
		return err
	}
	err = a.rTokenRepo.Delete(refresh.RefreshToken)
	if err != nil {
		return err
	}
	return nil
}
