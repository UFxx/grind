import { type Season, type FormattedSeason } from '~/types/leaderboard';

export default (seasons: Season[]): FormattedSeason[] =>
{
	return seasons.map(season =>
		{
			return {
				id        : season.id,
				name      : season.name,
				periodStart : season.period_start,
				periodEnd   : season.period_end,
			};
		}
	)
};