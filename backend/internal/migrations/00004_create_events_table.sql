-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	title TEXT NOT NULL,
	user_id UUID NOT NULL
		REFERENCES users(id) ON DELETE CASCADE,
	source TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
