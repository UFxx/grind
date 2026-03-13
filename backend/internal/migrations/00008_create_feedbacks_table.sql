-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS feedbacks (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	sender TEXT NOT NULL,
	rating INT NOT NULL,
	message TEXT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS feedbacks;
-- +goose StatementEnd
