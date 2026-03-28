import type { FormattedSeason, Season } from "~/types/leaderboard";

export default (season: Season): FormattedSeason =>
{
	return {
		id: season.id,
		name: season.name,
		periodStart: season.period_start,
		periodEnd: season.period_end
	};
};