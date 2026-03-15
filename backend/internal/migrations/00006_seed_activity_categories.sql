-- +goose Up
-- +goose StatementBegin
INSERT INTO activity_categories (code, display_name) VALUES
	('learning', 'Learning'),
	('deep_work', 'Deep Work'),
	('project_work', 'Project Work'),
	('creative_work', 'Creative Work'),
	('reading', 'Reading'),
	('writing', 'Writing'),
	('communication', 'Communication'),
	('networking', 'Networking'),
	('health_workout', 'Health Workout'),
	('health_recovery', 'Health Recovery'),
	('nutrition', 'Nutrition'),
	('routine_household', 'Routine Household'),
	('planning', 'Planning'),
	('reflection_review', 'Reflection Review'),
	('finance_admin', 'Finance Admin'),
	('digital_hygiene', 'Digital Hygiene'),
	('career_growth', 'Career Growth'),
	('contribution_community', 'Contribution Community')
ON CONFLICT (code) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM activity_categories WHERE code IN (
	'learning',
	'deep_work',
	'project_work',
	'creative_work',
	'reading',
	'writing',
	'communication',
	'networking',
	'health_workout',
	'health_recovery',
	'nutrition',
	'routine_household',
	'planning',
	'reflection_review',
	'finance_admin',
	'digital_hygiene',
	'career_growth',
	'contribution_community'
);
-- +goose StatementEnd
