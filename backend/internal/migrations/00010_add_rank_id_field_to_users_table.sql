-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS users
ADD COLUMN IF NOT EXISTS rank_id UUID NOT NULL
	REFERENCES ranks(id) ON DELETE RESTRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS users
DROP COLUMN IF EXISTS rank_id;
-- +goose StatementEnd
