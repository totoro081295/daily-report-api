package usecase

import "github.com/totoro081295/daily-report-api/dailycontent"

func (d *dailyContentUsecase) Update(payload dailycontent.UpdatePayload) error {
	existsDailyContent, err := d.dailyContentRepo.GetByID(payload.ID)
	if err != nil {
		return err
	}
	dailyContent := dailycontent.DailyContent{
		ID:         existsDailyContent.ID,
		TargetDate: existsDailyContent.TargetDate,
		Text:       payload.Text,
		CreatedBy:  existsDailyContent.CreatedBy,
		UpdatedBy:  payload.UpdatedBy,
	}
	err = d.dailyContentRepo.Update(&dailyContent)
	if err != nil {
		return err
	}
	return nil
}
