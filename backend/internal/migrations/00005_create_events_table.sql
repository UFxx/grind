-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	title TEXT NOT NULL,
	user_id UUID NOT NULL
		REFERENCES users(id) ON DELETE CASCADE,
	source TEXT NOT NULL,
	event_type_id UUID NOT NULL
		REFERENCES event_types(id) ON DELETE CASCADE,
	has_impact BOOLEAN NOT NULL,
	is_hard BOOLEAN NOT NULL,
	is_new BOOLEAN NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
