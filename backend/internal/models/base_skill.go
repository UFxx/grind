package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseSkill struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
