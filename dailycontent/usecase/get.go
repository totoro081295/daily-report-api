package usecase

import (
	"time"

	"github.com/totoro081295/daily-report-api/dailycontent"
)

func (d *dailyContentUsecase) GetByTargetDate(date time.Time) (*dailycontent.Response, error) {
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
