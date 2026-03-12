package models

import (
	"time"

	"github.com/google/uuid"
)

type EventType struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
