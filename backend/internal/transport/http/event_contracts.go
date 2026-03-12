package http

import "github.com/google/uuid"

type (
	GetEventTypesResponseItem struct {
		ID    uuid.UUID `json:"id"`
		Title string    `json:"title"`
	}
)
