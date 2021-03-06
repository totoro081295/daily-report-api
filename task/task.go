package task

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/totoro081295/daily-report-api/category"
	"github.com/totoro081295/daily-report-api/project"
)

// Task task struct
type Task struct {
	ID          uuid.UUID  `json:"id" sql:"type:uuid" gorm:"primary_key"`
	Name        string     `json:"name"`
	AccountID   uuid.UUID  `json:"accountId" sql:"type:uuid"`
	CategoryID  uuid.UUID  `json:"categoryId" sql:"type:uuid"`
	ProjectID   uuid.UUID  `json:"projectId" sql:"type:uuid"`
	Cost        int        `json:"cost"`
	Description string     `json:"description" xml:"description" name:"説明"`
	Problem     string     `json:"problem"`
	Done        bool       `json:"isDone"`
	CreatedBy   uuid.UUID  `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy   uuid.UUID  `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt   time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt   time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt   *time.Time `json:"deletedAt" name:"削除日"`
}

// Collection task arrat
type Collection []*Task

// GetIDs idの配列を返す
func (t Collection) GetIDs() []uuid.UUID {
	IDs := make([]uuid.UUID, len(t))
	for i := 0; i < len(t); i++ {
		IDs[i] = t[i].ID
	}
	return IDs
}

// GetCategoryIDs idの配列を返す
func (t Collection) GetCategoryIDs() []uuid.UUID {
	categoryIDs := make([]uuid.UUID, len(t))
	for i := 0; i < len(t); i++ {
		categoryIDs[i] = t[i].CategoryID
	}
	return categoryIDs
}

// GetProjectIDs idの配列を返す
func (t Collection) GetProjectIDs() []uuid.UUID {
	projectIDs := make([]uuid.UUID, len(t))
	for i := 0; i < len(t); i++ {
		projectIDs[i] = t[i].ProjectID
	}
	return projectIDs
}

// Response task response
type Response struct {
	ID         uuid.UUID         `json:"id"`
	Name       string            `json:"name"`
	TargetDate time.Time         `json:"target_date"`
	Category   category.Response `json:"category"`
	Project    project.Response  `json:"project"`
}
