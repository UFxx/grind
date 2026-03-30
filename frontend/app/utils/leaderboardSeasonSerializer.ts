import type { FormattedSeason, Season } from "~/types/leaderboard";

export default (season: Season): FormattedSeason =>
{
	return {
		id          : season.id,
		name        : season.name,
		periodEnd   : season.period_end,
		periodStart : season.period_start
	};
};