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
)
