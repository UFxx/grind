package http

import (
	"time"

	"github.com/google/uuid"
	"github.com/sunsetsavorer/grind/internal/skill"
)

type (
	Profile struct {
		ID         uuid.UUID `json:"id"`
		TelegramID *int64    `json:"telegram_id"`
		Name       string    `json:"name"`
		AvatarURL  string    `json:"avatar_url"`
		Level      int       `json:"level"`
		CreatedAt  time.Time `json:"created_at"`
	}

	GetProfileResponse struct {
		Profile
		Inviter *Profile `json:"inviter"`
	}

	InviteCode struct {
		ID        uuid.UUID `json:"id"`
		Code      string    `json:"code"`
		Uses      int       `json:"uses"`
		MaxUses   int       `json:"max_uses"`
		CreatedAt time.Time `json:"created_at"`
	}

	CreateInviteCodeRequest struct {
		Code    string `json:"code" validate:"required"`
		MaxUses int    `json:"max_uses" validate:"required,min=1"`
	}

	SkillProgress struct {
		CurrentLevel        int `json:"current_level"`
		NextLevel           int `json:"next_level"`
		TotalXP             int `json:"total_xp"`
		CurrentLevelStartXP int `json:"current_level_start_xp"`
		NextLevelStartXP    int `json:"next_level_start_xp"`
		XPToNextLevel       int `json:"xp_to_next_level"`
	}

	Subskill struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		SkillProgress
	}

	RootSkill struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		SkillProgress
		Items []Subskill `json:"items"`
	}

	ManuallyCreatedActivityRequest struct {
		Mode           string              `json:"mode" validate:"required"`
		Description    string              `json:"description" validate:"required"`
		ActivityTypeID uuid.UUID           `json:"activity_type_id" validate:"required"`
		HasImpact      bool                `json:"has_impact"`
		IsNew          bool                `json:"is_new"`
		IsHard         bool                `json:"is_hard"`
		SkillWeights   []skill.SkillWeight `json:"skill_weights" validate:"dive"`
	}

	GetSkillsResponseItem struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}

	GetActivitiesRequest struct {
		PaginationRequest
	}

	ActivityType struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}

	ActivityReward struct {
		SkillID   uuid.UUID `json:"skill_id"`
		SkillName string    `json:"skill_name"`
		XPAmount  int       `json:"xp_amount"`
	}

	GetActivitiesResponseItem struct {
		ID           uuid.UUID        `json:"id"`
		Description  string           `json:"description"`
		ActivityType ActivityType     `json:"activity_type"`
		HasImpact    bool             `json:"has_impact"`
		IsNew        bool             `json:"is_new"`
		IsHard       bool             `json:"is_hard"`
		Rewards      []ActivityReward `json:"rewards"`
		CreatedAt    time.Time        `json:"created_at"`
	}

	GetActivitiesResponse struct {
		PaginationResponse `json:"pagination"`
		Items              []GetActivitiesResponseItem `json:"items"`
	}
)
