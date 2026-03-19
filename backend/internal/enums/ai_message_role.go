package enums

type AIMessageRole = string

const (
	aiMessageRoleSystem AIMessageRole = "system"
	aiMessageRoleUser   AIMessageRole = "user"
)

var AIMessageRoles = struct {
	System AIMessageRole
	User   AIMessageRole
}{
	System: aiMessageRoleSystem,
	User:   aiMessageRoleUser,
}
