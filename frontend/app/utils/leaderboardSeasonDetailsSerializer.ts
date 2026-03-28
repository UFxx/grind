import type { FormattedSeasonDetails, SeasonDetails } from "~/types/leaderboard";

export default (seasonDetails: SeasonDetails): FormattedSeasonDetails =>
{
	return {
		season: leaderboardSeasonSerializer(seasonDetails.season),
		cta: seasonDetails.cta,
		entries: leaderboardEntriesSerializer(seasonDetails.entries),
		me: leaderboardMyEntrySerializer(seasonDetails.me)
	}
};