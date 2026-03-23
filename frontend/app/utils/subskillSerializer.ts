import { type Subskill, type FormattedSubskill } from "~/types/skill";

export default (subskills: Subskill[]): FormattedSubskill[] =>
{
	return subskills.map(subskill =>
		(
			{
				id                  : subskill.id,
				name                : subskill.name,
				totalXp             : subskill.total_xp,
				nextLevel           : subskill.next_level,
				currentLevel        : subskill.current_level,
				xpToNextLevel       : subskill.xp_to_next_level,
				nextLevelStartXp    : subskill.next_level_start_xp,
				currentLevelStartXp : subskill.current_level_start_xp
			}
		)
	)
};