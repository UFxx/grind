-- +goose Up
-- +goose StatementBegin
INSERT INTO base_skills (name) VALUES
	('Intellect'),
	('Vitality'),
	('Charisma'),
	('Discipline'),
	('Craft'),
	('Wisdom')
ON CONFLICT (name) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM base_skills WHERE name IN (
	'Intellect',
	'Vitality',
	'Charisma',
	'Discipline',
	'Craft',
	'Wisdom'
);
-- +goose StatementEnd
