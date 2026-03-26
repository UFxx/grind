package models

import (
	"time"

	"github.com/google/uuid"
)

type LeaderboardEntry struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	SeasonID  uuid.UUID `gorm:"column:season_id" json:"season_id"`
	UserID    uuid.UUID `gorm:"column:user_id" json:"user_id"`
	Score     int       `gorm:"column:score" json:"score"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
