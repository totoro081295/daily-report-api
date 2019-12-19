package project

import (
	"time"

	"github.com/gofrs/uuid"
)

// Project project struct
type Project struct {
	ID          uuid.UUID  `json:"id" xml:"id" gorm:"primary_key" sql:"type:uuid" name:"id"`
	Name        string     `json:"name" xml:"name" name:"name"`
	Description string     `json:"description" xml:"description" name:"説明"`
	CreatedBy   uuid.UUID  `json:"createdBy" name:"作成者" sql:"type:uuid"`
	UpdatedBy   uuid.UUID  `json:"updatedBy" name:"更新者" sql:"type:uuid"`
	CreatedAt   time.Time  `json:"createdAt" name:"作成日"`
	UpdatedAt   time.Time  `json:"updatedAt" name:"更新日"`
	DeletedAt   *time.Time `json:"deletedAt" name:"削除日"`
}

// Response project response struct
type Response struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// Payload project payload struct
type Payload struct {
	ID          *uuid.UUID `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	UpdatedBy   uuid.UUID
}
