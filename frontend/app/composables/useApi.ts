import auth from "~/api/auth";
import leaderboard from "~/api/leaderboard";
import user from "~/api/user";

export const useApi = () =>
{
	return {
		auth,
		user,
		leaderboard,
	};
}