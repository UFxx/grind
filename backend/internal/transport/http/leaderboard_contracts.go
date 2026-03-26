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

	MyLeaderboardEntry struct {
		Position        *int   `json:"position"`
		Name            string `json:"name"`
		AvatarURL       string `json:"avatar_url"`
		Score           *int   `json:"score"`
		IsInLeaderboard bool   `json:"is_in_leaderboard"`
	}

	LeaderboardEntry struct {
		Position  int    `json:"position"`
		Name      string `json:"name"`
		AvatarURL string `json:"avatar_url"`
		Score     int    `json:"score"`
	}

	LeaderboardSeasonDetailResponse struct {
		Season  LeaderboardSeasonResponseItem `json:"season"`
		CTA     string                        `json:"cta"`
		Entries []LeaderboardEntry            `json:"entries"`
		Me      MyLeaderboardEntry            `json:"me"`
	}
)
