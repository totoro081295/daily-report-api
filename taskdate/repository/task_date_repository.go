package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/totoro081295/daily-report-api/taskdate"
)

type taskDateRepository struct {
	Conn *gorm.DB
}

// NewTaskDateRepository mount taskDate repository
func NewTaskDateRepository(db *gorm.DB) TaskDateRepository {
	return &taskDateRepository{
		Conn: db,
	}
}

// TaskDateRepository repository interface
type TaskDateRepository interface {
	ListByTaskIDs(ids []uuid.UUID) ([]*taskdate.TaskDate, error)
}

func (m *taskDateRepository) ListByTaskIDs(ids []uuid.UUID) ([]*taskdate.TaskDate, error) {
	var taskDates = []*taskdate.TaskDate{}
	err := m.Conn.Model(&taskDates).Where("task_id in (?)", ids).Find(&taskDates).Error
	if err != nil {
		return nil, err
	}
	return taskDates, nil
}
