-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS rewards(
	event_id UUID NOT NULL
		REFERENCES events(id) ON DELETE CASCADE,
	skill_id UUID NOT NULL
		REFERENCES skills(id) ON DELETE CASCADE,
	xp_amount INTEGER NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	PRIMARY KEY (event_id, skill_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rewards;
-- +goose StatementEnd
