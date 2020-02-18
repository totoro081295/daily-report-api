package usecase

import (
	model "github.com/totoro081295/daily-report-api/category"
)

func (c *categoryUsecase) List() ([]*model.Response, error) {
	categories, err := c.categoryRepo.List()
	if err != nil {
		return nil, err
	}
	res := []*model.Response{}
	for _, category := range categories {
		response := model.Response{
			ID:          category.ID,
			Name:        category.Name,
			Color:       category.Color,
			Description: category.Description,
		}
		res = append(res, &response)
	}
	return res, nil
}
