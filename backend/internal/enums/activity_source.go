package enums

type ActivitySource = string

const (
	activitySourceManual ActivitySource = "manual"
)

var ActivitySources = struct {
	Manual ActivitySource
}{
	Manual: activitySourceManual,
}
