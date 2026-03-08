package models

import (
	"time"

	"github.com/google/uuid"
)

type Skill struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
