-- +goose Up
-- +goose StatementBegin
INSERT INTO base_skills (code, display_name) VALUES
	('intellect', 'Intellect'),
	('vitality', 'Vitality'),
	('charisma', 'Charisma'),
	('discipline', 'Discipline'),
	('craft', 'Craft'),
	('wisdom', 'Wisdom')
ON CONFLICT (code) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM base_skills WHERE code IN (
	'intellect',
	'vitality',
	'charisma',
	'discipline',
	'craft',
	'wisdom'
);
-- +goose StatementEnd
