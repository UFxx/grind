-- +goose Up
-- +goose StatementBegin
INSERT INTO ranks (code, display_name) VALUES
	('novice', 'Novice')
ON CONFLICT (code) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM ranks WHERE code IN (
	'novice'
);
-- +goose StatementEnd
