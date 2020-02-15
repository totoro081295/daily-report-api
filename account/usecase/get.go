package usecase

import (
	"github.com/gofrs/uuid"
	model "github.com/totoro081295/daily-report-api/account"
)

func (a *accountUsecase) Get(id uuid.UUID) (*model.Response, error) {
	account, err := a.accountRepo.Get(id)
	if err != nil {
		return nil, err
	}
	res := model.Response{
		ID:   account.ID,
		Name: account.Name,
	}
	return &res, nil
}
