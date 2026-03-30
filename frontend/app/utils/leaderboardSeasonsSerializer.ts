import { type Season, type FormattedSeason } from '~/types/leaderboard';

export default (seasons: Season[]): FormattedSeason[] =>
{
	return seasons.map(season => leaderboardSeasonSerializer(season));
};