import { type LevelProgress } from "~/types/level";

export default (level: LevelProgress): number =>
{
	const { totalXp, currentLevelStartXp, nextLevelStartXp } = level;

	const progress   = totalXp - currentLevelStartXp;
	const levelRange = nextLevelStartXp - currentLevelStartXp;

	return Math.ceil((progress / levelRange) * 100);
};