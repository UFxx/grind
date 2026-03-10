package http

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserProfile struct {
		ID         uuid.UUID `json:"id"`
		TelegramID *int64    `json:"telegram_id"`
		Nickname   string    `json:"nickname"`
		AvatarURL  string    `json:"avatar_url"`
		CreatedAt  time.Time `json:"created_at"`
	}

	GetUserProfileResponse struct {
		UserProfile
		Inviter *UserProfile `json:"inviter"`
	}

	UserInviteCode struct {
		ID        uuid.UUID `json:"id"`
		Code      string    `json:"code"`
		Uses      int       `json:"uses"`
		MaxUses   int       `json:"max_uses"`
		CreatedAt time.Time `json:"created_at"`
	}

	GetUserInviteCodesResponse = []UserInviteCode
)
