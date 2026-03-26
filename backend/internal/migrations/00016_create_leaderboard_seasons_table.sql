-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS leaderboard_seasons (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	period_start DATE,
	period_end DATE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS leaderboard_seasons;
-- +goose StatementEnd
