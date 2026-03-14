-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS activity_rewards(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	activity_id UUID NOT NULL
		REFERENCES activities(id) ON DELETE CASCADE,
	user_skill_id UUID NOT NULL
		REFERENCES user_skills(id) ON DELETE CASCADE,
	xp_amount INTEGER NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS activity_rewards;
-- +goose StatementEnd
