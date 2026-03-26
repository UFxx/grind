-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS leaderboard_entries(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	season_id UUID NOT NULL
		REFERENCES leaderboard_seasons(id) ON DELETE CASCADE,
	user_id UUID NOT NULL
		REFERENCES users(id) ON DELETE CASCADE,
	score INTEGER NOT NULL DEFAULT 0,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

	UNIQUE(season_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS leaderboard_entries;
-- +goose StatementEnd
