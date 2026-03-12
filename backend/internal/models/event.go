package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	UserID      uuid.UUID `gorm:"column:user_id" json:"user_id"`
	Source      string    `gorm:"column:source" json:"source"`
	EventTypeID uuid.UUID `gorm:"column:event_type_id" json:"event_type_id"`
	HasImpact   bool      `gorm:"column:has_impact" json:"has_impact"`
	IsHard      bool      `gorm:"column:is_hard" json:"is_hard"`
	IsNew       bool      `gorm:"column:is_new" json:"is_new"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`

	EventType EventType     `gorm:"foreignKey:EventTypeID" json:"event_type"`
	Rewards   []EventReward `gorm:"foreignKey:EventID" json:"rewards"`
}
