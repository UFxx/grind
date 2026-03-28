import { type MyEntry, type FormattedMyEntry } from "~/types/leaderboard";

export default (entry: MyEntry): FormattedMyEntry =>
{
	return {
		position: entry.position,
		name: entry.name,
		avatarURL: entry.avatar_url,
		score: entry.score,
		isInLeaderboard: entry.is_in_leaderboard
	};
};