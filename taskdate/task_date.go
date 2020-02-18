package taskdate

import (
	"time"

	"github.com/gofrs/uuid"
)

// TaskDate task_date struct
type TaskDate struct {
	TaskID     uuid.UUID  `json:"taskId" sql:"type:uuid"`
	TargetDate time.Time  `json:"targetDate"`
	CreatedBy  uuid.UUID  `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy  uuid.UUID  `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt  time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt  time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt  *time.Time `json:"deletedAt" name:"削除日"`
}

// TableName overrides the table name setting in Gorm to force a specific table name in database
func (t TaskDate) TableName() string {
	return "task_date"
}
