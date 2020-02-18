package dailycontent

import (
	"time"

	"github.com/gofrs/uuid"
)

// DailyContent struct daily contents
type DailyContent struct {
	ID         uuid.UUID  `json:"id" gorm:"primary_key" sql:"type:uuid" name:"id"`
	TargetDate time.Time  `json:"targetDate" name:"対象日"`
	Text       string     `json:"text" name:"内容"`
	CreatedBy  uuid.UUID  `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy  uuid.UUID  `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt  time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt  time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt  *time.Time `json:"deletedAt" name:"削除日"`
}

// Response struct daily content response
type Response struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

// CreatePayload struct daily content payload
type CreatePayload struct {
	Text       string    `json:"text"`
	TargetDate time.Time `json:"targetDate"`
	CreatedBy  uuid.UUID
}

// CreateResponse struct daily content create response
type CreateResponse struct {
	ID         uuid.UUID `json:"id"`
	Text       string    `json:"text"`
	TargetDate time.Time `json:"targetDate"`
}
