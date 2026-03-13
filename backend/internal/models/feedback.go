package models

import (
	"time"

	"github.com/google/uuid"
)

type Feedback struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Sender    string    `gorm:"column:sender" json:"sender"`
	Rating    int       `gorm:"column:rating" json:"rating"`
	Message   string    `gorm:"column:message" json:"message"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
