package http

import "github.com/google/uuid"

type (
	GetActivityTypesResponseItem struct {
		ID   uuid.UUID `json:"id"`
		Code string    `json:"code"`
		Name string    `json:"name"`
	}
)
