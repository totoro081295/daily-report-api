package usecase

import (
	model "github.com/totoro081295/daily-report-api/category"
	"github.com/totoro081295/daily-report-api/category/repository"
)

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

// NewCategoryUsecase mount category usecase
func NewCategoryUsecase(category repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: category,
	}
}

// CategoryUsecase usecase interface
type CategoryUsecase interface {
	List() ([]*model.Response, error)
}
