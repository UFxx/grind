-- +goose Up
-- +goose StatementBegin
INSERT INTO ai_prompts (code, system_prompt, output_schema_prompt) VALUES
	('activity_evaluation', 'You must decide whether the activity is new, challenging, or useful based on other similar activities.', '{"is_new":bool,"is_hard":bool,"has_impact":bool}'),
	('skill_weights_distribution', 'You must select the most appropriate skills and weights from the list based on the input. Skills must not be repeated. The sum of the weights must strictly equal 1.', '{"weights":[{"code":string,"weight":float}]}'),
	('select_activity_category', 'You must choose the most suitable option based on the input.', '{"code":string}')
ON CONFLICT (code) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM ai_prompts WHERE code IN (
	'activity_evaluation',
	'skill_weights_distribution',
	'select_activity_category'
);
-- +goose StatementEnd
