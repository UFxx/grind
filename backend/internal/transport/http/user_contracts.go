package http

import (
	"time"

	"github.com/google/uuid"
	"github.com/sunsetsavorer/grind/internal/skill"
)

type (
	Rank struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}

	Profile struct {
		ID         uuid.UUID     `json:"id"`
		TelegramID *int64        `json:"telegram_id"`
		Name       string        `json:"name"`
		Rank       Rank          `json:"rank"`
		AvatarURL  string        `json:"avatar_url"`
		Level      SkillProgress `json:"level"`
		CreatedAt  time.Time     `json:"created_at"`
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
		Code    string `json:"code" validate:"required,max=10,woSpaces"`
		MaxUses int    `json:"max_uses" validate:"required,min=1,max=100"`
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
		Mode               string              `json:"mode" validate:"required"`
		Description        string              `json:"description" validate:"required"`
		ActivityCategoryID uuid.UUID           `json:"activity_category_id" validate:"required"`
		HasImpact          bool                `json:"has_impact"`
		IsNew              bool                `json:"is_new"`
		IsHard             bool                `json:"is_hard"`
		SkillWeights       []skill.SkillWeight `json:"skill_weights" validate:"dive"`
	}

	SelectActivityCategoryResponse struct {
		Code string `json:"code"`
	}

	EvaluateActivityResponse struct {
		HasImpact bool `json:"has_impact"`
		IsNew     bool `json:"is_new"`
		IsHard    bool `json:"is_hard"`
	}

	SkillWeightsDistributionResponse struct {
		Weights []struct {
			Code   string  `json:"code"`
			Weight float64 `json:"weight"`
		} `json:"weights"`
	}

	AICreatedActivityRequest struct {
		Mode        string `json:"mode" validate:"required"`
		Description string `json:"description" validate:"required,min=3,max=150"`
	}

	CreateActivityDTO struct {
		Mode               string              `json:"mode" validate:"required"`
		Description        string              `json:"description" validate:"required"`
		ActivityCategoryID uuid.UUID           `json:"activity_category_id" validate:"required"`
		HasImpact          bool                `json:"has_impact"`
		IsNew              bool                `json:"is_new"`
		IsHard             bool                `json:"is_hard"`
		SkillWeights       []skill.SkillWeight `json:"skill_weights" validate:"required,dive"`
	}

	BaseCreateActivityRequest struct {
		Mode string `json:"mode" validate:"required"`
	}

	GetSkillsResponseItem struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}

	GetActivitiesRequest struct {
		PaginationRequest
	}

	ActivityCategory struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}

	ActivityReward struct {
		SkillID   uuid.UUID `json:"skill_id"`
		SkillName string    `json:"skill_name"`
		XPAmount  int       `json:"xp_amount"`
	}

	GetActivitiesResponseItem struct {
		ID               uuid.UUID        `json:"id"`
		Description      string           `json:"description"`
		ActivityCategory ActivityCategory `json:"activity_category"`
		HasImpact        bool             `json:"has_impact"`
		IsNew            bool             `json:"is_new"`
		IsHard           bool             `json:"is_hard"`
		Rewards          []ActivityReward `json:"rewards"`
		CreatedAt        time.Time        `json:"created_at"`
	}

	GetActivitiesResponse struct {
		PaginationResponse `json:"pagination"`
		Items              []GetActivitiesResponseItem `json:"items"`
	}
)
