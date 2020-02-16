package usecase

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/totoro081295/daily-report-api/dailycontent"

	"github.com/totoro081295/daily-report-api/dailycontent/repository"
)

type dailyContentUsecase struct {
	dailyContentRepo repository.DailyContentRepository
}

// NewDailyContentUsecase mount dailyContent usecase
func NewDailyContentUsecase(dailyContet repository.DailyContentRepository) DailyContentUsecase {
	return &dailyContentUsecase{
		dailyContentRepo: dailyContet,
	}
}

// DailyContentUsecase usecase interface
type DailyContentUsecase interface {
	GetByTargetDate(date time.Time) (*dailycontent.Response, error)
}

func (d *dailyContentUsecase) GetByTargetDate(date time.Time) (*dailycontent.Response, error) {
	pp.Println("usecase", date)
	dailyContent, err := d.dailyContentRepo.GetByTargetDate(date)
	if err != nil {
		return nil, err
	}
	res := dailycontent.Response{
		ID:   dailyContent.ID,
		Text: dailyContent.Text,
	}
	return &res, nil
}
