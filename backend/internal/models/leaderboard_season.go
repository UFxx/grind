package models

import (
	"time"

	"github.com/google/uuid"
)

type LeaderboardSeason struct {
	ID          uuid.UUID  `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	PeriodStart *time.Time `gorm:"column:period_start" json:"period_start"`
	PeriodEnd   *time.Time `gorm:"column:period_end" json:"period_end"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
}
