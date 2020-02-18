package repository

import (
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	model "github.com/totoro081295/daily-report-api/dailycontent"
	"github.com/totoro081295/daily-report-api/status"
)

type dailyContentRepository struct {
	Conn *gorm.DB
}

// NewDailyContentRepository mount daily_content repository
func NewDailyContentRepository(db *gorm.DB) DailyContentRepository {
	return &dailyContentRepository{
		Conn: db,
	}
}

// DailyContentRepository repository interface
type DailyContentRepository interface {
	GetByID(id uuid.UUID) (*model.DailyContent, error)
	GetByTargetDate(targetDate time.Time) (*model.DailyContent, error)
	Create(dailyContent *model.DailyContent) (*model.DailyContent, error)
	Update(dailyContent *model.DailyContent) error
}

func (m *dailyContentRepository) GetByID(id uuid.UUID) (*model.DailyContent, error) {
	var dailyContent = model.DailyContent{}
	err := m.Conn.Model(&dailyContent).Where("id = ?", id).Find(&dailyContent).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &dailyContent, nil

}

func (m *dailyContentRepository) GetByTargetDate(targetDate time.Time) (*model.DailyContent, error) {
	var dailyContent = model.DailyContent{}
	err := m.Conn.Model(&dailyContent).Where("target_date = ?", targetDate).Find(&dailyContent).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &dailyContent, nil
}

func (m *dailyContentRepository) Create(dailyContent *model.DailyContent) (*model.DailyContent, error) {
	err := m.Conn.Create(dailyContent).Error
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return dailyContent, nil
}

func (m *dailyContentRepository) Update(dailyContent *model.DailyContent) error {
	err := m.Conn.Model(dailyContent).Update(dailyContent).Error
	if err != nil {
		return err
	}
	return nil
}
