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
		Nickname   string    `json:"nickname"`
		AvatarURL  string    `json:"avatar_url"`
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
		ID    uuid.UUID `json:"id"`
		Title string    `json:"title"`
		SkillProgress
	}

	RootSkill struct {
		ID    uuid.UUID `json:"id"`
		Title string    `json:"title"`
		SkillProgress
		Items []Subskill `json:"items"`
	}

	ManuallyCreatedEventRequest struct {
		Mode         string              `json:"mode" validate:"required"`
		Title        string              `json:"title" validate:"required"`
		EventTypeID  uuid.UUID           `json:"event_type_id" validate:"required"`
		HasImpact    bool                `json:"has_impact"`
		IsNew        bool                `json:"is_new"`
		IsHard       bool                `json:"is_hard"`
		SkillWeights []skill.SkillWeight `json:"skill_weights" validate:"dive"`
	}

	GetSkillsResponseItem struct {
		ID    uuid.UUID `json:"id"`
		Title string    `json:"title"`
	}
)
