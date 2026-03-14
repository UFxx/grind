package models

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID             uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Description    string    `gorm:"column:description" json:"description"`
	UserID         uuid.UUID `gorm:"column:user_id" json:"user_id"`
	Source         string    `gorm:"column:source" json:"source"`
	ActivityTypeID uuid.UUID `gorm:"column:activity_type_id" json:"activity_type_id"`
	HasImpact      bool      `gorm:"column:has_impact" json:"has_impact"`
	IsHard         bool      `gorm:"column:is_hard" json:"is_hard"`
	IsNew          bool      `gorm:"column:is_new" json:"is_new"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`

	ActivityType ActivityType     `gorm:"foreignKey:ActivityTypeID" json:"activity_type"`
	Rewards      []ActivityReward `gorm:"foreignKey:ActivityID" json:"rewards"`
}
