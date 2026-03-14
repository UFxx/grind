package http

import "github.com/google/uuid"

type (
	GetActivityTypesResponseItem struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}
)
