package http

type (
	CreateInviteCodeRequest struct {
		Code    string `json:"code" validate:"required"`
		MaxUses int    `json:"max_uses" validate:"required,min=1"`
	}
)
