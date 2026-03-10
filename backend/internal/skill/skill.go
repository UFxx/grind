package skill

import "math"

type Progress struct {
	CurrentLevel        int `json:"current_level"`
	NextLevel           int `json:"next_level"`
	TotalXP             int `json:"total_xp"`
	CurrentLevelStartXP int `json:"current_level_start_xp"`
	NextLevelStartXP    int `json:"next_level_start_xp"`
	XPToNextLevel       int `json:"xp_to_next_level"`
}

const (
	defaultStartCost             float64 = 100
	defaultAdditionalCoefficient float64 = 1.13
)

type SkillService struct {
	startCost             float64
	additionalCoefficient float64
}

func NewSkillService(startCost float64, additionalCoefficient float64) *SkillService {

	if startCost < 1 {
		startCost = defaultStartCost
	}

	if additionalCoefficient <= 1 {
		additionalCoefficient = defaultAdditionalCoefficient
	}

	return &SkillService{
		startCost:             startCost,
		additionalCoefficient: additionalCoefficient,
	}
}

func (service *SkillService) xpForLevel(level int) float64 {

	if level <= 1 {
		return 0
	}

	xp := service.startCost * math.Pow(service.additionalCoefficient, float64(level-2))

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
