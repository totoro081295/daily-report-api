package category

import (
	"time"

	"github.com/gofrs/uuid"
)

// Category category struct
type Category struct {
	ID          uuid.UUID  `json:"id" sql:"type:uuid" gorm:"primary_key"`
	Name        string     `json:"name"`
	Color       string     `json:"color"`
	Description string     `json:"description" xml:"description" name:"説明"`
	CreatedBy   uuid.UUID  `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy   uuid.UUID  `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt   time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt   time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt   *time.Time `json:"deletedAt" name:"削除日"`
}

// Response category response
type Response struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	Description string    `json:"description"`
}
