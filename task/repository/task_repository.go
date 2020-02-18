package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	model "github.com/totoro081295/daily-report-api/task"
)

type taskRepository struct {
	Conn *gorm.DB
}

// NewTaskRepository mount task repository
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		Conn: db,
	}
}

// TaskRepository repository interface
type TaskRepository interface {
	ListByAccountID(accountID uuid.UUID) (model.Collection, error)
}

func (m *taskRepository) ListByAccountID(accountID uuid.UUID) (model.Collection, error) {
	var tasks = model.Collection{}
	err := m.Conn.Model(&tasks).Where("account_id = ?", accountID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
