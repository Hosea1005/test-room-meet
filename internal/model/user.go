package model

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(300)" json:"name"`
	Email     string    `gorm:"type:varchar(300);unique" json:"username"`
	Role      string    `gorm:"type:varchar(300);" json:"role"`
	Password  string    `gorm:"type:varchar(300)" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
