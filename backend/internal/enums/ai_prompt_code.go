package enums

type AIPromptCode = string

const (
	aiPromptSelectActivityCategory   AIPromptCode = "select_activity_category"
	aiPromptEvaluateActivity         AIPromptCode = "activity_evaluation"
	aiPromptSkillWeightsDistribution AIPromptCode = "skill_weights_distribution"
)

var AIPromptCodes = struct {
	SelectActivityCategory   AIPromptCode
	EvaluateActivity         AIPromptCode
	SkillWeightsDistribution AIPromptCode
}{
	SelectActivityCategory:   aiPromptSelectActivityCategory,
	EvaluateActivity:         aiPromptEvaluateActivity,
	SkillWeightsDistribution: aiPromptSkillWeightsDistribution,
}
