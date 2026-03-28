import { type Entry, type FormattedEntry } from "~/types/leaderboard";

export default (entries: Entry[]): FormattedEntry[] =>
{
	return entries.map(entry =>
		{
			return {
				position: entry.position,
				name: entry.name,
				avatarURL: entry.avatar_url,
				score: entry.score
			};
		}
	)
};