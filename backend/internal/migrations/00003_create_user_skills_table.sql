-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_skills(
	user_id UUID NOT NULL
		REFERENCES users(id) ON DELETE CASCADE,
	skill_id UUID NOT NULL
		REFERENCES skills(id) ON DELETE CASCADE,
	total_xp INTEGER NOT NULL DEFAULT 0,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	PRIMARY KEY (user_id, skill_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_skills;
-- +goose StatementEnd
