package http

import (
	"time"

	"github.com/google/uuid"
)

type (
	Profile struct {
		ID         uuid.UUID `json:"id"`
		TelegramID *int64    `json:"telegram_id"`
		Nickname   string    `json:"nickname"`
		AvatarURL  string    `json:"avatar_url"`
		CreatedAt  time.Time `json:"created_at"`
	}

	GetProfileResponse struct {
		Profile
		Inviter *Profile `json:"inviter"`
	}

	UserInviteCode struct {
		ID        uuid.UUID `json:"id"`
		Code      string    `json:"code"`
		Uses      int       `json:"uses"`
		MaxUses   int       `json:"max_uses"`
		CreatedAt time.Time `json:"created_at"`
	}

	GetInviteCodesResponse = []UserInviteCode

	CreateInviteCodeRequest struct {
		Code    string `json:"code" validate:"required"`
		MaxUses int    `json:"max_uses" validate:"required,min=1"`
	}
)
