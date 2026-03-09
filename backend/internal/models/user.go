package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID     `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	TelegramID *int64        `gorm:"column:telegram_id" json:"telegram_id"`
	Nickname   string        `gorm:"column:nickname" json:"nickname"`
	InvitedBy  uuid.NullUUID `gorm:"column:invited_by" json:"invited_by"`
	CreatedAt  time.Time     `gorm:"column:created_at" json:"created_at"`

	Inviter     *User        `gorm:"foreignKey:InvitedBy;references:ID" json:"inviter"`
	Skills      []UserSkill  `gorm:"foreignKey:UserID;references:ID" json:"skills"`
	InviteCodes []InviteCode `gorm:"foreignKey:CreatedBy;references:ID" json:"invite-codes"`
}
