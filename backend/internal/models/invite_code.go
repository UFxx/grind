package models

import (
	"time"

	"github.com/google/uuid"
)

type InviteCode struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Code      string    `gorm:"column:code" json:"code"`
	CreatedBy uuid.UUID `gorm:"column:created_by" json:"created_by"`
	Uses      int       `gorm:"column:uses" json:"uses"`
	MaxUses   int       `gorm:"column:max_uses" json:"max_uses"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
