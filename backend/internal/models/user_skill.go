package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSkill struct {
	UserID    uuid.UUID `gorm:"column:user_id" json:"user_id"`
	SkillID   uuid.UUID `gorm:"column:skill_id" json:"skill_id"`
	TotalXP   int       `gorm:"column:total_xp" json:"total_xp"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	Skill Skill `gorm:"foreignKey:SkillID;references:ID" json:"skill"`
}
