package repository

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/totoro081295/daily-report-api/dailycontent"
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
	GetByTargetDate(targetDate time.Time) (*dailycontent.DailyContent, error)
}

func (m *dailyContentRepository) GetByTargetDate(targetDate time.Time) (*dailycontent.DailyContent, error) {
	var dailyContent = dailycontent.DailyContent{}
	err := m.Conn.Model(&dailyContent).Where("target_date = ?", targetDate).Find(&dailyContent).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.Wrap(status.ErrNotFound, err.Error())
	} else if err != nil {
		log.Println(err)
		return nil, errors.Wrap(status.ErrInternalServer, err.Error())
	}
	return &dailyContent, nil
}
