package http

type (
	OtherError struct {
		Other string `json:"other"`
	}

	ErrorResp[T OtherError | map[string]string] struct {
		Errors T `json:"errors"`
	}

	SuccessDataResponse struct {
		Data any `json:"data"`
	}

	PaginationResponse struct {
		TotalPages  int `json:"total_pages"`
		CurrentPage int `json:"current_page"`
		NextPage    int `json:"next_page"`
		Limit       int `json:"limit"`
	}

	PaginationRequest struct {
		Page  int `form:"page" validate:"required,min=1"`
		Limit int `form:"limit" validate:"required,min=1,max=100"`
	}
)
