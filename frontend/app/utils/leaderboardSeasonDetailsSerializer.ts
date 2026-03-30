import type { FormattedSeasonDetails, SeasonDetails } from "~/types/leaderboard";

export default (seasonDetails: SeasonDetails): FormattedSeasonDetails =>
{
	return {
		me      : leaderboardMyEntrySerializer(seasonDetails.me),
		cta     : seasonDetails.cta,
		season  : leaderboardSeasonSerializer(seasonDetails.season),
		entries : leaderboardEntriesSerializer(seasonDetails.entries)
	}
};