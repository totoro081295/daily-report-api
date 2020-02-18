package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/totoro081295/daily-report-api/dailycontent"
)

func (d *dailyContentUsecase) Create(payload *dailycontent.CreatePayload) (*dailycontent.CreateResponse, error) {
	dailyContentID, _ := uuid.NewV4()
	dailyContent := dailycontent.DailyContent{
		ID:         dailyContentID,
		Text:       payload.Text,
		TargetDate: payload.TargetDate,
		CreatedBy:  payload.CreatedBy,
		UpdatedBy:  payload.CreatedBy,
	}
	_, err := d.dailyContentRepo.Create(&dailyContent)
	if err != nil {
		return nil, err
	}
	res := dailycontent.CreateResponse{
		ID:         dailyContent.ID,
		Text:       dailyContent.Text,
		TargetDate: dailyContent.TargetDate,
	}
	return &res, nil
}
