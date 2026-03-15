package http

import "github.com/google/uuid"

type (
	GetActivityCategoriesResponseItem struct {
		ID   uuid.UUID `json:"id"`
		Code string    `json:"code"`
		Name string    `json:"name"`
	}
)
