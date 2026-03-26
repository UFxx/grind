package models

import (
	"time"

	"github.com/google/uuid"
)

type AiPrompt struct {
	ID           uuid.UUID `gorm:"column:id;default:gen_random_uuid();primaryKey" json:"id"`
	Code         string    `gorm:"column:code" json:"code"`
	SystemPrompt string    `gorm:"column:system_prompt" json:"system_prompt"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}
