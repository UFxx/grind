package http

type (
	SendFeedbackRequest struct {
		Sender  string `json:"sender" validate:"required"`
		Rating  int    `json:"rating" validate:"required,min=1,max=5"`
		Message string `json:"message"`
	}
)
