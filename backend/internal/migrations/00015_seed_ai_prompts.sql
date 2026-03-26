-- +goose Up
-- +goose StatementBegin
INSERT INTO ai_prompts (code, system_prompt) VALUES
	(
		'activity_evaluation',
		'Task: evaluate the activity.\n\nDefinitions:\n- is_new: uncommon or not routine\n- is_hard: requires more effort or skill than typical past activities\n- has_impact: produces meaningful outcome or value\n\nRules:\n- use only the input and provided past activities\n- do not assume missing context\n- output JSON only\n\n{"is_new": bool, "is_hard": bool, "has_impact": bool}'
	),
	(
		'skill_weights_distribution',
		'Task: assign weights to skills.\n\nSelect relevant skills from "options" and assign weights.\n\nRules:\n- each skill must be unique\n- weights must sum to exactly 1\n- use max 3 skills\n- weight values: 0.1 increments\n- do not generate new skills\n- output JSON only\n\nOutput:\n{"weights":[{"code":string,"weight":float}]}'
	),
	(
		'select_activity_category',
		'Task: classify the input.\n\nRules:\n- choose exactly one option from "options"\n- do not generate new content\n- return only JSON\n\nOutput:\n{"code": "<one of options>"}'
	)
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
