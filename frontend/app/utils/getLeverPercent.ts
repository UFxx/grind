import { type FormattedUserLevel } from "~/types/user";

export default (level: FormattedUserLevel): number =>
{
	const { totalXp, currentLevelStartXp, nextLevelStartXp } = level;

	const progress   = totalXp - currentLevelStartXp;
	const levelRange = nextLevelStartXp - currentLevelStartXp;

	return Math.ceil((progress / levelRange) * 100);
};