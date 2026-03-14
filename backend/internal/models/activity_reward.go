package models

import (
	"time"

	"github.com/google/uuid"
)

type ActivityReward struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	ActivityID  uuid.UUID `gorm:"column:activity_id" json:"activity_id"`
	UserSkillID uuid.UUID `gorm:"column:user_skill_id" json:"user_skill_id"`
	XPAmount    int       `gorm:"column:xp_amount" json:"xp_amount"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`

	UserSkill UserSkill `gorm:"foreignKey:UserSkillID" json:"user_skill"`
	Activity  Activity  `gorm:"foreignKey:ActivityID" json:"activity"`
}
