package models

import (
	"time"

	"github.com/google/uuid"
)

type EventReward struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	EventID     uuid.UUID `gorm:"column:event_id" json:"event_id"`
	UserSkillID uuid.UUID `gorm:"column:user_skill_id" json:"user_skill_id"`
	XPAmount    int       `gorm:"column:xp_amount" json:"xp_amount"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}
