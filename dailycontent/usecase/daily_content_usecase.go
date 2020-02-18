package usecase

import (
	"time"

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
	Create(payload *dailycontent.CreatePayload) (*dailycontent.CreateResponse, error)
}
