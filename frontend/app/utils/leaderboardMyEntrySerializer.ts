import { type MyEntry, type FormattedMyEntry } from "~/types/leaderboard";

export default (entry: MyEntry): FormattedMyEntry =>
{
	return {
		name            : entry.name,
		score           : entry.score,
		position        : entry.position,
		avatarURL       : entry.avatar_url,
		isInLeaderboard : entry.is_in_leaderboard
	};
};