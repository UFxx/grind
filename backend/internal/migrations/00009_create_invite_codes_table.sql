-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS invite_codes(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	code TEXT NOT NULL UNIQUE,
	created_by UUID NOT NULL
		REFERENCES users(id) ON DELETE CASCADE,
	uses INT NOT NULL DEFAULT 0,
	max_uses INT NOT NULL DEFAULT 1,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS invite_codes;
-- +goose StatementEnd
