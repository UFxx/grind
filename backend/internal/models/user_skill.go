package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSkill struct {
	ID            uuid.UUID     `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name          string        `gorm:"column:name" json:"name"`
	UserID        uuid.UUID     `gorm:"column:user_id" json:"user_id"`
	BaseSkillID   uuid.NullUUID `gorm:"column:base_skill_id" json:"base_skill_id"`
	ParentSkillID uuid.NullUUID `gorm:"column:parent_skill_id" json:"parent_skill_id"`
	TotalXP       int           `gorm:"column:total_xp" json:"total_xp"`
	CreatedAt     time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time     `gorm:"column:updated_at" json:"updated_at"`

	BaseSkill   *BaseSkill `gorm:"foreignKey:BaseSkillID;references:ID" json:"base_skill"`
	ParentSkill *UserSkill `gorm:"foreignKey:ParentSkillID;references:ID" json:"parent_skill"`
}
