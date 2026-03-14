-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS activities(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	description TEXT NOT NULL,
	user_id UUID NOT NULL
		REFERENCES users(id) ON DELETE CASCADE,
	source TEXT NOT NULL,
	activity_type_id UUID NOT NULL
		REFERENCES activity_types(id) ON DELETE CASCADE,
	has_impact BOOLEAN NOT NULL,
	is_hard BOOLEAN NOT NULL,
	is_new BOOLEAN NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS activities;
-- +goose StatementEnd
