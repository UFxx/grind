import auth from "~/api/auth";
import user from "~/api/user";
import activity from "~/api/activity";
import leaderboard from "~/api/leaderboard";

export const useApi = () =>
{
	return {
		auth,
		user,
		activity,
		leaderboard
	};
}