package models

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID                 uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Description        string    `gorm:"column:description" json:"description"`
	UserID             uuid.UUID `gorm:"column:user_id" json:"user_id"`
	Source             string    `gorm:"column:source" json:"source"`
	ActivityCategoryID uuid.UUID `gorm:"column:activity_category_id" json:"activity_category_id"`
	HasImpact          bool      `gorm:"column:has_impact" json:"has_impact"`
	IsHard             bool      `gorm:"column:is_hard" json:"is_hard"`
	IsNew              bool      `gorm:"column:is_new" json:"is_new"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at" json:"updated_at"`

	ActivityCategory ActivityCategory `gorm:"foreignKey:ActivityCategoryID" json:"activity_category"`
	Rewards          []ActivityReward `gorm:"foreignKey:ActivityID" json:"rewards"`
}
