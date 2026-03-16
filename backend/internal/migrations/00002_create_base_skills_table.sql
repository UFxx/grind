-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS base_skills(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	code TEXT NOT NULL UNIQUE,
	display_name TEXT NOT NULL UNIQUE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS base_skills;
-- +goose StatementEnd
