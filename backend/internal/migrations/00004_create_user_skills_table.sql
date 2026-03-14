-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_skills(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	user_id UUID NOT NULL
		REFERENCES users(id) ON DELETE CASCADE,
	base_skill_id UUID
		REFERENCES base_skills(id) ON DELETE CASCADE,
	parent_skill_id UUID
		REFERENCES user_skills(id) ON DELETE CASCADE,
	total_xp INTEGER NOT NULL DEFAULT 0,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_skills;
-- +goose StatementEnd
