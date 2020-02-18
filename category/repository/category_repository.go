package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	model "github.com/totoro081295/daily-report-api/category"
)

type categoryRepository struct {
	Conn *gorm.DB
}

// NewCategoryRepository mount category repository
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		Conn: db,
	}
}

// CategoryRepository repository interface
type CategoryRepository interface {
	List() ([]*model.Category, error)
	ListByIDs(ids []uuid.UUID) ([]*model.Category, error)
}

func (m *categoryRepository) List() ([]*model.Category, error) {
	var categories = []*model.Category{}
	err := m.Conn.Model(&categories).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (m *categoryRepository) ListByIDs(ids []uuid.UUID) ([]*model.Category, error) {
	var categories = []*model.Category{}
	err := m.Conn.Model(&categories).Where("id in (?)", ids).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
