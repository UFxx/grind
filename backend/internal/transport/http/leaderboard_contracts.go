package http

import (
	"time"

	"github.com/google/uuid"
)

type (
	LeaderboardSeasonResponseItem struct {
		ID          uuid.UUID  `json:"id"`
		Name        string     `json:"name"`
		PeriodStart *time.Time `json:"period_start"`
		PeriodEnd   *time.Time `json:"period_end"`
	}
)
