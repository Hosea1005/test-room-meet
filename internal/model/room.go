package model

import (
	"github.com/google/uuid"
	"time"
)

type Room struct {
	ID        uuid.UUID  `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(100)" json:"name"`
	Location  string     `gorm:"type:varchar(20)" json:"location"`
	Status    string     `gorm:"type:varchar(20)" json:"status"`
	Capacity  int64      `gorm:"type:float" json:"capacity"`
	CreatedAt *time.Time ` json:"created_at"`
	UpdatedAt *time.Time ` json:"updated_at"`
	DeletedAt *time.Time ` json:"deleted_at_at"`
}
