package usecase

import (
	"os"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/k0kubun/pp"
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
}

func (a *authUsecase) Login(l *auth.Login) (*auth.Token, error) {
	account, err := a.accountRepo.GetByEmail(l.Email)
	if err != nil {
		pp.Println("account, err := a.accountRepo.GetByEmail(l.Email)")
		return nil, err
	}
	refresh, err := a.rTokenRepo.GetByAccountID(account.ID)
	if refresh != nil {
		err := a.rTokenRepo.Delete(refresh.RefreshToken)
		if err != nil {
			pp.Println("err := a.rTokenRepo.Delete(refresh.RefreshToken)")
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
		pp.Println("err = a.rTokenRepo.Create(insert)")
		return nil, err
	}

	token, err := a.token.GenerateJWT(account.ID, false)
	if err != nil {
		pp.Println("token, err := a.token.GenerateJWT(account.ID, false)")
		return nil, err
	}

	res := auth.Token{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}
	return &res, nil
}
