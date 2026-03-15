package skill

import (
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/sunsetsavorer/grind/internal/config"
	"github.com/sunsetsavorer/grind/internal/exceptions"
)

type (
	Progress struct {
		CurrentLevel        int `json:"current_level"`
		NextLevel           int `json:"next_level"`
		TotalXP             int `json:"total_xp"`
		CurrentLevelStartXP int `json:"current_level_start_xp"`
		NextLevelStartXP    int `json:"next_level_start_xp"`
		XPToNextLevel       int `json:"xp_to_next_level"`
	}

	SkillWeight struct {
		SkillID uuid.UUID `json:"skill_id" validate:"required"`
		Weight  float64   `json:"weight" validate:"required,min=0,max=1"`
	}

	SkillReward struct {
		SkillID  uuid.UUID `json:"skill_id"`
		XPAmount int       `json:"xp_amount"`
	}

	ActivityReward struct {
		TotalXP      int           `json:"total_xp"`
		SkillRewards []SkillReward `json:"skill_rewards"`
	}
)

const (
	defaultStartCost             float64 = 100
	defaultAdditionalCoefficient float64 = 1.13
	defaultBaseXP                float64 = 30
	defaultImpactMultiplier      float64 = 1.4
	defaultNewMultiplier         float64 = 1.35
	defaultHardMultiplier        float64 = 1.3
)

type SkillService struct {
	startLevelCost        float64
	baseXP                float64
	additionalCoefficient float64
	impactMultiplier      float64
	newMultiplier         float64
	hardMultiplier        float64
}

func NewSkillService(skillConfig config.SkillConfig) *SkillService {

	skillService := &SkillService{
		startLevelCost:        skillConfig.StartLevelCost,
		baseXP:                skillConfig.BaseXP,
		additionalCoefficient: skillConfig.AdditionalCoefficient,
		impactMultiplier:      skillConfig.ImpactMultiplier,
		newMultiplier:         skillConfig.NewMultiplier,
		hardMultiplier:        skillConfig.HardMultiplier,
	}

	if skillConfig.StartLevelCost <= 0 {
		skillService.startLevelCost = defaultStartCost
	}

	if skillConfig.BaseXP <= 0 {
		skillService.baseXP = defaultBaseXP
	}

	if skillConfig.AdditionalCoefficient <= 1 {
		skillService.additionalCoefficient = defaultAdditionalCoefficient
	}

	if skillConfig.ImpactMultiplier <= 0 {
		skillService.impactMultiplier = defaultImpactMultiplier
	}

	if skillConfig.NewMultiplier <= 0 {
		skillService.newMultiplier = defaultNewMultiplier
	}

	if skillConfig.HardMultiplier <= 0 {
		skillService.hardMultiplier = defaultHardMultiplier
	}

	return skillService
}

func (service *SkillService) xpForLevel(level int) float64 {

	if level <= 1 {
		return 0
	}

	xp := service.startLevelCost * math.Pow(service.additionalCoefficient, float64(level-2))

	return xp
}

func (service *SkillService) CalcProgress(totalXP int) Progress {

	if totalXP < 0 {
		totalXP = 0
	}

	currentLevel := 1
	currentLevelStartXP := 0

	for {
		nextLevelStartXP := currentLevelStartXP + int(math.Round(service.xpForLevel(currentLevel+1)))

		if totalXP < nextLevelStartXP {
			return Progress{
				CurrentLevel:        currentLevel,
				NextLevel:           currentLevel + 1,
				TotalXP:             totalXP,
				CurrentLevelStartXP: currentLevelStartXP,
				NextLevelStartXP:    nextLevelStartXP,
				XPToNextLevel:       nextLevelStartXP - totalXP,
			}
		}

		currentLevel++
		currentLevelStartXP = nextLevelStartXP
	}
}

func (service *SkillService) CalcActivityReward(
	skillWeights []SkillWeight,
	hasImpact bool,
	isHard bool,
	isNew bool,
) (ActivityReward, error) {

	if len(skillWeights) == 0 {
		return ActivityReward{}, exceptions.NewServiceError(fmt.Errorf("no skill weights provided"))
	}

	// check weights sum & duplicates
	skillIDsMap := make(map[uuid.UUID]struct{})
	var skillWeightSum float64

	for _, skillWeight := range skillWeights {
		if _, exists := skillIDsMap[skillWeight.SkillID]; exists {
			return ActivityReward{}, exceptions.NewServiceError(fmt.Errorf("skill weights has duplicates"))
		}

		skillIDsMap[skillWeight.SkillID] = struct{}{}
		skillWeightSum += skillWeight.Weight
	}

	if math.Abs(skillWeightSum-1) > 1e-9 {
		return ActivityReward{}, exceptions.NewServiceError(fmt.Errorf("invalid skill weights"))
	}

	rawTotalXP := service.baseXP

	if hasImpact {
		rawTotalXP *= service.impactMultiplier
	}

	if isHard {
		rawTotalXP *= service.hardMultiplier
	}

	if isNew {
		rawTotalXP *= service.newMultiplier
	}

	skillRewards := make([]SkillReward, 0, len(skillWeights))
	var totalXP int

	for _, skillWeight := range skillWeights {
		skillXP := int(math.Round(rawTotalXP * skillWeight.Weight))

		skillRewards = append(skillRewards, SkillReward{
			SkillID:  skillWeight.SkillID,
			XPAmount: skillXP,
		})

		totalXP += skillXP
	}

	return ActivityReward{
		TotalXP:      totalXP,
		SkillRewards: skillRewards,
	}, nil
}
