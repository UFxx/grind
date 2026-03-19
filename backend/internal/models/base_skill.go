package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseSkill struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Code        string    `gorm:"column:code" json:"code"`
	DisplayName string    `gorm:"column:display_name" json:"display_name"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
