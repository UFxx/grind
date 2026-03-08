package http

type (
	TelegramAuthRequest struct {
		InviteCode string `json:"invite_code"`
	}

	TokenResponse struct {
		Token string `json:"token"`
	}
)
