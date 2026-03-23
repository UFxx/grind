import { type Skill, type FormattedSkill } from "~/types/skill"

export default (skills: Skill[]): FormattedSkill[] =>
{
	return skills.map(skill =>
		(
			{
				id                  : skill.id,
				name                : skill.name,
				totalXp             : skill.total_xp,
				nextLevel           : skill.next_level,
				currentLevel        : skill.current_level,
				xpToNextLevel       : skill.xp_to_next_level,
				nextLevelStartXp    : skill.next_level_start_xp,
				currentLevelStartXp : skill.current_level_start_xp,
				subskills           : skill.items?.length ? subskillSerializer(skill.items) : undefined
			}
		)
	)
}