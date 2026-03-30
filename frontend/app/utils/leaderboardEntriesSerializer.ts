import { type Entry, type FormattedEntry } from "~/types/leaderboard";

export default (entries: Entry[]): FormattedEntry[] =>
{
	return entries.map(entry =>
		{
			return {
				name      : entry.name,
				score     : entry.score,
				position  : entry.position,
				avatarURL : entry.avatar_url
			};
		}
	)
};