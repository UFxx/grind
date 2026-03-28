import { type SuccessResponse } from "~/types/common";
import { type Season, type SeasonDetail } from "~/types/leaderboard";

export default {
	// Fetchs
	fetchSeasons		: async () => await useRequest<SuccessResponse<Season[]>>('/leaderboards/seasons'),
	fetchSeasonDetail	: async (id: string) => await useRequest<SuccessResponse<SeasonDetail>>(`/leaderboards/seasons/${id}`)
};