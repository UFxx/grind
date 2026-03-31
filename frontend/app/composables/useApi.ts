import auth from "~/api/auth";
import user from "~/api/user";
import other from "~/api/other";
import activity from "~/api/activity";
import leaderboard from "~/api/leaderboard";

export const useApi = () =>
{
	return {
		auth,
		user,
		other,
		activity,
		leaderboard
	};
}