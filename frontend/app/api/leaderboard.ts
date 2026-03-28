import { type SuccessResponse } from "~/types/common";
import { type Season, type SeasonDetails } from "~/types/leaderboard";

export default {
	// Fetchs
	fetchSeasons		: async () => await useRequest<SuccessResponse<Season[]>>('/leaderboards/seasons'),
	fetchSeasonDetails	: async (id: string) => await useRequest<SuccessResponse<SeasonDetails>>(`/leaderboards/seasons/${id}`)
};