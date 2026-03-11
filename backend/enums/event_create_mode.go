package enums

type EventCreateMode = string

var (
	eventCreateModeManually EventCreateMode = "manually"
	eventCreateModeAI       EventCreateMode = "ai"
)

var EventCreateModes = struct {
	Manually EventCreateMode
	AI       EventCreateMode
}{
	Manually: eventCreateModeManually,
	AI:       eventCreateModeAI,
}
