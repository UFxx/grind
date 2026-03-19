package enums

type ActivityCreateMode = string

const (
	activityCreateModeManually ActivityCreateMode = "manually"
	activityCreateModeAI       ActivityCreateMode = "ai"
)

var ActivityCreateModes = struct {
	Manually ActivityCreateMode
	AI       ActivityCreateMode
}{
	Manually: activityCreateModeManually,
	AI:       activityCreateModeAI,
}
