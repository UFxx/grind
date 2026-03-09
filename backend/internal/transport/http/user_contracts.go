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
		CreatedAt  time.Time `json:"created_at"`
	}

	GetUserProfileResponse struct {
		UserProfile
		Inviter *UserProfile `json:"inviter"`
	}
)
