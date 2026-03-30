import auth from "~/api/auth";
import user from "~/api/user";
import leaderboard from "~/api/leaderboard";

export const useApi = () =>
{
	return {
		auth,
		user,
		leaderboard,
	};
}